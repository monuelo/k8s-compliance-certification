package compliance

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Function to create a new Ethereum client instance
func NewClient() (*ethclient.Client, error) {
	// Connect to the local Ganache instance
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		return nil, err
	}

	// Get the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to local Ganache instance (block number: %d)\n", blockNumber)

	// Return the client instance
	return client, nil
}
