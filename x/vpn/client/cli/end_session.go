package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func EndSessionTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end",
		Short: "End session ",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := hub.NewSubscriptionIDFromString(viper.GetString(flagSubscriptionID))
			if err != nil {
				return err
			}

			msg := types.NewMsgEndSession(ctx.FromAddress, id)

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagSubscriptionID, "", "Subscription ID")

	_ = cmd.MarkFlagRequired(flagSubscriptionID)

	return cmd
}

func EndFreeSessionTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end-free-session",
		Short: "End free session ",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			clientID := viper.GetString(flagClientID)

			nodeID, err := hub.NewNodeIDFromString(viper.GetString(flagNodeID))
			if err != nil {
				return err
			}

			msg := types.NewMsgEndFreeSessionBandwidth(ctx.FromAddress, nodeID, clientID)

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagClientID, "", "Client ID")
	cmd.Flags().String(flagNodeID, "", "Node ID")

	_ = cmd.MarkFlagRequired(flagClientID)
	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}
