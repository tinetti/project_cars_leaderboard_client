package main

import (
    "encoding/json"
    "fmt"
)

type DebugHandler struct {
    LastLapTime  float32
    Participants Participants
}

func (handler *DebugHandler) HandlePacket(packet *Packet) {
    //fmt.Printf("handling packet: %v (%v)\n", packet.Header.GetPacketType(), packet.Header.GetSequenceNumber())

    switch (packet.Header.GetPacketType()) {
    case PacketType_TELEMETRY:
        if packet.Telemetry.LastLapTime != handler.LastLapTime {
            lapTime := NewLapTime(packet.Telemetry, handler.Participants)
            lapTime.LogLapTime()
        }
        handler.LastLapTime = packet.Telemetry.LastLapTime

    case PacketType_PARTICIPANT:
        handler.Participants = packet.Participants
        break

    case PacketType_PARTICIPANT_ADDITIONAL:
    default:
        break
    }
}

func (lapTime *LapTime) LogLapTime() {
    j, err := json.Marshal(lapTime)
    LogError(err, "marshaling lap time")
    fmt.Println(string(j))
}