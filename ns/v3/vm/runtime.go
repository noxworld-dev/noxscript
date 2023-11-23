package vm

var impl Implementation

// Runtime returns implementation of all the accessible functions as an interface.
func Runtime() Implementation {
	return impl
}

// Game is an optional interface for the engine that exposes NoxScript runtime.
type Game interface {
	// NoxScriptVM returns implementation of NoxScript VM.
	NoxScriptVM() Implementation
}

// Implementation of the script VM. Only used in the engine itself.
type Implementation interface {
	NewString(s string) uint32
	GetString(val uint32) string

	GetFuncInd(fnc string) int
	GetFuncVar(fnc int, vari int) uint32
	SetFuncVar(fnc int, vari int, val uint32)

	CallFunc(fnc int, args []uint32) uint32
}
