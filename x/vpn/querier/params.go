package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func queryParameters(ctx sdk.Context, k keeper.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)
	
	res, err := types.ModuleCdc.MarshalJSON(params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}
