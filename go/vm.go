package main

import (
	"fmt"
	"math"
	"math/bits"

	"github.com/holiman/uint256"
)

var (
	pc uint64
)

type VM struct {
	Context        *executionContext
	EVMInterpreter *Interpreter
}

func NewVM(ctx *executionContext) *VM {
	vm := &VM{
		Context: ctx,
	}
	vm.EVMInterpreter = NewInterpreter(vm)
	return vm
}

// returns the current intererpreter
func (vm *VM) Interpreter() *Interpreter {
	return vm.EVMInterpreter
}

func (vm *VM) execute(code []byte) ([]uint256.Int, string, bool) {
	var stack []uint256.Int
	in := len(code)

	for !vm.Context.halt && vm.Context.pc < uint64(in) {
		opCode, n := decodeOp(vm.Context)
		op := vm.EVMInterpreter.instructionSet[OpCode(opCode)]

		var memorySize uint64

		if op.memorySize != nil {
			memSize, overflow := op.memorySize(vm.Context.stack)
			if overflow {
				fmt.Println("memory size overflow")
				return stack, "", false
			}

			if memorySize, overflow = SafeMul(toWordSize(memSize), 32); overflow {
				fmt.Println("memory size overflow")
				return stack, "", false
			}

			if memorySize > 0 {
				vm.Context.memory.resize(memorySize)
			}
		}

		// execute the instruction
		stack = op.execute(vm.Context.pc, vm.Context, vm.EVMInterpreter)
		vm.Context.pc += n
	}

	if vm.Context.pc >= uint64(in) && vm.Context.halt {
		vm.Context.done = true
	}

	return stack, vm.Context.returnData, vm.Context.done
}

func decodeOp(ctx *executionContext) (byte, uint64) {
	// decode the code
	code, n := ctx.readCode(1)

	return code[0], n
}

// SafeMul returns x*y and checks for overflow.
func SafeMul(x, y uint64) (uint64, bool) {
	hi, lo := bits.Mul64(x, y)
	return lo, hi != 0
}

func toWordSize(size uint64) uint64 {
	if size > math.MaxUint64-31 {
		return math.MaxUint64/32 + 1
	}

	return (size + 31) / 32
}
