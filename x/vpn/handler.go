package vpn

import (
	"bytes"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		
		switch msg := msg.(type) {
		case types.MsgRegisterNode:
			return handleRegisterNode(ctx, k, msg)
		case types.MsgUpdateNodeInfo:
			return handleUpdateNodeInfo(ctx, k, msg)
		case types.MsgAddFreeClient:
			return handleAddFreeClient(ctx, k, msg)
		case types.MsgRemoveFreeClient:
			return handleRemoveFreeClient(ctx, k, msg)
		case types.MsgRegisterVPNOnResolver:
			return handleRegisterVPNOnResolver(ctx, k, msg)
		case types.MsgDeregisterVPNOnResolver:
			return handleDeregisterVPNOnResolver(ctx, k, msg)
		case types.MsgDeregisterNode:
			return handleDeregisterNode(ctx, k, msg)
		case types.MsgStartSubscription:
			return handleStartSubscription(ctx, k, msg)
		case types.MsgEndSubscription:
			return handleEndSubscription(ctx, k, msg)
		case types.MsgUpdateSessionInfo:
			return handleUpdateSessionInfo(ctx, k, msg)
		case types.MsgEndSession:
			return handleEndSession(ctx, k, msg)
		case types.MsgRegisterResolver:
			return handleRegisterResolver(ctx, k, msg)
		case types.MsgUpdateResolverInfo:
			return handleUpdateResolverInfo(ctx, k, msg)
		case types.MsgDeregisterResolver:
			return handleDeregisterResolver(ctx, k, msg)
		
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func EndBlock(ctx sdk.Context, k keeper.Keeper) {
	height := ctx.BlockHeight()
	_height := height - k.SessionInactiveInterval(ctx)
	
	ids := k.GetActiveSessionIDs(ctx, _height)
	for _, id := range ids {
		session, _ := k.GetSession(ctx, id.(hub.SessionID))
		subscription, _ := k.GetSubscription(ctx, session.SubscriptionID)
		
		bandwidth := session.Bandwidth.CeilTo(hub.GB.Quo(subscription.PricePerGB.Amount))
		
		freeClients := k.GetFreeClientsOfNode(ctx, subscription.NodeID)
		
		pay := sdk.NewInt(0)
		if !types.IsFreeClient(freeClients, subscription.Client) {
			amount := bandwidth.Sum().Mul(subscription.PricePerGB.Amount).Quo(hub.GB)
			payCoin := sdk.NewCoin(subscription.PricePerGB.Denom, amount)
			
			pay = payCoin.Amount
			if !pay.IsZero() {
				node, _ := k.GetNode(ctx, subscription.NodeID)
				
				_resolver, found := k.GetResolver(ctx, subscription.ResolverID)
				if !found {
					panic("no resolver found")
				}
				
				commission := _resolver.GetCommission(payCoin)
				
				if commission.IsPositive() {
					if err := k.SendDeposit(ctx, subscription.Client, _resolver.Owner, commission); err != nil {
						panic(err)
					}
					
					if err := k.SendDeposit(ctx, subscription.Client, node.Owner, payCoin.Sub(commission)); err != nil {
						panic(err)
					}
				}
				
				if commission.IsZero() {
					if err := k.SendDeposit(ctx, subscription.Client, node.Owner, payCoin); err != nil {
						panic(err)
					}
				}
			}
		}
		session.Status = types.StatusInactive
		session.StatusModifiedAt = height
		k.SetSession(ctx, session)
		
		subscription.RemainingDeposit.Amount = subscription.RemainingDeposit.Amount.Sub(pay)
		subscription.RemainingBandwidth = subscription.RemainingBandwidth.Sub(bandwidth)
		k.SetSubscription(ctx, subscription)
		
		scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
		k.SetSessionsCountOfSubscription(ctx, subscription.ID, scs+1)
	}
	
	k.DeleteActiveSessionIDs(ctx, _height)
}

func handleRegisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterNode) (*sdk.Result, error) {
	
	nc := k.GetNodesCount(ctx)
	node := types.Node{
		ID:               hub.NewNodeID(nc),
		Owner:            msg.From,
		Deposit:          sdk.NewInt64Coin(k.Deposit(ctx).Denom, 0),
		Type:             msg.T,
		Version:          msg.Version,
		Moniker:          msg.Moniker,
		PricesPerGB:      msg.PricesPerGB,
		InternetSpeed:    msg.InternetSpeed,
		Encryption:       msg.Encryption,
		Status:           types.StatusRegistered,
		StatusModifiedAt: ctx.BlockHeight(),
	}
	
	nca := k.GetNodesCountOfAddress(ctx, node.Owner)
	if nca >= k.FreeNodesCount(ctx) {
		node.Deposit = k.Deposit(ctx)
		
		if err := k.AddDeposit(ctx, node.Owner, node.Deposit); err != nil {
			return nil, err
		}
	}
	
	k.SetNode(ctx, node)
	k.SetNodeIDByAddress(ctx, node.Owner, nca, node.ID)
	
	k.SetNodesCount(ctx, nc+1)
	k.SetNodesCountOfAddress(ctx, node.Owner, nca+1)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgRegisterNode,
			sdk.NewAttribute(AttributeKeyFromAddress, node.Owner.String()),
			sdk.NewAttribute(AttributeKeyNodeID, node.ID.String()),
			sdk.NewAttribute(AttributeKeyStatus, node.Status),
		),
	)
	
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleUpdateNodeInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateNodeInfo) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.ID)
	if !found {
		return nil, types.ErrorNodeDoesNotExist()
	}
	if !msg.From.Equals(node.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	if node.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidNodeStatus()
	}
	
	_node := types.Node{
		Type:          msg.T,
		Version:       msg.Version,
		Moniker:       msg.Moniker,
		PricesPerGB:   msg.PricesPerGB,
		InternetSpeed: msg.InternetSpeed,
		Encryption:    msg.Encryption,
	}
	node = node.UpdateInfo(_node)
	
	k.SetNode(ctx, node)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgUpdateNodeInfo,
			sdk.NewAttribute(AttributeKeyNodeID, msg.ID.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
		),
	)
	
	return &sdk.Result{Events: ctx.EventManager().Events(),
		Data: types.ModuleCdc.MustMarshalJSON(node)}, nil
}

func handleAddFreeClient(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddFreeClient) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return nil, types.ErrorNodeDoesNotExist()
	}
	if !msg.From.Equals(node.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	if node.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidNodeStatus()
	}
	
	k.SetFreeNodeOfClient(ctx, msg.Client, msg.NodeID)
	k.SetFreeClientOfNode(ctx, msg.NodeID, msg.Client)
	
	freeClient := types.FreeClient{
		NodeID: msg.NodeID,
		Client: msg.Client,
	}
	k.SetFreeClient(ctx, freeClient)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgAddFreeClient,
			sdk.NewAttribute(AttributeKeyNodeID, msg.NodeID.String()),
			sdk.NewAttribute(AttributeKeyClientAddress, msg.Client.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.Client.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleRemoveFreeClient(ctx sdk.Context, k keeper.Keeper, msg types.MsgRemoveFreeClient) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return nil, types.ErrorNodeDoesNotExist()
	}
	if !msg.From.Equals(node.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	if node.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidNodeStatus()
	}
	
	_, found = k.GetFreeClientOfNode(ctx, msg.NodeID, msg.Client)
	if !found {
		return nil, types.ErrorFreeClientDoesNotExist()
	}
	
	k.RemoveFreeClientOfNode(ctx, msg.NodeID, msg.Client)
	k.RemoveFreeClient(ctx, msg.NodeID)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgRemoveFreeClient,
			sdk.NewAttribute(AttributeKeyNodeID, msg.NodeID.String()),
			sdk.NewAttribute(AttributeKeyClientAddress, msg.Client.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
		))
	
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleRegisterVPNOnResolver(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterVPNOnResolver) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return nil, types.ErrorNodeDoesNotExist()
	}
	if !msg.From.Equals(node.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	if node.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidNodeStatus()
	}
	
	resolver, found := k.GetResolver(ctx, msg.ResolverID)
	if !found {
		return nil, types.ErrorResolverDoesNotExist()
	}
	if resolver.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidResolverStatus()
	}
	
	k.SetResolverOfNode(ctx, node.ID, resolver.ID)
	k.SetNodeOfResolver(ctx, resolver.ID, node.ID)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgRegisterVPNOnResolver,
			sdk.NewAttribute(AttributeKeyNodeID, msg.NodeID.String()),
			sdk.NewAttribute(AttributeKeyResolverID, msg.ResolverID.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
		),
	)
	
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleDeregisterVPNOnResolver(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeregisterVPNOnResolver) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return nil, types.ErrorNodeDoesNotExist()
	}
	if !msg.From.Equals(node.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	if node.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidNodeStatus()
	}
	
	resolver, found := k.GetResolverOfNode(ctx, msg.NodeID, msg.ResolverID)
	if !found {
		return nil, types.ErrorResolverDoesNotExist()
	}
	
	k.RemoveVPNNodeOnResolver(ctx, msg.NodeID, resolver)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgDeregisterVPNOnResolver,
			sdk.NewAttribute(AttributeKeyNodeID, msg.NodeID.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
			sdk.NewAttribute(AttributeKeyResolverID, msg.ResolverID.String()),
		),
	)
	
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleDeregisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeregisterNode) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.ID)
	if !found {
		return nil, types.ErrorNodeDoesNotExist()
	}
	if !msg.From.Equals(node.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	if node.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidNodeStatus()
	}
	
	if node.Deposit.IsPositive() {
		if err := k.SubtractDeposit(ctx, node.Owner, node.Deposit); err != nil {
			return nil, err
		}
	}
	
	node.Status = types.StatusDeRegistered
	node.StatusModifiedAt = ctx.BlockHeight()
	
	k.SetNode(ctx, node)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgDeregisterNode,
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
			sdk.NewAttribute(AttributeKeyStatus, node.Status),
			sdk.NewAttribute(AttributeKeyNodeID, msg.ID.String()),
		))
	
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
	
}

// nolint:funlen
func handleStartSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartSubscription) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return nil, types.ErrorNodeDoesNotExist()
	}
	if node.Status != types.StatusRegistered {
		return nil, types.ErrorInvalidNodeStatus()
	}
	
	_, found = k.GetResolverOfNode(ctx, msg.NodeID, msg.ResolverID)
	if !found {
		return nil, types.ErrorResolverDoesNotExist()
	}
	
	freeClients := k.GetFreeClientsOfNode(ctx, msg.NodeID)
	if !types.IsFreeClient(freeClients, msg.From) {
		if err := k.AddDeposit(ctx, msg.From, msg.Deposit); err != nil {
			return nil, err
		}
	}
	
	bandwidth, err := node.DepositToBandwidth(msg.Deposit)
	if err != nil {
		return nil, err
	}
	pricePerGB := node.FindPricePerGB(msg.Deposit.Denom)
	
	sc := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:                 hub.NewSubscriptionID(sc),
		ResolverID:         msg.ResolverID,
		NodeID:             node.ID,
		Client:             msg.From,
		PricePerGB:         pricePerGB,
		TotalDeposit:       msg.Deposit,
		RemainingDeposit:   msg.Deposit,
		RemainingBandwidth: bandwidth,
		Status:             types.StatusActive,
		StatusModifiedAt:   ctx.BlockHeight(),
	}
	
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionsCount(ctx, sc+1)
	
	nsc := k.GetSubscriptionsCountOfNode(ctx, node.ID)
	k.SetSubscriptionIDByNodeID(ctx, node.ID, nsc, subscription.ID)
	k.SetSubscriptionsCountOfNode(ctx, node.ID, nsc+1)
	
	sca := k.GetSubscriptionsCountOfAddress(ctx, subscription.Client)
	k.SetSubscriptionIDByAddress(ctx, subscription.Client, sca, subscription.ID)
	k.SetSubscriptionsCountOfAddress(ctx, subscription.Client, sca+1)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgStartSubscription,
			sdk.NewAttribute(AttributeSubscriptionID, subscription.ID.String()),
			sdk.NewAttribute(AttributeKeyNodeID, subscription.NodeID.String()),
			sdk.NewAttribute(AttributeKeyResolverID, subscription.ResolverID.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
			sdk.NewAttribute(AttributeKeyStatus, subscription.Status),
			sdk.NewAttribute(AttributeKeyDeposit, msg.Deposit.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleEndSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgEndSubscription) (*sdk.Result, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist()
	}
	if !msg.From.Equals(subscription.Client) {
		return nil, types.ErrorUnauthorized()
	}
	if subscription.Status != types.StatusActive {
		return nil, types.ErrorInvalidSubscriptionStatus()
	}
	
	scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	
	_, found = k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
	if found {
		return nil, types.ErrorSessionAlreadyExists()
	}
	
	freeClients := k.GetFreeClientsOfNode(ctx, subscription.NodeID)
	
	if !types.IsFreeClient(freeClients, msg.From) && !subscription.RemainingDeposit.IsZero() {
		if err := k.SubtractDeposit(ctx, subscription.Client, subscription.RemainingDeposit); err != nil {
			return nil, err
		}
	}
	
	subscription.Status = types.StatusInactive
	subscription.StatusModifiedAt = ctx.BlockHeight()
	k.SetSubscription(ctx, subscription)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgEndSubscription,
			sdk.NewAttribute(AttributeSubscriptionID, subscription.ID.String()),
			sdk.NewAttribute(AttributeKeyStatus, subscription.Status),
		),
	)
	
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
	
}

func handleUpdateSessionInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateSessionInfo) (*sdk.Result, error) {
	subscription, found := k.GetSubscription(ctx, msg.SubscriptionID)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist()
	}
	if subscription.Status == types.StatusInactive {
		return nil, types.ErrorInvalidSubscriptionStatus()
	}
	if !bytes.Equal(msg.ClientSignature.PubKey.Address(), subscription.Client.Bytes()) {
		return nil, types.ErrorUnauthorized()
	}
	
	node, _ := k.GetNode(ctx, subscription.NodeID)
	if !bytes.Equal(msg.NodeOwnerSignature.PubKey.Address(), node.Owner.Bytes()) {
		return nil, types.ErrorUnauthorized()
	}
	
	scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	data := hub.NewBandwidthSignatureData(subscription.ID, scs, msg.Bandwidth).Bytes()
	if !msg.NodeOwnerSignature.VerifyBytes(data, msg.NodeOwnerSignature.Signature) {
		return nil, types.ErrorInvalidBandwidthSignature()
	}
	if !msg.ClientSignature.VerifyBytes(data, msg.ClientSignature.Signature) {
		return nil, types.ErrorInvalidBandwidthSignature()
	}
	if subscription.RemainingBandwidth.AnyLT(msg.Bandwidth) {
		return nil, types.ErrorInvalidBandwidth()
	}
	var session types.Session
	
	id, found := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
	if !found {
		sc := k.GetSessionsCount(ctx)
		session = types.Session{
			ID:             hub.NewSessionID(sc),
			SubscriptionID: subscription.ID,
			Bandwidth:      hub.NewBandwidthFromInt64(0, 0),
		}
		k.SetSessionsCount(ctx, sc+1)
		k.SetSessionIDBySubscriptionID(ctx, subscription.ID, scs, session.ID)
	} else {
		session, _ = k.GetSession(ctx, id)
	}
	
	k.RemoveSessionIDFromActiveList(ctx, session.StatusModifiedAt, session.ID)
	k.AddSessionIDToActiveList(ctx, ctx.BlockHeight(), session.ID)
	session.Bandwidth = msg.Bandwidth
	session.Status = types.StatusActive
	session.StatusModifiedAt = ctx.BlockHeight()
	
	k.SetSession(ctx, session)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgUpdateSessionInfo,
			sdk.NewAttribute(AttributeSubscriptionID, session.SubscriptionID.String()),
			sdk.NewAttribute(AttributeSessionID, session.ID.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events(), Data: types.ModuleCdc.MustMarshalJSON(session.Bandwidth)}, nil
	
}

func handleEndSession(ctx sdk.Context, k keeper.Keeper, msg types.MsgEndSession) (*sdk.Result, error) {
	subscription, found := k.GetSubscription(ctx, msg.SubscriptionID)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist()
	}
	if subscription.Status == types.StatusInactive {
		return nil, types.ErrorInvalidSubscriptionStatus()
	}
	
	node, _ := k.GetNode(ctx, subscription.NodeID)
	if !msg.From.Equals(node.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	
	scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	
	id, found := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
	if !found {
		return nil, types.ErrorInvalidSessionStatus()
	}
	
	session, _ := k.GetSession(ctx, id)
	
	session.Status = types.StatusInactive
	session.StatusModifiedAt = ctx.BlockHeight()
	
	freeClients := k.GetFreeClientsOfNode(ctx, subscription.NodeID)
	
	bandwidth := session.Bandwidth.CeilTo(hub.GB.Quo(subscription.PricePerGB.Amount))
	
	pay := sdk.NewInt(0)
	if !types.IsFreeClient(freeClients, subscription.Client) {
		amount := bandwidth.Sum().Mul(subscription.PricePerGB.Amount).Quo(hub.GB)
		payCoin := sdk.NewCoin(subscription.PricePerGB.Denom, amount)
		
		pay = payCoin.Amount
		if !pay.IsZero() {
			node, _ := k.GetNode(ctx, subscription.NodeID)
			
			_resolver, found := k.GetResolver(ctx, subscription.ResolverID)
			if !found {
				panic("no resolver found")
			}
			
			commission := _resolver.GetCommission(payCoin)
			
			if commission.IsPositive() {
				if err := k.SendDeposit(ctx, subscription.Client, _resolver.Owner, commission); err != nil {
					panic(err)
				}
				
				if err := k.SendDeposit(ctx, subscription.Client, node.Owner, payCoin.Sub(commission)); err != nil {
					panic(err)
				}
			}
			
			if commission.IsZero() {
				if err := k.SendDeposit(ctx, subscription.Client, node.Owner, payCoin); err != nil {
					panic(err)
				}
			}
		}
	}
	
	subscription.RemainingDeposit.Amount = subscription.RemainingDeposit.Amount.Sub(pay)
	subscription.RemainingBandwidth = subscription.RemainingBandwidth.Sub(bandwidth)
	k.SetSubscription(ctx, subscription)
	
	k.SetSession(ctx, session)
	k.SetSessionsCountOfSubscription(ctx, subscription.ID, scs+1)
	k.RemoveSessionIDFromActiveList(ctx, session.StatusModifiedAt, session.ID)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgUpdateSessionInfo,
			sdk.NewAttribute(AttributeSubscriptionID, session.SubscriptionID.String()),
			sdk.NewAttribute(AttributeSessionID, session.ID.String()),
			sdk.NewAttribute(AttributeKeyFromAddress, msg.From.String()),
		),
	)
	
	return &sdk.Result{Events: ctx.EventManager().Events(), Data: types.ModuleCdc.MustMarshalJSON(session.Status)}, nil
}

func handleRegisterResolver(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterResolver) (*sdk.Result, error) {
	rc := k.GetResolverCount(ctx)
	
	resolver := types.Resolver{
		ID:               hub.NewResolverID(rc),
		Owner:            msg.From,
		Commission:       msg.Commission,
		Status:           types.StatusRegistered,
		StatusModifiedAt: ctx.BlockHeight(),
	}
	
	rca := k.GetResolversCountOfAddress(ctx, resolver.Owner)
	k.SetResolver(ctx, resolver)
	k.SetResolverIDByAddress(ctx, resolver.Owner, rca, resolver.ID)
	
	k.SetResolverCountOfAddress(ctx, resolver.Owner, rca+1)
	k.SetResolverCount(ctx, rc+1)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgRegisterResolver,
			sdk.NewAttribute(AttributeKeyClientAddress, resolver.Owner.String()),
			sdk.NewAttribute(AttributeKeyResolverID, resolver.ID.String()),
			sdk.NewAttribute(AttributeKeyStatus, resolver.Status),
		))
	
	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
	
}

func handleUpdateResolverInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateResolverInfo) (*sdk.Result, error) {
	resolver, found := k.GetResolver(ctx, msg.ResolverID)
	if !found {
		return nil, types.ErrorResolverDoesNotExist()
	}
	if !msg.From.Equals(resolver.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	if resolver.Status == types.StatusDeRegistered {
		return nil, types.ErrorInvalidResolverStatus()
	}
	
	_resolver := types.Resolver{
		Commission: msg.Commission,
	}
	
	resolver = resolver.UpdateInfo(_resolver)
	k.SetResolver(ctx, resolver)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgUpdateResolverInfo,
			sdk.NewAttribute(AttributeKeyClientAddress, msg.From.String()),
			sdk.NewAttribute(AttributeKeyResolverID, msg.ResolverID.String()),
			sdk.NewAttribute(AttributeKeyCommission, msg.Commission.String()),
		),
	)
	
	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleDeregisterResolver(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeregisterResolver) (*sdk.Result, error) {
	resolver, found := k.GetResolver(ctx, msg.ResolverID)
	if !found {
		return nil, types.ErrorResolverDoesNotExist()
	}
	if !msg.From.Equals(resolver.Owner) {
		return nil, types.ErrorUnauthorized()
	}
	
	if resolver.Status != types.StatusRegistered {
		return nil, types.ErrorInvalidResolverStatus()
	}
	
	resolver.Status = types.StatusDeRegistered
	resolver.StatusModifiedAt = ctx.BlockHeight()
	
	k.SetResolver(ctx, resolver)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgDeregisterResolver,
			sdk.NewAttribute(AttributeKeyClientAddress, msg.From.String()),
			sdk.NewAttribute(AttributeKeyResolverID, msg.ResolverID.String()),
			sdk.NewAttribute(AttributeKeyStatus, resolver.Status),
		))
	
	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}
