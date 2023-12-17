package ns

import "github.com/noxworld-dev/opennox-lib/wall"

// NoWallSound enables or disables wall sounds.
func NoWallSound(noWallSound bool) {
	if impl == nil {
		return
	}
	impl.NoWallSound(noWallSound)
}

// WallObj is a pointer to a wall.
type WallObj interface {
	WallHandle
	// IsEnabled checks if object is currently enabled.
	IsEnabled() bool
	// Enable or disable the object.
	Enable(enable bool)
	// Toggle the object's enabled state.
	// It returns the new state after toggling the object.
	Toggle() bool
	// Flags returns wall flags.
	Flags() wall.Flags

	// Pos returns position of the wall in object coordinates.
	Pos() Pointf
	// GridPos returns wall grid position.
	GridPos() Point

	// Destroy breaks a wall.
	Destroy()
}

// Wall gets a wall by its grid coordinates.
func Wall(x int, y int) WallObj {
	if impl == nil {
		return nil
	}
	return impl.Wall(x, y)
}

// WallGroupObj is a group of walls.
type WallGroupObj interface {
	WallGroupHandle
	// Name returns wall group name.
	Name() string
	// Enable or disable the object.
	Enable(enable bool)
	// Toggle the object's enabled state.
	// It returns the new state after toggling the object.
	Toggle() bool

	// Destroy breaks walls in a group.
	Destroy()

	// EachWall calls fnc for all walls in the group.
	// If fnc returns false, the iteration stops.
	// If recursive is true, iteration will include items from nested groups.
	EachWall(recursive bool, fnc func(obj WallObj) bool)
}

// WallGroup lookups wall group by name.
func WallGroup(name string) WallGroupObj {
	if impl == nil {
		return nil
	}
	return impl.WallGroup(name)
}
