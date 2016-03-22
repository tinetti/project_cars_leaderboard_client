package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type ServerWriterHandler struct {
    ServerUrl string
}

func (serverWriter ServerWriterHandler) Handle(msg []byte) {
    packet, err := Parse(msg)
    CheckError(err)

    json, err := json.Marshal(packet)
    CheckError(err)

    req, err := http.NewRequest("POST", serverWriter.ServerUrl, bytes.NewBuffer(json))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    CheckError(err)
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
