package main

import (
    "fmt"
    "os"
)

func ExitOnError(err error) {
    if LogError(err) {
        os.Exit(0)
    }
}

func LogError(err error) bool {
    if err != nil {
        fmt.Println("Warn: ", err)
        return true
    }

    return false
}
