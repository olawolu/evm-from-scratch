package main

import (
	"github.com/holiman/uint256"
)

type stackStruct struct {
	n    int64
	data []uint256.Int
}

func newStack() *stackStruct {
	return &stackStruct{
		data: make([]uint256.Int, 0),
	}
}

func (s *stackStruct) push(val uint256.Int) {
	if s.n > 0 {
		s.data = append([]uint256.Int{val}, s.data...)
	} else {
		s.data = append(s.data, val)
	}
	s.n++
}

func (s *stackStruct) pop() uint256.Int {
	val := s.data[0]
	s.data = s.data[1:s.n]
	s.n--
	return val
}

func (s *stackStruct) peek() *uint256.Int {
	return &s.data[0]
}

func (s *stackStruct) peekN(n int64) uint256.Int {
	return s.data[n]
}

func (s *stackStruct) swap(n int64) {
	s.data[0], s.data[n] = s.data[n], s.data[0]
}

func (s *stackStruct) Back(n int64) *uint256.Int {
	return &s.data[n]
}
