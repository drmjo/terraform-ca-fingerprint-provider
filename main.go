package main

import (
	"crypto/tls"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

func main () {
	conf := &tls.Config{}
	
	// conn, err := tls.Dial("tcp", "99.84.224.106:443", conf)
	conn, err := tls.Dial("tcp", "mjo.io:443", conf)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	certs := conn.ConnectionState().PeerCertificates

	// the root CA is the last Cert
	rootCA := certs[len(certs)-1]

	sum := sha1.Sum(rootCA.Raw)
	hex := hex.EncodeToString(sum[:])
	fmt.Println(strings.ToUpper(hex))
}