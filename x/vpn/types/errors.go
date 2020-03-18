package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

const (
	errCodeUnknownMsgType            = 101
	errCodeUnknownQueryType          = 102
	errCodeInvalidField              = 103
	errCodeUnauthorized              = 104
	errCodeNodeDoesNotExist          = 105
	errCodeInvalidNodeStatus         = 106
	errCodeInvalidDeposit            = 107
	errCodeSubscriptionDoesNotExist  = 108
	errCodeSubscriptionAlreadyExists = 109
	errCodeInvalidSubscriptionStatus = 110
	errCodeInvalidBandwidth          = 111
	errCodeInvalidBandwidthSignature = 112
	errCodeSessionAlreadyExists      = 113
	errCodeInvalidSessionStatus      = 114
	errCodeResolverAlreadyExist      = 121
	errCodeResolverDoesNotExist      = 122
	errCodeInvalidResolverStatus     = 123
	errCodeFreeClientDoesNotExist    = 115

	errMsgUnknownMsgType            = "Unknown message type: "
	errMsgUnknownQueryType          = "Invalid query type: "
	errMsgInvalidField              = "Invalid field: "
	errMsgUnauthorized              = "Unauthorized"
	errMsgNodeDoesNotExist          = "Node does not exist"
	errMsgInvalidNodeStatus         = "Invalid node status"
	errMsgInvalidDeposit            = "Invalid deposit"
	errMsgSubscriptionDoesNotExist  = "Subscription does not exist"
	errMsgSubscriptionAlreadyExists = "Subscription already exists"
	errMsgInvalidSubscriptionStatus = "Invalid subscription status"
	errMsgInvalidBandwidth          = "Invalid bandwidth"
	errMsgInvalidBandwidthSignature = "Invalid bandwidth signature"
	errMsgSessionAlreadyExists      = "Session is active"
	errMsgInvalidSessionStatus      = "Invalid session status"
	errMsgResolverAlreadyExist      = "Resolver already exist"
	errMsgResolverDoesNotExist      = "Resolver does not exist"
	errMsgInvalidResolverStatus     = "Invalid resolver status"
	errMsgFreeClientDoesNotExist    = "Free client does not exist"
)

func ErrorMarshal() error {
	return sdkerrors.Register(ModuleName, hub.ErrCodeMarshal, hub.ErrMsgMarshal)
}

func ErrorUnmarshal() error {
	return sdkerrors.Register(ModuleName, hub.ErrCodeUnmarshal, hub.ErrMsgUnmarshal)
}

func ErrorUnknownMsgType(msgType string) error {
	return sdkerrors.Register(ModuleName, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func ErrorInvalidQueryType(queryType string) error {
	return sdkerrors.Register(ModuleName, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInvalidField(field string) error {
	return sdkerrors.Register(ModuleName, errCodeInvalidField, errMsgInvalidField+field)
}

func ErrorUnauthorized() error {
	return sdkerrors.Register(ModuleName, errCodeUnauthorized, errMsgUnauthorized)
}

func ErrorNodeDoesNotExist() error {
	return sdkerrors.Register(ModuleName, errCodeNodeDoesNotExist, errMsgNodeDoesNotExist)
}

func ErrorInvalidNodeStatus() error {
	return sdkerrors.Register(ModuleName, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}

func ErrorInvalidDeposit() error {
	return sdkerrors.Register(ModuleName, errCodeInvalidDeposit, errMsgInvalidDeposit)
}

func ErrorSubscriptionDoesNotExist() error {
	return sdkerrors.Register(ModuleName, errCodeSubscriptionDoesNotExist, errMsgSubscriptionDoesNotExist)
}

func ErrorSubscriptionAlreadyExists() error {
	return sdkerrors.Register(ModuleName, errCodeSubscriptionAlreadyExists, errMsgSubscriptionAlreadyExists)
}

func ErrorInvalidSubscriptionStatus() error {
	return sdkerrors.Register(ModuleName, errCodeInvalidSubscriptionStatus, errMsgInvalidSubscriptionStatus)
}

func ErrorInvalidBandwidth() error {
	return sdkerrors.Register(ModuleName, errCodeInvalidBandwidth, errMsgInvalidBandwidth)
}

func ErrorInvalidBandwidthSignature() error {
	return sdkerrors.Register(ModuleName, errCodeInvalidBandwidthSignature, errMsgInvalidBandwidthSignature)
}

func ErrorSessionAlreadyExists() error {
	return sdkerrors.Register(ModuleName, errCodeSessionAlreadyExists, errMsgSessionAlreadyExists)
}

func ErrorInvalidSessionStatus() error {
	return sdkerrors.Register(ModuleName, errCodeInvalidSessionStatus, errMsgInvalidSessionStatus)
}

func ErrorResolverAlreadyExist() error {
	return sdkerrors.Register(ModuleName, errCodeResolverAlreadyExist, errMsgResolverAlreadyExist)
}

func ErrorResolverDoesNotExist() error {
	return sdkerrors.Register(ModuleName, errCodeResolverDoesNotExist, errMsgResolverDoesNotExist)
}

func ErrorInvalidResolverStatus() error {
	return sdkerrors.Register(ModuleName, errCodeInvalidResolverStatus, errMsgInvalidResolverStatus)
}

func ErrorFreeClientDoesNotExist() error {
	return sdkerrors.Register(ModuleName, errCodeFreeClientDoesNotExist, errMsgFreeClientDoesNotExist)
}
