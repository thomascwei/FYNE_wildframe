package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}

	test.Type(in, "Andy")
	if out.Text != "Hello Andy!" {
		t.Error("Incorrect user greeting")
	}

}
