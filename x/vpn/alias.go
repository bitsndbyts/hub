// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/vpn/types/
// ALIASGEN: github.com/sentinel-official/hub/x/vpn/keeper/
// ALIASGEN: github.com/sentinel-official/hub/x/vpn/querier/
package vpn

import (
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/querier"
	"github.com/sentinel-official/hub/x/vpn/types"
)

const (
	ModuleName                       = types.ModuleName
	QuerierRoute                     = types.QuerierRoute
	RouterKey                        = types.RouterKey
	StoreKeySession                  = types.StoreKeySession
	StoreKeyResolver                 = types.StoreKeyResolver
	StoreKeyNode                     = types.StoreKeyNode
	StoreKeySubscription             = types.StoreKeySubscription
	StatusRegistered                 = types.StatusRegistered
	StatusActive                     = types.StatusActive
	StatusInactive                   = types.StatusInactive
	StatusDeRegistered               = types.StatusDeRegistered
	QueryParams                      = types.QueryParams
	QueryNode                        = types.QueryNode
	QueryNodesOfAddress              = types.QueryNodesOfAddress
	QueryAllNodes                    = types.QueryAllNodes
	QueryFreeNodesOfClient           = types.QueryFreeNodesOfClient
	QueryFreeClientsOfNode           = types.QueryFreeClientsOfNode
	QueryResolversOfNode             = types.QueryResolversOfNode
	QueryNodesOfResolver             = types.QueryNodesOfResolver
	QueryResolvers                   = types.QueryResolvers
	QuerySubscription                = types.QuerySubscription
	QuerySubscriptionsOfNode         = types.QuerySubscriptionsOfNode
	QuerySubscriptionsOfAddress      = types.QuerySubscriptionsOfAddress
	QueryAllSubscriptions            = types.QueryAllSubscriptions
	QuerySessionsCountOfSubscription = types.QuerySessionsCountOfSubscription
	QuerySession                     = types.QuerySession
	QuerySessionOfSubscription       = types.QuerySessionOfSubscription
	QuerySessionsOfSubscription      = types.QuerySessionsOfSubscription
	QueryAllSessions                 = types.QueryAllSessions
	DefaultParamspace                = keeper.DefaultParamspace
)

var (
	// functions aliases
	RegisterCodec                             = types.RegisterCodec
	ErrorMarshal                              = types.ErrorMarshal
	ErrorUnmarshal                            = types.ErrorUnmarshal
	ErrorUnknownMsgType                       = types.ErrorUnknownMsgType
	ErrorInvalidQueryType                     = types.ErrorInvalidQueryType
	ErrorInvalidField                         = types.ErrorInvalidField
	ErrorUnauthorized                         = types.ErrorUnauthorized
	ErrorNodeDoesNotExist                     = types.ErrorNodeDoesNotExist
	ErrorInvalidNodeStatus                    = types.ErrorInvalidNodeStatus
	ErrorInvalidDeposit                       = types.ErrorInvalidDeposit
	ErrorSubscriptionDoesNotExist             = types.ErrorSubscriptionDoesNotExist
	ErrorSubscriptionAlreadyExists            = types.ErrorSubscriptionAlreadyExists
	ErrorInvalidSubscriptionStatus            = types.ErrorInvalidSubscriptionStatus
	ErrorInvalidBandwidth                     = types.ErrorInvalidBandwidth
	ErrorInvalidBandwidthSignature            = types.ErrorInvalidBandwidthSignature
	ErrorSessionAlreadyExists                 = types.ErrorSessionAlreadyExists
	ErrorInvalidSessionStatus                 = types.ErrorInvalidSessionStatus
	NewGenesisState                           = types.NewGenesisState
	DefaultGenesisState                       = types.DefaultGenesisState
	NodeKey                                   = types.NodeKey
	NodesCountOfAddressKey                    = types.NodesCountOfAddressKey
	NodeIDByAddressKey                        = types.NodeIDByAddressKey
	SubscriptionKey                           = types.SubscriptionKey
	SubscriptionsCountOfNodeKey               = types.SubscriptionsCountOfNodeKey
	SubscriptionIDByNodeIDKey                 = types.SubscriptionIDByNodeIDKey
	SubscriptionsCountOfAddressKey            = types.SubscriptionsCountOfAddressKey
	SubscriptionIDByAddressKey                = types.SubscriptionIDByAddressKey
	SessionKey                                = types.SessionKey
	SessionsCountOfSubscriptionKey            = types.SessionsCountOfSubscriptionKey
	SessionIDBySubscriptionIDKey              = types.SessionIDBySubscriptionIDKey
	ActiveNodeIDsKey                          = types.ActiveNodeIDsKey
	ActiveSessionIDsKey                       = types.ActiveSessionIDsKey
	NewMsgRegisterNode                        = types.NewMsgRegisterNode
	NewMsgAddFreeClient                       = types.NewMsgAddFreeClient
	NewMsgRemoveFreeClient                    = types.NewMsgRemoveFreeClient
	NewMsgRegisterVPNOnResolver               = types.NewMsgRegisterVPNOnResolver
	NewMsgDeregisterVPNOnResolver             = types.NewMsgDeregisterVPNOnResolver
	NewMsgUpdateNodeInfo                      = types.NewMsgUpdateNodeInfo
	NewMsgDeregisterNode                      = types.NewMsgDeregisterNode
	NewMsgRegisterResolver                    = types.NewMsgRegisterResolver
	NewMsgUpdateResolverInfo                  = types.NewMsgUpdateResolverInfo
	NewMsgUpdateSessionInfo                   = types.NewMsgUpdateSessionInfo
	NewMsgStartSubscription                   = types.NewMsgStartSubscription
	NewMsgEndSubscription                     = types.NewMsgEndSubscription
	NewMsgEndSession                          = types.NewMsgEndSession
	NewMsgDeregisterResolver                  = types.NewMsgDeregisterResolver
	NewParams                                 = types.NewParams
	DefaultParams                             = types.DefaultParams
	NewQueryNodeParams                        = types.NewQueryNodeParams
	NewQueryNodesOfAddressParams              = types.NewQueryNodesOfAddressParams
	NewQueryFreeClientsOfNodeParams           = types.NewQueryFreeClientsOfNodeParams
	NewQueryNodesOfFreeClientPrams            = types.NewQueryNodesOfFreeClientPrams
	NewQueryResolversOfNodeParams             = types.NewQueryResolversOfNodeParams
	NewQueryNodesOfResolverPrams              = types.NewQueryNodesOfResolverPrams
	NewQuerySubscriptionParams                = types.NewQuerySubscriptionParams
	NewQuerySubscriptionsOfNodePrams          = types.NewQuerySubscriptionsOfNodePrams
	NewQuerySubscriptionsOfAddressParams      = types.NewQuerySubscriptionsOfAddressParams
	NewQuerySessionsCountOfSubscriptionParams = types.NewQuerySessionsCountOfSubscriptionParams
	NewQuerySessionParams                     = types.NewQuerySessionParams
	NewQuerySessionOfSubscriptionPrams        = types.NewQuerySessionOfSubscriptionPrams
	NewQuerySessionsOfSubscriptionPrams       = types.NewQuerySessionsOfSubscriptionPrams
	NewKeeper                                 = keeper.NewKeeper
	ParamKeyTable                             = keeper.ParamKeyTable
	NewQuerier                                = querier.NewQuerier
	RandomNode                                = keeper.RandomNode
	RandomSubscription                        = keeper.RandomSubscription
	RandomSession                             = keeper.RandomSession
	RandomResolver                            = keeper.RandomResolver

	// variable aliases
	ModuleCdc                            = types.ModuleCdc
	NodesCountKey                        = types.NodesCountKey
	NodeKeyPrefix                        = types.NodeKeyPrefix
	NodesCountOfAddressKeyPrefix         = types.NodesCountOfAddressKeyPrefix
	NodeIDByAddressKeyPrefix             = types.NodeIDByAddressKeyPrefix
	SubscriptionsCountKey                = types.SubscriptionsCountKey
	SubscriptionKeyPrefix                = types.SubscriptionKeyPrefix
	SubscriptionsCountOfNodeKeyPrefix    = types.SubscriptionsCountOfNodeKeyPrefix
	SubscriptionIDByNodeIDKeyPrefix      = types.SubscriptionIDByNodeIDKeyPrefix
	SubscriptionsCountOfAddressKeyPrefix = types.SubscriptionsCountOfAddressKeyPrefix
	SubscriptionIDByAddressKeyPrefix     = types.SubscriptionIDByAddressKeyPrefix
	SessionsCountKey                     = types.SessionsCountKey
	SessionKeyPrefix                     = types.SessionKeyPrefix
	SessionsCountOfSubscriptionKeyPrefix = types.SessionsCountOfSubscriptionKeyPrefix
	SessionIDBySubscriptionIDKeyPrefix   = types.SessionIDBySubscriptionIDKeyPrefix
	DefaultFreeNodesCount                = types.DefaultFreeNodesCount
	DefaultDeposit                       = types.DefaultDeposit
	DefaultSessionInactiveInterval       = types.DefaultSessionInactiveInterval
	KeyFreeNodesCount                    = types.KeyFreeNodesCount
	KeyDeposit                           = types.KeyDeposit
	KeySessionInactiveInterval           = types.KeySessionInactiveInterval

	EventTypeMsgRegisterNode            = types.EventTypeMsgRegisterNode
	EventTypeMsgUpdateNodeInfo          = types.EventTypeMsgUpdateNodeInfo
	EventTypeMsgDeregisterNode          = types.EventTypeMsgDeregisterNode
	EventTypeMsgRegisterResolver        = types.EventTypeMsgRegisterResolver
	EventTypeMsgUpdateResolverInfo      = types.EventTypeMsgUpdateResolverInfo
	EventTypeMsgDeregisterResolver      = types.EventTypeMsgDeregisterResolver
	EventTypeMsgAddFreeClient           = types.EventTypeMsgAddFreeClient
	EventTypeMsgRemoveFreeClient        = types.EventTypeMsgRemoveFreeClient
	EventTypeMsgRegisterVPNOnResolver   = types.EventTypeMsgRegisterVPNOnResolver
	EventTypeMsgDeregisterVPNOnResolver = types.EventTypeMsgDeregisterVPNOnResolver
	EventTypeMsgStartSubscription       = types.EventTypeMsgStartSubscription
	EventTypeMsgEndSubscription         = types.EventTypeMsgEndSubscription
	EventTypeMsgUpdateSessionInfo       = types.EventTypeMsgUpdateSessionInfo

	AttributeKeyClientAddress = types.AttributeKeyClientAddress
	AttributeKeyFromAddress   = types.AttributeKeyFromAddress
	AttributeKeyNodeID        = types.AttributeKeyNodeID
	AttributeKeyResolverID    = types.AttributeKeyResolverID
	AttributeSubscriptionID   = types.AttributeSubscriptionID
	AttributeSessionID        = types.AttributeSessionID
	AttributeKeyStatus        = types.AttributeKeyStatus
	AttributeKeyCommission    = types.AttributeKeyCommission
	AttributeKeyDeposit       = types.AttributeKeyDeposit
)

type (
	GenesisState                           = types.GenesisState
	Node                                   = types.Node
	MsgRegisterNode                        = types.MsgRegisterNode
	MsgUpdateNodeInfo                      = types.MsgUpdateNodeInfo
	MsgDeregisterNode                      = types.MsgDeregisterNode
	Params                                 = types.Params
	QueryNodeParams                        = types.QueryNodeParams
	QueryNodesOfAddressPrams               = types.QueryNodesOfAddressPrams
	QuerySubscriptionParams                = types.QuerySubscriptionParams
	QuerySubscriptionsOfNodePrams          = types.QuerySubscriptionsOfNodePrams
	QuerySubscriptionsOfAddressParams      = types.QuerySubscriptionsOfAddressParams
	QuerySessionsCountOfSubscriptionParams = types.QuerySessionsCountOfSubscriptionParams
	QuerySessionParams                     = types.QuerySessionParams
	QuerySessionOfSubscriptionPrams        = types.QuerySessionOfSubscriptionPrams
	QuerySessionsOfSubscriptionPrams       = types.QuerySessionsOfSubscriptionPrams
	Session                                = types.Session
	MsgUpdateSessionInfo                   = types.MsgUpdateSessionInfo
	Subscription                           = types.Subscription
	MsgStartSubscription                   = types.MsgStartSubscription
	MsgEndSubscription                     = types.MsgEndSubscription
	MsgEndSession                          = types.MsgEndSession
	Keeper                                 = keeper.Keeper
)
