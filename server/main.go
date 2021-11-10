package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	TLSKey  = "tls-key.pem"
	TLSCert = "tls-cert.pem"
	TLSCA   = "ca.pem"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})

	b, err := ioutil.ReadFile(TLSCA)
	if err != nil {
		panic(err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(b)
	if !ok {
		panic("failed to parse root CA")
	}

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		RootCAs:                  roots,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                roots,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         "localhost:4040",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS(TLSCert, TLSKey))
}
