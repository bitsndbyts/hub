package cli

import (
	"bufio"

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

func RemoveFreeClientTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-free-client",
		Short: "Removing free client",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txb := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			nodeID, err := hub.NewNodeIDFromString(viper.GetString(flagNodeID))
			if err != nil {
				return err
			}

			address, err := sdk.AccAddressFromBech32(viper.GetString(flagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveFreeClient(ctx.FromAddress, nodeID, address)

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "VPN node id")
	cmd.Flags().String(flagAddress, "", "Client address")

	_ = cmd.MarkFlagRequired(flagNodeID)
	_ = cmd.MarkFlagRequired(flagAddress)

	return cmd
}
