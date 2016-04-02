package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "fmt"
)

type ServerWriterHandler struct {
    URL          string

    LastLapTime  float32
    Participants Participants
}

func NewServerWriterHandler(URL string) ServerWriterHandler {
    return ServerWriterHandler{
        URL:URL,
    }
}

func (handler *ServerWriterHandler) HandlePacket(packet *Packet) {
    switch (packet.Header.GetPacketType()) {
    case PacketType_TELEMETRY:
        if packet.Telemetry.LastLapTime != handler.LastLapTime {
            lapTime := NewLapTime(packet.Telemetry, handler.Participants)
            _, err := handler.PostLapTime(lapTime)
            LogError(err, "posting lap time")
        }
        handler.LastLapTime = packet.Telemetry.LastLapTime

    case PacketType_PARTICIPANT:
        fmt.Printf("participant packet -> car_name:%v, car_class:%v, track_location:%v, track_variation:%v\n", packet.Participants.GetCarName(), packet.Participants.GetCarClassName(), packet.Participants.GetTrackLocation(), packet.Participants.GetTrackVariation())
        handler.Participants = packet.Participants
        break

    case PacketType_PARTICIPANT_ADDITIONAL:
    default:
        break
    }
}

func (handler *ServerWriterHandler) PostLapTime(lapTime LapTime) ([]byte, error) {
    jsonBytes, err := json.Marshal(lapTime)
    if err != nil {
        return jsonBytes, err
    }

    fmt.Println("posting lap time", string(jsonBytes))

    req, err := http.NewRequest("POST", handler.URL, bytes.NewBuffer(jsonBytes))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()

    return jsonBytes, err
}