package types

import (
	"encoding/json"
	"testing"

	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func TestMsgRegisterResolver_GetSignBytes(t *testing.T) {
	msg := *NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterResolver_GetSigners(t *testing.T) {
	msg := *NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterResolver_Route(t *testing.T) {
	msg := *NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgRegisterResolver_Type(t *testing.T) {
	msg := *NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	require.Equal(t, "register_resolver", msg.Type())
}

func TestMsgRegisterResolver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterResolver
		want error
	}{
		{
			"from is nil",
			*NewMsgRegisterResolver(nil, sdk.NewDecWithPrec(2, 1)),
			ErrInvalidField,
		}, {
			"from is empty",
			*NewMsgRegisterResolver([]byte(""), sdk.OneDec()),
			ErrInvalidField,
		}, {
			"commission is negative",
			*NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(-1, 0)),
			ErrInvalidField,
		}, {
			"commission is grater than 1",
			*NewMsgRegisterResolver(TestAddress2, sdk.NewDecWithPrec(2, 0)),
			ErrInvalidField,
		}, {
			"commission with zero",
			*NewMsgRegisterResolver(TestAddress2, sdk.NewDecWithPrec(0, 0)),
			nil,
		}, {
			"commission with one",
			*NewMsgRegisterResolver(TestAddress2, sdk.NewDecWithPrec(1, 0)),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := errors.Cause(tc.msg.ValidateBasic()); got != tc.want {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgUpdateResolverInfo_GetSignBytes(t *testing.T) {
	msg := *NewMsgUpdateResolverInfo(TestAddress1, hub.NewResolverID(0), sdk.OneDec())
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateResolverInfo_GetSigners(t *testing.T) {
	msg := *NewMsgUpdateResolverInfo(TestAddress1, hub.NewResolverID(0), sdk.OneDec())
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateResolverInfo_Route(t *testing.T) {
	msg := *NewMsgUpdateResolverInfo(TestAddress1, hub.NewResolverID(0), sdk.OneDec())
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateResolverInfo_Type(t *testing.T) {
	msg := *NewMsgUpdateResolverInfo(TestAddress1, hub.NewResolverID(0), sdk.OneDec())
	require.Equal(t, "update_resolver_info", msg.Type())
}

func TestMsgUpdateResolverInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateResolverInfo
		want error
	}{
		{
			"from is nil",
			*NewMsgUpdateResolverInfo(nil, hub.NewResolverID(0), sdk.NewDecWithPrec(2, 1)),
			ErrInvalidField,
		}, {
			"from is empty",
			*NewMsgUpdateResolverInfo([]byte(""), hub.NewResolverID(0), sdk.OneDec()),
			ErrInvalidField,
		}, {
			"id is nil",
			*NewMsgUpdateResolverInfo(nil, nil, sdk.NewDecWithPrec(2, 1)),
			ErrInvalidField,
		}, {
			"commission is negative",
			*NewMsgUpdateResolverInfo(TestAddress1, hub.NewResolverID(0), sdk.NewDecWithPrec(-1, 0)),
			ErrInvalidField,
		}, {
			"commission is grater than 1",
			*NewMsgUpdateResolverInfo(TestAddress2, hub.NewResolverID(0), sdk.NewDecWithPrec(2, 0)),
			ErrInvalidField,
		}, {
			"commission with zero",
			*NewMsgUpdateResolverInfo(TestAddress2, hub.NewResolverID(0), sdk.NewDecWithPrec(0, 0)),
			nil,
		}, {
			"commission with one",
			*NewMsgUpdateResolverInfo(TestAddress2, hub.NewResolverID(0), sdk.NewDecWithPrec(1, 0)),
			nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := errors.Cause(tc.msg.ValidateBasic()); got != tc.want {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgDeregisterResolver_GetSignBytes(t *testing.T) {
	msg := *NewMsgDeregisterResolver(TestAddress1, hub.NewResolverID(0))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterResolver_GetSigners(t *testing.T) {
	msg := *NewMsgDeregisterResolver(TestAddress1, hub.NewResolverID(0))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterResolver_Route(t *testing.T) {
	msg := *NewMsgDeregisterResolver(TestAddress1, hub.NewResolverID(0))
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgDeregisterResolver_Type(t *testing.T) {
	msg := *NewMsgDeregisterResolver(TestAddress1, hub.NewResolverID(0))
	require.Equal(t, "deregister_resolver", msg.Type())
}

func TestMsgDeregisterResolver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeregisterResolver
		want error
	}{
		{
			"from is nil",
			*NewMsgDeregisterResolver(nil, hub.NewResolverID(0)),
			ErrInvalidField,
		}, {
			"from is empty",
			*NewMsgDeregisterResolver([]byte(""), hub.NewResolverID(0)),
			ErrInvalidField,
		}, {
			"id is empty",
			*NewMsgDeregisterResolver([]byte(""), nil),
			ErrInvalidField,
		}, {
			"valid input",
			*NewMsgDeregisterResolver(types.TestAddress1, hub.NewResolverID(0)),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := errors.Cause(tc.msg.ValidateBasic()); got != tc.want {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}
