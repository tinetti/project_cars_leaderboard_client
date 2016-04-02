package main

import (
    "io/ioutil"
    "fmt"
)

type DirectoryPacketReader struct {
    Directory string
}

func (reader *DirectoryPacketReader) ReadPackets(onRead OnRead) error {
    return ReadPacketsFromDirectory(reader.Directory, onRead)
}

func ReadPacketsFromDirectory(directory string, onRead OnRead) error {
    fmt.Println("Reading packets from directory:", directory)

    files, err := ioutil.ReadDir(directory)
    if err != nil {
        return err
    }

    for _, f := range files {
        path := fmt.Sprintf("%v/%v", directory, f.Name())
        if f.IsDir() {
            dirError := ReadPacketsFromDirectory(path, onRead)
            if dirError != nil {
                LogError(dirError, "reading packets from directory", path)
                err = dirError
            }
        } else {
            contents, readErr := ioutil.ReadFile(path)
            if readErr != nil {
                LogError(readErr, "reading file", path)
                err = readErr
            }

            packet, umErr := Unmarshal(contents)
            if umErr != nil {
                LogError(umErr, "unmarshaling packet", path)
                err = umErr
            }

            onRead(packet)
        }
    }

    return err
}