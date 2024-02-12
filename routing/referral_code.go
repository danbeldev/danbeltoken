package routing

import (
	"danbeltoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
)

func SetRefCode(ctx *gin.Context) {
	code := ctx.Request.FormValue("ref_code")
	address := common.HexToAddress(ctx.Query("address"))
	t, _ := variables.TransactOpts(address, big.NewInt(0))
	_, err := variables.Contract.SetReferral(t, address, code)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}
