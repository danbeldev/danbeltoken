package routing

import (
	"danbeltoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"time"
)

type NFT struct {
	Id                        *big.Int
	Addr                      common.Address
	Name                      string
	Desc                      string
	Price                     *big.Int
	CreateDateInUnixTimestamp uint64
	Count                     *big.Int
}

func CreateNft(ctx *gin.Context) {
	nBigInt := new(big.Int)
	name := ctx.Request.FormValue("nft_name")
	desc := ctx.Request.FormValue("nft_desc")
	price, _ := nBigInt.SetString(ctx.Request.FormValue("nft_price"), 0)
	count, _ := nBigInt.SetString(ctx.Request.FormValue("nft_count"), 0)
	currentTime := uint64(time.Now().Unix())

	_, err := variables.Contract.CreateNft(variables.DefaultTransactOpts(), variables.Owner, name, desc, price, currentTime, count)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}

	RedirectFromRequestToRolePage(ctx)
}

func AllNft() []*NFT {
	nfts := make([]*NFT, 0)
	var index int64
	for {
		nft, err := GetNftById(big.NewInt(index))
		index++
		if err != nil {
			break
		}
		nfts = append(nfts, &nft)
	}
	return nfts
}

func GetNftById(id *big.Int) (NFT, error) {
	nft, err := variables.Contract.Nfts(variables.DefaultCallOpts(), id)
	if err != nil {
		return NFT{}, err
	}
	return nft, nil
}
