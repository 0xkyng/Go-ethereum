package main

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "log"
)

var (
    ctx         = context.Background()
    url         = "https://mainnet.infura.io/v3/407f60619ec14e538991ba8f9e0f4237"
    client, err = ethclient.DialContext(ctx, url)
)


func currentBlock() {
    block, err := client.BlockByNumber(ctx, nil)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(block.Number())
}

func createWallet() (string, string) {
    getPrivateKey, err := crypto.GenerateKey()

    if err != nil {
        log.Println(err)
    }

    getPublicKey := crypto.FromECDSA(getPrivateKey)
    thePublicKey := hexutil.Encode(getPublicKey)

}



func main() {
	currentBlock()

    // Querying Ethereum wallet balances with Geth

    address := common.HexToAddress("0x8335659d19e46e720e7894294630436501407c3e")

    balance, err := client.BalanceAt(ctx, address, nil)
    if err != nil {
        log.Print("There was an error", err)
    }
    fmt.Println("The balance aof the provided wallet address is", balance)

    pubAddress, pubKey := createWallet()

    fmt.Printf("Public address:%s, public key: %s\n", pubAddress, pubKey)
}