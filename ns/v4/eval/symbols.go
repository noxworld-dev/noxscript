package eval

import "reflect"

//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4
//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4/audio
//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4/class
//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4/damage
//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4/effect
//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4/enchant
//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4/spell
//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v4/subclass
//go:generate goimports -w .

var Symbols = make(map[string]map[string]reflect.Value)
