package main

import (
	"fmt"

	"github.com/holiman/uint256"
)

type VM struct {
	Context     *executionContext
	Interpreter *Interpreter
}

func NewVM(ctx *executionContext) *VM {
	vm := &VM{
		Context: ctx,
	}
	vm.Interpreter = NewInterpreter()
	return vm
}

func (vm *VM) execute(code []byte) []uint256.Int {
	var stack []uint256.Int
	in := len(code)

	fmt.Println("code length: ", in)
	for  !vm.Context.halt && vm.Context.pc < in {
		fmt.Println("pc: ", vm.Context.pc)

		opCode := decodeOp(vm.Context)

		// get the instruction by the opcode
		op := vm.Interpreter.instructionSet[OpCode(opCode)]

		// execute the instruction
		stack = op.execute(vm.Context)
	}

	return stack
}

func decodeOp(ctx *executionContext) byte {
	// decode the code
	code, _ := ctx.readCode(1)

	return code[0]
}
