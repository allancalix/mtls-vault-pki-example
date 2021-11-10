# mTLS With Vault Certificate Authority

An example of how to perform mutual TLS authentication using Vault's built in
certificate authority. By following the steps outlined here you will:
  * Start a local Vault server

  * Create a root certificate to use in Vault's CA

  * Prepare Vault to be able to issue certificates using our root certificate

  * Issue certificates for an example **client** and **server** binary written in
    Go.

  * Execute a mutually authenticated request on localhost from our client to our
    server.

## Vault Root Certificates

Vault provides you with two options for root certificates: you can pass in your
own certificate (like we are doing here) or you can have Vault generate a self-signed
certificate. Vault's documentation recommends passing in a certificate you own
so that's what we do here.

In a real use, you would not want to use your root certificate directly in Vault.
Instead, you should take the extra step of creating an intermediate certificate
with a stricter TTL.

## Prerequisites

The scripts in this repository depend on a few binaries to run:

  * Vault - used as a server and as our interface for interacting with Vault
  * jq - used to parse Vault secret payloads into TLS files
  * openssl - Used to create root certificate
  * Go (optional) - used to execute example binaries

**All scripts in this repository are meant to be run from the root directory. Running
  scripts from some other directories might result in things not ending up where they
  are supposed to.**

We'll need a couple open ports for everything to work correctly.

  * localhost:8200 - Vault HTTP server
  * localhost:4040 - TLS server listener

## Steps


1. Start a local Vault development server
```bash
./scripts/init.sh
```

While the server is running you'll see a ROOT_TOKEN displayed in the logs. You
should set a couple environments for the Vault client to authenticate for future
requests using that token. In your current terminal window, run the following to
set the required environment variables and replace the root token provided in
Vault's startup output.

```bash
export VAULT_TOKEN="<ROOT_TOKEN_FROM_OUTPUT>"
export VAULT_ADDR="http://127.0.0.1:8200"
```

2. Create TLS certificates and initialize Vault's PKI secret backend
```bash
./scripts/generate-root.sh
./scripts/install-root.sh
```

3. Create certificates for client and server on localhost
```bash
./scripts/create-client-certs.sh
./scripts/create-server-certs.sh
```

4. Run example binaries
```bash
# In one terminal window.
cd server && go run main.go

cd client && go run main.go
```
