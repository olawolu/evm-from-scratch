package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/holiman/uint256"
)

type Interpreter struct {
	vm             *VM
	instructionSet ISet
}

func NewInterpreter(vm *VM) *Interpreter {
	instructionSet := loadInstructionSet()
	return &Interpreter{
		vm:             vm,
		instructionSet: instructionSet,
	}
}


type executionContext struct {
	pc          uint64
	halt        bool
	done        bool
	code        []byte
	block       *Block
	state       interface{}
	stack       *stackStruct
	memory      *memoryStruct
	storage     *storageStruct
	returnData  string
	transaction *Transaction
}

type contract struct {
	CallerAddress string
	OriginAddress string
	CallValue     *uint256.Int
	GasLimit      uint64
	GasPrice      *uint256.Int
	Code          []byte
}

// read n number of bytes from the code
func (ctx *executionContext) readCode(n uint64) ([]byte, uint64) {
	pc := ctx.pc
	value := ctx.code[pc : pc+n]
	return value, n
}

func (ctx *executionContext) GetBalance(address string) *uint256.Int {
	if ctx.state == nil {
		return uint256.NewInt(0)
	}
	obj, _ := ctx.state.(map[string]interface{})

	balanceI, _ := obj[address].(map[string]interface{})

	balance := balanceI["balance"].(string)

	uBalance, err := strconv.ParseUint(balance, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}

	uBalance256 := new(uint256.Int).SetUint64(uBalance)

	fmt.Println(balance)
	return uBalance256
}

func (ctx *executionContext) GetCodeSize(address string) uint64 {
	if ctx.state == nil {
		return 0
	}
	obj, _ := ctx.state.(map[string]interface{})

	codeI, _ := obj[address].(map[string]interface{})

	code := codeI["code"].(map[string]interface{})

	bin := code["bin"].(string)

	binBytes, err := hex.DecodeString(bin)
	if err != nil {
		fmt.Println("Error", err)
	}

	return uint64(len(binBytes))
}

func (ctx *executionContext) GetCode(address string) []byte {
	if ctx.state == nil {
		return []byte{}
	}

	obj, _ := ctx.state.(map[string]interface{})

	codeI, _ := obj[address].(map[string]interface{})

	code := codeI["code"].(map[string]interface{})

	bin := code["bin"].(string)

	binBytes, err := hex.DecodeString(bin)
	if err != nil {
		fmt.Println("Error", err)
	}

	return binBytes
}
