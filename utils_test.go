package main

import (
    "testing"
    "fmt"
)

func Error(t *testing.T, err error) bool {
    if err != nil {
        t.Error(err)
        return true
    }

    return false
}

func AssertEquals(actual, expected interface{}) interface{} {
    if actual == expected {
        return nil
    }

    return fmt.Sprintf("actual: %v, expected: %v", actual, expected)
}
