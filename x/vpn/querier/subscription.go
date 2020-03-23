package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func querySubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	
	subscription, found := k.GetSubscription(ctx, params.ID)
	if !found {
		return nil, nil
	}
	
	res, err := types.ModuleCdc.MarshalJSON(subscription)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func querySubscriptionsOfNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionsOfNodePrams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	
	subscriptions := k.GetSubscriptionsOfNode(ctx, params.ID)
	
	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func querySubscriptionsOfAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySubscriptionsOfAddressParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	
	subscriptions := k.GetSubscriptionsOfAddress(ctx, params.Address)
	
	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func queryAllSubscriptions(ctx sdk.Context, k keeper.Keeper) ([]byte, error) {
	subscriptions := k.GetAllSubscriptions(ctx)
	
	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func querySessionsCountOfSubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QuerySessionsCountOfSubscriptionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	count := k.GetSessionsCountOfSubscription(ctx, params.ID)
	
	res, err := types.ModuleCdc.MarshalJSON(count)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	
	return res, nil
}
