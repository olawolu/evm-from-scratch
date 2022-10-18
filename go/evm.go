package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
)

type code struct {
	Bin string
	Asm string
}

type expect struct {
	Stack   []string
	Success bool
	Return  string
}

type TestCase struct {
	Name   string
	Code   code
	Expect expect
}

func evm(code []byte) []big.Int {
	var stack []big.Int
	stack = execute(code)
	return stack
}

func execute(code []byte) []big.Int {
	var stack []big.Int
	var op, val byte
	fmt.Println("code: ", code)

	// iterate through opcodes
	for i := 0; i < len(code); i += 2 {
		if len(code) == 1 {
			op, _ = interpret(code[i:])
		} else {
			op, val = interpret(code[i : i+2])
		}

		switch op {
		case 0x00:
			// STOP
			return stack
		case 0x60:
			stack = push(stack, *new(big.Int).SetBytes([]byte{val}))
		case 0x50:
			// POP
			stack = pop(stack)
		default:
			fmt.Println("unimplemented opcode: ", op)
			return stack
		}
	}
	return stack
}

func interpret(code []byte) (byte, byte) {
	if len(code) != 1 {
		return code[0], code[1]
	}
	return code[0], 0
}

func push(stack []big.Int, value big.Int) []big.Int {
	if len(stack) > 0 {
		stack = append([]big.Int{value}, stack...)
	} else {
		stack = append(stack, value)
	}
	return stack
}

func pop(stack []big.Int) []big.Int {
	return stack[1:]
}

func main() {
	content, err := ioutil.ReadFile("../evm.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []TestCase
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json.Unmarshal(): ", err)
	}

	for index, test := range payload {
		fmt.Printf("Test #%v of %v: %v\n", index+1, len(payload), test.Name)

		bin, err := hex.DecodeString(test.Code.Bin)
		if err != nil {
			log.Fatal("Error during hex.DecodeString(): ", err)
		}

		var expectedStack []big.Int
		for _, s := range test.Expect.Stack {
			i, ok := new(big.Int).SetString(s, 0)
			if !ok {
				log.Fatal("Error during big.Int.SetString(): ", err)
			}
			expectedStack = append(expectedStack, *i)
		}

		// Note: as the test cases get more complex, you'll need to modify this
		// to pass down more arguments to the evm function and return more than
		// just the stack.
		stack := evm(bin)

		match := len(stack) == len(expectedStack)
		if match {
			for i, s := range stack {
				match = match && (s.Cmp(&expectedStack[i]) == 0)
			}
		}

		if !match {
			fmt.Printf("Instructions: \n%v\n", test.Code.Asm)
			fmt.Printf("Expected: %v\n", toStrings(expectedStack))
			fmt.Printf("Got: %v\n\n", toStrings(stack))
			fmt.Printf("Progress: %v/%v\n\n", index, len(payload))
			log.Fatal("Stack mismatch")
		}
	}
}

func toStrings(stack []big.Int) []string {
	var strings []string
	for _, s := range stack {
		strings = append(strings, s.String())
	}
	return strings
}
