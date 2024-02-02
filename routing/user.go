package routing

import (
	"danbeltoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	RefCode string
}

func GetUser(ctx *gin.Context) {
	address := common.HexToAddress("0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151")
	refCode := variables.Contract.GetReferralCode(address)
	user := User{RefCode: refCode}
	ctx.JSON(http.StatusOK, user)
}
