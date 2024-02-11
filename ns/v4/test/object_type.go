package nstest

import (
	"sync/atomic"

	"github.com/noxworld-dev/opennox-lib/object"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/class"
	"github.com/noxworld-dev/noxscript/ns/v4/subclass"
)

type ObjectTypes struct {
	last   uint32
	ByID   map[uint32]*ObjectType
	ByName map[string]*ObjectType

	Player *ObjectType
	NPC    *ObjectType
}

func (s *ObjectTypes) New(r *Runtime, cl object.Class, name string) *ObjectType {
	id := atomic.AddUint32(&s.last, 1)
	t := &ObjectType{r: r, id: id, name: name, class: cl}
	if s.ByID == nil {
		s.ByID = make(map[uint32]*ObjectType)
		s.ByName = make(map[string]*ObjectType)
	}
	s.ByID[t.id] = t
	s.ByName[name] = t
	return t
}

func (r *Runtime) NewObjectType(cl object.Class, name string) *ObjectType {
	return r.Types.New(r, cl, name)
}

type ObjectType struct {
	r        *Runtime
	id       uint32
	name     string
	class    object.Class
	FlagsVal object.Flags
}

func (t *ObjectType) asObjType() ns.ObjType {
	if t == nil {
		return nil
	}
	return t
}

func (t *ObjectType) Name() string {
	return t.name
}

func (t *ObjectType) Index() int {
	return int(t.id)
}

func (t *ObjectType) Class() object.Class {
	return t.class
}

func (t *ObjectType) HasClass(class class.Class) bool {
	//TODO implement me
	panic("implement me")
}

func (t *ObjectType) HasSubclass(subclass subclass.SubClass) bool {
	//TODO implement me
	panic("implement me")
}

func (t *ObjectType) Flags() object.Flags {
	return t.FlagsVal
}

func (t *ObjectType) Create(pos ns.Positioner) ns.Obj {
	if pos == nil {
		return nil
	}
	return t.r.NewObject(t, "", pos.Pos())
}

func (r *Runtime) ObjectType(name string) ns.ObjType {
	return r.Types.ByName[name].asObjType()
}

func (r *Runtime) ObjectTypeByInd(ind int) ns.ObjType {
	return r.Types.ByID[uint32(ind)].asObjType()
}
