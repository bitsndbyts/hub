package simapp

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	simapparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func AppStateFn(cdc *codec.Codec, simManager *module.SimulationManager) simulation.AppStateFn {
	return func(r *rand.Rand, accs []simulation.Account, config simulation.Config,
	) (appState json.RawMessage, simAccs []simulation.Account, chainID string, genesisTimestamp time.Time) {

		if FlagGenesisTimeValue == 0 {
			genesisTimestamp = simulation.RandTimestamp(r)
		} else {
			genesisTimestamp = time.Unix(FlagGenesisTimeValue, 0)
		}

		chainID = config.ChainID
		switch {
		case config.ParamsFile != "" && config.GenesisFile != "":
			panic("cannot provide both a genesis file and a params file")

		case config.GenesisFile != "":
			genesisDoc, accounts := AppStateFromGenesisFileFn(r, cdc, config.GenesisFile)

			if FlagGenesisTimeValue == 0 {
				genesisTimestamp = genesisDoc.GenesisTime
			}

			appState = genesisDoc.AppState
			chainID = genesisDoc.ChainID
			simAccs = accounts

		case config.ParamsFile != "":
			appParams := make(simulation.AppParams)
			bz, err := ioutil.ReadFile(config.ParamsFile)
			if err != nil {
				panic(err)
			}

			cdc.MustUnmarshalJSON(bz, &appParams)
			appState, simAccs = AppStateRandomizedFn(simManager, r, cdc, accs, genesisTimestamp, appParams)

		default:
			appParams := make(simulation.AppParams)
			appState, simAccs = AppStateRandomizedFn(simManager, r, cdc, accs, genesisTimestamp, appParams)
		}

		return appState, simAccs, chainID, genesisTimestamp
	}
}

func AppStateRandomizedFn(
	simManager *module.SimulationManager, r *rand.Rand, cdc *codec.Codec,
	accs []simulation.Account, genesisTimestamp time.Time, appParams simulation.AppParams,
) (json.RawMessage, []simulation.Account) {
	numAccs := int64(len(accs))
	genesisState := NewDefaultGenesisState()

	var initialStake, numInitiallyBonded int64
	appParams.GetOrGenerate(
		cdc, simapparams.StakePerAccount, &initialStake, r,
		func(r *rand.Rand) { initialStake = r.Int63n(1e12) },
	)
	appParams.GetOrGenerate(
		cdc, simapparams.InitiallyBondedValidators, &numInitiallyBonded, r,
		func(r *rand.Rand) { numInitiallyBonded = int64(r.Intn(300)) },
	)

	if numInitiallyBonded > numAccs {
		numInitiallyBonded = numAccs
	}

	fmt.Printf(
		`Selected randomly generated parameters for simulated genesis:
{
  stake_per_account: "%d",
  initially_bonded_validators: "%d"
}
`, initialStake, numInitiallyBonded,
	)

	simState := &module.SimulationState{
		AppParams:    appParams,
		Cdc:          cdc,
		Rand:         r,
		GenState:     genesisState,
		Accounts:     accs,
		InitialStake: initialStake,
		NumBonded:    numInitiallyBonded,
		GenTimestamp: genesisTimestamp,
	}

	simManager.GenerateGenesisStates(simState)

	appState, err := cdc.MarshalJSON(genesisState)
	if err != nil {
		panic(err)
	}

	return appState, accs
}

func AppStateFromGenesisFileFn(r io.Reader, cdc *codec.Codec, genesisFile string) (tmtypes.GenesisDoc, []simulation.Account) {
	bytes, err := ioutil.ReadFile(genesisFile)
	if err != nil {
		panic(err)
	}

	var genesis tmtypes.GenesisDoc
	cdc.MustUnmarshalJSON(bytes, &genesis)

	var appState GenesisState
	cdc.MustUnmarshalJSON(genesis.AppState, &appState)

	var authGenesis auth.GenesisState
	if appState[auth.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[auth.ModuleName], &authGenesis)
	}

	newAccs := make([]simulation.Account, len(authGenesis.Accounts))
	for i, acc := range authGenesis.Accounts {
		privkeySeed := make([]byte, 15)
		if _, err := r.Read(privkeySeed); err != nil {
			panic(err)
		}

		privKey := secp256k1.GenPrivKeySecp256k1(privkeySeed)
		simAcc := simulation.Account{PrivKey: privKey, PubKey: privKey.PubKey(), Address: acc.GetAddress()}
		newAccs[i] = simAcc
	}

	return genesis, newAccs
}
