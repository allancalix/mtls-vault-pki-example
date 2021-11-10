#!/usr/bin/env bash

set -e

source scripts/__env.sh

if [[ -z $(vault secrets list | grep "pki") ]]; then
  vault secrets enable pki
fi

# Install the root cert/key as the PKI CA root.
vault write pki/config/ca \
    pem_bundle=@"$ROOT_PEM_BUNDLE"

vault write pki/config/urls \
    issuing_certificates="http://127.0.0.1:8200/v1/pki/ca" \
    crl_distribution_points="http://127.0.0.1:8200/v1/pki/crl"

vault write pki/roles/tester \
  allowed_domains=localhost \
  allow_subdomains=true \
  max_ttl=72h
