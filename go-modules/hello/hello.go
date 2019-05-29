package hello

import (
    // "fmt"
    "github.com/satori/go.uuid"
)
func Hello() string {
    uid := uuid.NewV4()
    return uid.String()
}