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

func StartSubscriptionTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start subscription",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txb := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			resolver, err := hub.NewResolverIDFromString(viper.GetString(flagResolverID))
			if err != nil {
				return err
			}
			
			nodeID, err := hub.NewNodeIDFromString(viper.GetString(flagNodeID))
			if err != nil {
				return err
			}
			
			deposit := viper.GetString(flagDeposit)
			
			parsedDeposit, err := sdk.ParseCoin(deposit)
			if err != nil {
				return err
			}
			
			fromAddress := ctx.GetFromAddress()
			
			msg := types.NewMsgStartSubscription(fromAddress, resolver, nodeID, parsedDeposit)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
	
	cmd.Flags().String(flagResolverID, "", "Resolver")
	cmd.Flags().String(flagNodeID, "", "Node ID")
	cmd.Flags().String(flagDeposit, "", "Deposit")
	
	_ = cmd.MarkFlagRequired(flagResolverID)
	_ = cmd.MarkFlagRequired(flagNodeID)
	_ = cmd.MarkFlagRequired(flagDeposit)
	
	return cmd
}
