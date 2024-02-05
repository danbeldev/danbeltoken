package routing

import (
	"danbeltoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func SetRefCode(ctx *gin.Context) {
	code := ctx.Request.FormValue("ref_code")
	address := common.HexToAddress(ctx.Query("address"))
	err := variables.Contract.SetReferral(variables.DefaultCallOpts(), address, code)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
}
