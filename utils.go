package main

import (
    "fmt"
    "os"
)

/* A Simple function to verify error */
func CheckError(err error) {
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(0)
    }
}

