package ns

import "github.com/noxworld-dev/opennox-lib/object"

// ObjCond checks if object matches a certain condition.
type ObjCond interface {
	// Matches returns true if objects matches a condition.
	// It must not modify the object.
	Matches(obj Obj) bool
}

// ObjCondFunc is a function that implements ObjCond.
type ObjCondFunc func(obj Obj) bool

func (f ObjCondFunc) Matches(obj Obj) bool {
	return f(obj)
}

var _ ObjCond = AND{}

// AND matches an object if all conditions match an object.
type AND []ObjCond

func (arr AND) Matches(obj Obj) bool {
	for _, c := range arr {
		if !c.Matches(obj) {
			return false
		}
	}
	return true
}

var _ ObjCond = OR{}

// OR matches an object if any of the conditions matches an object.
type OR []ObjCond

func (arr OR) Matches(obj Obj) bool {
	for _, c := range arr {
		if c.Matches(obj) {
			return true
		}
	}
	return len(arr) == 0
}

var _ ObjCond = NOT{}

// NOT matches an object if any other condition returns false.
type NOT struct {
	Cond ObjCond
}

func (c NOT) Matches(obj Obj) bool {
	return !c.Cond.Matches(obj)
}

var _ ObjCond = InRectf{}

// InRectf filters objects that are inside a given map rectangle.
type InRectf struct {
	Min, Max Pointf
}

func (c InRectf) Matches(obj Obj) bool {
	if c.Min.X > c.Max.X {
		c.Min.X, c.Max.X = c.Max.X, c.Min.X
	}
	if c.Min.Y > c.Max.Y {
		c.Min.Y, c.Max.Y = c.Max.Y, c.Min.Y
	}
	pos := obj.Pos()
	return pos.X >= c.Min.X && pos.X <= c.Max.X &&
		pos.Y >= c.Min.Y && pos.Y <= c.Max.Y
}

var _ ObjCond = InCirclef{}

// InCirclef filters objects that are inside a circle on the map.
type InCirclef struct {
	Center Positioner
	R      float64
}

func (c InCirclef) Matches(obj Obj) bool {
	if c.Center == nil {
		return false
	}
	return obj.Pos().Sub(c.Center.Pos()).Len() <= c.R
}

var _ ObjCond = EqualClass(0)

// EqualClass checks that class of object exactly matches this value.
type EqualClass object.Class

func (c EqualClass) Matches(obj Obj) bool {
	return obj.Class() == object.Class(c)
}

var _ ObjCond = HasClass(0)

// HasClass checks that class of object matches any of provided classes.
type HasClass object.Class

func (c HasClass) Matches(obj Obj) bool {
	c2 := object.Class(c)
	return obj.Class()&c2 != 0
}

var _ ObjCond = HasAllClasses(0)

// HasAllClasses checks that class of object matches all of provided classes.
type HasAllClasses object.Class

func (c HasAllClasses) Matches(obj Obj) bool {
	c2 := object.Class(c)
	return obj.Class()&c2 == c2
}

var _ ObjCond = HasObjFlags(0)

// HasObjFlags checks that object's flags match any of provided flags.
type HasObjFlags object.Flags

func (c HasObjFlags) Matches(obj Obj) bool {
	c2 := object.Flags(c)
	return obj.Flags()&c2 != 0
}

var _ ObjCond = HasAllClasses(0)

// HasAllObjFlags checks that object's flags match all of provided flags.
type HasAllObjFlags object.Flags

func (c HasAllObjFlags) Matches(obj Obj) bool {
	c2 := object.Flags(c)
	return obj.Flags()&c2 == c2
}

var _ ObjCond = HasTypeName{}

// HasTypeName checks that object matches any of provided type names.
type HasTypeName []string

func (arr HasTypeName) Matches(obj Obj) bool {
	typ := obj.Type().Name()
	for _, t := range arr {
		if typ == t {
			return true
		}
	}
	return false
}

var _ ObjCond = HasType{}

// HasType checks that object matches any of provided types.
type HasType []ObjType

func (arr HasType) Matches(obj Obj) bool {
	typ := obj.Type()
	for _, t := range arr {
		if t != nil && typ.Index() == t.Index() {
			return true
		}
	}
	return false
}

// HasTeam checks that object matches any of provided teams.
type HasTeam []Team

func (arr HasTeam) Matches(obj Obj) bool {
	for _, t := range arr {
		// don't check for nil, because it's a valid query for HasTeam
		if obj.HasTeam(t) {
			return true
		}
	}
	return false
}

// TODO: InPolygon, need to support polygons in the library
