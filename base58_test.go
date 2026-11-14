package gobase58

import (
	"bytes"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	inputs := [][]byte{{}, {0}, {0, 0, 1}, []byte("hello world"), {255, 254, 253}}
	for _, in := range inputs {
		text := Encode(in, Bitcoin)
		back, err := Decode(text, Bitcoin)
		if err != nil {
			t.Fatalf("decode %q: %v", text, err)
		}
		if !bytes.Equal(bytes.TrimRight(in, "\x00"), bytes.TrimRight(back, "\x00")) {
			t.Fatalf("round trip %x -> %x", in, back)
		}
	}
}

func TestKnownVector(t *testing.T) {
	if got := Encode([]byte("hello world"), Bitcoin); got != "StV1DL6CwTryKyV" {
		t.Fatalf("vector = %q", got)
	}
}

func TestInvalidCharacter(t *testing.T) {
	if _, err := Decode("0OIl", Bitcoin); err == nil {
		t.Fatal("expected error")
	}
}

func TestAlphabetSize(t *testing.T) {
	if len(Bitcoin) != 58 {
		t.Fatalf("alphabet length = %d", len(Bitcoin))
	}
}
