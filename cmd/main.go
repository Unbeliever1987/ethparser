package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Unbeliever1987/ethparser"
)

func main() {
	fmt.Println("Eth Parser")
	ctx := context.Background()

	ethParser := ethparser.New(ctx, "db-connection", "https://ethereum-rpc.publicnode.com", "0x14912e1")
	if success := ethParser.Subscribe(ctx, "0x2aca0c7ed4d5eb4a2116a3bc060a2f264a343357"); !success {
		log.Fatal("Failed to subscribe the address")
	}

	time.Sleep(5 * time.Second)

	fmt.Printf("CurrentBlock: %d\n", ethParser.GetCurrentBlock(ctx))
	fmt.Printf("Transactions: %v\n", ethParser.GetTransactions(ctx, "0x2aca0c7ed4d5eb4a2116a3bc060a2f264a343357"))
}
