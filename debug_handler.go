package main

import (
    "encoding/json"
    "fmt"
)

type DebugHandler struct {
}

func (handler *DebugHandler) HandlePacket(packet *Packet) {
    fmt.Printf("handling packet: %v (%v)\n", packet.Header.GetPacketType(), packet.Header.GetSequenceNumber())
    j, err := json.Marshal(packet)
    LogError(err)
    fmt.Println("json:", string(j))
}
