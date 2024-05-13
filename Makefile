# Target: pushd
# Description: Creates the "certs" directory if it doesn't exist.
pushd:
	@echo "Creating certs directory if it doesn't exist"
	@mkdir -p certs

# Target: ca-gen
# Description: Generates the CA (Certificate Authority) certificate and key.
ca-gen: pushd
	@echo "Generating CA certificate"
	@openssl req -new -x509 -days 1024 -nodes -keyout certs/ca.key -out certs/ca.crt -subj "/C=NL/ST=North-Holland/L=Amsterdam/O=My Company/OU=DevOps/CN=localhost"

# Target: server-gen
# Description: Generates the server certificate and key signed by the CA.
server-gen: ca-gen
	@echo "Generating server certificate"
	@openssl req -new -nodes -out certs/server.csr -keyout certs/server.key -subj "/C=NL/ST=North-Holland/L=Amsterdam/O=My Company/OU=DevOps/CN=localhost"
	@openssl x509 -req -in certs/server.csr -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/server.crt -days 1024

# Target: client-gen
# Description: Generates the client certificate and key signed by the CA.
client-gen: ca-gen
	@echo "Generating client certificate"
	@openssl req -new -nodes -out certs/client.csr -keyout certs/client.key -subj "/C=NL/ST=North-Holland/L=Amsterdam/O=My Company/OU=DevOps/CN=localhost"
	@openssl x509 -req -in certs/client.csr -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/client.crt -days 1024

# Target: all
# Description: Generates all certificates.
all: pushd ca-gen server-gen client-gen
	@echo "All certificates generated"

# Target: clean
# Description: Removes all generated certificates.
clean:
	@echo "Removing all generated certificates"
	@rm -rf certs

.PHONY: pushd ca-gen server-gen client-gen all
