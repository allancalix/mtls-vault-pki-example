#!/usr/bin/env bash

set -e

source scripts/__env.sh

clean

openssl genrsa -out "$ROOT_KEY_FILE" 2048
openssl req -x509 -sha256 -new -nodes \
  -config "$OPENSSL_CONFIG" \
  -key "$ROOT_KEY_FILE" -days 3650 \
  -out "$ROOT_CERT_FILE"

echo "$ROOT_PEM_BUNDLE"
cat "$ROOT_KEY_FILE" >> "$ROOT_PEM_BUNDLE"
cat "$ROOT_CERT_FILE" >> "$ROOT_PEM_BUNDLE"
