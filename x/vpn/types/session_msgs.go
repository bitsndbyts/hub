package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	hub "github.com/sentinel-official/hub/types"
)

var _ sdk.Msg = (*MsgUpdateSessionInfo)(nil)

type MsgUpdateSessionInfo struct {
	From               sdk.AccAddress     `json:"from"`
	SubscriptionID     hub.SubscriptionID `json:"subscription_id"`
	Bandwidth          hub.Bandwidth      `json:"bandwidth"`
	NodeOwnerSignature auth.StdSignature  `json:"node_owner_signature"`
	ClientSignature    auth.StdSignature  `json:"client_signature"`
}

func (msg MsgUpdateSessionInfo) Type() string {
	return "update_session_info"
}

func (msg MsgUpdateSessionInfo) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if !msg.Bandwidth.AllPositive() {
		return ErrorInvalidField("bandwidth")
	}
	if msg.NodeOwnerSignature.Signature == nil || msg.NodeOwnerSignature.PubKey == nil {
		return ErrorInvalidField("node_owner_signature")
	}
	if msg.ClientSignature.Signature == nil || msg.ClientSignature.PubKey == nil {
		return ErrorInvalidField("client_signature")
	}

	return nil
}

func (msg MsgUpdateSessionInfo) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateSessionInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgUpdateSessionInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionInfo(from sdk.AccAddress,
	subscriptionID hub.SubscriptionID, bandwidth hub.Bandwidth,
	nodeOwnerSignature, clientSignature auth.StdSignature) *MsgUpdateSessionInfo {
	return &MsgUpdateSessionInfo{
		From:               from,
		SubscriptionID:     subscriptionID,
		Bandwidth:          bandwidth,
		NodeOwnerSignature: nodeOwnerSignature,
		ClientSignature:    clientSignature,
	}
}

var _ sdk.Msg = (*MsgEndSession)(nil)

type MsgEndSession struct {
	From           sdk.AccAddress     `json:"from"`
	SubscriptionID hub.SubscriptionID `json:"subscription_id"`
}

func (msg MsgEndSession) Type() string {
	return "end_session"
}

func (msg MsgEndSession) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	return nil
}

func (msg MsgEndSession) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgEndSession) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgEndSession) Route() string {
	return RouterKey
}

func NewMsgEndSession(from sdk.AccAddress, subscriptionID hub.SubscriptionID) *MsgEndSession {
	return &MsgEndSession{
		From:           from,
		SubscriptionID: subscriptionID,
	}
}

var _ sdk.Msg = (*MsgUpdateFreeSessionBandwidth)(nil)

func NewMsgUpdateFreeSessionBandwidth(from sdk.AccAddress, nodeID hub.NodeID, client string, bandwidth hub.Bandwidth) MsgUpdateFreeSessionBandwidth {
	return MsgUpdateFreeSessionBandwidth{
		From:      from,
		NodeID:    nodeID,
		ClientID:  client,
		BandWidth: bandwidth,
	}
}

type MsgUpdateFreeSessionBandwidth struct {
	From      sdk.AccAddress `json:"from"`
	NodeID    hub.NodeID     `json:"node_id"`
	ClientID  string         `json:"client_id"`
	BandWidth hub.Bandwidth  `json:"band_width"`
}

func (msg MsgUpdateFreeSessionBandwidth) Route() string {
	return RouterKey
}

func (msg MsgUpdateFreeSessionBandwidth) Type() string {
	return "update_free_session_bandwidth"
}

func (msg MsgUpdateFreeSessionBandwidth) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	if msg.ClientID == "" {
		return ErrorInvalidField("client_id")
	}

	if msg.NodeID == nil {
		return ErrorInvalidField("node_id")
	}

	if msg.BandWidth.AnyNil() || !msg.BandWidth.AllPositive() {
		return ErrorInvalidField("band_width")
	}

	return nil
}

func (msg MsgUpdateFreeSessionBandwidth) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateFreeSessionBandwidth) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

var _ sdk.Msg = (*MsgEndFreeSessionBandwidth)(nil)

func NewMsgEndFreeSessionBandwidth(from sdk.AccAddress, nodeID hub.NodeID, client string) MsgEndFreeSessionBandwidth {
	return MsgEndFreeSessionBandwidth{
		From:     from,
		NodeID:   nodeID,
		ClientID: client,
	}
}

type MsgEndFreeSessionBandwidth struct {
	From     sdk.AccAddress `json:"from"`
	NodeID   hub.NodeID     `json:"node_id"`
	ClientID string         `json:"client_id"`
}

func (msg MsgEndFreeSessionBandwidth) Route() string {
	return RouterKey
}

func (msg MsgEndFreeSessionBandwidth) Type() string {
	return "end_free_session_bandwidth"
}

func (msg MsgEndFreeSessionBandwidth) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	if msg.ClientID == "" {
		return ErrorInvalidField("client_id")
	}

	if msg.NodeID == nil {
		return ErrorInvalidField("node_id")
	}

	return nil
}

func (msg MsgEndFreeSessionBandwidth) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgEndFreeSessionBandwidth) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}
