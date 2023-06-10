package compliance

import (
	"context"
	"math/big"

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

	// Load the contract instance
	contractAddress := common.HexToAddress("0xBDbbaA60717E4A749415Aa97c879e4442072D3dE") // replace with actual contract address
	certificateEmiter, err := NewCertificateEmiter(contractAddress, client)
	if err != nil {
		return "", err
	}

	// *ecdsa.PrivateKey from string (replace with actual private key)
	privateKey, err := crypto.HexToECDSA("578fb20cc9ca9662d07ff4b5d33e22c4bda2697cd3fd33f7d3b5278ac6c343be")
	if err != nil {
		return "", err
	}

	// Receiver public key
	receiverPublicKey := "0xcD6b5Cf0E7eE343d319Aa93B9939222898109036" // replace with actual public key

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return "", err
	}

	receiverAddress := common.HexToAddress(receiverPublicKey)
	tx, err := certificateEmiter.EmitCertificate(auth, receiverAddress, name, date, issuer, status)

	if err != nil {
		return "", err
	}

	// Wait for the transaction to be mined
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return "", err
	}

	// Return the transaction hash
	return receipt.TxHash.String(), nil
}
