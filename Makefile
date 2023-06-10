# Compile the Solidity contract and generate the ABI and Go bindings
sol:
	solc --optimize --abi ./internal/contracts/CertificateEmitter.sol -o build 
	solc --optimize --bin ./internal/contracts/CertificateEmitter.sol -o build 
	abigen --abi=./build/CertificateEmitter.abi --bin=./build/CertificateEmitter.bin --pkg=compliance --out=./internal/compliance/CertificateEmitter.go

# Build the Go binary
build:
	make sol
	go build -o ./bin/compliance ./cmd/compliance

# Clean the Solidity build artifacts
clean-sol:
	rm -rf ./build/*

# Clean the Go build artifacts
clean-go:
	rm -rf ./bin/*

# Clean all build artifacts
clean: clean-sol clean-go

.PHONY: sol build clean clean-sol clean-go