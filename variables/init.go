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
	Owner    = common.HexToAddress("0x47C57dC89dE3C53AeaBE5B0edAFf5F4E3C3F0AB1")
	OwnerPWD = "owner"
	Tom      = common.HexToAddress("0x9a794BDBAfB0FacAAEE920461D7aC60c512EAD11")
	TomPWD   = "tom"
	Max      = common.HexToAddress("0x77a82cb37120C39B9F2a215B30a869D9dbF72c96")
	MaxPWD   = "max"
	Jack     = common.HexToAddress("0x0f7dF5FD4B743a0fad2e00BCd90428CE147be186")
	JackPWD  = "jack"
	Miner    = common.HexToAddress("0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151")
	MinerPWD = "miner"
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
	_, _, cn, err := contract.DeployContract(DefaultTransactOpts(), Client, Owner, Tom, Max, Jack)
	if err != nil {
		log.Fatal(err)
	}
	Contract = cn
	k.Unlock(*ImportAccount(Tom), TomPWD)
	k.Unlock(*ImportAccount(Max), MaxPWD)
	k.Unlock(*ImportAccount(Jack), JackPWD)
}
