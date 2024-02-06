package routing

import (
	"danbeltoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

func UserPage(ctx *gin.Context) {
	address := common.HexToAddress(ctx.Query("address"))
	role := ctx.Query("role")
	user, _ := variables.Contract.Users(variables.DefaultCallOpts(), address)

	blockNumber, _ := variables.Client.BlockNumber(ctx)
	eth, _ := variables.Client.BalanceAt(ctx, address, big.NewInt(int64(blockNumber)))

	templateDict := map[string]string{
		"BASE_USER": "user.html",
		"OWNER":     "owner.html",
	}

	ctx.HTML(http.StatusOK, templateDict[role], gin.H{
		"UserData": user,
		"Eth":      eth,
	})
}
