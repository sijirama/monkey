package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}
		operands, read := ReadOperands(def, ins[i+1:])
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)
	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}
	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])

	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

type Opcode byte

const (
	OpConstant Opcode = iota //VM executes OpConstant it gets the constant using the operand as an index, that's why the number of operand for this is 1, it's just the constant index from the constant pool
	OpAdd
	OpPop
	OpSub
	OpMul
	OpDiv
	OpTrue
	OpFalse
	OpEqual
	OpNotEqual
	OpGreaterThan
	OpMinus
	OpBang
	OpJumpNotTruthy
	OpJump
	OpNull
	OpGetGlobal
	OpSetGlobal
	OpArray
	OpHash
	OpIndex
)

var opcodeNames = map[Opcode]string{
	OpConstant:      "OpConstant",
	OpAdd:           "OpAdd",
	OpPop:           "OpPop",
	OpSub:           "OpSub",
	OpMul:           "OpMul",
	OpDiv:           "OpDiv",
	OpTrue:          "OpTrue",
	OpFalse:         "OpFalse",
	OpEqual:         "OpEqual",
	OpNotEqual:      "OpNotEqual",
	OpGreaterThan:   "OpGreaterThan",
	OpMinus:         "OpMinus",
	OpBang:          "OpBang",
	OpJumpNotTruthy: "OpJumpNotTruthy",
	OpJump:          "OpJump",
	OpNull:          "OpNull",
	OpGetGlobal:     "OpGetGlobal",
	OpSetGlobal:     "OpSetGlobal",
	OpArray:         "OpArray",
	OpHash:          "OpHash",
	OpIndex:         "OpIndex",
}

func GetOpcodeName(op Opcode) string {
	name := opcodeNames[op]
	if name == "" {
		return fmt.Sprintf("Unknown opcode %d", op)
	}
	return name
}

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
	OpConstant:      {"OpConstant", []int{2}},
	OpAdd:           {"OpAdd", []int{}},
	OpPop:           {"OpPop", []int{}},
	OpSub:           {"OpSub", []int{}},
	OpMul:           {"OpMul", []int{}},
	OpDiv:           {"OpDiv", []int{}},
	OpTrue:          {"OpTrue", []int{}},
	OpFalse:         {"OpFalse", []int{}},
	OpEqual:         {"OpEqual", []int{}},
	OpNotEqual:      {"OpNotEqual", []int{}},
	OpGreaterThan:   {"OpGreaterThan", []int{}},
	OpMinus:         {"OpMinus", []int{}},
	OpBang:          {"OpBang", []int{}},
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}},
	OpJump:          {"OpJump", []int{2}},
	OpNull:          {"OpNull", []int{}},
	OpGetGlobal:     {"OpGetGlobal", []int{2}},
	OpSetGlobal:     {"OpSetGlobal", []int{2}},
	OpArray:         {"OpArray", []int{2}},
	OpHash:          {"OpHash", []int{2}},
	OpIndex:         {"OpIndex", []int{}},
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

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0
	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}
		offset += width
	}
	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
