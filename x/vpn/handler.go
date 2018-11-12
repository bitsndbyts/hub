package vpn

import (
	"encoding/json"
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func NewHandler(k Keeper, ik ibc.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgRegisterVpn:
			return handleRegisterVpn(ctx, k, ik, msg)
		case MsgNodeStatus:
			return handleSetNodeStatus(ctx, k, msg)
		default:
			errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(msg).Name()

			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleRegisterVpn(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgRegisterVpn) csdkTypes.Result {

	vpnId := msg.From.String()
	cdc := codec.New()
	vpnData, err := k.GetVpnDetails(ctx, vpnId)

	if err != nil {
		panic(err)
	}

	if vpnData != nil {
		panic("Already registered")
	}

	err = k.SetVpnDetails(ctx, vpnId, msg.Details)

	if err != nil {
		panic(err)
	}

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  "sentinel-vpn",
		DestChainId: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerId: vpnId,
			Address:  msg.From,
			Coins:    msg.Coins,
		},
	}

	err = ik.PostIBCPacket(ctx, ibcPacket)

	if err != nil {
		panic(err)
	}

	tags := csdkTypes.NewTags("Registered Vpn address:", []byte(msg.From.String()))
	data, _ := cdc.MarshalJSON(msg)

	return csdkTypes.Result{
		Tags: tags,
		Data: data,
	}
}

func handleSetNodeStatus(ctx csdkTypes.Context, k Keeper, msg MsgNodeStatus) csdkTypes.Result {
	var Data sdkTypes.VpnDetails

	vpnId := msg.VpnId
	status := msg.Status
	vpnData, err := k.GetVpnDetails(ctx, vpnId)

	if err != nil {
		panic(err)
	}

	if vpnData != nil {
		panic("Already registered")
	}

	err = json.Unmarshal(vpnData, &Data)

	if err != nil {
		panic(err)
	}

	err = k.SetVpnStatus(ctx, vpnId, Data, status)

	if err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}
