package app

import (
	"encoding/json"
	"io"
	"os"
	
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	db "github.com/tendermint/tm-db"
	
	"github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/version"
	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn"
)

const (
	appName = "Sentinel Hub App"
)

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.sentinel-hubcli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.sentinel-hubd")
	
	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(client.ProposalHandler, distribution.ProposalHandler),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
		deposit.AppModuleBasic{},
		vpn.AppModuleBasic{},
	)
	
	moduleAccountPermissions = map[string][]string{
		auth.FeeCollectorName:     nil,
		distribution.ModuleName:   nil,
		mint.ModuleName:           {supply.Minter},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		gov.ModuleName:            {supply.Burner},
		deposit.ModuleName:        nil,
	}
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	
	sdk.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
	ModuleBasics.RegisterCodec(cdc)
	
	return cdc
}

type HubApp struct {
	*baseapp.BaseApp
	cdc *codec.Codec
	
	invCheckPeriod uint
	
	keys          map[string]*sdk.KVStoreKey
	transientKeys map[string]*sdk.TransientStoreKey
	
	subspaces map[string]params.Subspace
	
	accountKeeper      auth.AccountKeeper
	bankKeeper         bank.Keeper
	supplyKeeper       supply.Keeper
	stakingKeeper      staking.Keeper
	slashingKeeper     slashing.Keeper
	mintKeeper         mint.Keeper
	distributionKeeper distribution.Keeper
	govKeeper          gov.Keeper
	crisisKeeper       crisis.Keeper
	paramsKeeper       params.Keeper
	depositKeeper      deposit.Keeper
	vpnKeeper          vpn.Keeper
	
	mm *module.Manager
}

// nolint:funlen
func NewHubApp(logger log.Logger, db db.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*baseapp.BaseApp)) *HubApp {
	cdc := MakeCodec()
	
	bApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	
	keys := sdk.NewKVStoreKeys(
		baseapp.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distribution.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, deposit.StoreKey,
		vpn.StoreKeyNode, vpn.StoreKeySubscription, vpn.StoreKeySession, vpn.StoreKeyResolver,
	)
	
	transientKeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)
	
	var app = &HubApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		transientKeys:  transientKeys,
		subspaces:      make(map[string]params.Subspace),
	}
	
	app.paramsKeeper = params.NewKeeper(app.cdc,
		keys[params.StoreKey],
		transientKeys[params.TStoreKey], )
	
	app.subspaces[auth.ModuleName] = app.paramsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.paramsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.paramsKeeper.Subspace(staking.DefaultParamspace)
	app.subspaces[mint.ModuleName] = app.paramsKeeper.Subspace(mint.DefaultParamspace)
	app.subspaces[distribution.ModuleName] = app.paramsKeeper.Subspace(distribution.DefaultParamspace)
	app.subspaces[slashing.ModuleName] = app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	app.subspaces[gov.ModuleName] = app.paramsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
	app.subspaces[crisis.ModuleName] = app.paramsKeeper.Subspace(crisis.DefaultParamspace)
	app.subspaces[vpn.ModuleName] = app.paramsKeeper.Subspace(vpn.ModuleName)
	
	app.accountKeeper = auth.NewAccountKeeper(app.cdc,
		keys[auth.StoreKey],
		app.subspaces[auth.ModuleName],
		auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper,
		app.subspaces[bank.ModuleName],
		app.ModuleAccountAddrs())
	app.supplyKeeper = supply.NewKeeper(app.cdc,
		keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		moduleAccountPermissions)
	stakingKeeper := staking.NewKeeper(app.cdc,
		keys[staking.StoreKey],
		app.supplyKeeper,
		app.subspaces[staking.ModuleName], )
	app.mintKeeper = mint.NewKeeper(app.cdc,
		keys[mint.StoreKey],
		app.subspaces[mint.ModuleName],
		&stakingKeeper,
		app.supplyKeeper,
		auth.FeeCollectorName)
	app.distributionKeeper = distribution.NewKeeper(app.cdc,
		keys[distribution.StoreKey],
		app.subspaces[distribution.ModuleName],
		&stakingKeeper,
		app.supplyKeeper,
		auth.FeeCollectorName,
		app.ModuleAccountAddrs())
	app.slashingKeeper = slashing.NewKeeper(app.cdc,
		keys[slashing.StoreKey],
		&stakingKeeper,
		app.subspaces[slashing.ModuleName], )
	app.crisisKeeper = crisis.NewKeeper(app.subspaces[crisis.ModuleName],
		invCheckPeriod,
		app.supplyKeeper,
		auth.FeeCollectorName)
	
	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper))
	
	app.govKeeper = gov.NewKeeper(app.cdc,
		keys[gov.StoreKey],
		app.subspaces[gov.ModuleName],
		app.supplyKeeper,
		&stakingKeeper,
		govRouter)
	
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.distributionKeeper.Hooks(), app.slashingKeeper.Hooks()))
	
	app.depositKeeper = deposit.NewKeeper(app.cdc,
		keys[deposit.StoreKey],
		app.supplyKeeper)
	
	app.vpnKeeper = vpn.NewKeeper(app.cdc,
		keys[vpn.StoreKeyNode],
		keys[vpn.StoreKeySubscription],
		keys[vpn.StoreKeySession],
		keys[vpn.StoreKeyResolver],
		app.subspaces[vpn.ModuleName],
		app.depositKeeper)
	
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		crisis.NewAppModule(&app.crisisKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distribution.NewAppModule(app.distributionKeeper, app.accountKeeper, app.supplyKeeper, app.stakingKeeper),
		gov.NewAppModule(app.govKeeper, app.accountKeeper, app.supplyKeeper),
		mint.NewAppModule(app.mintKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),
		deposit.NewAppModule(app.depositKeeper),
		vpn.NewAppModule(app.vpnKeeper),
	)
	
	app.mm.SetOrderBeginBlockers(mint.ModuleName, distribution.ModuleName, slashing.ModuleName)
	app.mm.SetOrderEndBlockers(crisis.ModuleName, gov.ModuleName, staking.ModuleName, vpn.ModuleName)
	app.mm.SetOrderInitGenesis(
		distribution.ModuleName, staking.ModuleName,
		auth.ModuleName, bank.ModuleName, slashing.ModuleName, gov.ModuleName,
		mint.ModuleName, supply.ModuleName, crisis.ModuleName, genutil.ModuleName,
		deposit.ModuleName, vpn.ModuleName,
	)
	
	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
	app.MountKVStores(keys)
	app.MountTransientStores(transientKeys)
	
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		auth.NewAnteHandler(app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer))
	app.SetEndBlocker(app.EndBlocker)
	
	if loadLatest {
		if err := app.LoadLatestVersion(app.keys[baseapp.MainStoreKey]); err != nil {
			tmos.Exit(err.Error())
		}
	}
	
	return app
}

func (app *HubApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *HubApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *HubApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &state)
	
	return app.mm.InitGenesis(ctx, state)
}

func (app *HubApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[baseapp.MainStoreKey])
}

func (app *HubApp) ModuleAccountAddrs() map[string]bool {
	moduleAccounts := make(map[string]bool)
	for acc := range moduleAccountPermissions {
		moduleAccounts[supply.NewModuleAddress(acc).String()] = true
	}
	
	return moduleAccounts
}
