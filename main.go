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

    thePublicAddress := crypto.PubkeyToAddress(getPrivateKey.PublicKey).Hex()
    return thePublicAddress, thePublicKey

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

    // Making Ethereum transactions in Go using Go-Ethereum

    RecipientAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")

    privateKey, err := crypto.HexToECDSA("The Hexadecimal Private Key ")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Public Key Error")
    }

    SenderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := client.PendingNonceAt(ctx, SenderAddress)
    if err != nil {

        log.Println(err)
    }

    amount := big.NewInt("amount    In Wei")
    gasLimit := 3600
    gas, err := client.SuggestGasPrice(ctx)


    if err != nil {
        log.Println(err)
    }

    ChainID, err := client.NetworkID(ctx)
    if err != nil {
        log.Println(err)
    }

    transaction := types.NewTransaction(nonce, RecipientAddress, amount, uint64(gasLimit), gas, nil)
    signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(ChainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }
    err = client.SendTransaction(ctx, signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("transaction sent: %s", signedTx.Hash().Hex())

}