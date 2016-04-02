package main

import (
    "fmt"
    "os"
    "time"
    "io/ioutil"
)

type FileWriterHandler struct {
    OutputDir string
}

func (fileWriter FileWriterHandler) HandlePacket(packet *Packet) {
    now := time.Now()
    filename := fmt.Sprintf("%s/%s.pcars_bin", fileWriter.OutputDir, now)
    mode := os.FileMode(0644)
    err := ioutil.WriteFile(filename, packet.Marshal(), mode)
    ExitOnError(err, "wirting file", filename)

    fmt.Println("Wrote", filename)
}