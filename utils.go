package main

import (
    "fmt"
    "os"
)

func ExitOnError(err error, context ...string) {
    if doLogError(err, context) {
        os.Exit(0)
    }
}

func LogError(err error, context ...string) bool {
    return doLogError(err, context)
}

func WrapError(err error, context ...string) error {
    if err == nil {
        return nil
    }

    return fmt.Errorf("%v %v", err, context)
}

func doLogError(err error, context []string) bool {
    if err != nil {
        fmt.Println("Warn:", err, context)
        return true
    }

    return false
}
