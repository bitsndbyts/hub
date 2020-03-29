package simapp

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

type App interface {
	Name() string

	Codec() *codec.Codec

	BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock

	EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock

	InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain

	LoadHeight(height int64) error

	ExportAppStateAndValidators(
		forZeroHeight bool, jailWhiteList []string,
	) (json.RawMessage, []tmtypes.GenesisValidator, error)

	ModuleAccountAddrs() map[string]bool

	SimulationManager() *module.SimulationManager
}
