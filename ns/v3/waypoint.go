package ns

import ns4 "github.com/noxworld-dev/noxscript/ns/v4"

type WaypointID int

func (id WaypointID) WaypointScriptID() int {
	return int(id)
}

type WaypointGroupID int

func (id WaypointGroupID) WaypointGroupScriptID() int {
	return int(id)
}

// Waypoint looks up waypoint by name.
func Waypoint(name string) WaypointID {
	return asWaypoint(ns4.Waypoint(name))
}

// WayPointOn enables a waypoint.
func WayPointOn(waypoint WaypointID) {
	ns4.AsWaypoint(waypoint).Enable(true)
}

// WayPointOff disables a waypoint.
func WayPointOff(waypoint WaypointID) {
	ns4.AsWaypoint(waypoint).Enable(false)
}

// WayPointToggle toggles waypoint.
func WayPointToggle(waypoint WaypointID) {
	ns4.AsWaypoint(waypoint).Toggle()
}

// IsWaypointOn gets whether waypoint is enabled.
func IsWaypointOn(waypoint WaypointID) bool {
	return ns4.AsWaypoint(waypoint).IsEnabled()
}

// GetWaypointX gets waypoint X coordinate.
func GetWaypointX(waypoint WaypointID) float32 {
	return ns4.AsWaypoint(waypoint).Pos().X
}

// GetWaypointY gets waypoint Y coordinate.
func GetWaypointY(waypoint WaypointID) float32 {
	return ns4.AsWaypoint(waypoint).Pos().Y
}

// MoveWaypoint sets a waypoint's location.
func MoveWaypoint(waypoint WaypointID, x, y float32) {
	ns4.AsWaypoint(waypoint).SetPos(ns4.Pointf{X: x, Y: y})
}

// WaypointGroup looks up waypoint group by name.
func WaypointGroup(name string) WaypointGroupID {
	return asWaypointGroup(ns4.WaypointGroup(name))
}

// WayPointGroupOn enables waypoints in a group.
func WayPointGroupOn(waypointGroup WaypointGroupID) {
	ns4.AsWaypointGroup(waypointGroup).Enable(true)
}

// WayPointGroupOff disables waypoints in a group.
func WayPointGroupOff(waypointGroup WaypointGroupID) {
	ns4.AsWaypointGroup(waypointGroup).Enable(false)
}

// WayPointGroupToggle toggles whether objects in group are disabled.
func WayPointGroupToggle(waypointGroup WaypointGroupID) {
	ns4.AsWaypointGroup(waypointGroup).Toggle()
}
