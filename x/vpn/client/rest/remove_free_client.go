package rest

import (
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

type msgRemoveFreeClient struct {
	BaseReq rest.BaseReq `json:"base_req"`
	NodeID  string       `json:"node_id"`
}

func removeFreeClientHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgRemoveFreeClient
		
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		
		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		nodeID, err := hub.NewNodeIDFromString(req.NodeID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		vars := mux.Vars(r)
		client, err := sdk.AccAddressFromBech32(vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		msg := types.NewMsgRemoveFreeClient(fromAddress, nodeID, client)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}
