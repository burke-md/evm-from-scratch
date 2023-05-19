package evm

import (
	"math/big"
    "fmt"
)

func Evm(code []byte) ([]*big.Int, bool) {
    fmt.Println("Code:", code)

	var stack []*big.Int
	pc := 0

	for pc < len(code) {
		op := code[pc]

        switch op {
        case byte(0)://STOP
            break

        case byte(0x5f)://PUSH0
            opBigInt := new(big.Int)
            opBigInt.SetString("0x00", 16)
            stack = append(stack, opBigInt)
        
            pc++

        case byte(0x60)://PUSH1
            stack = append(stack, parseBigInt(code, pc, 1))
            pc++

        case byte(0x61)://PUSH2
            stack = append(stack, parseBigInt(code, pc, 2))
            pc += 3

        case byte(0x63)://PUSH4
            stack = append(stack, parseBigInt(code, pc, 4))
            pc += 4

        default:
            fmt.Println("Does not reconize opcode")
        }
	
        pc++
	}

	return stack, true
}

func parseBigInt(code []byte, pc int, size int) *big.Int {
    opBigInt := new(big.Int)
    opBigInt.SetBytes(code[pc+1:pc+1+size])
    return opBigInt
}
