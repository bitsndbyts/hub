package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	
	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryDepositOfAddress:
			return queryDepositOfAddress(ctx, req, k)
		case types.QueryAllDeposits:
			return queryAllDeposits(ctx, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}
