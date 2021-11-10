export CERTS_DIR="certs/"

export OPENSSL_CONFIG="scripts/openssl.cnf"

export ROOT_KEY_FILE="${CERTS_DIR}root-ca-key.pem"
export ROOT_CERT_FILE="${CERTS_DIR}root-ca-cert.pem"

export ROOT_PEM_BUNDLE="${CERTS_DIR}bundle.pem"

function clean() {
  if [[ -f "$ROOT_KEY_FILE" ]]; then
    rm "$ROOT_KEY_FILE"
  fi

  if [[ -f "$ROOT_CERT_FILE" ]]; then
    rm "$ROOT_CERT_FILE"
  fi

  if [[ -f "$ROOT_PEM_BUNDLE" ]]; then
    rm "$ROOT_PEM_BUNDLE"
  fi
}
