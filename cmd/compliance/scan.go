package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/monuelo/compliance-cli/internal/compliance"
)

func init() {
	// Define "scan" command
	cmd := flag.NewFlagSet("scan", flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Fprintf(cmd.Output(), "Usage: %s scan [flags]\n", os.Args[0])
		cmd.PrintDefaults()
	}
	server := cmd.String("server", "", "Kubernetes API server address and port")
	token := cmd.String("token", "", "Bearer token for authentication to the API server")
	output := cmd.String("output", "json", "Output format: json or table")
	cmd.Parse(os.Args[2:])

	// Call Marvin and handle errors
	result, err := compliance.Scan(*server, *token)
	if err != nil {
		log.Fatal(err)
	}

	// Print output in the specified format
	switch *output {
	case "json":
		_, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(jsonResult))
	case "table":
		// TODO: Implement table output
	default:
		log.Fatalf("Invalid output format: %s", *output)
	}
}
