package main

import (
	"github.com/holiman/uint256"
)

type memoryStruct struct {
	data []byte
}

func newMemory() *memoryStruct {
	return &memoryStruct{
		data: make([]byte, 0),
	}
}

func (m *memoryStruct) get(offset, size uint64) []byte {
	if len(m.data) > int(offset) {
		return m.data[offset : offset+size]
	}
	return nil
}

func (m *memoryStruct) set(offset, size uint64, value []byte) {
	if offset+size > uint64(len(m.data)) {
		panic("invalid memory: store empty")
	}

	if len(value) > len(m.data) {
		m.resize(uint64(len(value)))
	}
	copy(m.data[offset:], value[:])
}

func (m *memoryStruct) set32(offset uint64, val *uint256.Int) {
	// length of store may never be less than offset + size.
	// The store should be resized PRIOR to setting the memory
	if offset+32 > uint64(len(m.data)) {
		panic("invalid memory: store empty")
	}
	// Fill in relevant bits
	b32 := val.Bytes32()
	copy(m.data[offset:], b32[:])
}

func (m *memoryStruct) resize(size uint64) {
	if uint64(len(m.data)) <= size {
		m.data = append(m.data, make([]byte, size)...)
	}
}
