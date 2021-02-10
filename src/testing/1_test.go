package testing

import "testing"

func Test1(t *testing.T) {
	if 6 != 6 {
		t.Error("Wrong result")
	}
}
