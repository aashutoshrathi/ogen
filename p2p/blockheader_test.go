package p2p

import (
	"bytes"
	"github.com/olympus-protocol/ogen/utils/chainhash"
	"testing"
	"time"
)

var (
	merkleRootTest = chainhash.Hash([chainhash.HashSize]byte{
		0xfc, 0x4b, 0x8c, 0xb9, 0x03, 0xae, 0xd5, 0x4e,
		0x11, 0xe1, 0xae, 0x8a, 0x5b, 0x7a, 0xd0, 0x97,
		0xad, 0xe3, 0x49, 0x88, 0xa8, 0x45, 0x00, 0xad,
		0x2d, 0x80, 0xe4, 0xd1, 0xf5, 0xbc, 0xc9, 0x5d,
	})
	blockHeaderTest = BlockHeader{
		Version:       1,
		PrevBlockHash: chainhash.Hash{},
		MerkleRoot:    merkleRootTest,
		Timestamp:     time.Unix(0x5A3BB72B, 0),
	}
)

func TestBlock_Serialize(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	err := blockHeaderTest.Serialize(buf)
	if err != nil {

	}
	var blockHeader BlockHeader
	err = blockHeader.Deserialize(buf)
	if err != nil {

	}
	hash, err := blockHeader.Hash()
	if err != nil {

	}
	oldhash, err := blockHeaderTest.Hash()
	if err != nil {

	}
	if hash != oldhash {
		t.Error("error headers hash doesn't match")
		return
	}
}
