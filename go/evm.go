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
        /*
        case byte(0x60)://PUSH1
            opBigInt := new(big.Int)
            opBigInt.SetString("0x1", 16)
            stack = append(stack, opBigInt)
        */
        default:
            fmt.Println("Does not reconize opcode")
        }
	
        pc++
	}

	return stack, true
}
