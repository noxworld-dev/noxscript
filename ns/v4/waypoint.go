package ns

type WaypointObj interface {
	WaypointHandle
	// Name returns waypoint name.
	Name() string
	// Pos returns current position of the object.
	Pos() Pointf
	// SetPos instantly moves object to a given position.
	SetPos(p Pointf)
	// IsEnabled checks if object is currently enabled.
	IsEnabled() bool
	// Enable or disable the object.
	Enable(enable bool)
	// Toggle the object's enabled state.
	// It returns the new state after toggling the object.
	Toggle() bool
}

type WaypointGroupObj interface {
	WaypointGroupHandle
	// Name returns waypoint group name.
	Name() string
	// Enable or disable the object.
	Enable(enable bool)
	// Toggle the object's enabled state.
	// It returns the new state after toggling the object.
	Toggle() bool

	// EachWaypoint calls fnc for all waypoints in the group.
	// If fnc returns false, the iteration stops.
	// If recursive is true, iteration will include items from nested groups.
	EachWaypoint(recursive bool, fnc func(obj WaypointObj) bool)
}

// NewWaypoint creates a new waypoint with a given name and position.
func NewWaypoint(name string, pos Pointf) WaypointObj {
	if impl == nil {
		return nil
	}
	return impl.NewWaypoint(name, pos)
}

// Waypoints returns a list of all waypoints.
func Waypoints() []WaypointObj {
	if impl == nil {
		return nil
	}
	return impl.Waypoints()
}

// Waypoint looks up waypoint by name.
func Waypoint(name string) WaypointObj {
	if impl == nil {
		return nil
	}
	return impl.Waypoint(name)
}

// WaypointGroup looks up waypoint group by name.
func WaypointGroup(name string) WaypointGroupObj {
	if impl == nil {
		return nil
	}
	return impl.WaypointGroup(name)
}
