package main

import (
	"testing"
)
func TestHello(t *testing.T) {
    msg := Hello()
    if msg != "Hello world" {
        t.Fatalf(`Hello() = %q, needed "Hello world", error`, msg)
    }
}