# K8s Compliance Certification System

This repository holds the implementation of a Compliance Certification System for Kubernetes clusters. It leverages advanced scanning techniques, that runs in a Trusted Execution Environment (TEE), and Blockchain technology to ensure the safety, security, and compliance of Kubernetes environments.

## Technologies Utilized

- **SCONE**: This implementation utilizes SCONE, a Intel SGX based Trusted Execution Environment (TEE) solution, to establish a secure environment for the cluster scanning process. SCONE ensures the integrity and confidentiality of sensitive data by encrypting and isolating it from potential threats.

- **Marvin**: The scanning engine is powered by [Marvin](https://undistro.io/marvin), a CLI tool designed to help Kubernetes cluster administrators ensure the security and reliability of their environments. Marvin performs extensive checks on cluster resources and configurations to identify potential issues, misconfigurations, and vulnerabilities.

- **Ethereum**: The compliance certification system integrates with the Ethereum blockchain technology. It leverages Ethereum smart contracts, specifically the `CertificateEmitter` contract, to achieve transparency, traceability, and reliability in the certification process.

## Smart Contract and Build Process

The Ethereum smart contract utilized in this system is called CertificateEmitter. It is defined in the `./internal/contracts/CertificateEmitter.sol` file. To compile the Solidity contract and generate the ABI and Go bindings, run the following command:

```
make sol
```

This command uses the Solidity compiler (`solc`) and the Go compiler (`abigen`) to generate the necessary artifacts. The ABI and Go bindings will be placed in the `./build` directory, and the Go bindings will be generated in the `./internal/compliance` directory.

## Getting Started

To build and execute the system, follow these steps:

1. Clone the repository:

```
git clone https://github.com/your-username/k8s-compliance-certification.git compliance
```

2. Navigate to the project directory:

```
cd compliance
```

3. Compile the Solidity contract and generate the ABI and Go bindings:

```
make sol
```

4. Build the compliance binary:

```
make build
```

5. Execute the compliance certification system:

```
./bin/compliance scan
```

Ensure you have the necessary dependencies installed and properly configured before running the build and execution commands.

## Environment Variables

To configure and run the compliance certification system effectively, you need to set the following environment variables:

- `KUBE_SERVER`: The address of the Kubernetes API server.
- `KUBE_TOKEN`: The token used for authentication with the Kubernetes API server.
- `CONTRACT_ADDRESS`: The Ethereum contract address for the compliance certification system.
- `PRIVATE_KEY`: The private key associated with the Ethereum account used for contract interaction.
- `RECEIVER_PUBLIC_KEY`: The public key of the receiver Ethereum account used for certificate issuance.

Make sure to provide the appropriate values for these variables to ensure the proper functioning of the compliance certification system. You can set these environment variables in your shell environment or in a configuration file, depending on your preference and deployment setup.

Please ensure that you keep the sensitive information, such as private keys and tokens, confidential and secure.

## SCONE Integration

The system integrates with SCONE, a Trusted Execution Environment (TEE) solution, to establish a secure environment for the cluster scanning process. SCONE ensures the integrity and confidentiality of sensitive data by encrypting and isolating it from potential threats.

### Setting Up SCONE

To facilitate the usage of SCONE, we recommend creating an alias called `scone` for the SCONE CLI. You can create this alias by adding the following line to your shell configuration file (e.g., `.bashrc` or `.zshrc`):

```
alias scone="docker run -it --rm \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v \"$HOME/.docker:/root/.docker\" \
    -v \"$HOME/.cas:/root/.cas\" \
    -v \"$HOME/.scone:/root/.scone\" \
    -v \"\$PWD:/root\" \
    -w /root \
    registry.scontain.com/community/cli scone"
```

### Attesting SCONE CAS

Before utilizing SCONE, it is crucial to attest the SCONE CAS (Configuration and Attestation Service) to ensure its integrity and authenticity. You can attest the SCONE CAS by running the following command:

```
scone cas attest $CAS_ADDR $CAS_MRENCLAVE -GCS --only_for_testing-debug --only_for_testing-ignore-signer
```

Replace `$CAS_ADDR` with the address of the SCONE CAS and `$CAS_MRENCLAVE` with the MRENCLAVE of the CAS.

### Creating a Session

To create a session, use the SCONE CLI with a session configuration file (`session.yml`). This file specifies the necessary configurations for the session. You can create a session by running the following command:

```
scone session create session.yml
```

Make sure to customize the `session.yml` file according to your specific requirements and configurations.

By following these steps, you can set up and utilize SCONE within compliance system to establish a secure and trusted execution environment for the scanning process.
