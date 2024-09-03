package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota // Push a constant by its index in the constant pool
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {
		Name:          "OpConstant",
		OperandWidths: []int{2},
	},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]

	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]

	if !ok {
		return []byte{}
	}

	instructionLen := 1 // Opcode size in bytes

	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1 // Skip opcode

	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2: // 2 bytes -> 16 bits
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}

		offset += width
	}

	return instruction
}