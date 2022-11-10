package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/blocktree/openwallet/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
)

func invalidOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	fmt.Println("invalid opcode", ctx.code)
	return ctx.stack.data
}

func addOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.AddOverflow(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func popOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	ctx.stack.pop()
	return ctx.stack.data
}

func push1Op(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var value uint256.Int

	ctx.pc += 1
	codeLen := uint64(len(ctx.code))
	if ctx.pc < codeLen {
		data := ctx.code[ctx.pc : ctx.pc+1]
		value.SetBytes(data)
		ctx.stack.push(value)
	}
	return ctx.stack.data
}

func makePush(n uint64) func(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var value uint256.Int
	return func(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
		data := ctx.code[ctx.pc+1 : ctx.pc+n+1]
		value.SetBytes(data)
		ctx.stack.push(value)
		ctx.pc += n
		return ctx.stack.data
	}
}

func stopOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	ctx.halt = true
	return ctx.stack.data
}

func mulOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.MulOverflow(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func subOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.SubOverflow(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func divOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Div(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func sdivOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.SDiv(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func modOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Mod(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func smodOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.SMod(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func ltOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
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

func sltOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
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

func gtOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
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

func sgtOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
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

func eqOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
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

func iszeroOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
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

func andOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.And(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func orOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Or(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func xorOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result.Xor(&a, &b)
	ctx.stack.push(result)
	return ctx.stack.data
}

func notOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a uint256.Int

	a = ctx.stack.pop()
	result.Not(&a)
	ctx.stack.push(result)
	return ctx.stack.data
}

func byteOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.pop()
	b = ctx.stack.pop()
	result = *b.Byte(&a)
	ctx.stack.push(result)
	ctx.pc++
	return ctx.stack.data
}

func dup1Op(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	result := ctx.stack.peek()
	ctx.stack.push(*result)
	return ctx.stack.data
}

func dup3Op(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	result := ctx.stack.peekN(2)
	ctx.stack.push(result)
	return ctx.stack.data
}

func swap1Op(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	ctx.stack.swap(1)
	return ctx.stack.data
}

func swap3Op(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	ctx.stack.swap(3)
	return ctx.stack.data
}

func jumpOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	ctx.stack.pop()
	ctx.pc += 2
	return ctx.stack.data
}

func jumpiOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	pos, cond := ctx.stack.pop(), ctx.stack.pop()
	if !cond.IsZero() {
		if !validJumpDest(ctx.code, &pos) {
			return ctx.stack.data
		}
		ctx.pc = pos.Uint64()
	}
	return ctx.stack.data
}

func validJumpDest(code []byte, dest *uint256.Int) bool {
	udest, overflow := dest.Uint64WithOverflow()
	if overflow || udest >= uint64(len(code)) {
		return false
	}
	if OpCode(code[udest]) != JUMPDEST {
		return false
	}
	return true
}

func jumpDestOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	return ctx.stack.data
}

func pcOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	ctx.stack.push(*new(uint256.Int).SetUint64(ctx.pc))
	return ctx.stack.data
}

func mstoreOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	_offset, _val := ctx.stack.pop(), ctx.stack.pop()
	offset := _offset.Uint64()

	ctx.memory.set32(offset, &_val)

	return ctx.stack.data
}
func mloadOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result uint256.Int
	offset := ctx.stack.pop()
	content := ctx.memory.get(offset.Uint64(), 32)
	result.SetBytes(content)
	ctx.stack.push(result)

	return ctx.stack.data
}

func mstore8Op(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	offset, val := ctx.stack.pop(), ctx.stack.pop()
	ctx.memory.data[offset.Uint64()] = byte(val.Uint64())
	return ctx.stack.data
}

func msizeOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	ctx.stack.push(*new(uint256.Int).SetUint64(uint64(len(ctx.memory.data))))
	return ctx.stack.data
}

func sha3Op(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result uint256.Int
	offset, size := ctx.stack.pop(), ctx.stack.pop()
	content := ctx.memory.get(offset.Uint64(), size.Uint64())
	result.SetBytes(crypto.Keccak256(content))
	ctx.stack.push(result)
	return ctx.stack.data
}

func addressOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	val, err := uint256.FromHex(ctx.transaction.To)
	if err != nil {
		fmt.Println("Error", err)
	}

	ctx.stack.push(*val)

	return ctx.stack.data
}

func callerOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	val, err := uint256.FromHex(ctx.transaction.From)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*val)
	return ctx.stack.data
}

func balanceOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	address := ctx.stack.pop()
	balance := ctx.GetBalance(address.String())
	ctx.stack.push(*balance)

	fmt.Println("stack", ctx.stack.data)
	return ctx.stack.data
}

func originOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	val, err := uint256.FromHex(ctx.transaction.Origin)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*val)
	return ctx.stack.data
}

func coinbaseOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	val, err := uint256.FromHex(ctx.block.Coinbase)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*val)
	return ctx.stack.data
}

func timestampOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	timeStampUint, err := strconv.ParseUint(ctx.block.Timestamp, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*new(uint256.Int).SetUint64(timeStampUint))
	return ctx.stack.data
}

func numberOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	numberUint, err := strconv.ParseUint(ctx.block.Number, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*new(uint256.Int).SetUint64(numberUint))
	return ctx.stack.data
}

func difficultyOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	difficultyUint, err := uint256.FromHex(ctx.block.Difficulty)
	if err != nil {
		fmt.Println("Error", err)
	}

	ctx.stack.push(*difficultyUint)
	return ctx.stack.data
}

func gaslimitOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	gasLimitUint, err := uint256.FromHex(ctx.block.GasLimit)
	if err != nil {
		fmt.Println("Error", err)
	}

	ctx.stack.push(*gasLimitUint)
	return ctx.stack.data
}

func gaspriceOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	gasPriceUint, err := strconv.ParseUint(ctx.transaction.GasPrice, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*new(uint256.Int).SetUint64(gasPriceUint))
	return ctx.stack.data
}

func chainidOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	chainIdUint, err := strconv.ParseUint(ctx.block.ChainId, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*new(uint256.Int).SetUint64(chainIdUint))
	return ctx.stack.data
}

func callvalueOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	callValueUint, err := strconv.ParseUint(ctx.transaction.Value, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.push(*new(uint256.Int).SetUint64(callValueUint))
	return ctx.stack.data
}

func calldataloadOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	x := ctx.stack.peek()

	if o, overflow := x.Uint64WithOverflow(); !overflow {
		b, err := hex.DecodeString(ctx.transaction.Data)
		if err != nil {
			fmt.Println("Error", err)
		}
		dt := getData(b, o, 32)
		x.SetBytes(dt)
	} else {
		x.Clear()
	}

	return ctx.stack.data
}

func getData(data []byte, offset, size uint64) []byte {
	dataLen := uint64(len(data))
	if offset > dataLen {
		offset = dataLen
	}

	end := offset + size
	if end > dataLen {
		end = dataLen
	}

	return rightPadData(data[offset:end], int(size))

}

func rightPadData(data []byte, size int) []byte {
	if size < len(data) {
		return data
	}
	padded := make([]byte, size)
	copy(padded, data)
	return padded
}

func calldatasizeOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	data, err := hex.DecodeString(ctx.transaction.Data)
	if err != nil {
		fmt.Println("Error", err)
	}
	dataLen := uint64(len(data))
	ctx.stack.push(*new(uint256.Int).SetUint64(dataLen))
	return ctx.stack.data
}

func calldatacopyOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	mOffset := ctx.stack.pop()
	offset := ctx.stack.pop()
	mSize := ctx.stack.pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		b, err := hex.DecodeString(ctx.transaction.Data)
		if err != nil {
			fmt.Println("Error", err)
		}
		ctx.memory.set(mOffset.Uint64(), mSize.Uint64(), getData(b, o, mSize.Uint64()))
	} else {
		mOffset.Clear()
	}

	return ctx.stack.data
}

func codesizeOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	codeSize := uint64(len(ctx.code))
	ctx.stack.push(*new(uint256.Int).SetUint64(codeSize))
	return ctx.stack.data
}

func codecopyOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	mOffset := ctx.stack.pop()
	offset := ctx.stack.pop()
	mSize := ctx.stack.pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		ctx.memory.set(mOffset.Uint64(), mSize.Uint64(), getData(ctx.code, o, mSize.Uint64()))
	} else {
		mOffset.Clear()
	}

	return ctx.stack.data
}

func extcodesizeOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	address := ctx.stack.pop()
	addressHex := address.String()

	codeSize := ctx.GetCodeSize(addressHex)
	ctx.stack.push(*new(uint256.Int).SetUint64(codeSize))
	return ctx.stack.data
}

func extcodecopyOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	address := ctx.stack.pop()
	mOffset := ctx.stack.pop()
	offset := ctx.stack.pop()
	mSize := ctx.stack.pop()

	addressHex := address.String()
	code := ctx.GetCode(addressHex)

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		ctx.memory.set(mOffset.Uint64(), mSize.Uint64(), getData(code, o, mSize.Uint64()))
	} else {
		mOffset.Clear()
	}

	return ctx.stack.data
}

func selfbalanceOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	balance := ctx.GetBalance(ctx.transaction.To)
	ctx.stack.push(*balance)
	return ctx.stack.data
}

func sstoreOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	key := ctx.stack.pop()
	value := ctx.stack.pop()

	keyBytes := key.Bytes32()
	valueBytes := value.Bytes()

	ctx.storage.set(keyBytes, valueBytes)
	// ctx.SetStorage(key.String(), value.String())
	return ctx.stack.data
}

func sloadOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var result uint256.Int
	key := ctx.stack.pop()

	keyBytes := key.Bytes32()
	valueBytes := ctx.storage.get(keyBytes)

	result.SetBytes(valueBytes)
	ctx.stack.push(result)
	return ctx.stack.data
}

func returnOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var returnData []byte
	offset := ctx.stack.pop()
	size := ctx.stack.pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		if s, overflow := size.Uint64WithOverflow(); !overflow {
			returnData = ctx.memory.get(o, s)
		}
	}

	trf := hex.EncodeToString(returnData)

	ctx.returnData = trf

	ctx.done = true
	return ctx.stack.data
}

func revertOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	var returnData []byte
	offset := ctx.stack.pop()
	size := ctx.stack.pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		if s, overflow := size.Uint64WithOverflow(); !overflow {
			returnData = ctx.memory.get(o, s)
		}
	}

	ctx.returnData = hex.EncodeToString(returnData)
	return ctx.stack.data
}

func callOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	_ = ctx.stack.pop()
	to := ctx.stack.pop()
	value := ctx.stack.pop()
	inOffset := ctx.stack.pop()
	inSize := ctx.stack.pop()
	outOffset := ctx.stack.pop()
	outSize := ctx.stack.pop()
	fmt.Println("outsize", outSize)
	toHex := to.String()
	_ = value.String()
	_ = inOffset.String()
	_ = inSize.String()
	_ = outOffset.String()
	_ = outSize.String()

	stateObj, _ := ctx.state.(map[string]interface{})
	addrObj, _ := stateObj[toHex].(map[string]interface{})
	codeObj, _ := addrObj["code"].(map[string]interface{})
	codeStr, _ := codeObj["bin"].(string)

	fromAddress := ctx.transaction.To

	bin, err := hex.DecodeString(codeStr)
	if err != nil {
		panic(err)
	}
	vm := interpreter.vm

	// pause the current context and pass execution to a new subcontext
	_ = ctx
	pc = 0

	newCtx := &executionContext{
		pc:      pc,
		memory:  newMemory(),
		stack:   newStack(),
		storage: newStorage(),
		code:    bin,
		state:   ctx.state,
		transaction: &Transaction{
			From: fromAddress,
		},
	}

	vm.Context = newCtx

	_, _, success := vm.execute(bin)
	switch success {
	case true:
		ctx.stack.push(*uint256.NewInt(1))
	case false:
		ctx.stack.push(*uint256.NewInt(0))
	}

	// resume the parent context
	memSize := len(vm.Context.memory.data)
	ctx.memory.resize(uint64(memSize))
	ctx.memory = vm.Context.memory
	vm.Context = ctx
	return ctx.stack.data
}

func createOp(pc uint64, ctx *executionContext, interpreter *Interpreter) []uint256.Int {
	value := ctx.stack.pop()
	offset := ctx.stack.pop()
	size := ctx.stack.pop()

	senderAddress := ctx.transaction.To

	rlpA, err := rlp.EncodeToBytes([]string{senderAddress, "0x0"})
	if err != nil {
		panic(err)
	}

	bin := ctx.memory.get(offset.Uint64(), size.Uint64())

	addr, err := uint256.FromHex(senderAddress)
	if err != nil {
		panic(err)
	}

	if len(bin) == 0 {
		ctx.state = map[string]interface{}{
			senderAddress: map[string]interface{}{
				"balance": strconv.FormatUint(value.Uint64(), 10),
			},
		}
		ctx.stack.push(*addr)
		return ctx.stack.data
	}

	address := crypto.Keccak256(rlpA)[12:]
	newAddressStr := hex.EncodeToString(address)

	newTx := &Transaction{
		From:  senderAddress,
		To:    newAddressStr,
		Value: value.String(),
	}

	vm := interpreter.vm

	newCtx := &executionContext{
		pc:          0,
		memory:      newMemory(),
		stack:       newStack(),
		storage:     newStorage(),
		code:        bin,
		state:       ctx.state,
		transaction: newTx,
	}

	vm.Context = newCtx

	_, returnValue, _ := vm.execute(bin)
	returnUint256, err := uint256.FromHex(fmt.Sprintf("%v%v", "0x", returnValue))
	if err != nil {
		panic(err)
	}

	ctx.stack.push(*addr)
	ctx.memory.set(offset.Uint64(), size.Uint64(), returnUint256.Bytes())
	ctx.state = map[string]interface{}{
		senderAddress: map[string]interface{}{
			"code": map[string]interface{}{
				"bin": returnValue,
			},
		},
	}
	vm.Context = ctx
	fmt.Println(ctx.stack.data)
	return ctx.stack.data
}
