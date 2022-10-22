package main

import "github.com/holiman/uint256"

type instruction struct {
	execute executionFunc
}

type (
	executionFunc  func(ctx *executionContext) []uint256.Int
	InstructionSet [256]*instruction
	ISet           map[OpCode]*instruction
)

func loadInstructionSet() ISet {
	instructionSet := ISet{
		STOP: {
			execute: stopOp,
		},
		ADD: {
			execute: addOp,
		},
		MUL: {
			execute: mulOp,
		},
		SUB: {
			execute: subOp,
		},
		DIV: {
			execute: divOp,
		},
		SDIV: {
			execute: sdivOp,
		},
		MOD: {
			execute: modOp,
		},
		SMOD: {
			execute: smodOp,
		},
		LT: {
			execute: ltOp,
		},
		GT: {
			execute: gtOp,
		},
		SLT: {
			execute: sltOp,
		},
		SGT: {
			execute: sgtOp,
		},
		EQ: {
			execute: eqOp,
		},
		ISZERO: {
			execute: iszeroOp,
		},
		AND: {
			execute: andOp,
		},
		OR: {
			execute: orOp,
		},
		XOR: {
			execute: xorOp,
		},
		NOT: {
			execute: notOp,
		},
		BYTE: {
			execute: byteOp,
		},
		POP: {
			execute: popOp,
		},
		PUSH1: {
			execute: makePush(1),
		},
		PUSH2: {
			execute: makePush(2),
		},
		PUSH3: {
			execute: makePush(3),
		},
		PUSH32: {
			execute: makePush(32),
		},
		DUP1: {
			execute: dup1Op,
		},
		DUP3: {
			execute: dup3Op,
		},
		SWAP1: {
			execute: swap1Op,
		},
		SWAP3: {
			execute: swap3Op,
		},
	}

	return validateInstructionSet(instructionSet)
}

func validateInstructionSet(is ISet) ISet {
	for i := 0; i < 256; i++ {
		if is[OpCode(i)] == nil {
			is[OpCode(i)] = &instruction{
				execute: invalidOp,
			}
		}
	}
	return is
}