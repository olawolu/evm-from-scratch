package main

type storageStruct struct {
	// key  [32]byte
	// data []byte
	store map[[32]byte][]byte
}

func newStorage() *storageStruct {
	return &storageStruct{
		store: make(map[[32]byte][]byte),
	}
}

func (s *storageStruct) set(key [32]byte, data []byte) {
	s.store[key] = data
}

func (s *storageStruct) get(key [32]byte) []byte {
	return s.store[key]
}
