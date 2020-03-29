package simapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	tmkv "github.com/tendermint/tendermint/libs/kv"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func SetupSimulation(dirPrefix, dbName string) (simulation.Config, dbm.DB, string, log.Logger, bool, error) {
	if !FlagEnabledValue {
		return simulation.Config{}, nil, "", nil, true, nil
	}

	config := NewConfigFromFlags()
	config.ChainID = helpers.SimAppChainID

	var logger log.Logger
	if FlagVerboseValue {
		logger = log.TestingLogger()
	} else {
		logger = log.NewNopLogger()
	}

	dir, err := ioutil.TempDir("", dirPrefix)
	if err != nil {
		return simulation.Config{}, nil, "", nil, false, err
	}

	db, err := sdk.NewLevelDB(dbName, dir)
	if err != nil {
		return simulation.Config{}, nil, "", nil, false, err
	}

	return config, db, dir, logger, false, nil
}

func SimulationOperations(app App, cdc *codec.Codec, config simulation.Config) []simulation.WeightedOperation {
	simState := module.SimulationState{
		AppParams: make(simulation.AppParams),
		Cdc:       cdc,
	}

	if config.ParamsFile != "" {
		bz, err := ioutil.ReadFile(config.ParamsFile)
		if err != nil {
			panic(err)
		}

		app.Codec().MustUnmarshalJSON(bz, &simState.AppParams)
	}

	simState.ParamChanges = app.SimulationManager().GenerateParamChanges(config.Seed)
	simState.Contents = app.SimulationManager().GetProposalContents(simState)
	return app.SimulationManager().WeightedOperations(simState)
}

func CheckExportSimulation(
	app App, config simulation.Config, params simulation.Params,
) error {
	if config.ExportStatePath != "" {
		fmt.Println("exporting app state...")
		appState, _, err := app.ExportAppStateAndValidators(false, nil)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(config.ExportStatePath, []byte(appState), 0644); err != nil {
			return err
		}
	}

	if config.ExportParamsPath != "" {
		fmt.Println("exporting simulation params...")
		paramsBz, err := json.MarshalIndent(params, "", " ")
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(config.ExportParamsPath, paramsBz, 0644); err != nil {
			return err
		}
	}
	return nil
}

func PrintStats(db dbm.DB) {
	fmt.Println("\nLevelDB Stats")
	fmt.Println(db.Stats()["leveldb.stats"])
	fmt.Println("LevelDB cached block size", db.Stats()["leveldb.cachedblock"])
}

func GetSimulationLog(storeName string, sdr sdk.StoreDecoderRegistry, cdc *codec.Codec, kvAs, kvBs []tmkv.Pair) (log string) {
	for i := 0; i < len(kvAs); i++ {

		if len(kvAs[i].Value) == 0 && len(kvBs[i].Value) == 0 {
			// skip if the value doesn't have any bytes
			continue
		}

		decoder, ok := sdr[storeName]
		if ok {
			log += decoder(cdc, kvAs[i], kvBs[i])
		} else {
			log += fmt.Sprintf("store A %X => %X\nstore B %X => %X\n", kvAs[i].Key, kvAs[i].Value, kvBs[i].Key, kvBs[i].Value)
		}
	}

	return
}
