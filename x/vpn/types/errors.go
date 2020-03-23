package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func ErrorUnknownMsgType(msgType string) error {
	return sdkerrors.Wrapf(ErrUnknownMsgType, msgType)
}

func ErrorInvalidQueryType(queryType string) error {
	return sdkerrors.Wrapf(ErrInvalidQueryType, queryType)
}

func ErrorInvalidField(field string) error {
	return sdkerrors.Wrapf(ErrInvalidField, field)
}

func ErrorUnauthorized() error {
	return ErrUnauthorized
}

func ErrorNodeDoesNotExist() error {
	return ErrNodeDoesNotExist
}

func ErrorInvalidNodeStatus() error {
	return ErrInvalidNodeStatus
}

func ErrorInvalidDeposit() error {
	return ErrInvalidDeposit
}

func ErrorSubscriptionDoesNotExist() error {
	return ErrSubscriptionDoesNotExist
}

func ErrorSubscriptionAlreadyExists() error {
	return ErrSubscriptionAlreadyExists
}

func ErrorInvalidSubscriptionStatus() error {
	return ErrInvalidSubscriptionStatus
}

func ErrorInvalidBandwidth() error {
	return ErrInvalidBandwidth
}

func ErrorInvalidBandwidthSignature() error {
	return ErrInvalidBandwidthSignature
}

func ErrorSessionAlreadyExists() error {
	return ErrSessionAlreadyExists
}

func ErrorInvalidSessionStatus() error {
	return ErrInvalidSessionStatus
}

func ErrorResolverAlreadyExist() error {
	return ErrResolverAlreadyExist
}

func ErrorResolverDoesNotExist() error {
	return ErrResolverDoesNotExist
}

func ErrorInvalidResolverStatus() error {
	return ErrInvalidResolverStatus
}

func ErrorFreeClientDoesNotExist() error {
	return ErrFreeClientDoesNotExist
}

var (
	ErrUnknownMsgType            = sdkerrors.Register(ModuleName, errCodeUnknownMsgType, errMsgUnknownMsgType)
	ErrInvalidQueryType          = sdkerrors.Register(ModuleName, errCodeUnknownQueryType, errMsgUnknownQueryType)
	ErrInvalidField              = sdkerrors.Register(ModuleName, errCodeInvalidField, errMsgInvalidField)
	ErrUnauthorized              = sdkerrors.Register(ModuleName, errCodeUnauthorized, errMsgUnauthorized)
	ErrNodeDoesNotExist          = sdkerrors.Register(ModuleName, errCodeNodeDoesNotExist, errMsgNodeDoesNotExist)
	ErrInvalidNodeStatus         = sdkerrors.Register(ModuleName, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
	ErrInvalidDeposit            = sdkerrors.Register(ModuleName, errCodeInvalidDeposit, errMsgInvalidDeposit)
	ErrSubscriptionDoesNotExist  = sdkerrors.Register(ModuleName, errCodeSubscriptionDoesNotExist, errMsgSubscriptionDoesNotExist)
	ErrSubscriptionAlreadyExists = sdkerrors.Register(ModuleName, errCodeSubscriptionAlreadyExists, errMsgSubscriptionAlreadyExists)
	ErrInvalidSubscriptionStatus = sdkerrors.Register(ModuleName, errCodeInvalidSubscriptionStatus, errMsgInvalidSubscriptionStatus)
	ErrInvalidBandwidth          = sdkerrors.Register(ModuleName, errCodeInvalidBandwidth, errMsgInvalidBandwidth)
	ErrInvalidBandwidthSignature = sdkerrors.Register(ModuleName, errCodeInvalidBandwidthSignature, errMsgInvalidBandwidthSignature)
	ErrSessionAlreadyExists      = sdkerrors.Register(ModuleName, errCodeSessionAlreadyExists, errMsgSessionAlreadyExists)
	ErrInvalidSessionStatus      = sdkerrors.Register(ModuleName, errCodeInvalidSessionStatus, errMsgInvalidSessionStatus)
	ErrResolverAlreadyExist      = sdkerrors.Register(ModuleName, errCodeResolverAlreadyExist, errMsgResolverAlreadyExist)
	ErrResolverDoesNotExist      = sdkerrors.Register(ModuleName, errCodeResolverDoesNotExist, errMsgResolverDoesNotExist)
	ErrInvalidResolverStatus     = sdkerrors.Register(ModuleName, errCodeInvalidResolverStatus, errMsgInvalidResolverStatus)
	ErrFreeClientDoesNotExist    = sdkerrors.Register(ModuleName, errCodeFreeClientDoesNotExist, errMsgFreeClientDoesNotExist)
)
