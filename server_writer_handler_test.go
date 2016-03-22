package main

import (
    "net/http"
    "testing"
    "net/http/httptest"
    "io/ioutil"
    "encoding/json"
)

func TestHandle(t *testing.T) {
    bodies := []string{}
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        contents, err := ioutil.ReadAll(r.Body)
        if err != nil {
            t.Error(err)
        }

        text := string(contents)
        bodies = append(bodies, text)
    }))
    defer ts.Close()

    handler := NewServerWriterHandler(ts.URL)

    // handle participant packet (should not trigger post)
    pPacket, err := ReadFile("test/pcars_udp_1.bin")
    if Error(t, err) {
        return
    }

    handler.HandlePacket(pPacket)
    if m := AssertEquals(len(bodies), 0); m != nil {
        t.Error(m)
    }

    // handle a telemetry packet (should trigger post)
    tPacket1, err := ReadFile("test/pcars_udp_0.bin")
    if Error(t, err) {
        return
    }
    handler.HandlePacket(tPacket1)
    if m := AssertEquals(len(bodies), 1); m != nil {
        t.Error(m)
    } else {
        expectedLapTime := LapTime{
            LapTime:-1,
            CarName:"Formula C",
            CarClassName:"FC",
            TrackLocation:"Brands Hatch",
            TrackVariation:"Indy",
        }
        expectedJsonBytes, err := json.Marshal(expectedLapTime)
        Error(t, err)
        expectedJson := string(expectedJsonBytes)
        if m := AssertEquals(bodies[0], expectedJson); m != nil {
            t.Error(m)
        }
    }

    // handle the same packet data (should not trigger post)
    tPacket2, err := ReadFile("test/pcars_udp_0.bin")
    if Error(t, err) {
        return
    }
    handler.HandlePacket(tPacket2)
    if m := AssertEquals(len(bodies), 1); m != nil {
        t.Error(m)
    }

    // modify lap time, handle a different packet data (should trigger post)
    tPacket3, err := ReadFile("test/pcars_udp_0.bin")
    if Error(t, err) {
        return
    }
    tPacket3.Telemetry.LastLapTime = 120
    handler.HandlePacket(tPacket3)
    if m := AssertEquals(len(bodies), 2); m != nil {
        t.Error(m)
    } else {
        expectedLapTime := LapTime{
            LapTime:120,
            CarName:"Formula C",
            CarClassName:"FC",
            TrackLocation:"Brands Hatch",
            TrackVariation:"Indy",
        }
        expectedJsonBytes, err := json.Marshal(expectedLapTime)
        Error(t, err)
        expectedJson := string(expectedJsonBytes)
        if m := AssertEquals(bodies[1], expectedJson); m != nil {
            t.Error(m)
        }
    }
}
