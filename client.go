package main

import (
    "flag"
    "fmt"
    "net"
)

func main() {
    handlers := make([]PacketHandler, 0)

    var outputDir string
    flag.StringVar(&outputDir, "output-dir", "", "output directory")

    var serverUrl string
    flag.StringVar(&serverUrl, "server-url", "", "server url")

    flag.Parse()

    if len(outputDir) > 0 {
        handler := FileWriterHandler{outputDir}
        handlers = append(handlers, handler)
    }
    if len(serverUrl) > 0 {
        handler := ServerWriterHandler{serverUrl}
        handlers = append(handlers, handler)
    }

    if (len(handlers) == 0) {
        println("no handlers defined.  run again with --help")
        return
    }

    fmt.Println("Sending data to", handlers)


    /* Prepare address at port 5606 */
    addr, err := net.ResolveUDPAddr("udp", ":5606")
    CheckError(err)

    /* Now listen at selected port */
    serverConn, err := net.ListenUDP("udp", addr)
    CheckError(err)
    defer serverConn.Close()

    fmt.Println("Started listening on port", addr)

    buf := make([]byte, 2048)
    for {
        _, _, err := serverConn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("Error: ", err)
            continue
        }

        for i := 0; i < len(handlers); i++ {
            handler := handlers[i]
            handler.Handle(buf)
        }
    }
}

type PacketHandler interface {
    Handle(msg []byte)
}