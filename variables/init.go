package variables

import (
	"danbeltoken/contract"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Miner    = common.HexToAddress("0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151")
	MinerPWD = "miner"
)

var (
	Client   *ethclient.Client
	Keystore *keystore.KeyStore
	Contract *contract.Contract
)

func Init() {

}
