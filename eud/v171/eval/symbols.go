package eval

import "reflect"

//go:generate yaegi extract github.com/noxworld-dev/noxscript/eud/v171
//go:generate goimports -w .

var Symbols = make(map[string]map[string]reflect.Value)
