package ns

import ns4 "github.com/noxworld-dev/noxscript/ns/v4"

type Handle = ns4.Handle

func asTimer(obj ns4.Timer) TimerID {
	if obj == nil {
		return 0
	}
	return TimerID(obj.TimerScriptID())
}

func asObj(obj ns4.Obj) ObjectID {
	if obj == nil {
		return 0
	}
	return ObjectID(obj.ObjScriptID())
}

func asObjGroup(obj ns4.ObjGroup) ObjectGroupID {
	if obj == nil {
		return 0
	}
	return ObjectGroupID(obj.ObjGroupScriptID())
}

func asWall(obj ns4.WallObj) WallID {
	if obj == nil {
		return 0
	}
	return WallID(obj.WallScriptID())
}

func asWallGroup(obj ns4.WallGroupObj) WallGroupID {
	if obj == nil {
		return 0
	}
	return WallGroupID(obj.WallGroupScriptID())
}

func asWaypoint(obj ns4.WaypointObj) WaypointID {
	if obj == nil {
		return 0
	}
	return WaypointID(obj.WaypointScriptID())
}

func asWaypointGroup(obj ns4.WaypointGroupObj) WaypointGroupID {
	if obj == nil {
		return 0
	}
	return WaypointGroupID(obj.WaypointGroupScriptID())
}
