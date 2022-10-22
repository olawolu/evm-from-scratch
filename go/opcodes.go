package main

// evm opcode
type OpCode byte

// 0x00 range - arithmetic ops
const (
	STOP       OpCode = 0x00
	ADD        OpCode = 0x01
	MUL        OpCode = 0x02
	SUB        OpCode = 0x03
	DIV        OpCode = 0x04
	SDIV       OpCode = 0x05
	MOD        OpCode = 0x06
	SMOD       OpCode = 0x07
	ADDMOD     OpCode = 0x08
	MULMOD     OpCode = 0x09
	EXP        OpCode = 0x0a
	SIGNEXTEND OpCode = 0x0b
)

// 0x10 range - comparison ops
const (
	LT     OpCode = 0x10
	GT     OpCode = 0x11
	SLT    OpCode = 0x12
	SGT    OpCode = 0x13
	EQ     OpCode = 0x14
	ISZERO OpCode = 0x15
	AND    OpCode = 0x16
	OR     OpCode = 0x17
	XOR    OpCode = 0x18
	NOT    OpCode = 0x19
	BYTE   OpCode = 0x1A
)

// 0x20 range - crypto
const (
	SHA3 OpCode = 0x20
)

// 0x30 range - closure state
const (
	ADDRESS        OpCode = 0x30
	BALANCE        OpCode = 0x31
	ORIGIN         OpCode = 0x32
	CALLER         OpCode = 0x33
	CALLVALUE      OpCode = 0x34
	CALLDATALOAD   OpCode = 0x35
	CALLDATASIZE   OpCode = 0x36
	CALLDATACOPY   OpCode = 0x37
	CODESIZE       OpCode = 0x38
	CODECOPY       OpCode = 0x39
	GASPRICE       OpCode = 0x3a
	EXTCODESIZE    OpCode = 0x3b
	EXTCODECOPY    OpCode = 0x3c
	RETURNDATASIZE OpCode = 0x3d
	RETURNDATACOPY OpCode = 0x3e
)

// 0x40 range - block operations
const (
	BLOCKHASH   OpCode = 0x40
	COINBASE    OpCode = 0x41
	TIMESTAMP   OpCode = 0x42
	NUMBER      OpCode = 0x43
	DIFFICULTY  OpCode = 0x44
	GASLIMIT    OpCode = 0x45
	CHAINID     OpCode = 0x46
	SELFBALANCE OpCode = 0x47
)

// 0x50 range - 'storage' and execution
const (
	POP      OpCode = 0x50
	MLOAD    OpCode = 0x51
	MSTORE   OpCode = 0x52
	MSTORE8  OpCode = 0x53
	SLOAD    OpCode = 0x54
	SSTORE   OpCode = 0x55
	JUMP     OpCode = 0x56
	JUMPI    OpCode = 0x57
	PC       OpCode = 0x58
	MSIZE    OpCode = 0x59
	GAS      OpCode = 0x5a
	JUMPDEST OpCode = 0x5b
)

// 0x60 range - stack ops
const (
	PUSH1  OpCode = 0x60
	PUSH2  OpCode = 0x61
	PUSH3  OpCode = 0x62
	PUSH4  OpCode = 0x63
	PUSH5  OpCode = 0x64
	PUSH6  OpCode = 0x65
	PUSH7  OpCode = 0x66
	PUSH8  OpCode = 0x67
	PUSH9  OpCode = 0x68
	PUSH10 OpCode = 0x69
	PUSH11 OpCode = 0x6a
	PUSH12 OpCode = 0x6b
	PUSH13 OpCode = 0x6c
	PUSH14 OpCode = 0x6d
	PUSH15 OpCode = 0x6e
	PUSH16 OpCode = 0x6f
	PUSH17 OpCode = 0x70
	PUSH18 OpCode = 0x71
	PUSH19 OpCode = 0x72
	PUSH20 OpCode = 0x73
	PUSH21 OpCode = 0x74
	PUSH22 OpCode = 0x75
	PUSH23 OpCode = 0x76
	PUSH24 OpCode = 0x77
	PUSH25 OpCode = 0x78
	PUSH26 OpCode = 0x79
	PUSH27 OpCode = 0x7a
	PUSH28 OpCode = 0x7b
	PUSH29 OpCode = 0x7c
	PUSH30 OpCode = 0x7d
	PUSH31 OpCode = 0x7e
	PUSH32 OpCode = 0x7F
)

// 0x80 range - storage ops
const (
	DUP1 OpCode = 0x80
	DUP2 OpCode = 0x81
	DUP3 OpCode = 0x82
	DUP4 OpCode = 0x83
	DUP5 OpCode = 0x84
	DUP6 OpCode = 0x85
	DUP7 OpCode = 0x86
	DUP8 OpCode = 0x87
	DUP9 OpCode = 0x88
	DUP10 OpCode = 0x89
	DUP11 OpCode = 0x8a
	DUP12 OpCode = 0x8b
	DUP13 OpCode = 0x8c
	DUP14 OpCode = 0x8d
	DUP15 OpCode = 0x8e
	DUP16 OpCode = 0x8f
)

// 0x90 range - swapping ops
const (
	SWAP1 OpCode = 0x90
	SWAP2 OpCode = 0x91
	SWAP3 OpCode = 0x92
	SWAP4 OpCode = 0x93
	SWAP5 OpCode = 0x94
	SWAP6 OpCode = 0x95
	SWAP7 OpCode = 0x96
	SWAP8 OpCode = 0x97
	SWAP9 OpCode = 0x98
	SWAP10 OpCode = 0x99
	SWAP11 OpCode = 0x9a
	SWAP12 OpCode = 0x9b
	SWAP13 OpCode = 0x9c
	SWAP14 OpCode = 0x9d
	SWAP15 OpCode = 0x9e
	SWAP16 OpCode = 0x9f
)

// 0xa0 range - logging ops
const (
	LOG0 OpCode = 0xa0
	LOG1 OpCode = 0xa1
	LOG2 OpCode = 0xa2
	LOG3 OpCode = 0xa3
	LOG4 OpCode = 0xa4
)

// 0xf0 range - closures
const (
	CREATE OpCode = 0xf0
	CALL   OpCode = 0xf1
	CALLCODE OpCode = 0xf2
	RETURN OpCode = 0xf3
	DELEGATECALL OpCode = 0xf4
	CREATE2 OpCode = 0xf5
	STATICCALL OpCode = 0xfa
	REVERT OpCode = 0xfd
	INVALID OpCode = 0xfe
	SELFDESTRUCT OpCode = 0xff
)
