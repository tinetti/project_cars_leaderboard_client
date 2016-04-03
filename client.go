package main

import (
    "flag"
    "fmt"
)

type Client struct {
    Reader   PacketReader
    Handlers []PacketHandler
}

func main() {
    client, err := createClient()
    if err != nil {
        LogError(err, "creating client")
        return
    }

    err = client.Reader.ReadPackets(func(packet *Packet) {
        client.HandlePacket(packet)
    })

    LogError(err, "reading packets")
}

func (client *Client) HandlePacket(packet *Packet) {
    for i := 0; i < len(client.Handlers); i++ {
        handler := client.Handlers[i]
        handler.HandlePacket(packet)
    }
}

func createClient() (*Client, error) {

    var inputDir string
    flag.StringVar(&inputDir, "input-dir", "", "input directory")

    var outputDir string
    flag.StringVar(&outputDir, "output-dir", "", "output directory")

    var serverUrl string
    flag.StringVar(&serverUrl, "server-url", "", "server url")
    var username string
    flag.StringVar(&username, "username", "", "username")
    var password string
    flag.StringVar(&password, "password", "", "password")

    var debugMode bool
    flag.BoolVar(&debugMode, "debug", false, "debug mode")

    flag.Parse()

    var reader PacketReader
    if len(inputDir) > 0 {
        reader = &DirectoryPacketReader{Directory:inputDir}
    } else {
        reader = &NetworkPacketReader{}
    }

    handlers := []PacketHandler{}
    if len(outputDir) > 0 {
        handler := &FileWriterHandler{outputDir}
        handlers = append(handlers, handler)
    }
    if len(serverUrl) > 0 {
        handler := &ServerWriterHandler{URL:serverUrl, Username:username, Password:password}
        handlers = append(handlers, handler)
    }
    if debugMode {
        handler := &DebugHandler{}
        handlers = append(handlers, handler)
    }

    var err error
    if len(handlers) == 0 {
        err = fmt.Errorf("no handlers defined - run again with --help for usage")
    }

    return &Client{Reader:reader, Handlers:handlers}, err
}

type OnRead func(*Packet)

type PacketReader interface {
    ReadPackets(onRead OnRead) error
}

type NetworkPacketReader struct {

}

type PacketHandler interface {
    HandlePacket(packet *Packet)
}
