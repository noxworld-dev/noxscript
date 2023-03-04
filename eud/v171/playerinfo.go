package eud

import (
	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

func GetPlayerIndex(unit ns3.ObjectID) int {
	panic("not implemented")
}

func PlayerGetEquipedWeapon(unit ns3.ObjectID) ns3.ObjectID {
	panic("not implemented")
}

func PlayerGetNextWeapon(unit ns3.ObjectID) ns3.ObjectID {
	panic("not implemented")
}

func PlayerGetStrength(unit ns3.ObjectID) int {
	panic("not implemented")
}

func PlayerSetStrength(unit ns3.ObjectID, val int) {
	panic("not implemented")
}

func PlayerGetMaxSpeed(unit ns3.ObjectID) int {
	panic("not implemented")
}

func PlayerGetMaxWeight(unit ns3.ObjectID) int {
	panic("not implemented")
}

func PlayerGetArmorValue(unit ns3.ObjectID) int {
	panic("not implemented")
}

func GetPlayerMouseX(unit ns3.ObjectID) int {
	panic("not implemented")
}

func GetPlayerMouseY(unit ns3.ObjectID) int {
	panic("not implemented")
}

func GetPlayerMouseXY(unit ns3.ObjectID, xpos, ypos *int) bool {
	panic("not implemented")
}

func GiveCreatureToPlayer(owner int, unit ns3.ObjectID) {
	panic("not implemented")
}

func PlayerGetCurrentManaAmount(unit ns3.ObjectID) int {
	return ns4.AsObj(unit).CurrentMana()
}

func PlayerSetCurrentManaAmount(unit ns3.ObjectID, val int) {
	ns4.AsObj(unit).SetMana(val)
}

func PlayerGetMaximumManaAmount(unit ns3.ObjectID) int {
	return ns4.AsObj(unit).MaxMana()
}

func PlayerSetMaximumManaAmount(unit ns3.ObjectID, val int) {
	ns4.AsObj(unit).SetMaxMana(val)
}

func PlayerGetCurrentLevel(unit ns3.ObjectID) int {
	return ns4.AsObj(unit).GetLevel()
}

func PlayerIngameNick(unit ns3.ObjectID) string {
	if u := ns4.AsObj(unit); u != nil {
		if p := u.Player(); p != nil {
			return p.Name()
		}
	}
	return "Jack" // TODO: not sure if it expects this
}
