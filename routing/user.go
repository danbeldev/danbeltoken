package routing

import (
	_ "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string
}

func GetUser(ctx *gin.Context) {
	user := User{"Danila"}
	ctx.JSON(http.StatusOK, user)
}
