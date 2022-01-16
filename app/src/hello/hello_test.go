package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, World!"

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHelloImport(t *testing.T) {
	want := "Ahoy, world!"

	if got := HelloImport(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}