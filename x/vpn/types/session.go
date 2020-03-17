package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Session struct {
	ID               hub.SessionID      `json:"id"`
	SubscriptionID   hub.SubscriptionID `json:"subscription_id"`
	Bandwidth        hub.Bandwidth      `json:"bandwidth"`
	Status           string             `json:"status"`
	StatusModifiedAt int64              `json:"status_modified_at"`
}

func (s Session) String() string {
	return fmt.Sprintf(`Session
  ID:                   %s
  Subscription ID:      %s
  Bandwidth:            %s
  Status:               %s
  Status Modified At:   %d`, s.ID, s.SubscriptionID, s.Bandwidth, s.Status, s.StatusModifiedAt)
}

func (s Session) IsValid() error {
	if s.Bandwidth.AnyNil() {
		return fmt.Errorf("invalid bandwidth")
	}
	if s.Status != StatusRegistered && s.Status != StatusDeRegistered {
		return fmt.Errorf("invalid status")
	}

	return nil
}

type FreeSessionBandwidth struct {
	NodeID      hub.NodeID     `json:"node_id"`
	NodeAddress sdk.AccAddress `json:"node_address"`
	ClientID    string         `json:"client_id"`
	Bandwidth   hub.Bandwidth  `json:"bandwidth"`
}

func (s FreeSessionBandwidth) String() string {
	return fmt.Sprintf(`
NodeID: %s
NodeAddress: %s
ClientID:%s
Bandwidth:%s`, s.NodeID.String(), s.NodeAddress.String(), s.ClientID, s.Bandwidth.String())
}
