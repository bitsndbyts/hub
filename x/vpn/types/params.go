package types

import (
	"errors"
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
)

var (
	DefaultFreeNodesCount          uint64 = 5
	DefaultDeposit                        = sdk.NewInt64Coin("stake", 100)
	DefaultSessionInactiveInterval int64  = 25
)

var (
	KeyFreeNodesCount          = []byte("FreeNodesCount")
	KeyDeposit                 = []byte("Deposit")
	KeySessionInactiveInterval = []byte("SessionInactiveInterval")
)

var _ params.ParamSet = (*Params)(nil)

type Params struct {
	FreeNodesCount          uint64   `json:"free_nodes_count"`
	Deposit                 sdk.Coin `json:"deposit"`
	SessionInactiveInterval int64    `json:"session_inactive_interval"`
}

func NewParams(freeNodesCount uint64, deposit sdk.Coin, sessionInactiveInterval int64) Params {
	return Params{
		FreeNodesCount:          freeNodesCount,
		Deposit:                 deposit,
		SessionInactiveInterval: sessionInactiveInterval,
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`Params
  Free Nodes Count:          %d
  Deposit:                   %s
  Session Inactive Interval: %d`, p.FreeNodesCount, p.Deposit, p.SessionInactiveInterval)
}

func (p *Params) ParamSetPairs() subspace.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyFreeNodesCount, Value: &p.FreeNodesCount, ValidatorFn: validateFreeNodeCount},
		{Key: KeyDeposit, Value: &p.Deposit, ValidatorFn: validateDeposit},
		{Key: KeySessionInactiveInterval, Value: &p.SessionInactiveInterval, ValidatorFn: validateSessionInactiveInterval},
	}
}

func DefaultParams() Params {
	return Params{
		FreeNodesCount:          DefaultFreeNodesCount,
		Deposit:                 DefaultDeposit,
		SessionInactiveInterval: DefaultSessionInactiveInterval,
	}
}

func (p Params) Validate() error {
	if !p.Deposit.IsValid() {
		return fmt.Errorf("deposit is invalid: %s ", p.Deposit.String())
	}
	if p.SessionInactiveInterval < 0 {
		return fmt.Errorf("SessionInactiveInterval: %d should be positive interger", p.SessionInactiveInterval)
	}
	
	return nil
}

func validateFreeNodeCount(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}
	
	return nil
}

func validateSessionInactiveInterval(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}
	
	return nil
}

func validateDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if v.IsZero() || v.IsNegative() {
		return errors.New("deposit amount cannot be blank")
	}
	
	if v.Amount.LT(sdk.NewInt(100)) { // TODO : verify amount ??
		return errors.New("deposit amount should be greater than 100")
	}
	
	return nil
}
