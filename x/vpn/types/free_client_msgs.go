package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var _ sdk.Msg = (*MsgAddFreeClient)(nil)

type MsgAddFreeClient struct {
	From   sdk.AccAddress `json:"from"`
	NodeID hub.NodeID     `json:"node_id"`
	Client sdk.AccAddress `json:"client"`
}

func (msg MsgAddFreeClient) Type() string {
	return "add_free_client"
}

func (msg MsgAddFreeClient) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	return nil
}

func (msg MsgAddFreeClient) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgAddFreeClient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgAddFreeClient) Route() string {
	return RouterKey
}

func NewMsgAddFreeClient(from sdk.AccAddress, nodeID hub.NodeID, client sdk.AccAddress) *MsgAddFreeClient {
	return &MsgAddFreeClient{
		From:   from,
		NodeID: nodeID,
		Client: client,
	}
}