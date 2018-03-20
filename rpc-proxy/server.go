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
const defaultKodiHttpPort = 8080
const defaultKodiRpcPort = 9090

type request struct {
	Url string
}

var kodiHttpPort *int
var kodiRpcPort *int

func main() {
	serverPort := flag.Int("server", defaultProxyServerPort, "Proxy server port")
	kodiHttpPort = flag.Int("http", defaultKodiHttpPort, "Kodi web port")
	kodiRpcPort = flag.Int("rpc", defaultKodiRpcPort, "Kodi rpc port")
	flag.Parse()

	fmt.Println("Starting server on port", *serverPort)
	http.HandleFunc("/jsonrpc", HandleRpc)
	http.HandleFunc("/", rootpage)
	http.ListenAndServe(fmt.Sprintf(":%d", *serverPort), nil)
}

func HandleRpc(w http.ResponseWriter, req *http.Request) {
	s := request{}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &s); err != nil || s.Url == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Send request to kodi
	payload := bytes.NewBufferString(generatePayload(s.Url))
	rsp, err := http.Post(fmt.Sprintf("http://localhost:%d/jsonrpc", *kodiHttpPort), "application/json", payload)
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
