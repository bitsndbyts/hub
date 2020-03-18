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

func RemoveVPNOnResolverTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deregister-vpn-on-resolver",
		Short: "Deregister vpn node on resolver node",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txb := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			nodeID, err := hub.NewNodeIDFromString(viper.GetString(flagNodeID))
			if err != nil {
				return err
			}

			resolver, err := hub.NewResolverIDFromString(viper.GetString(flagResolverID))
			if err != nil {
				return err
			}

			msg := types.NewMsgDeregisterVPNOnResolver(ctx.FromAddress, nodeID, resolver)

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "VPN node id")
	cmd.Flags().String(flagResolverID, "", "Resolver node address")

	_ = cmd.MarkFlagRequired(flagNodeID)
	_ = cmd.MarkFlagRequired(flagResolverID)

	return cmd
}
