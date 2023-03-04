package eud

// SetMemory writes value to memory at a given address.
//
// Deprecated: OpenNox cannot support direct memory access.
func SetMemory(addr int, value int) {
	panic("direct memory access is not possible in OpenNox")
}

// GetMemory reads value from memory at a given address.
//
// Deprecated: OpenNox cannot support direct memory access.
func GetMemory(addr int) int {
	panic("direct memory access is not possible in OpenNox")
}

// GetMemoryFloat reads float value from memory at a given address.
//
// Deprecated: OpenNox cannot support direct memory access.
func GetMemoryFloat(addr int) float32 {
	panic("direct memory access is not possible in OpenNox")
}
