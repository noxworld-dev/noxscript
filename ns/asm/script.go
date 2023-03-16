package asm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"io"
	"strconv"
	"strings"
)

// NCobjName is a name of a temporary file with script section, created by Nox when decoding the map.
const NCobjName = "nc.obj"

// Script is a decoded NoxScript.
type Script struct {
	// Strings table. Used by all string-related instructions.
	Strings []string
	// Funcs is a list of function definitions.
	//
	// First two functions are usually called GLOBAL and only define and initialize global variables.
	Funcs []FuncDef
}

// VarDef is a variable definition.
type VarDef struct {
	// Size of a variable. It is > 1 for array variables.
	Size int
	// Offs is an offset for a given variable.
	Offs int
}

// FuncDef is a NoxScript function definition.
type FuncDef struct {
	// Name of the function.
	Name string
	// Args is the number of arguments.
	Args int
	// Return is the number of returns. Only 1 return max is supported.
	Return int
	// Unused stores some unused values in script encoding.
	// It will be written back during encoding.
	Unused int
	// Vars is a list of variable definitions (args followed by locals).
	Vars []VarDef
	// VarsSz is a total size of memory required for all variables.
	VarsSz int
	// Code of the function. See Disasm.
	Code []uint32
	// Rest is a trailing bytes after a Code section.
	// It will be written back during encoding.
	Rest []byte
	// NamePref is appended to all object name lookups during script runtime.
	//
	// TODO: rename to NameSuff
	NamePref string
	// PosOff is added to all position-related operations during script runtime.
	PosOff image.Point
}

// Disasm decodes assembly instruction from the function binary code.
func (f *FuncDef) Disasm() ([]Instr, error) {
	return Decode(f.Code)
}

// ReadScript reads the NoxScript binary section.
func ReadScript(rd io.Reader) (*Script, error) {
	readInt := func() (int, error) {
		var b [4]byte
		_, err := io.ReadFull(rd, b[:])
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		} else if err != nil {
			return 0, err
		}
		v := int32(binary.LittleEndian.Uint32(b[:]))
		return int(v), nil
	}
	readString := func(sz int) (string, error) {
		buf := make([]byte, sz)
		_, err := io.ReadFull(rd, buf)
		if err == io.EOF {
			return "", io.ErrUnexpectedEOF
		} else if err != nil {
			return "", err
		}
		return string(buf), nil
	}
	readStringN := func() (string, error) {
		n, err := readInt()
		if err != nil {
			return "", err
		}
		return readString(n)
	}
	expect := func(exp string) error {
		got, err := readString(len(exp))
		if err != nil {
			return err
		}
		if exp != got {
			return fmt.Errorf("unexpected token: %q", got)
		}
		return nil
	}

	if err := expect("SCRIPT03"); err != nil {
		return nil, err
	}
	if err := expect("STRG"); err != nil {
		return nil, err
	}
	sc := new(Script)
	cnt, err := readInt()
	if err != nil {
		return nil, err
	} else if cnt < 0 {
		return nil, errors.New("negative strings count")
	}
	sc.Strings = make([]string, 0, cnt)
	for i := 0; i < cnt; i++ {
		s, err := readStringN()
		if err != nil {
			return nil, err
		}
		sc.Strings = append(sc.Strings, s)
	}
	if err := expect("CODE"); err != nil {
		return sc, err
	}
	fcnt, err := readInt()
	if err != nil {
		return sc, err
	} else if fcnt < 0 {
		return sc, errors.New("negative func count")
	}
	sc.Funcs = make([]FuncDef, 0, fcnt)
	for i := 0; i < fcnt; i++ {
		if err := expect("FUNC"); err != nil {
			return sc, err
		}

		name, err := readStringN()
		if err != nil {
			return sc, err
		}
		def := FuncDef{Name: name}
		if sub := strings.SplitN(name, "%", 4); len(sub) > 1 {
			def.NamePref = "%" + sub[1]
			if len(sub) > 2 {
				def.PosOff.X, _ = strconv.Atoi(sub[2])
			}
			if len(sub) > 3 {
				def.PosOff.X, _ = strconv.Atoi(sub[3])
			}
		}
		def.Return, err = readInt()
		if err != nil {
			return sc, err
		}
		def.Args, err = readInt()
		if err != nil {
			return sc, err
		}

		if err := expect("SYMB"); err != nil {
			return sc, err
		}
		locals, err := readInt()
		if err != nil {
			return sc, err
		}
		localsCap := locals

		isGlobal := i == 0
		if isGlobal {
			localsCap++
		}

		def.Unused, err = readInt()
		if err != nil {
			return sc, err
		}
		def.Vars = make([]VarDef, 0, localsCap)
		if isGlobal {
			def.Vars = append(def.Vars, VarDef{})
		}
		sum := 0
		for j := 0; j < locals; j++ {
			sz, err := readInt()
			if err != nil {
				return sc, err
			}
			def.Vars = append(def.Vars, VarDef{Size: sz, Offs: sum})
			sum += sz
		}
		def.VarsSz = sum
		if err := expect("DATA"); err != nil {
			return sc, err
		}
		n, err := readInt()
		if err != nil {
			return sc, err
		}
		code := make([]byte, n)
		_, err = io.ReadFull(rd, code)
		if err != nil {
			return sc, err
		}
		def.Code = make([]uint32, 0, n/4)
		def.Rest = make([]byte, 0, n%4)
		for len(code) >= 4 {
			def.Code = append(def.Code, binary.LittleEndian.Uint32(code))
			code = code[4:]
		}
		def.Rest = append(def.Rest, code...)
		sc.Funcs = append(sc.Funcs, def)
	}
	if err := expect("DONE"); err != nil && err != io.ErrUnexpectedEOF {
		return sc, err
	}
	return sc, nil
}
