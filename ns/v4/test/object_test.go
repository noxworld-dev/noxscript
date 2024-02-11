package nstest

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/types"
	"github.com/shoenig/test/must"
)

func TestObject(t *testing.T) {
	r := NewRuntime()

	obj := r.NewObject(r.Types.NPC, "", types.Ptf(0, 0))

	// Object moves according to velocity and resets it
	obj.ApplyForce(types.Ptf(1, 0))
	must.EqOp(t, types.Ptf(0, 0), obj.PosVec)
	obj.Update()
	must.EqOp(t, types.Ptf(1, 0), obj.PosVec)
	must.EqOp(t, types.Ptf(0, 0), obj.VelVec)

	// Can pick up and equip items
	weapon := r.NewObjectType(object.ClassWeapon, "Sword")
	w1 := r.NewObject(weapon, "", types.Ptf(0, 0))
	w2 := r.NewObject(weapon, "", types.Ptf(0, 0))
	must.False(t, obj.HasItem(w1))
	must.False(t, obj.HasEquipment(w1))

	must.True(t, obj.Pickup(w1))
	must.True(t, obj.Equip(w2))
	must.True(t, obj.HasItem(w1))
	must.False(t, obj.HasEquipment(w1))
	must.True(t, obj.HasItem(w2))
	must.True(t, obj.HasEquipment(w2))

	// Equipping one weapon de-equips another
	must.True(t, obj.Equip(w1))
	must.True(t, obj.HasItem(w1))
	must.True(t, obj.HasEquipment(w1))
	must.True(t, obj.HasItem(w2))
	must.False(t, obj.HasEquipment(w2))
}
