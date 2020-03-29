package simapp

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func TestSimAppExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewSimApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, 0)

	genesisState := NewDefaultGenesisState()
	stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
	require.NoError(t, err)

	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	app2 := NewSimApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, 0)
	_, _, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
