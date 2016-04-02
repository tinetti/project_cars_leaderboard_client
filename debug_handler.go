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
        fmt.Printf("participant packet -> car_name:%v, car_class:%v, track_location:%v, track_variation:%v\n", packet.Participants.GetCarName(), packet.Participants.GetCarClassName(), packet.Participants.GetTrackLocation(), packet.Participants.GetTrackVariation())
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