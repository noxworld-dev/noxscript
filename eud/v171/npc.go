package eud

import (
	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

func CreateNewNPC(x, y float32, dest *ns3.ObjectID) {
	obj := ns4.CreateObject("NPC", ns4.Ptf(x, y))
	if dest != nil {
		*dest = ns3.ObjectID(obj.ObjScriptID())
	}
	// FIXME: set colors, and other values
}

func CreateSingleColorMaidenAt(r, g, b int, x, y float32) ns3.ObjectID {
	obj := ns4.CreateObject("Maiden", ns4.Ptf(x, y))
	// FIXME: set color and other things
	// TODO: original code uses "Bear2" to fix death animation and maybe something else
	return ns3.ObjectID(obj.ObjScriptID())
}

func NPCDressupEquipment(npc ns3.ObjectID, item ns3.ObjectID, mode bool) {
	panic("not implemented")
}
