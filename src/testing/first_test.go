package testing

import "testing"

func TestFoo(t *testing.T) {
    if 6 != 6 {
        t.Error("Wrong result")
    }
}