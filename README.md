# Mutual TLS Authentication

This is a simple example of mutual certificates authentication using golang and openssl.

## How to run

- [ ] Generate the certificates using the `Makefile`: It will generate the certificates for the server and the client in the `certs` directory.
   ```bash
   make all
   ```

- [ ] Run the server:
   ```bash
   go run server.go
   ```

- [ ] Run the client:
   ```bash
   go run client.go
   ```

## How it works

1. Certificate Authority (CA):
   - A Certificate Authority (CA) is a trusted entity that issues digital certificates.
   - In the current scenario, the server acts as the CA and generates its own CA certificate. This CA certificate is used to sign both the server and client certificates.

2. Server Certificate:
    - The server generates its own certificate signed by the CA.
    - This certificate includes the server's public key and identifying information (e.g., Common Name, Organization) to establish its identity.

3. Client Certificate:
   - The client generates its own certificate signed by the same CA used by the server.
   - This certificate includes the client's public key and identifying information.
      
4. Certificate Authentication:
   - During the TLS handshake, the server presents its certificate to the client.
   - The client verifies the server's certificate using the CA certificate it trusts. If the server's certificate is signed by the trusted CA, the client trusts the server's identity.
   - Similarly, the client presents its certificate to the server.
   - The server verifies the client's certificate using the same CA certificate. If the client's certificate is signed by the trusted CA, the server trusts the client's identity.
   
5. Establishing Secure Connection:
   - After successful certificate authentication, both the server and the client establish a secure TLS connection.
   - This secure connection encrypts data transmitted between the server and the client, ensuring confidentiality and integrity.
   
By using mutual TLS authentication, both the server and the client authenticate each other using their respective certificates signed by the trusted CA. This helps to establish a secure and trusted communication channel between them.