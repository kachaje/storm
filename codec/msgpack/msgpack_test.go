package msgpack

import (
	"testing"

	"github.com/kachaje/storm/v3/codec/internal"
)

func TestMsgpack(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
