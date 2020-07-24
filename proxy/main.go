package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	client *http.Client
)

func main() {
	client = &http.Client{}
	port, ok := os.LookupEnv("PROXY_PORT")
	if !ok {
		port = "8080"
	}
	log.Printf("Proxy listening on port: %s", port)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "proxy returned")

	// req, err := CallPython()
	// if err != nil {
	// 	fmt.Fprintf(w, "proxy error calling python: %s", err)
	// }
	// fmt.Fprintf(w, "proxy returned: %s", req)

}

type Request struct {
	Name string `json:"name"`
}

func CallPython() (*Request, error) {
	httpAddr := "http://127.0.0.1"
	httpPort := 54321
	req := Request{Name: "annie"}
	serverURL := fmt.Sprintf("%s:%d", httpAddr, httpPort)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling execute custom function request: %w", err)
	}

	resp, err := client.Post(serverURL, "application/json", bytes.NewBuffer(reqBody))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading execute custom function response: %w", err)
	}

	var execResp Request
	if err := json.Unmarshal(body, &execResp); err != nil {
		return nil, fmt.Errorf("error unmarshaling execute custom function response: %w", err)
	}

	return &execResp, nil
}
