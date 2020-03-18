package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

const (
	errCodeUnknownQueryType         = 101
	errCodeInsufficientDepositFunds = 102

	errMsgUnknownQueryType         = "Invalid query type: "
	errMsgInsufficientDepositFunds = "insufficient deposit funds: %s < %s"
)

func ErrorMarshal() error {
	return sdkerrors.Register(ModuleName, hub.ErrCodeMarshal, hub.ErrMsgMarshal)
}

func ErrorUnmarshal() error {
	return sdkerrors.Register(ModuleName, hub.ErrCodeUnmarshal, hub.ErrMsgUnmarshal)
}

func ErrorInvalidQueryType(queryType string) error {
	return sdkerrors.Register(ModuleName, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInsufficientDepositFunds(x, y sdk.Coins) error {
	return sdkerrors.Register(ModuleName, errCodeInsufficientDepositFunds, fmt.Sprintf(errMsgInsufficientDepositFunds, x, y))
}
