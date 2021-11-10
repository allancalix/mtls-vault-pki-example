#!/usr/bin/env bash

set -e

source scripts/__env.sh

tls=$(vault write pki/issue/tester common_name=localhost --format=json)

echo "$tls" | jq --raw-output '.data.certificate' > server/tls-cert.pem
echo "$tls" | jq --raw-output '.data.private_key' > server/tls-key.pem
echo "$tls" | jq --raw-output '.data.issuing_ca' > server/ca.pem
