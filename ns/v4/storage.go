package ns

type StorageType interface {
	isStorageType()
}

// Session is a game session storage with a given name. It will persist as long as a single game session runs.
type Session struct {
	Name string
}

func (Session) isStorageType() {}

// Persistent returns a persistent storage with a given name. It will persist between OpenNox restarts.
type Persistent struct {
	Name string
}

func (Persistent) isStorageType() {}

// Store returns a global storage with a given type.
func Store(typ StorageType) Storage {
	if impl == nil {
		return nil
	}
	return impl.Store(typ)
}

type Storage interface {
	// Get fetches a custom object or value from the storage with a given key.
	// Argument must be a pointer to a value of the same type as used during Set.
	Get(key string, obj any) error
	// Set stores a custom object or value into the storage with a given key.
	Set(key string, obj any) error
}
