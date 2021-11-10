package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TLSKey  = "tls-key.pem"
	TLSCert = "tls-cert.pem"
	TLSCA   = "ca.pem"
)

func NewSecureClient() *http.Client {
	rootPEM, err := ioutil.ReadFile(TLSCA)
	if err != nil {
		panic(err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(TLSCert, TLSKey)
	if err != nil {
		panic(err)
	}

	tlsConf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      roots,
	}
	tr := &http.Transport{TLSClientConfig: tlsConf}
	return &http.Client{Transport: tr}
}

func main() {
	client := NewSecureClient()
	res, err := client.Get("https://localhost:4040")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
