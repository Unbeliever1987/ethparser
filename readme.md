In the `main.go` file, there is only simple test because of test limitation. I think it is enough to show case the three interface functions.

```
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
```
We can change the 4th parameter `0x14912e1` to `latest` [ref](https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_getblockbynumber), it will fetch the latest block and write into the repository automatically, but this will make the testing hard. So for testing purpose, we can hardcode a past block number and find a address from the response and use the selected address to subcribe and test. But in the real use scenario, we should use `latest` parameter all the time.
