package gob

import (
	"testing"

	"github.com/kachaje/storm/v3/codec/internal"
)

func TestGob(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
