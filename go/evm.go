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
        fmt.Println("op:", op)

        switch op {
        case byte(0):
            break
        default:
            fmt.Println("Does not reconize opcode")
        }
	
        pc++
	}

	return stack, true
}
            /*
            fmt.Println("Opcode is 0", op)
            
            opBigInt := new(big.Int)
            // Must convert byte to string
            opBigInt.SetString("0x00", 16)

            fmt.Println("opBigInt:", opBigInt)

            stack = append(stack, opBigInt)
            */
