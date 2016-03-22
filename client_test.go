package main

import (
    "testing"
)

func TestHandleBytes(t *testing.T) {
    mockHandler := &MockHandler{}
    client := &Client{Handlers:[]PacketHandler{mockHandler}}

    packet, err := ReadFile("test/pcars_udp_0.bin")
    if err != nil {
        t.Error(err)
        return
    }

    client.HandlePacket(packet)

    if m := AssertEquals(len(mockHandler.Packets), 1); m != nil {
        t.Error(m)
    }
}

type MockHandler struct {
    Packets []*Packet
}

func (mockHandler *MockHandler) HandlePacket(packet *Packet) {
    mockHandler.Packets = append(mockHandler.Packets, packet)
}
