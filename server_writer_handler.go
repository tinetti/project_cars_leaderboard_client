package main

import (
    "fmt"
)

type ServerWriterHandler struct {
    ServerUrl string
    //state     State
}

type State struct {
    LastParticipantPacket Packet
    LastTelemetryPacket   Packet
}

func (serverWriter ServerWriterHandler) Handle(msg []byte) {
    //packet, err := Parse(msg)
    //CheckError(err)
    //
    //switch (packet.Header.GetPacketType()) {
    //case PacketType_PARTICIPANT:
    //    serverWriter.state.LastParticipantPacket = packet
    //
    //case PacketType_TELEMETRY:
    //    serverWriter.state.LastTelemetryPacket = packet
    //}

    fmt.Println("writing packet to server", serverWriter.ServerUrl)
}
