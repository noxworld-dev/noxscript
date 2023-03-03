package ns

type Handle interface {
	// ScriptID returns internal script ID for the object.
	ScriptID() int
}

type TimerHandle interface {
	Handle
	// TimerScriptID returns internal script ID for the timer.
	TimerScriptID() int
}

// AsTimer looks up a Timer for a given handle.
func AsTimer(h TimerHandle) Timer {
	if impl == nil || h == nil {
		return nil
	}
	if obj, ok := h.(Timer); ok {
		return obj
	}
	return impl.TimerByHandle(h)
}

type ObjHandle interface {
	Handle
	// ObjScriptID returns internal script ID for the object.
	ObjScriptID() int
}

// AsObj looks up an Obj for a given handle.
func AsObj(h ObjHandle) Obj {
	if impl == nil || h == nil {
		return nil
	}
	if obj, ok := h.(Obj); ok {
		return obj
	}
	return impl.ObjectByHandle(h)
}

type ObjGroupHandle interface {
	Handle
	// ObjGroupScriptID returns internal script ID for the object group.
	ObjGroupScriptID() int
}

// AsObjGroup looks up a ObjGroup for a given handle.
func AsObjGroup(h ObjGroupHandle) ObjGroup {
	if impl == nil || h == nil {
		return nil
	}
	if obj, ok := h.(ObjGroup); ok {
		return obj
	}
	return impl.ObjectGroupByHandle(h)
}

type WallHandle interface {
	Handle
	// WallScriptID returns internal script ID for the wall.
	WallScriptID() int
}

// AsWall looks up a WallObj for a given handle.
func AsWall(h WallHandle) WallObj {
	if impl == nil || h == nil {
		return nil
	}
	if obj, ok := h.(WallObj); ok {
		return obj
	}
	return impl.WallByHandle(h)
}

type WallGroupHandle interface {
	Handle
	// WallGroupScriptID returns internal script ID for the wall group.
	WallGroupScriptID() int
}

// AsWallGroup looks up a WallGroupObj for a given handle.
func AsWallGroup(h WallGroupHandle) WallGroupObj {
	if impl == nil || h == nil {
		return nil
	}
	if obj, ok := h.(WallGroupObj); ok {
		return obj
	}
	return impl.WallGroupByHandle(h)
}

type WaypointHandle interface {
	Handle
	// WaypointScriptID returns internal script ID for the waypoint.
	WaypointScriptID() int
}

// AsWaypoint looks up a WaypointObj for a given handle.
func AsWaypoint(h WaypointHandle) WaypointObj {
	if impl == nil || h == nil {
		return nil
	}
	if obj, ok := h.(WaypointObj); ok {
		return obj
	}
	return impl.WaypointByHandle(h)
}

type WaypointGroupHandle interface {
	Handle
	// WaypointGroupScriptID returns internal script ID for the waypoint.
	WaypointGroupScriptID() int
}

// AsWaypointGroup looks up a WaypointGroupObj for a given handle.
func AsWaypointGroup(h WaypointGroupHandle) WaypointGroupObj {
	if impl == nil || h == nil {
		return nil
	}
	if obj, ok := h.(WaypointGroupObj); ok {
		return obj
	}
	return impl.WaypointGroupByHandle(h)
}
