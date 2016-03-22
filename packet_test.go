package main

import (
    "fmt"
    "io/ioutil"
    "testing"
    "unsafe"
    "encoding/json"
    "bytes"
    "os"
)

func compare(t *testing.T, name string, actual interface{}, expected interface{}) {
    if actual != expected {
        t.Error(fmt.Sprintf("%v: actual %T(%v) != expected %T(%v)", name, actual, actual, expected, expected))
    }
}

func TestParse(t *testing.T) {
    return

    filenames := []string{
        //"test/pcars_udp_0.bin",
        "test/pcars_udp_1.bin",
    }
    for i := 0; i < len(filenames); i++ {
        filename := filenames[i]
        contents, err := ioutil.ReadFile(filename)
        if err != nil {
            t.Error("read error", err)
        }
        fmt.Println("read bytes", len(contents))
        fmt.Println("size of struct", unsafe.Sizeof(Packet{}))

        packet, err := Unmarshal(contents)
        if err != nil {
            t.Error("parse error", err)
        }

        fmt.Printf("packet: %v\n", packet)
        b, err := json.Marshal(packet)
        if err != nil {
            t.Error("json error", err)
        }
        //os.Stdout.Write(bytes)
        var out bytes.Buffer
        json.Indent(&out, b, "", " ")
        out.WriteTo(os.Stdout)

        lap := NewLapTime(Telemetry{}, packet.Participants)
        fmt.Printf("lap: %v\n", lap)

    }
}
