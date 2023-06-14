package compliance

import (
	"context"
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Function to create a new CertificateEmiter contract instance
func NewCertificateEmiter(contractAddress common.Address, client *ethclient.Client) (*Compliance, error) {
	// Load the contract instance
	instance, err := NewCompliance(contractAddress, client)
	if err != nil {
		return nil, err
	}

	// Return the contract instance
	return instance, nil
}

// Function to emit a certificate
func EmitCertificate(name string, date string, issuer string, status string) (string, error) {
	// Create a new client instance
	client, err := NewClient()
	if err != nil {
		return "", err
	}

	contract := os.Getenv("CONTRACT_ADDRESS")
	if contract == "" {
		return "", errors.New("CONTRACT_ADDRESS env is required")
	}

	pkey := os.Getenv("PRIVATE_KEY")
	if pkey == "" {
		return "", errors.New("PRIVATE_KEY env is required")
	}

	receiverPublicKey := os.Getenv("RECEIVER_PUBLIC_KEY")
	if receiverPublicKey == "" {
		return "", errors.New("RECEIVER_PUBLIC_KEY env is required")
	}

	log.Printf("Loading contract instance...")
	// Load the contract instance
	contractAddress := common.HexToAddress(contract) // replace with actual contract address
	certificateEmiter, err := NewCertificateEmiter(contractAddress, client)
	if err != nil {
		return "", err
	}

	// *ecdsa.PrivateKey from string (replace with actual private key)
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return "", err
	}

	receiverAddress := common.HexToAddress(receiverPublicKey)

	log.Printf("Starting certificate emission...")
	tx, err := certificateEmiter.EmitCertificate(auth, receiverAddress, name, date, issuer, status)

	if err != nil {
		return "", err
	}

	// Wait for the transaction to be mined
	log.Printf("Waiting for the transaction to be mined...")
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return "", err
	}

	// Return the transaction hash
	return receipt.TxHash.String(), nil
}
