package ns

import "math"

// ObjSearcher is an interface for iterating and filtering over objects.
type ObjSearcher interface {
	// FindObjects calls fnc for all objects in a set matching all the conditions.
	// It returns a number of objects matched. If fnc returns false, the function stops the search.
	// If fnc is nil, the function only counts a number of objects matching a condition.
	//
	// Example:
	//
	//	// Find and damage monsters in a 100px circle around obj.
	//	g.FindObjects(
	//		func(it ns.Obj) bool {
	//			it.Damage(obj, 10, damage.FLAME)
	//			return true
	//		},
	//		ns.InCirclef{Center: obj, R: 100},
	//		ns.HasClass(object.ClassMonster),
	//	)
	FindObjects(fnc func(it Obj) bool, conditions ...ObjCond) int
}

// FindObject searches the map for any object that matches all conditions.
// If object is not found it returns nil.
//
// Example:
//
//	// Find exit on the map.
//	found := ns.FindObject(ns.EqualClass(object.ClassExit))
func FindObject(conditions ...ObjCond) Obj {
	return FindObjectIn(nil, conditions...)
}

// FindObjectIn searches the set for any object that matches all conditions.
// If object is not found it returns nil.
//
// Example:
//
//	// Find exit on the map.
//	found := ns.FindObject(ns.EqualClass(object.ClassExit))
func FindObjectIn(s ObjSearcher, conditions ...ObjCond) Obj {
	var found Obj
	findObjectsIn(s, func(it Obj) bool {
		found = it
		return false
	}, conditions...)
	return found
}

// FindClosestObject searches the map for closest object that matches all conditions.
// If no objects are found, it returns nil.
//
// Example:
//
//	// Find monster or player closest to obj.
//	found := ns.FindClosestObject(obj, ns.HasClass(object.ClassMonster | object.ClassPlayer))
func FindClosestObject(from Positioner, conditions ...ObjCond) Obj {
	return FindClosestObjectIn(from, nil, conditions...)
}

// FindClosestObjectIn searches the set for closest object that matches all conditions.
// If no objects are found, it returns nil.
//
// Example:
//
//	// Find monster or player closest to obj.
//	found := ns.FindClosestObjectIn(obj, g, ns.HasClass(object.ClassMonster | object.ClassPlayer))
func FindClosestObjectIn(from Positioner, s ObjSearcher, conditions ...ObjCond) Obj {
	if from == nil {
		return nil
	}
	fromObj, _ := from.(Obj)
	var (
		found Obj
		min   float64 = math.MaxFloat64
	)
	pos := from.Pos()
	findObjectsIn(s, func(it Obj) bool {
		if it == fromObj {
			return true // skip
		}
		dist := it.Pos().Sub(pos).Len()
		if dist < min {
			min = dist
			found = it
		}
		return true
	}, conditions...)
	return found
}

// FindObjects calls fnc for all map objects matching all the conditions.
// It returns a number of objects matched. If fnc returns false, the function stops the search.
// If fnc is nil, the function only counts a number of objects matching a condition.
//
// Example:
//
//	// Find and damage all monsters in a 100px circle around obj.
//	ns.FindObjects(
//		func(it ns.Obj) bool {
//			it.Damage(obj, 10, damage.FLAME)
//			return true
//		},
//		ns.InCirclef{Center: obj, R: 100},
//		ns.HasClass(object.ClassMonster),
//	)
func FindObjects(fnc func(it Obj) bool, conditions ...ObjCond) int {
	return findObjectsIn(nil, fnc, conditions...)
}

func findObjectsIn(s ObjSearcher, fnc func(it Obj) bool, conditions ...ObjCond) int {
	if s == nil {
		s = impl
	}
	if s == nil {
		return 0
	}
	return s.FindObjects(fnc, conditions...)
}

// FindAllObjects is similar to FindObjects, but returns all objects as an array/slice.
//
// Example:
//
//	// Find and return all missiles in a 50px circle around obj.
//	arr := ns.FindAllObjects(
//		ns.InCirclef{Center: obj, R: 50},
//		ns.HasClass(object.ClassMissile),
//	)
func FindAllObjects(conditions ...ObjCond) []Obj {
	return FindAllObjectsIn(nil, conditions...)
}

// FindAllObjectsIn iterates over objects, and returns all matched objects as an array/slice.
//
// Example:
//
//	// Find and return missiles in a 50px circle around obj.
//	arr := ns.FindAllObjects(g,
//		ns.InCirclef{Center: obj, R: 50},
//		ns.HasClass(object.ClassMissile),
//	)
func FindAllObjectsIn(s ObjSearcher, conditions ...ObjCond) []Obj {
	var out []Obj
	findObjectsIn(s, func(it Obj) bool {
		out = append(out, it)
		return true
	}, conditions...)
	return out
}

// WallSearcher is an interface for iterating and filtering over walls.
type WallSearcher interface {
	// FindWalls calls fnc for all walls in a set matching all the conditions.
	// It returns a number of walls matched. If fnc returns false, the function stops the search.
	// If fnc is nil, the function only counts a number of walls matching a condition.
	//
	// Example:
	//
	//	// Find and disable walls with a window in a 100px circle around obj.
	//	g.FindWalls(
	//		func(it ns.WallObj) bool {
	//			it.Enable(false)
	//			return true
	//		},
	//		ns.InCirclef{Center: obj, R: 100},
	//		ns.HasWallFlags(ns.WallFlagWindow),
	//	)
	FindWalls(fnc func(it WallObj) bool, conditions ...WallCond) int
}

// FindWall searches the map for any wall that matches all conditions.
// If wall is not found it returns nil.
//
// Example:
//
//	// Find window on the map.
//	found := ns.FindWall(ns.HasWallFlags(ns.WallFlagWindow))
func FindWall(conditions ...WallCond) WallObj {
	return FindWallIn(nil, conditions...)
}

// FindWallIn searches the set for any walls that matches all conditions.
// If wall is not found it returns nil.
//
// Example:
//
//	// Find window on the map.
//	found := ns.FindWall(ns.HasWallFlags(ns.WallFlagWindow))
func FindWallIn(s WallSearcher, conditions ...WallCond) WallObj {
	var found WallObj
	findWallsIn(s, func(it WallObj) bool {
		found = it
		return false
	}, conditions...)
	return found
}

// FindClosestWall searches the map for closest wall that matches all conditions.
// If no walls are found, it returns nil.
//
// Example:
//
//	// Find window closest to obj.
//	found := ns.FindClosestWall(obj, ns.HasWallFlags(ns.WallFlagWindow))
func FindClosestWall(from Positioner, conditions ...WallCond) WallObj {
	return FindClosestWallIn(from, nil, conditions...)
}

// FindClosestWallIn searches the set for closest wall that matches all conditions.
// If no walls are found, it returns nil.
//
// Example:
//
//	// Find window closest to obj.
//	found := ns.FindClosestWall(obj, ns.HasWallFlags(ns.WallFlagWindow))
func FindClosestWallIn(from Positioner, s WallSearcher, conditions ...WallCond) WallObj {
	if from == nil {
		return nil
	}
	fromWall, _ := from.(WallObj)
	var (
		found WallObj
		min   float64 = math.MaxFloat64
	)
	pos := from.Pos()
	findWallsIn(s, func(it WallObj) bool {
		if it == fromWall {
			return true // skip
		}
		dist := it.Pos().Sub(pos).Len()
		if dist < min {
			min = dist
			found = it
		}
		return true
	}, conditions...)
	return found
}

// FindWalls calls fnc for all map walls matching all the conditions.
// It returns a number of walls matched. If fnc returns false, the function stops the search.
// If fnc is nil, the function only counts a number of walls matching a condition.
//
// Example:
//
//	// Find and disable walls with a window in a 100px circle around obj.
//	g.FindWalls(
//		func(it ns.WallObj) bool {
//			it.Enable(false)
//			return true
//		},
//		ns.InCirclef{Center: obj, R: 100},
//		ns.HasWallFlags(ns.WallFlagWindow),
//	)
func FindWalls(fnc func(it WallObj) bool, conditions ...WallCond) int {
	return findWallsIn(nil, fnc, conditions...)
}

func findWallsIn(s WallSearcher, fnc func(it WallObj) bool, conditions ...WallCond) int {
	if s == nil {
		s = impl
	}
	if s == nil {
		return 0
	}
	return s.FindWalls(fnc, conditions...)
}

// FindAllWalls is similar to FindWalls, but returns all walls as an array/slice.
//
// Example:
//
//	// Find and return all walls in a 50px circle around obj.
//	arr := ns.FindAllWalls(
//		ns.InCirclef{Center: obj, R: 50},
//	)
func FindAllWalls(conditions ...WallCond) []WallObj {
	return FindAllWallsIn(nil, conditions...)
}

// FindAllWallsIn iterates over walls, and returns all matched walls as an array/slice.
//
// Example:
//
//	// Find and return missiles in a 50px circle around obj.
//	arr := ns.FindAllWalls(
//		ns.InCirclef{Center: obj, R: 50},
//	)
func FindAllWallsIn(s WallSearcher, conditions ...WallCond) []WallObj {
	var out []WallObj
	findWallsIn(s, func(it WallObj) bool {
		out = append(out, it)
		return true
	}, conditions...)
	return out
}
