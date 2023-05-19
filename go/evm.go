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
        case byte(0):
            break
        case byte(0x5f):
            opBigInt := new(big.Int)
            opBigInt.SetString("0x5f", 16)
            stack = append(stack, opBigInt)
        default:
            fmt.Println("Does not reconize opcode")
        }
	
        pc++
	}

	return stack, true
}
