package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func queryDepositOfAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryDepositOfAddressPrams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	deposit, found := k.GetDeposit(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(deposit)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAllDeposits(ctx sdk.Context, k keeper.Keeper) ([]byte, error) {
	deposits := k.GetAllDeposits(ctx)

	res, err := types.ModuleCdc.MarshalJSON(deposits)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
