package asm

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

// Op is an operation code.
type Op uint32

func (op Op) String() string {
	if int(op) < len(opNames) {
		return opNames[op]
	}
	return "0x" + strconv.FormatUint(uint64(op), 16)
}

const (
	OpLoadVarInt     = Op(0x00)
	OpLoadVarFloat   = Op(0x01)
	OpLoadVarPtr     = Op(0x02)
	OpLoadVarString  = Op(0x03)
	OpPushInt        = Op(0x04)
	OpPushFloat      = Op(0x05)
	OpPushString     = Op(0x06)
	OpIntAdd         = Op(0x07)
	OpFloatAdd       = Op(0x08)
	OpIntSub         = Op(0x09)
	OpFloatSub       = Op(0x0A)
	OpIntMul         = Op(0x0B)
	OpFloatMul       = Op(0x0C)
	OpIntDiv         = Op(0x0D)
	OpFloatDiv       = Op(0x0E)
	OpIntMod         = Op(0x0F)
	OpIntAnd         = Op(0x10)
	OpIntOr          = Op(0x11)
	OpIntXOr         = Op(0x12)
	OpJump           = Op(0x13)
	OpJumpIf         = Op(0x14)
	OpJumpIfNot      = Op(0x15)
	OpStoreInt       = Op(0x16)
	OpStoreFloat     = Op(0x17)
	OpStoreString    = Op(0x18)
	OpStoreIntMul    = Op(0x19)
	OpStoreFloatMul  = Op(0x1A)
	OpStoreIntDiv    = Op(0x1B)
	OpStoreFloatDiv  = Op(0x1C)
	OpStoreIntAdd    = Op(0x1D)
	OpStoreFloatAdd  = Op(0x1E)
	OpStoreStringAdd = Op(0x1F)
	OpStoreIntSub    = Op(0x20)
	OpStoreFloatSub  = Op(0x21)
	OpStoreIntMod    = Op(0x22)
	OpIntEq          = Op(0x23)
	OpFloatEq        = Op(0x24)
	OpStringEq       = Op(0x25)
	OpIntLSh         = Op(0x26)
	OpIntRSh         = Op(0x27)
	OpIntLt          = Op(0x28)
	OpFloatLt        = Op(0x29)
	OpStringLt       = Op(0x2A)
	OpIntGt          = Op(0x2B)
	OpFloatGt        = Op(0x2C)
	OpStringGt       = Op(0x2D)
	OpIntLte         = Op(0x2E)
	OpFloatLte       = Op(0x2F)
	OpStringLte      = Op(0x30)
	OpIntGte         = Op(0x31)
	OpFloatGte       = Op(0x32)
	OpStringGte      = Op(0x33)
	OpIntNeq         = Op(0x34)
	OpFloatNeq       = Op(0x35)
	OpStringNeq      = Op(0x36)
	OpBoolAnd        = Op(0x37)
	OpBoolOr         = Op(0x38)
	OpStoreIntLSh    = Op(0x39)
	OpStoreIntRSh    = Op(0x3A)
	OpStoreIntAnd    = Op(0x3B)
	OpStoreIntOr     = Op(0x3C)
	OpStoreIntXOr    = Op(0x3D)
	OpBoolNot        = Op(0x3E)
	OpIntNot         = Op(0x3F)
	OpIntNeg         = Op(0x40)
	OpFloatNeg       = Op(0x41)
	OpIndexInt       = Op(0x42)
	OpIndexFloat     = Op(0x43)
	OpIndexPtr       = Op(0x44)
	OpCallBuiltin    = Op(0x45)
	OpCallScript     = Op(0x46)
	OpReturn0        = Op(0x47)
	OpReturn         = Op(0x48)
	OpStringAdd      = Op(0x49)
)

var opNames = [74]string{
	OpLoadVarInt:     "LoadVarInt",
	OpLoadVarFloat:   "LoadVarFloat",
	OpLoadVarPtr:     "LoadVarPtr",
	OpLoadVarString:  "LoadVarString",
	OpPushInt:        "PushInt",
	OpPushFloat:      "PushFloat",
	OpPushString:     "PushString",
	OpIntAdd:         "IntAdd",
	OpFloatAdd:       "FloatAdd",
	OpIntSub:         "IntSub",
	OpFloatSub:       "FloatSub",
	OpIntMul:         "IntMul",
	OpFloatMul:       "FloatMul",
	OpIntDiv:         "IntDiv",
	OpFloatDiv:       "FloatDiv",
	OpIntMod:         "IntMod",
	OpIntAnd:         "IntAnd",
	OpIntOr:          "IntOr",
	OpIntXOr:         "IntXOr",
	OpJump:           "Jump",
	OpJumpIf:         "JumpIf",
	OpJumpIfNot:      "JumpIfNot",
	OpStoreInt:       "StoreInt",
	OpStoreFloat:     "StoreFloat",
	OpStoreString:    "StoreString",
	OpStoreIntMul:    "StoreIntMul",
	OpStoreFloatMul:  "StoreFloatMul",
	OpStoreIntDiv:    "StoreIntDiv",
	OpStoreFloatDiv:  "StoreFloatDiv",
	OpStoreIntAdd:    "StoreIntAdd",
	OpStoreFloatAdd:  "StoreFloatAdd",
	OpStoreStringAdd: "StoreStringAdd",
	OpStoreIntSub:    "StoreIntSub",
	OpStoreFloatSub:  "StoreFloatSub",
	OpStoreIntMod:    "StoreIntMod",
	OpIntEq:          "IntEq",
	OpFloatEq:        "FloatEq",
	OpStringEq:       "StringEq",
	OpIntLSh:         "IntLSh",
	OpIntRSh:         "IntRSh",
	OpIntLt:          "IntLt",
	OpFloatLt:        "FloatLt",
	OpStringLt:       "StringLt",
	OpIntGt:          "IntGt",
	OpFloatGt:        "FloatGt",
	OpStringGt:       "StringGt",
	OpIntLte:         "IntLte",
	OpFloatLte:       "FloatLte",
	OpStringLte:      "StringLte",
	OpIntGte:         "IntGte",
	OpFloatGte:       "FloatGte",
	OpStringGte:      "StringGte",
	OpIntNeq:         "IntNeq",
	OpFloatNeq:       "FloatNeq",
	OpStringNeq:      "StringNeq",
	OpBoolAnd:        "BoolAnd",
	OpBoolOr:         "BoolOr",
	OpStoreIntLSh:    "StoreIntLSh",
	OpStoreIntRSh:    "StoreIntRSh",
	OpStoreIntAnd:    "StoreIntAnd",
	OpStoreIntOr:     "StoreIntOr",
	OpStoreIntXOr:    "StoreIntXOr",
	OpBoolNot:        "BoolNot",
	OpIntNot:         "IntNot",
	OpIntNeg:         "IntNeg",
	OpFloatNeg:       "FloatNeg",
	OpIndexInt:       "IndexInt",
	OpIndexFloat:     "IndexFloat",
	OpIndexPtr:       "IndexPtr",
	OpCallBuiltin:    "CallBuiltin",
	OpCallScript:     "CallScript",
	OpReturn0:        "Return0",
	OpReturn:         "Return",
	OpStringAdd:      "StringAdd",
}

// Instr is an instruction for NoxScript assembly.
type Instr interface {
	OpCode() Op
	Len() int
	String() string
	EncodeTo(out []uint32) []uint32
}

type LoadVar struct {
	Op       Op
	IsGlobal int32
	Index    int32
}

func (v LoadVar) OpCode() Op {
	return v.Op
}
func (v LoadVar) Len() int {
	return 3
}
func (v LoadVar) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op), uint32(v.IsGlobal), uint32(v.Index))
}
func (v LoadVar) String() string {
	kind := "local"
	if v.IsGlobal == 1 {
		kind = "global"
	} else if v.IsGlobal != 0 {
		kind = fmt.Sprintf("global%d", v.IsGlobal)
	}
	name, typ := "LOAD", ""
	switch v.Op {
	case OpLoadVarInt:
		// no type = int
	case OpLoadVarFloat:
		typ = "(float)"
	case OpLoadVarString:
		typ = "(string)"
	case OpLoadVarPtr:
		kind = "&" + kind
	default:
		name = v.Op.String()
	}
	return fmt.Sprintf("%s %s_%d %s", name, kind, v.Index, typ)
}

type Push struct {
	Op  Op
	Val int32
}

func (v Push) OpCode() Op {
	return v.Op
}
func (v Push) Len() int {
	return 2
}
func (v Push) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op), uint32(v.Val))
}
func (v Push) String() string {
	name, typ := "PUSH", ""
	val := strconv.FormatUint(uint64(v.Val), 10)
	if v.Op == OpPushFloat {
		val = strconv.FormatFloat(float64(math.Float32frombits(uint32(v.Val))), 'g', -1, 32)
	}
	switch v.Op {
	case OpPushInt:
		// no type = int
	case OpPushFloat:
		typ = "(float)"
		val = strconv.FormatFloat(float64(math.Float32frombits(uint32(v.Val))), 'g', -1, 32)
	case OpPushString:
		typ = "(string)"
	default:
		name = v.Op.String()
	}
	return fmt.Sprintf("%s %s %s", name, val, typ)
}

type UnaryOp struct {
	Op Op
}

func (v UnaryOp) OpCode() Op {
	return v.Op
}
func (v UnaryOp) Len() int {
	return 1
}
func (v UnaryOp) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op))
}
func (v UnaryOp) String() string {
	name, typ := v.Op.String(), ""
	switch v.Op {
	case OpBoolNot:
		name, typ = "NOT", "(bool)"
	case OpIntNot:
		name, typ = "NOT", "(int)"
	case OpIntNeg:
		name, typ = "NEG", "(int)"
	case OpFloatNeg:
		name, typ = "NEG", "(float)"
	}
	return fmt.Sprintf("%s %s", name, typ)
}

type BinaryOp struct {
	Op Op
}

func (v BinaryOp) OpCode() Op {
	return v.Op
}
func (v BinaryOp) Len() int {
	return 1
}
func (v BinaryOp) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op))
}
func (v BinaryOp) String() string {
	name, typ := v.Op.String(), ""
	switch v.Op {
	case OpIntAdd, OpIntSub, OpIntMul, OpIntDiv, OpIntMod,
		OpIntAnd, OpIntOr, OpIntXOr, OpIntLSh, OpIntRSh,
		OpIntEq, OpIntLt, OpIntGt, OpIntLte, OpIntGte, OpIntNeq:
		// no type = int
	case OpFloatAdd, OpFloatSub, OpFloatMul, OpFloatDiv,
		OpFloatEq, OpFloatLt, OpFloatGt, OpFloatLte, OpFloatGte, OpFloatNeq:
		typ = "(float)"
	case OpStringEq, OpStringLt, OpStringGt, OpStringLte, OpStringGte, OpStringNeq, OpStringAdd:
		typ = "(string)"
	case OpBoolAnd, OpBoolOr:
		typ = "(bool)"
	}
	switch v.Op {
	case OpIntAdd, OpFloatAdd, OpStringAdd:
		name = "ADD"
	case OpIntSub, OpFloatSub:
		name = "SUB"
	case OpIntMul, OpFloatMul:
		name = "MUL"
	case OpIntDiv, OpFloatDiv:
		name = "DIV"
	case OpIntMod:
		name = "MOD"
	case OpIntAnd, OpBoolAnd:
		name = "AND"
	case OpIntOr, OpBoolOr:
		name = "OR"
	case OpIntXOr:
		name = "XOR"
	case OpIntLSh:
		name = "LSH"
	case OpIntRSh:
		name = "RSH"
	case OpIntEq, OpFloatEq, OpStringEq:
		name = "EQ"
	case OpIntNeq, OpFloatNeq, OpStringNeq:
		name = "NEQ"
	case OpIntLt, OpFloatLt, OpStringLt:
		name = "LT"
	case OpIntGt, OpFloatGt, OpStringGt:
		name = "GT"
	case OpIntLte, OpFloatLte, OpStringLte:
		name = "LTE"
	case OpIntGte, OpFloatGte, OpStringGte:
		name = "GTE"
	}
	return fmt.Sprintf("%s %s", name, typ)
}

type Jump struct {
	Op  Op
	Off int32
}

func (v Jump) OpCode() Op {
	return v.Op
}
func (v Jump) Len() int {
	return 2
}
func (v Jump) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op), uint32(v.Off))
}
func (v Jump) String() string {
	cond := v.Op.String()
	switch v.Op {
	case OpJump:
		cond = ""
	case OpJumpIf:
		cond = "if"
	case OpJumpIfNot:
		cond = "if not"
	}
	return fmt.Sprintf("JUMP [%d] %s", v.Off, cond)
}

type StoreVar struct {
	Op Op
}

func (v StoreVar) OpCode() Op {
	return v.Op
}
func (v StoreVar) Len() int {
	return 1
}
func (v StoreVar) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op))
}
func (v StoreVar) String() string {
	name, typ := v.Op.String(), ""
	switch v.Op {
	case OpStoreInt, OpStoreIntMul, OpStoreIntDiv, OpStoreIntAdd, OpStoreIntSub, OpStoreIntMod,
		OpStoreIntLSh, OpStoreIntRSh, OpStoreIntAnd, OpStoreIntOr, OpStoreIntXOr:
		// no type = int
	case OpStoreFloat, OpStoreFloatMul, OpStoreFloatDiv, OpStoreFloatAdd, OpStoreFloatSub:
		typ = "(float)"
	case OpStoreString, OpStoreStringAdd:
		typ = "(string)"
	}
	switch v.Op {
	case OpStoreInt, OpStoreFloat, OpStoreString:
		name = ""
	case OpStoreIntMul, OpStoreFloatMul:
		name = " MUL"
	case OpStoreIntDiv, OpStoreFloatDiv:
		name = " DIV"
	case OpStoreIntAdd, OpStoreFloatAdd, OpStoreStringAdd:
		name = " ADD"
	case OpStoreIntSub, OpStoreFloatSub:
		name = " SUB"
	case OpStoreIntMod:
		name = " MOD"
	case OpStoreIntLSh:
		name = " LSH"
	case OpStoreIntRSh:
		name = " RSH"
	case OpStoreIntAnd:
		name = " AND"
	case OpStoreIntOr:
		name = " OR"
	case OpStoreIntXOr:
		name = " XOR"
	}
	return fmt.Sprintf("STORE%s %s", name, typ)
}

type Index struct {
	Op Op
}

func (v Index) OpCode() Op {
	return v.Op
}
func (v Index) Len() int {
	return 1
}
func (v Index) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op))
}
func (v Index) String() string {
	name, typ := "INDEX", ""
	switch v.Op {
	case OpIndexInt:
		// no type = int
	case OpIndexFloat:
		typ = "(float)"
	case OpIndexPtr:
		typ = "(ptr)"
	default:
		name = v.Op.String()
	}
	return fmt.Sprintf("%s %s", name, typ)
}

type CallBuiltin struct {
	Index Builtin
}

func (v CallBuiltin) OpCode() Op {
	return OpCallBuiltin
}
func (v CallBuiltin) Len() int {
	return 2
}
func (v CallBuiltin) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.OpCode()), uint32(v.Index))
}
func (v CallBuiltin) String() string {
	return v.OpCode().String() + " " + strconv.FormatUint(uint64(v.Index), 10)
}

type CallScript struct {
	Index int32
}

func (v CallScript) OpCode() Op {
	return OpCallScript
}
func (v CallScript) Len() int {
	return 2
}
func (v CallScript) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.OpCode()), uint32(v.Index))
}
func (v CallScript) String() string {
	return v.OpCode().String() + " " + strconv.FormatUint(uint64(v.Index), 10)
}

type Return struct {
	Op Op
}

func (v Return) OpCode() Op {
	return v.Op
}
func (v Return) Len() int {
	return 1
}
func (v Return) EncodeTo(out []uint32) []uint32 {
	return append(out, uint32(v.Op))
}
func (v Return) String() string {
	switch v.Op {
	case OpReturn:
		return "RETURN"
	}
	return v.Op.String()
}

// DecodeNext decodes one instruction from a binary form and returns it's length.
// It returns nil, 1 if instruction was not recognized.
func DecodeNext(code []uint32) (Instr, int) {
	if len(code) == 0 {
		return nil, 0
	}
	op := Op(code[0])
	code = code[1:]
	n := 1
	switch op {
	case OpLoadVarInt, OpLoadVarFloat, OpLoadVarString, OpLoadVarPtr:
		v := LoadVar{Op: op}
		if len(code) >= 1 {
			v.IsGlobal = int32(code[0])
			n++
		}
		if len(code) >= 2 {
			v.Index = int32(code[1])
			n++
		}
		return v, n
	case OpPushInt, OpPushFloat, OpPushString:
		v := Push{Op: op}
		if len(code) >= 1 {
			v.Val = int32(code[0])
			n++
		}
		return v, n
	case OpIntAdd, OpIntSub, OpIntMul, OpIntDiv, OpIntMod,
		OpIntAnd, OpIntOr, OpIntXOr, OpIntLSh, OpIntRSh,
		OpFloatAdd, OpFloatSub, OpFloatMul, OpFloatDiv,
		OpIntEq, OpIntLt, OpIntGt, OpIntLte, OpIntGte, OpIntNeq,
		OpFloatEq, OpFloatLt, OpFloatGt, OpFloatLte, OpFloatGte, OpFloatNeq,
		OpStringEq, OpStringLt, OpStringGt, OpStringLte, OpStringGte, OpStringNeq,
		OpBoolAnd, OpBoolOr, OpStringAdd:
		return BinaryOp{Op: op}, n
	case OpBoolNot, OpIntNot, OpIntNeg, OpFloatNeg:
		return UnaryOp{Op: op}, n
	case OpIndexInt, OpIndexFloat, OpIndexPtr:
		return Index{Op: op}, n
	case OpCallBuiltin:
		v := CallBuiltin{}
		if len(code) >= 1 {
			v.Index = Builtin(code[0])
			n++
		}
		return v, n
	case OpCallScript:
		v := CallScript{}
		if len(code) >= 1 {
			v.Index = int32(code[0])
			n++
		}
		return v, n
	case OpJump, OpJumpIf, OpJumpIfNot:
		v := Jump{Op: op}
		if len(code) >= 1 {
			v.Off = int32(code[0])
			n++
		}
		return v, n
	case OpStoreInt, OpStoreFloat, OpStoreString,
		OpStoreIntMul, OpStoreIntDiv, OpStoreIntAdd, OpStoreIntSub, OpStoreIntMod,
		OpStoreIntLSh, OpStoreIntRSh, OpStoreIntAnd, OpStoreIntOr, OpStoreIntXOr,
		OpStoreFloatMul, OpStoreFloatDiv, OpStoreFloatAdd, OpStoreFloatSub,
		OpStoreStringAdd:
		return StoreVar{Op: op}, n
	case OpReturn0, OpReturn:
		return Return{Op: op}, n
	}
	return nil, 1
}

// Decode instructions from a binary form.
func Decode(code []uint32) ([]Instr, error) {
	var out []Instr
	for len(code) > 0 {
		v, n := DecodeNext(code)
		if v == nil {
			return out, fmt.Errorf("cannot decode opcode 0x%x", code[0])
		}
		out = append(out, v)
		if len(code) < v.Len() {
			return out, fmt.Errorf("cannot fully decode opcode 0x%x", code[0])
		}
		code = code[n:]
	}
	return out, nil
}

// Encode instructions into a binary form.
func Encode(code []Instr) []uint32 {
	var out []uint32
	for _, v := range code {
		out = v.EncodeTo(out)
	}
	return out
}

// Print assembly instruction in a human-friendly way.
func Print(w io.Writer, code []Instr) error {
	line := 0
	for _, v := range code {
		if _, err := fmt.Fprintf(w, "%5d:  %s\n", line, v); err != nil {
			return err
		}
		line += v.Len()
	}
	return nil
}
