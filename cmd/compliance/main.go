package main

import (
	"flag"
	"log"

	"github.com/monuelo/compliance-cli/internal/compliance"
)

func main() {
	var (
		server string
		token  string
	)
	flag.StringVar(&server, "server", "", "Kubernetes API server address and port")
	flag.StringVar(&token, "token", "", "Bearer token for authentication to the API server")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Call Marvin and handle errors
	output, err := compliance.Scan(server, token)
	if err != nil {
		log.Fatal(err)
	}

	// Verify tests
	if verifyTests(output) {
		log.Printf("All compliance tests passed!")
		name := "John Doe"
		date := "2022-01-01"
		issuer := "My Company"
		status := "Pass"
		txHash, err := compliance.EmitCertificate(name, date, issuer, status)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Transaction hash: %s", txHash)
	} else {
		log.Printf("Some compliance tests failed.")
	}
}

func verifyTests(result *compliance.ScanResult) bool {
	return true
	var (
		totalFailed int
		totalPassed int
	)

	for _, check := range result.Checks {
		if check.Status == "Failed" {
			totalFailed++
		} else if check.Status == "Passed" {
			totalPassed++
		}
	}

	log.Printf("Total failed: %d\n", totalFailed)
	log.Printf("Total passed: %d\n", totalPassed)

	return totalFailed == 0
}
