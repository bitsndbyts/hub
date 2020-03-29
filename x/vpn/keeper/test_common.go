package keeper

import (
	"math/rand"
	"testing"
	
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stretchr/testify/require"
	
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func CreateTestInput(t *testing.T, isCheckTx bool) (sdk.Context, Keeper, deposit.Keeper, bank.Keeper) {
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	keyAccount := sdk.NewKVStoreKey(auth.StoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	keyDeposit := sdk.NewKVStoreKey(deposit.StoreKey)
	keyNode := sdk.NewKVStoreKey(types.StoreKeyNode)
	keySubscription := sdk.NewKVStoreKey(types.StoreKeySubscription)
	keySession := sdk.NewKVStoreKey(types.StoreKeySession)
	keyResolver := sdk.NewKVStoreKey(types.StoreKeyResolver)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)
	
	mdb := db.NewMemDB()
	ms := store.NewCommitMultiStore(mdb)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyAccount, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyDeposit, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyNode, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyResolver, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySubscription, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySession, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, mdb)
	require.Nil(t, ms.LoadLatestVersion())
	
	depositAccount := supply.NewEmptyModuleAccount(types.ModuleName)
	blacklist := make(map[string]bool)
	blacklist[depositAccount.String()] = true
	accountPermissions := map[string][]string{
		deposit.ModuleName: nil,
	}
	
	cdc := MakeTestCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, isCheckTx, log.NewNopLogger())
	
	pk := params.NewKeeper(cdc, keyParams, tkeyParams)
	ak := auth.NewAccountKeeper(cdc, keyAccount, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bk := bank.NewBaseKeeper(ak, pk.Subspace(bank.DefaultParamspace), blacklist)
	sk := supply.NewKeeper(cdc, keySupply, ak, bk, accountPermissions)
	dk := deposit.NewKeeper(cdc, keyDeposit, sk)
	vk := NewKeeper(cdc, keyNode, keySubscription, keySession, keyResolver, pk.Subspace(DefaultParamspace), dk)
	
	sk.SetModuleAccount(ctx, depositAccount)
	vk.SetParams(ctx, types.DefaultParams())
	
	return ctx, vk, dk, bk
}

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)
	return cdc
}

func RandomNode(r *rand.Rand, ctx sdk.Context, keeper Keeper) types.Node {
	nodes := keeper.GetAllNodes(ctx)
	if len(nodes) == 0 {
		return types.Node{}
	}
	i := r.Intn(len(nodes))
	
	return nodes[i]
}

func RandomNodeByAddress(r *rand.Rand, ctx sdk.Context, keeper Keeper, address sdk.AccAddress) types.Node {
	var randomID int
	count := keeper.GetNodesCountOfAddress(ctx, address)
	
	if count == 0 {
		randomID = 0
	} else {
		randomID = simulation.RandIntBetween(r, 0, int(count))
	}
	
	nodeID, found := keeper.GetNodeIDByAddress(ctx, address, uint64(randomID))
	if !found {
		return types.Node{}
	}
	
	node, found := keeper.GetNode(ctx, nodeID)
	if !found {
		return types.Node{}
	}
	if !node.Owner.Equals(address) {
		return types.Node{}
	}
	return node
}

func RandomFreeClientOfNode(r *rand.Rand, ctx sdk.Context, keeper Keeper, nodeID hub.NodeID) sdk.AccAddress {
	freeClients := keeper.GetFreeClientsOfNode(ctx, nodeID)
	if len(freeClients) == 0 {
		return nil
	}
	i := r.Intn(len(freeClients))
	return freeClients[i]
}

func RandomSubscription(r *rand.Rand, ctx sdk.Context, keeper Keeper) types.Subscription {
	subscriptions := keeper.GetAllSubscriptions(ctx)
	if len(subscriptions) == 0 {
		return types.Subscription{}
	}
	i := r.Intn(len(subscriptions))
	
	return subscriptions[i]
}

func RandomSubscriptionsOfAddress(r *rand.Rand, ctx sdk.Context, keeper Keeper, address sdk.AccAddress) types.Subscription {
	subscriptionsIDS := keeper.GetSubscriptionsOfAddress(ctx, address)
	if len(subscriptionsIDS) == 0 {
		return types.Subscription{}
	}
	
	i := r.Intn(len(subscriptionsIDS))
	subscription, found := keeper.GetSubscription(ctx, subscriptionsIDS[i].ID)
	if !found {
		return types.Subscription{}
	}
	return subscription
}

func RandomSession(r *rand.Rand, ctx sdk.Context, keeper Keeper) types.Session {
	sessions := keeper.GetAllSessions(ctx)
	if len(sessions) == 0 {
		return types.Session{}
	}
	i := r.Intn(len(sessions))
	
	return sessions[i]
}

func RandomResolver(r *rand.Rand, ctx sdk.Context, keeper Keeper) types.Resolver {
	resolvers := keeper.GetAllResolvers(ctx)
	if len(resolvers) == 0 {
		return types.Resolver{}
	}
	i := r.Intn(len(resolvers))
	
	return resolvers[i]
}

func RandomResolverByAddress(r *rand.Rand, ctx sdk.Context, keeper Keeper, address sdk.AccAddress) types.Resolver {
	resolvers := keeper.GetResolversOfAddress(ctx, address)
	if len(resolvers) == 0 {
		return types.Resolver{}
	}
	
	i := r.Intn(len(resolvers))
	return resolvers[i]
}

func RandomPairOfResolverAndNode(r *rand.Rand, ctx sdk.Context, keeper Keeper, node types.Node) hub.ResolverID {
	resolvers := keeper.GetResolversOfNode(ctx, node.ID)
	if len(resolvers) == 0 {
		return nil
	}
	i := r.Intn(len(resolvers))
	
	resolverID, found := keeper.GetResolverOfNode(ctx, node.ID, resolvers[i])
	if !found {
		return nil
	}
	
	return resolverID
}

func GetRandomCoin(r *rand.Rand, coins sdk.Coins) sdk.Coin {
	if len(coins) == 0 {
		return sdk.Coin{}
	}
	coins = simulation.RandSubsetCoins(r, coins)
	i := r.Intn(len(coins))
	
	coin := coins[i]
	if coin.IsZero() || coin.IsNegative() {
		return sdk.Coin{}
	}
	
	coin.Amount = sdk.NewInt(int64(simulation.RandIntBetween(r, 10, 10000)))
	return coin
}
