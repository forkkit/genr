// Code generated 'by go generate ./...'; DO NOT EDIT.

package intarray

import (
	"encoding/binary"
)

type U16 struct {
	vals    []uint16
	eltSize int
	first   []byte
}

func NewU16(elts []uint16) (a *U16, err error) {
	a = &U16{
		vals:    make([]uint16, 10),
		eltSize: 2,
		first:   make([]byte, 0),
	}

	binary.LittleEndian.PutUint16(a.first, elts[0])
	return a, nil
}

type I64 struct {
	vals    []int64
	eltSize int
	first   []byte
}

func NewI64(elts []int64) (a *I64, err error) {
	a = &I64{
		vals:    make([]int64, 10),
		eltSize: 8,
		first:   make([]byte, 0),
	}

	binary.LittleEndian.PutUint64(a.first, uint64(elts[0]))
	return a, nil
}
