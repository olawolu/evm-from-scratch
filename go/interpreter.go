package main

import (
	"fmt"
)

type Interpreter struct {
	instructionSet ISet
}

func NewInterpreter() *Interpreter {
	instructionSet := loadInstructionSet()
	return &Interpreter{
		instructionSet: instructionSet,
	}
}

type executionContext struct {
	pc    int
	stack *stackStruct
	code  []byte
	halt  bool
}

// read n number of bytes from the code
func (ctx *executionContext) readCode(n int) ([]byte, int) {
	value := ctx.code[ctx.pc : ctx.pc+n]

	fmt.Println("code: ", value)
	// fmt.Println("pc: ", ctx.pc)
	ctx.pc += n
	return value, n
}
