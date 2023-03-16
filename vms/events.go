package vms

type eventsCache struct {
	onFrame func()
	onEvent map[string]func()
}

// OnFrame calls OnFrame function in VM.
func OnFrame(vm VM) {
	b := vm.base()
	if b.events.onFrame != nil {
		b.events.onFrame()
		return
	}
	const name = "OnFrame"
	var out func()
	if fnc, err := GetVMSymbol[func()](vm, name); err == nil {
		out = fnc
	} else if fnc, err := GetVMSymbol[func(frame int)](vm, name); err == nil {
		e := vm.GetEngine()
		out = func() {
			fnc(e.Frame())
		}
	}
	if out == nil {
		return
	}
	b.events.onFrame = out
	out()
}

// OnEvent calls event-related functions in VM.
func OnEvent(vm VM, event string) {
	b := vm.base()
	if fnc := b.events.onEvent[event]; fnc != nil {
		fnc()
		return
	}
	var out func()
	if fnc, err := GetVMSymbol[func()](vm, event); err == nil {
		out = fnc
	} else if fnc, err := GetVMSymbol[func()](vm, "On"+event); err == nil {
		out = fnc
	} else if fnc, err := GetVMSymbol[func(ev string)](vm, "OnEvent"); err == nil {
		out = func() {
			fnc(event)
		}
	}
	if out == nil {
		return
	}
	if b.events.onEvent == nil {
		b.events.onEvent = make(map[string]func())
	}
	b.events.onEvent[event] = out
	out()
}
