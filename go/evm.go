package evm

import (
	"math/big"
    "fmt"
)

func Evm(code []byte) ([]*big.Int, bool) {
	var stack []*big.Int
	pc := 0

	for pc < len(code) {
		op := code[pc]
		pc++

        switch op {
        case byte(00):
            fmt.Println("Opcode is 00")
            stack = append(stack, op)
        default:
            fmt.Println("Does not reconize opcode")
        }
	}

	return stack, true
}
