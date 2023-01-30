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
    url         = "Your Infura URL here"
    client, err = ethclient.DialContext(ctx, url)
)