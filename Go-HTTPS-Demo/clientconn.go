package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "ca.pem"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}
	config := &tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{cliCrt},
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:8088", config)
	if err != nil {
		fmt.Printf("client: dial: %s\n", err)
	}
	defer conn.Close()
	fmt.Println("client: connected to: ", conn.RemoteAddr())

	req, _ := http.NewRequest("Get", "/", nil)
	reqBuff, _ := httputil.DumpRequest(req, false)
	conn.Write(reqBuff)

	reply := make([]byte, 256)
	n, err := conn.Read(reply)
	fmt.Printf("client read:(%d bytes)\n%q \n", n, string(reply[:n]))
}
