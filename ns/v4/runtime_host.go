package ns

// SetRuntime is used by the script host to bind runtime for the script package.
// Scripts must not call this functions.
func SetRuntime(r Implementation) {
	impl = r
}
