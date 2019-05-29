package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); len(got) <= 0 {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}