package types

import (
	"encoding/json"
	"testing"
	
	"github.com/pkg/errors"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
	
	hub "github.com/sentinel-official/hub/types"
)

func TestMsgUpdateSessionInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateSessionInfo
		want error
	}{
		{
			"from is nil",
			NewMsgUpdateSessionInfo(nil, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrInvalidField,
		}, {
			"from is empty",
			NewMsgUpdateSessionInfo([]byte(""), hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrInvalidField,
		}, {
			"bandwidth is zero",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthZero, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrInvalidField,
		}, {
			"bandwidth is neg",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthNeg, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrInvalidField,
		}, {
			"bandwidth is zero",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthZero, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrInvalidField,
		}, {
			"node owner sign is empty  ",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, auth.StdSignature{}, TestClientStdSignaturePos1),
			ErrInvalidField,
		}, {
			"client sign is empty  ",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, auth.StdSignature{}),
			ErrInvalidField,
		}, {
			"valid ",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
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

func TestMsgUpdateSessionInfo_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateSessionInfo_GetSigners(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateSessionInfo_Type(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	require.Equal(t, "update_session_info", msg.Type())
}

func TestMsgUpdateSessionInfo_Route(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	require.Equal(t, RouterKey, msg.Route())
}
