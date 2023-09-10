package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/Sanchit011/blockchaingo/types"
)

type Header struct {
	Version  uint32
	PrevBlock types.Hash
	TimeStamp int64
	Height uint32
	Nonce uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.PrevBlock)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.TimeStamp)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.Height)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.Nonce)
	if err != nil {
		return err
	}
	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	err := binary.Read(r, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &h.PrevBlock)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &h.TimeStamp)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &h.Height)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &h.Nonce)
	if err != nil {
		return err
	}
	return nil
}

type Block struct {
	Header
	Transactions []Transaction

	// cache of header hash
	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	if !b.hash.IsZero() {
		return b.hash
	}

	buff := &bytes.Buffer{}
	b.Header.EncodeBinary(buff)

	b.hash = types.Hash(sha256.Sum256(buff.Bytes()))
	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {
	err := b.Header.EncodeBinary(w)
	if err != nil {
		return err
	}
	for _, tx := range(b.Transactions) {
		err = tx.EncodeBinary(w)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	err := b.Header.DecodeBinary(r)
	if err != nil {
		return err
	}
	for _, tx := range(b.Transactions) {
		err = tx.DecodeBinary(r)
		if err != nil {
			return err
		}
	}

	return nil
}