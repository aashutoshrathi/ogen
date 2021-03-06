package p2p_test

import (
	"github.com/olympus-protocol/ogen/pkg/p2p"
	testdata "github.com/olympus-protocol/ogen/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMsgTx(t *testing.T) {
	v := new(p2p.MsgTx)
	v.Data = testdata.FuzzTx(1)[0]

	ser, err := v.Marshal()
	assert.NoError(t, err)

	desc := new(p2p.MsgTx)
	err = desc.Unmarshal(ser)
	assert.NoError(t, err)

	assert.Equal(t, v, desc)

	assert.Equal(t, p2p.MsgTxCmd, v.Command())
	assert.Equal(t, uint64(188), v.MaxPayloadLength())

}
