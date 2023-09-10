package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/Sanchit011/blockchaingo/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version: 1,
		PrevBlock: types.RandomHash(),
		TimeStamp: time.Now().UnixNano(),
		Height: 10,
		Nonce: 1234,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buff))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buff))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version: 1,
			PrevBlock: types.RandomHash(),
			TimeStamp: time.Now().UnixNano(),
			Height: 10,
			Nonce: 1234,
		},
		Transactions: nil,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buff))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buff))
	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version: 1,
			PrevBlock: types.RandomHash(),
			TimeStamp: time.Now().UnixNano(),
			Height: 10,
			Nonce: 1234,
		},
		Transactions: nil,
	}

	h := b.Hash()
	assert.False(t, h.IsZero())
}