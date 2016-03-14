package main

import (
    "fmt"
    "os"
    "time"
    "io/ioutil"
)

type PacketHandler interface {
    Handle(msg []byte)
}

type FileWriterHandler struct {
    OutputDir string
}

func (fileWriter FileWriterHandler) Handle(msg []byte) {
    now := time.Now()
    filename := fmt.Sprintf("/tmp/%s.pcars_bin", now)
    mode := os.FileMode(0644)
    err := ioutil.WriteFile(filename, msg, mode)
    CheckError(err)

    fmt.Println("Wrote", filename)

}

type ServerWriterHandler struct {
    ServerUrl string
}

func (serverWriter ServerWriterHandler) Handle(msg []byte) {
    fmt.Println("writing packet to server", serverWriter.ServerUrl)
}
