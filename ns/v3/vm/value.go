package vm

import (
	"math"

	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
)

// Value is an untyped value, as stored in NoxScript runtime.
type Value struct {
	data uint32
}

// String returns NoxScript value as string.
func (v Value) String() string {
	if impl == nil {
		return ""
	}
	return impl.GetString(v.data)
}

// Int returns NoxScript value as int.
func (v Value) Int() int {
	return int(int32(v.data))
}

// Float returns NoxScript value as float.
func (v Value) Float() float32 {
	return math.Float32frombits(v.data)
}

// Object returns NoxScript value as object.
func (v Value) Object() ns3.ObjectID {
	return ns3.ObjectID(v.data)
}

// ObjectGroup returns NoxScript value as object group.
func (v Value) ObjectGroup() ns3.ObjectGroupID {
	return ns3.ObjectGroupID(v.data)
}

// Waypoint returns NoxScript value as waypoint.
func (v Value) Waypoint() ns3.WaypointID {
	return ns3.WaypointID(v.data)
}

// WaypointGroup returns NoxScript value as waypoint group.
func (v Value) WaypointGroup() ns3.WaypointGroupID {
	return ns3.WaypointGroupID(v.data)
}

// Wall returns NoxScript value as wall.
func (v Value) Wall() ns3.WallID {
	return ns3.WallID(v.data)
}

// WallGroup returns NoxScript value as wall group.
func (v Value) WallGroup() ns3.WallGroupID {
	return ns3.WallGroupID(v.data)
}

// toValue converts some Go values into NoxScript compatible ones.
func toValue(val any) Value {
	var v Value
	switch val := val.(type) {
	case nil:
		v.data = 0
	case Value:
		v = val
	case int:
		v.data = uint32(int32(val))
	case int32:
		v.data = uint32(val)
	case int64:
		v.data = uint32(int32(val))
	case float32:
		v.data = math.Float32bits(val)
	case float64:
		v.data = math.Float32bits(float32(val))
	case string:
		if impl != nil {
			v.data = impl.NewString(val)
		} else {
			v.data = 0
		}
	case ns3.Handle:
		v.data = uint32(val.ScriptID())
	default:
		panic("unsupported value type")
	}
	return v
}
