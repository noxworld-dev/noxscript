package vms

import (
	"fmt"
	"reflect"
)

func NewBaseVM(g Engine) *BaseVM {
	return &BaseVM{g: g}
}

type BaseVM struct {
	g      Engine
	stdout Writer
	stderr Writer
	events eventsCache
}

func (vm *BaseVM) base() *BaseVM {
	return vm
}

func (vm *BaseVM) GetEngine() Engine {
	return vm.g
}

func (vm *BaseVM) Stdout() Writer {
	if vm.stdout == nil {
		vm.stdout = vm.g.Console(false)
	}
	return vm.stdout
}

func (vm *BaseVM) Stderr() Writer {
	if vm.stderr == nil {
		vm.stderr = vm.g.Console(true)
	}
	return vm.stderr
}

// VM is an interface for a virtual machine running the script engine.
type VM interface {
	base() *BaseVM
	// GetEngine returns Engine used by this VM.
	GetEngine() Engine
	// Stdout returns standard output used by this VM.
	Stdout() Writer
	// Stderr returns error output used by this VM.
	Stderr() Writer

	// Exec executes text as a script source code.
	Exec(s string) (reflect.Value, error)
	// ExecFile executes a script from the file or directory.
	ExecFile(path string) error
	// GetSymbol tries to find a function or variable with a given name and type.
	// If symbol is found, but type doesn't match, it returns an error.
	// If symbol is not found, it returns reflect.Value{}, false, nil.
	GetSymbol(name string, typ reflect.Type) (reflect.Value, bool, error)
	// Close the VM and free resources.
	Close() error
}

type ErrIncorrectSymbolType struct {
	Name     string
	Expected reflect.Type
	Actual   reflect.Type
}

func (e *ErrIncorrectSymbolType) Error() string {
	return fmt.Sprintf("unexpected type of %q: expected %v, actual %v", e.Name, e.Expected, e.Actual)
}

// GetVMSymbol is a generic helper for VM.GetSymbol.
// It is suitable for functions mostly, since it returns value directly, not a pointer.
func GetVMSymbol[T any](vm VM, name string) (T, error) {
	var zero T
	rv, ok, err := vm.GetSymbol(name, reflect.TypeOf(zero))
	if err != nil || !ok {
		return zero, err
	}
	return rv.Interface().(T), nil
}

// GetVMSymbolPtr is a generic helper for VM.GetSymbol.
// It is similar to GetVMSymbol, but returns a pointer to the value. Useful for variables.
func GetVMSymbolPtr[T any](vm VM, name string) (*T, error) {
	var zero T
	rv, ok, err := vm.GetSymbol(name, reflect.TypeOf(zero))
	if err != nil || !ok {
		return nil, err
	}
	return rv.Addr().Interface().(*T), nil
}
