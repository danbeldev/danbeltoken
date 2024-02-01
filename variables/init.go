package variables

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Client   *ethclient.Client
	Keystore *keystore.KeyStore
	//Contract *contract.Contract
)

func Init() {

}
