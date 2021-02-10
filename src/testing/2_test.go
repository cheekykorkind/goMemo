package testing

import (
	"jiho_local/basic"
	"testing"
)

func TestIf1(t *testing.T) {
	if basic.If1() != false {
		t.Error("Wrong if1()")
	}
}

func TestFor1(t *testing.T) {
	result := basic.For1(10)

	if result == 11 {
		t.Error("Wrong For1(10)")
	}
}

func TestFor2(t *testing.T) {
	result := basic.For2(20)

	if result == 21 {
		t.Error("Wrong For2(20)")
	}
}

func TestPointer1(t *testing.T) {
	msg := "myMessage"

	basic.Pointer1(&msg)

	if msg != "changed!" {
		t.Errorf("Wrong Pointer1 : %s", msg)
	}
}
