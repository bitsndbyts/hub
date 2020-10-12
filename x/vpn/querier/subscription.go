package querier

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func querySubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscription, found := k.GetSubscription(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(subscription)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsOfNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsOfNodePrams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptionsOfNode(ctx, params.ID)

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsOfAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsOfAddressParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptionsOfAddress(ctx, params.Address)

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryAllSubscriptions(ctx sdk.Context, k keeper.Keeper) ([]byte, sdk.Error) {
	subscriptions := k.GetAllSubscriptions(ctx)

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySessionsCountOfSubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionsCountOfSubscriptionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	count := k.GetSessionsCountOfSubscription(ctx, params.ID)

	res, err := types.ModuleCdc.MarshalJSON(count)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryFreeSessionsCountOfClient(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryFreeSessionsCountOfClientParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	count := k.GetFreeSessionsCountOfClient(ctx, params.Client, params.NodeID)

	res, err := types.ModuleCdc.MarshalJSON(count)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryFreeSessionsOfClient(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryFreeSessionsOfClientParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	fmt.Println("stage 1111")
	session := k.GetFreeSessionsByClientID(ctx, params.Client)

	res, err := types.ModuleCdc.MarshalJSON(session)
	if err != nil {
		fmt.Println("error stageee")
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryFreeSessionsOfNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryFreeSessionsOfNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	count := k.GetFreeSessionsByNodeAddress(ctx, params.NodeAddress)

	res, err := types.ModuleCdc.MarshalJSON(count)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
