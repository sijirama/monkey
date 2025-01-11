package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

func (ins Instructions) String() string {
	return ""
}

type Opcode byte

const (
	OpConstant Opcode = iota //VM executes OpConstant it gets the constant using the operand as an index

)

// it’s handy being able to lookup how many operands an opcode has and what its human-readable name is
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	/*
		The definition for OpConstant says that its only operand is two bytes wide, which makes it an
		uint16 and limits its maximum value to 65535. If we include 0 the number of representable values
		is then 65536. That should be enough for us, because I don’t think we’re going to reference
		more than 65536 constants in our Monkey programs
	*/
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

/*
INFO: convert an opcode and it's operands to bytecode format
*/
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
