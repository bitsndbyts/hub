package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*ID)(nil), nil)
	cdc.RegisterConcrete(NodeID{}, "types/nodeID", nil)
	cdc.RegisterConcrete(SessionID{}, "types/sessionID", nil)
	cdc.RegisterConcrete(SubscriptionID{}, "types/subscriptionID", nil)
}
