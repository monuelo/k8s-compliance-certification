package main

import (
	"log"
	"os"
	"time"

	"github.com/monuelo/compliance-cli/internal/compliance"
)

func main() {
	args := os.Args[1:]

	server := getFlagValue(args, "--server")
	token := getFlagValue(args, "--token")

	if server == "" {
		server = os.Getenv("KUBE_SERVER")
	}
	if token == "" {
		token = os.Getenv("KUBE_TOKEN")
	}

	if server == "" || token == "" {
		log.Fatal("Server address and token are required")
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Call Marvin and handle errors
	output, err := compliance.Scan(server, token)
	if err != nil {
		log.Fatal(err)
	}

	// Verify tests
	if verifyTests(output) {
		log.Printf("All compliance tests passed!")
		name := "Company Name"
		date := time.Now().Format("01-02-2006")
		issuer := "Zora Compliance"
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

func getFlagValue(args []string, flagName string) string {
	for i := 0; i < len(args); i++ {
		if args[i] == flagName && i+1 < len(args) {
			return args[i+1]
		}
	}
	return ""
}
