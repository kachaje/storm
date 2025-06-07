package codec_test

import (
	"fmt"
	"os"

	"github.com/kachaje/storm/v3"
	"github.com/kachaje/storm/v3/codec/gob"
	"github.com/kachaje/storm/v3/codec/json"
	"github.com/kachaje/storm/v3/codec/msgpack"
	"github.com/kachaje/storm/v3/codec/protobuf"
	"github.com/kachaje/storm/v3/codec/sereal"
)

func Example() {
	defer func() {
		for _, file := range []string{"gob.db", "json.db", "msgpack.db", "sereal.db", "protobuf.db"} {
			os.RemoveAll(file)
		}
	}()

	// The examples below show how to set up all the codecs shipped with Storm.
	// Proper error handling left out to make it simple.
	var gobDb, _ = storm.Open("gob.db", storm.Codec(gob.Codec))
	var jsonDb, _ = storm.Open("json.db", storm.Codec(json.Codec))
	var msgpackDb, _ = storm.Open("msgpack.db", storm.Codec(msgpack.Codec))
	var serealDb, _ = storm.Open("sereal.db", storm.Codec(sereal.Codec))
	var protobufDb, _ = storm.Open("protobuf.db", storm.Codec(protobuf.Codec))

	fmt.Printf("%T\n", gobDb.Codec())
	fmt.Printf("%T\n", jsonDb.Codec())
	fmt.Printf("%T\n", msgpackDb.Codec())
	fmt.Printf("%T\n", serealDb.Codec())
	fmt.Printf("%T\n", protobufDb.Codec())

	// Output:
	// *gob.gobCodec
	// *json.jsonCodec
	// *msgpack.msgpackCodec
	// *sereal.serealCodec
	// *protobuf.protobufCodec
}
