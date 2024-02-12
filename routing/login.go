package routing

import (
	"danbeltoken/variables"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RedirectToError(ctx *gin.Context, err error) {
	ctx.Redirect(http.StatusMovedPermanently, "/error?err="+err.Error())
}

func RedirectToRolePage(ctx *gin.Context, addr string, role string) {
	ctx.Redirect(http.StatusMovedPermanently, "/user-page"+"?address="+addr+"&role="+role)
}

func RedirectFromRequestToRolePage(ctx *gin.Context) {
	role := ctx.Query("role")
	address := ctx.Query("address")
	time.Sleep(time.Second)
	RedirectToRolePage(ctx, address, role)
}

func Login(ctx *gin.Context) {
	addr := common.HexToAddress(ctx.Request.FormValue("login"))
	password := ctx.Request.FormValue("password")
	user, err := variables.Contract.Users(variables.DefaultCallOpts(), addr)

	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	acc := variables.ImportAccount(addr)
	if acc == nil {
		RedirectToError(ctx, errors.New("invalid login"))
		return
	}
	if err := variables.Keystore.Unlock(*acc, password); err != nil {
		RedirectToError(ctx, errors.New("invalid password"))
		return
	}

	RedirectToRolePage(ctx, addr.String(), user.Role)
}
