package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var _ sdk.Msg = (*MsgRegisterNode)(nil)

type MsgRegisterNode struct {
	From          sdk.AccAddress `json:"from"`
	Type_         string         `json:"type"` // nolint:golint
	Version       string         `json:"version"`
	Moniker       string         `json:"moniker"`
	PricesPerGB   sdk.Coins      `json:"prices_per_gb"`
	InternetSpeed hub.Bandwidth  `json:"internet_speed"`
	Encryption    string         `json:"encryption"`
}

func (msg MsgRegisterNode) Type() string {
	return "MsgRegisterNode"
}

// nolint: gocyclo
func (msg MsgRegisterNode) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.Type_ == "" {
		return ErrorInvalidField("type")
	}
	if msg.Version == "" {
		return ErrorInvalidField("version")
	}
	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB == nil ||
		msg.PricesPerGB.Len() == 0 || !msg.PricesPerGB.IsValid() {

		return ErrorInvalidField("prices_per_gb")
	}
	if !msg.InternetSpeed.AllPositive() {
		return ErrorInvalidField("internet_speed")
	}
	if msg.Encryption == "" {
		return ErrorInvalidField("encryption")
	}

	return nil
}

func (msg MsgRegisterNode) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgRegisterNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgRegisterNode) Route() string {
	return RouterKey
}

func NewMsgRegisterNode(from sdk.AccAddress,
	_type, version, moniker string, pricesPerGB sdk.Coins,
	internetSpeed hub.Bandwidth, encryption string) *MsgRegisterNode {

	return &MsgRegisterNode{
		From:          from,
		Type_:         _type,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ sdk.Msg = (*MsgUpdateNodeInfo)(nil)

type MsgUpdateNodeInfo struct {
	From          sdk.AccAddress `json:"from"`
	ID            hub.ID         `json:"id"`
	Type_         string         `json:"type"` // nolint:golint
	Version       string         `json:"version"`
	Moniker       string         `json:"moniker"`
	PricesPerGB   sdk.Coins      `json:"prices_per_gb"`
	InternetSpeed hub.Bandwidth  `json:"internet_speed"`
	Encryption    string         `json:"encryption"`
}

func (msg MsgUpdateNodeInfo) Type() string {
	return "MsgUpdateNodeInfo"
}

func (msg MsgUpdateNodeInfo) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB != nil &&
		(msg.PricesPerGB.Len() == 0 || !msg.PricesPerGB.IsValid()) {

		return ErrorInvalidField("prices_per_gb")
	}
	if msg.InternetSpeed.AnyNegative() {
		return ErrorInvalidField("internet_speed")
	}

	return nil
}

func (msg MsgUpdateNodeInfo) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateNodeInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgUpdateNodeInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeInfo(from sdk.AccAddress, id hub.ID,
	_type, version, moniker string, pricesPerGB sdk.Coins,
	internetSpeed hub.Bandwidth, encryption string) *MsgUpdateNodeInfo {

	return &MsgUpdateNodeInfo{
		From:          from,
		ID:            id,
		Type_:         _type,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ sdk.Msg = (*MsgUpdateNodeStatus)(nil)

type MsgUpdateNodeStatus struct {
	From   sdk.AccAddress `json:"from"`
	ID     hub.ID         `json:"id"`
	Status string         `json:"status"`
}

func (msg MsgUpdateNodeStatus) Type() string {
	return "MsgUpdateNodeStatus"
}

func (msg MsgUpdateNodeStatus) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.Status != StatusActive && msg.Status != StatusInactive {
		return ErrorInvalidField("status")
	}

	return nil
}

func (msg MsgUpdateNodeStatus) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateNodeStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgUpdateNodeStatus) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeStatus(from sdk.AccAddress, id hub.ID, status string) *MsgUpdateNodeStatus {

	return &MsgUpdateNodeStatus{
		From:   from,
		ID:     id,
		Status: status,
	}
}

var _ sdk.Msg = (*MsgDeregisterNode)(nil)

type MsgDeregisterNode struct {
	From sdk.AccAddress `json:"from"`
	ID   hub.ID         `json:"id"`
}

func (msg MsgDeregisterNode) Type() string {
	return "MsgDeregisterNode"
}

func (msg MsgDeregisterNode) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	return nil
}

func (msg MsgDeregisterNode) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgDeregisterNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgDeregisterNode) Route() string {
	return RouterKey
}

func NewMsgDeregisterNode(from sdk.AccAddress, id hub.ID) *MsgDeregisterNode {
	return &MsgDeregisterNode{
		From: from,
		ID:   id,
	}
}
