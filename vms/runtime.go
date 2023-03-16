package vms

import (
	"sort"

	"golang.org/x/exp/maps"
)

var vms = make(map[string]RuntimeInfo)

// Register registers a new script VM runtime.
func Register(r RuntimeInfo) {
	if r.Name == "" {
		panic("name must be set")
	}
	if r.NewMap == nil {
		panic("new map function must be set")
	}
	if _, ok := vms[r.Name]; ok {
		panic("already registered")
	}
	vms[r.Name] = r
}

// VMRuntimes returns all registered VM runtimes.
func VMRuntimes() []RuntimeInfo {
	keys := maps.Keys(vms)
	sort.Strings(keys)
	var out []RuntimeInfo
	for _, k := range keys {
		out = append(out, vms[k])
	}
	return out
}

// RuntimeInfo is a type for registering new script runtime implementations.
type RuntimeInfo struct {
	// Name is a short name for the VM runtime. Must be unique.
	Name string
	// Title is a human-friendly name for the VM runtime.
	Title string
	// NewMap creates a new script VM for map scripts.
	// If there's no scripts for the map, function may return nil, nil.
	NewMap func(g Engine, maps string, name string) (VM, error)
}
