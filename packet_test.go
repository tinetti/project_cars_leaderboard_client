package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "testing"
    "bytes"
    "unsafe"
)

func compare(t *testing.T, name string, actual interface{}, expected interface{}) {
    if actual != expected {
        t.Error(fmt.Sprintf("%v: actual %T(%v) != expected %T(%v)", name, actual, actual, expected, expected))
    }
}

func TestParse(t *testing.T) {
    filenames := []string{
        "test/pcars_udp_0.bin",
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

        packet, err := Parse(contents)
        if err != nil {
            t.Error("parse error", err)
        }

        lapTime := CreateLapTime(packet)

        b, err := json.Marshal(lapTime)
        if err != nil {
            t.Error("json error", err)
        }
        //os.Stdout.Write(bytes)
        var out bytes.Buffer
        json.Indent(&out, b, "", "\t")
        out.WriteTo(os.Stdout)
    }
}