package main

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/ethclient"
    // "github.com/ethereum/go-ethereum/common/hexutil"
    // "github.com/ethereum/go-ethereum/crypto"
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

func main() {
	currentBlock()
}