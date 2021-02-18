package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

//go:generate go run gen.go

const defaultProxyServerPort = 8000
const defaultKodiHTTPPort = 8080
const defaultKodiRPCPort = 9090

type request struct {
	URL string
}

var kodiHTTPPort *int
var kodiRPCPort *int

func main() {
	serverPort := flag.Int("server", defaultProxyServerPort, "Proxy server port")
	kodiHTTPPort = flag.Int("http", defaultKodiHTTPPort, "Kodi web port")
	kodiRPCPort = flag.Int("rpc", defaultKodiRPCPort, "Kodi rpc port")
	flag.Parse()

	fmt.Printf("Starting server on http://localhost:%d", *serverPort)
	fmt.Println()
	http.HandleFunc("/jsonrpc", handleRPC)
	http.HandleFunc("/", rootpage)
	http.ListenAndServe(fmt.Sprintf(":%d", *serverPort), nil)
}

func handleRPC(w http.ResponseWriter, req *http.Request) {
	s := request{}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &s); err != nil || s.URL == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Send request to kodi
	payload := bytes.NewBufferString(generatePayload(s.URL))
	rsp, err := http.Post(fmt.Sprintf("http://localhost:%d/jsonrpc", *kodiHTTPPort), "application/json", payload)
	if err != nil {
		panic(err)
	}

	// Get response from request to kodi and return in to the original requester
	response, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	w.Write(response)
}
