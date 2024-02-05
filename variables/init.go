package variables

import (
	"danbeltoken/contract"
	"github.com/ethereum/go-ethereum/common"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	Node2         = "http://0.0.0.0:2222"
	Node2Keystore = "./network/node2/keystore"
)

var (
	Owner        = common.HexToAddress("0x47C57dC89dE3C53AeaBE5B0edAFf5F4E3C3F0AB1")
	OwnerPWD     = "owner"
	Investor1    = common.HexToAddress("0x873a20a8417530BfDfc7d0057fA68739BcE35438")
	Investor1PWD = "investor1"
	Investor2    = common.HexToAddress("0x20E47B326bb755963E6998a525E08eEFacA15718")
	Investor2PWD = "investor2"
	Miner        = common.HexToAddress("0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151")
	MinerPWD     = "miner"
)

var (
	Client   *ethclient.Client
	Keystore *keystore.KeyStore
	Contract *contract.Contract
)

func Init() {
	c, err := ethclient.Dial(Node2)
	if err != nil {
		log.Fatal(err)
	}
	Client = c
	k := keystore.NewKeyStore(Node2Keystore, keystore.StandardScryptN, keystore.StandardScryptP)
	Keystore = k
	Keystore.Unlock(*ImportAccount(Owner), OwnerPWD)
	_, _, cn, err := contract.DeployContract(DefaultTransactOpts(), Client, Owner, Investor1, Investor2)
	if err != nil {
		log.Fatal(err)
	}
	Contract = cn
	k.Unlock(*ImportAccount(Investor1), Investor1PWD)
	//k.Unlock(*ImportAccount(Investor2), Investor2PWD)
}
