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

func AddFreeClientTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-free-client",
		Short: "Adding free client",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			nodeID, err := hub.NewNodeIDFromString(viper.GetString(flagNodeID))
			if err != nil {
				return err
			}

			address, err := sdk.AccAddressFromBech32(viper.GetString(flagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgAddFreeClient(ctx.FromAddress, nodeID, address)

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "VPN node id")
	cmd.Flags().String(flagAddress, "", "Client address")

	_ = cmd.MarkFlagRequired(flagNodeID)
	_ = cmd.MarkFlagRequired(flagAddress)

	return cmd
}
