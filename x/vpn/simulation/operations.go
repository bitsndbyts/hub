package simulation

import (
	"fmt"
	"math/rand"
	"reflect"
	
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

const (
	OpWeightMsgRegisterNode            = "op_weight_msg_register_node"
	OpWeightMsgUpdateNodeInfo          = "op_weight_msg_update_node_info"
	OpWeightMsgAddFreeClient           = "op_weight_msg_add_free_client"
	OpWeightMsgRemoveFreeClient        = "op_weight_remove_free_client"
	OpWeightMsgRegisterVPNOnResolver   = "op_weight_regiser_vpn_on_resolver"
	OpWeightMsgDeregisterVPNOnResolver = "op_weight_deregister_vpn_on_resolver"
	OpWeightMsgDeregisterNode          = "op_weight_msg_deregister_node"
	OpWeightMsgStartSubscription       = "op_weight_msg_start_subscription"
	OpWeightMsgEndSubscription         = "op_weight_msg_end_subscription"
	OpWeightMsgUpdateSessionInfo       = "op_weight_msg_update_session"
	OpWeightMsgEndSession              = "op_weight_msg_end_session"
	OpWeightMsgRegisterResolver        = "op_weight_msg_register_resolver"
	OpWeightMsgUpdateResolverInfo      = "op_weight_msg_update_resolver_info"
	OpWeightMsgDeregisterResolver      = "op_weight_msg_deregister_resolver"
)

func WeightedOperations(
	appParams simulation.AppParams, cdc *codec.Codec, ak types.AccountKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgRegisterNode   int
		weightMsgUpdateNodeInfo int
		weightMsgDeregisterNode int
		
		weightMsgRegisterResolver   int
		weightMsgUpdateResolverInfo int
		weightMsgDeregisterResolver int
		
		weightMsgRegisterVPNOnResolver   int
		weightMsgDeregisterVPNOnResolver int
		
		weightMsgAddFreeClient     int
		weightMsgRemoveFreeClient  int
		weightMsgStartSubscription int
		weightMsgEndSubscription   int
		
		weightMsgUpdateSessionInfo int
		weightMsgEndSession        int
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgRegisterNode, &weightMsgRegisterNode, nil,
		func(r *rand.Rand) {
			weightMsgRegisterNode = DefaultWeightMsgRegisterNode
		},
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateNodeInfo, &weightMsgUpdateNodeInfo, nil,
		func(r *rand.Rand) {
			weightMsgUpdateNodeInfo = DefaultWeightMsgUpdateNodeInfo
		},
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgDeregisterNode, &weightMsgDeregisterNode, nil,
		func(r *rand.Rand) {
			weightMsgDeregisterNode = DefaultWeightMsgDeregisterNode
		},
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgRegisterResolver, &weightMsgRegisterResolver, nil,
		func(r *rand.Rand) {
			weightMsgRegisterResolver = DefaultWeightMsgRegisterResolver
		},
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateResolverInfo, &weightMsgUpdateResolverInfo, nil,
		func(r *rand.Rand) {
			weightMsgUpdateResolverInfo = DefaultWeightMsgUpdateResolverInfo
		},
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgDeregisterResolver, &weightMsgDeregisterResolver, nil,
		func(r *rand.Rand) {
			weightMsgDeregisterResolver = DefaultWeightMsgDeregisterResolver
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightMsgRegisterVPNOnResolver, &weightMsgRegisterVPNOnResolver, nil,
		func(r *rand.Rand) {
			weightMsgRegisterVPNOnResolver = DefaultWeightMsgRegisterVPNOnResolver
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightMsgDeregisterVPNOnResolver, &weightMsgDeregisterVPNOnResolver, nil,
		func(r *rand.Rand) {
			weightMsgDeregisterVPNOnResolver = DefaultWeightMsgDeregisterVPNOnResolver
		},
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgAddFreeClient, &weightMsgAddFreeClient, nil,
		func(r *rand.Rand) {
			weightMsgAddFreeClient = DefaultWeightMsgAddFreeClient
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveFreeClient, &weightMsgRemoveFreeClient, nil,
		func(r *rand.Rand) {
			weightMsgRemoveFreeClient = DefaultWeightMsgRemoveFreeClient
		},
	)
	
	appParams.GetOrGenerate(cdc, OpWeightMsgStartSubscription, &weightMsgStartSubscription, nil,
		func(r *rand.Rand) {
			weightMsgStartSubscription = DefaultWeightMsgStartSubscription
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightMsgEndSubscription, &weightMsgEndSubscription, nil,
		func(r *rand.Rand) {
			weightMsgEndSubscription = DefaultWeightMsgEndSubscription
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateSessionInfo, &weightMsgUpdateSessionInfo, nil,
		func(r *rand.Rand) {
			weightMsgUpdateSessionInfo = DefaultWeightMsgUpdateSessionInfo
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightMsgEndSession, &weightMsgEndSession, nil,
		func(r *rand.Rand) {
			weightMsgEndSession = DefaultWeightMsgEndSession
		},
	)
	
	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgRegisterNode,
			SimulateMsgRegisterNode(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateNodeInfo,
			SimulateMsgUpdateNodeInfo(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgDeregisterNode,
			SimulateMsgDeregisterNode(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRegisterResolver,
			SimulateMsgRegisterResolver(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateResolverInfo,
			SimulateMsgUpdateResolverInfo(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgDeregisterResolver,
			SimulateMsgDeregisterResolver(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRegisterVPNOnResolver,
			SimulateMsgRegisterVPNOnResolver(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgDeregisterVPNOnResolver,
			SimulateMsgDeregisterVPNOnResolver(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgAddFreeClient,
			SimulateMsgAddFreeClient(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRemoveFreeClient,
			SimulateMsgRemoveFreeClient(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgStartSubscription,
			SimulateMsgStartSubscription(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgEndSubscription,
			SimulateMsgEndSubscription(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateSessionInfo,
			SimulateMsgUpdateSessionInfo(ak, k),
		),
	}
}

func SimulateMsgRegisterNode(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account doesn't exist")
		}
		
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		coins = getRandomCoins(r)
		msg := *types.NewMsgRegisterNode(acc.GetAddress(),
			getRandomType(r), getRandomVersion(r), getRandomMoniker(r),
			coins, getRandomBandwidth(r), getRandomEncryption(r),
		)
		
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgUpdateNodeInfo(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if len(k.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		node := keeper.RandomNodeByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(node, types.Node{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if node.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		msg := *types.NewMsgUpdateNodeInfo(acc.GetAddress(), node.ID, getRandomType(r), getRandomVersion(r), getRandomMoniker(r),
			getRandomCoins(r), getRandomBandwidth(r), getRandomEncryption(r))
		
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgDeregisterNode(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account doesn't exist")
		}
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if len(k.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		node := keeper.RandomNodeByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(node, types.Node{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if node.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		msg := *types.NewMsgDeregisterNode(node.Owner, node.ID)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgRegisterResolver(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account doesn't exist")
		}
		
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		commission := sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 0, 100)), 2)
		msg := *types.NewMsgRegisterResolver(acc.GetAddress(), simulation.RandomDecAmount(r, commission))
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgUpdateResolverInfo(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		ak.SetAccount(ctx, acc)
		
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		if len(k.GetAllResolvers(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		resolver := keeper.RandomResolverByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(resolver, types.Resolver{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		commission := sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 0, 100)), 2)
		
		if resolver.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		msg := *types.NewMsgUpdateResolverInfo(acc.GetAddress(), resolver.ID, simulation.RandomDecAmount(r, commission))
		
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgDeregisterResolver(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account doesn't exist")
		}
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		if len(k.GetAllResolvers(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		resolver := keeper.RandomResolverByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(resolver, types.Resolver{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if resolver.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		msg := *types.NewMsgDeregisterResolver(acc.GetAddress(), resolver.ID)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgRegisterVPNOnResolver(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		ak.SetAccount(ctx, acc)
		
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		resolver := keeper.RandomResolver(r, ctx, k)
		if reflect.DeepEqual(resolver, types.Resolver{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if resolver.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		node := keeper.RandomNodeByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(node, types.Node{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if node.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		msg := *types.NewMsgRegisterVPNOnResolver(acc.GetAddress(), node.ID, resolver.ID)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgDeregisterVPNOnResolver(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account doesn't exist")
		}
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		if len(k.GetAllResolvers(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		node := keeper.RandomNodeByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(node, types.Node{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if node.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		resolverID := keeper.RandomPairOfResolverAndNode(r, ctx, k, node)
		if resolverID == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		msg := *types.NewMsgDeregisterVPNOnResolver(acc.GetAddress(), node.ID, resolverID)
		
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}
func SimulateMsgAddFreeClient(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		ak.SetAccount(ctx, acc)
		
		if len(k.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		node := keeper.RandomNodeByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(node, types.Node{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		clientAcc, _ := simulation.RandomAcc(r, accs)
		
		if node.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		msg := *types.NewMsgAddFreeClient(acc.GetAddress(), node.ID, clientAcc.Address)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgRemoveFreeClient(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if len(k.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		node := keeper.RandomNodeByAddress(r, ctx, k, acc.GetAddress())
		if reflect.DeepEqual(node, types.Node{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if node.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		client := keeper.RandomFreeClientOfNode(r, ctx, k, node.ID)
		if client.Empty() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		msg := *types.NewMsgRemoveFreeClient(acc.GetAddress(), node.ID, client)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgStartSubscription(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		coins := acc.SpendableCoins(ctx.BlockHeader().Time)
		
		node := keeper.RandomNode(r, ctx, k)
		if reflect.DeepEqual(node, types.Node{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if node.Status != types.StatusRegistered {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		resolver := keeper.RandomResolver(r, ctx, k)
		k.SetResolverOfNode(ctx, node.ID, resolver.ID)
		
		depositCoin := keeper.GetRandomCoin(r, coins)
		if reflect.DeepEqual(depositCoin, sdk.Coin{}) || depositCoin.IsZero() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		coin := node.FindPricePerGB(depositCoin.Denom)
		if reflect.DeepEqual(coin, sdk.Coin{}) || coin.IsZero() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		var (
			fees sdk.Coins
			err  error
		)
		coins, hasNeg := coins.SafeSub(sdk.Coins{depositCoin})
		if !hasNeg {
			fees, err = simulation.RandomFees(r, ctx, coins)
			if err != nil {
				return simulation.NoOpMsg(types.ModuleName), nil, err
			}
		}
		
		msg := *types.NewMsgStartSubscription(acc.GetAddress(), resolver.ID, node.ID, depositCoin)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgEndSubscription(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		
		if len(k.GetAllSubscriptions(ctx)) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		subscription := keeper.RandomSubscription(r, ctx, k)
		if reflect.DeepEqual(subscription, types.Subscription{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if !subscription.Client.Equals(acc.GetAddress()) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if subscription.Status != types.StatusActive {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
		_, found := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
		if found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		msg := *types.NewMsgEndSubscription(subscription.Client, subscription.ID)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}

func SimulateMsgUpdateSessionInfo(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		simAccount, _ := simulation.RandomAcc(r, accs)
		acc := ak.GetAccount(ctx, simAccount.Address)
		if acc == nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if err := acc.SetPubKey(simAccount.PubKey); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		ak.SetAccount(ctx, acc)
		
		coins := acc.SpendableCoins(ctx.BlockTime())
		var (
			fees sdk.Coins
			err  error
		)
		fees, err = simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		nodeOwnerAccount, _ := simulation.RandomAcc(r, accs)
		
		subscription := keeper.RandomSubscription(r, ctx, k, )
		if reflect.DeepEqual(subscription, types.Subscription{}) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		
		if !subscription.Client.Equals(acc.GetAddress()) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		subscription.Status = types.StatusActive
		subscription.Client = acc.GetAddress()
		k.SetSubscription(ctx, subscription)
		
		scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
		node, _ := k.GetNode(ctx, subscription.NodeID)
		node.Owner = nodeOwnerAccount.Address
		k.SetNode(ctx, node)
		
		bandwidth := getRandomBandwidth(r)
		
		bandWidthSignData := hub.NewBandwidthSignatureData(subscription.ID, scs, bandwidth)
		clientAccountSignedData, _ := simAccount.PrivKey.Sign(bandWidthSignData.Bytes())
		nodeOwnerAccountSignedData, _ := nodeOwnerAccount.PrivKey.Sign(bandWidthSignData.Bytes())
		
		clienStdSig := auth.StdSignature{
			PubKey:    simAccount.PubKey,
			Signature: clientAccountSignedData,
		}
		nodeOwnerStdSig := auth.StdSignature{
			PubKey:    nodeOwnerAccount.PubKey,
			Signature: nodeOwnerAccountSignedData,
		}
		
		if subscription.RemainingBandwidth.AnyLT(bandwidth) {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		msg := *types.NewMsgUpdateSessionInfo(acc.GetAddress(), subscription.ID, bandwidth,
			nodeOwnerStdSig, clienStdSig)
		if err := msg.ValidateBasic(); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("msg validation %v failed", err)
		}
		
		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAccount.PrivKey,
		)
		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		
		return simulation.NewOperationMsg(msg, true, "executed"), nil, nil
	}
}
