package main

import (
    "net"
    "fmt"
)

func (reader *NetworkPacketReader) ReadPackets(onRead OnRead) error {
    /* Prepare address at port 5606 */
    addr, err := net.ResolveUDPAddr("udp", ":5606")
    ExitOnError(err)

    /* Now listen at selected port */
    serverConn, err := net.ListenUDP("udp", addr)
    ExitOnError(err)
    defer serverConn.Close()

    fmt.Println("Started listening on port", addr)

    buf := make([]byte, 2048)
    for {
        _, _, err := serverConn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("Error reading UDP packet: ", err)
            continue
        }

        packet, err := Unmarshal(buf)
        if err != nil {
            fmt.Println("Error unmarshaling packet: ", err)
            continue
        }

        onRead(packet)
    }
}
