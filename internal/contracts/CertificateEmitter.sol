// SPDX-License-Identifier: UNLICENSE
pragma solidity ^0.8.0;

contract CertificateEmitter {
    
    // Define the certificate template
    struct CertificateTemplate {
        string name;
        string date;
        string issuer;
        string status;
        uint256 timestamp;
        bytes32 certificateHash;
    }
    
    // Define the mapping to store the certificates
    mapping (address => CertificateTemplate) public certificates;
    
    // Define the event that will be emitted when a certificate is generated
    event CertificateGenerated(address recipient, string name, string date, string issuer, string status, uint256 timestamp, bytes32 certificateHash);
    
    // Define the list of whitelisted addresses
    address[] public whitelist;
    
    // Function to add an address to the whitelist
    function addToWhitelist(address _address) public {
        whitelist.push(_address);
    }
    
    // Function to remove an address from the whitelist
    function removeFromWhitelist(address _address) public {
        for (uint i = 0; i < whitelist.length; i++) {
            if (whitelist[i] == _address) {
                whitelist[i] = whitelist[whitelist.length - 1];
                whitelist.pop();
                break;
            }
        }
    }
    
    // Function to check if an address is whitelisted
    function isWhitelisted(address _address) public view returns (bool) {
        for (uint i = 0; i < whitelist.length; i++) {
            if (whitelist[i] == _address) {
                return true;
            }
        }
        return false;
    }
    
    // Function to generate a certificate and store its hash
    function emitCertificate(address recipient, string memory name, string memory date, string memory issuer, string memory status) public {
        require(isWhitelisted(msg.sender), "Only whitelisted addresses can generate certificates.");
        
        // Create a new certificate using the certificate template
        CertificateTemplate memory certificate = CertificateTemplate({
            name: name,
            date: date,
            issuer: issuer,
            status: status,
            timestamp: block.timestamp,
            certificateHash: keccak256(abi.encodePacked(name, date, issuer, status, block.timestamp, recipient))
        });
        
        // Store the certificate in the mapping
        certificates[recipient] = certificate;
        
        // Emit the CertificateGenerated event
        emit CertificateGenerated(recipient, name, date, issuer, status, block.timestamp, certificate.certificateHash);
    }
    
    // Function to retrieve a certificate by address
    function getCertificate(address recipient) public view returns (string memory, string memory, string memory, string memory, uint256, bytes32) {
        CertificateTemplate memory certificate = certificates[recipient];
        return (certificate.name, certificate.date, certificate.issuer, certificate.status, certificate.timestamp, certificate.certificateHash);
    }
}