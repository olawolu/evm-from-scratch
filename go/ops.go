package main

import (
	"fmt"

	"github.com/holiman/uint256"
)

func invalidOp(ctx *executionContext) []uint256.Int {
	fmt.Println("invalid opcode")
	return ctx.stack.data
}

func addOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.AddOverflow(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func popOp(ctx *executionContext) []uint256.Int {
	ctx.stack.pop()
	return ctx.stack.data
}

func makePush(n int) func(ctx *executionContext) []uint256.Int {
	return func(ctx *executionContext) []uint256.Int {
		var value uint256.Int
		data, _ := ctx.readCode(n)
		value.SetBytes(data)
		ctx.stack.push(value)
		return ctx.stack.data
	}
}

func stopOp(ctx *executionContext) []uint256.Int {
	ctx.halt = true
	return ctx.stack.data
}

func mulOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.MulOverflow(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func subOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.SubOverflow(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func divOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Div(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func sdivOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.SDiv(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func modOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Mod(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func smodOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.SMod(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func ltOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()

	if a.Lt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.push(result)

	return ctx.stack.data
}

func sltOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()

	if a.Slt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.push(result)
	return ctx.stack.data
}

func gtOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	if a.Gt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.push(result)
	return ctx.stack.data
}

func sgtOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()

	if a.Sgt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.push(result)
	return ctx.stack.data
}

func eqOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()

	if a.Eq(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.push(result)

	return ctx.stack.data
}

func iszeroOp(ctx *executionContext) []uint256.Int {
	var result, a uint256.Int

	a = ctx.stack.pop()
	if a.IsZero() {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.push(result)
	return ctx.stack.data
}

func andOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.And(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func orOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Or(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func xorOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Xor(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func notOp(ctx *executionContext) []uint256.Int {
	var result, a uint256.Int

	a = ctx.stack.pop()
	result.Not(&a)
	ctx.stack.push(result)
	return ctx.stack.data
}

func byteOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result = *b.Byte(&a)
	ctx.stack.push(result)

	return ctx.stack.data
}

func dup1Op(ctx *executionContext) []uint256.Int {
	result := ctx.stack.peek()
	ctx.stack.push(result)
	return ctx.stack.data
}

func dup3Op(ctx *executionContext) []uint256.Int {
	result := ctx.stack.peekN(2)
	ctx.stack.push(result)
	return ctx.stack.data
}

func swap1Op(ctx *executionContext) []uint256.Int {
	ctx.stack.swap(1)
	return ctx.stack.data
}

func swap3Op(ctx *executionContext) []uint256.Int {
	ctx.stack.swap(3)
	return ctx.stack.data
}

func jumpOp(ctx *executionContext) []uint256.Int {
	return ctx.stack.data
}