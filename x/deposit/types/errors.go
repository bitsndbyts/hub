package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	errCodeUnknownQueryType         = 101
	errCodeInsufficientDepositFunds = 102
)

var (
	ErrorInvalidQueryType         = sdkerrors.Register(ModuleName, errCodeUnknownQueryType, "Invalid query type: ")
	ErrorInsufficientDepositFunds = sdkerrors.Register(ModuleName, errCodeInsufficientDepositFunds, "insufficient funds")
)
