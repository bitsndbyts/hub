package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	errCodeUnknownQueryType         = 101
	errCodeInsufficientDepositFunds = 102
)

var (
	ErrInvalidQueryType         = sdkerrors.Register(ModuleName, errCodeUnknownQueryType, "Invalid query type: ")
	ErrInsufficientDepositFunds = sdkerrors.Register(ModuleName, errCodeInsufficientDepositFunds, "insufficient funds")
)

func ErrorInvalidQueryType(queryType string) error {
	return sdkerrors.Wrapf(ErrInvalidQueryType, queryType)
}

func ErrorInsufficientDepositFunds(x, y sdk.Coins) error {
	return sdkerrors.Wrapf(ErrInsufficientDepositFunds, fmt.Sprintf("insufficient deposit funds: %s < %s", x, y))
}
