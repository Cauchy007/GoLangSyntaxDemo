package hello

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "hello world" {
		t.Fatal("hello world")
	}
}
