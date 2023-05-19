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

        case byte(0x64)://PUSH5
            stack = append(stack, parseBigInt(code, pc, 5))
            pc += 5

        case byte(0x65)://PUSH6
            stack = append(stack, parseBigInt(code, pc, 6))
            pc += 6

        case byte(0x66)://PUSH7
            stack = append(stack, parseBigInt(code, pc, 7))
            pc += 6

        case byte(0x67)://PUSH8
            stack = append(stack, parseBigInt(code, pc, 8))
            pc += 8

        case byte(0x68)://PUSH9
            stack = append(stack, parseBigInt(code, pc, 9))
            pc += 9

        case byte(0x69)://PUSH10
            stack = append(stack, parseBigInt(code, pc, 10))
            pc += 10

        case byte(0x6A)://PUSH11
            stack = append(stack, parseBigInt(code, pc, 11))
            pc += 11


        case byte(0x7F)://PUSH32
            fmt.Println("hit 7F")
            stack = append(stack, parseBigInt(code, pc, 32))
            pc += 32

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
