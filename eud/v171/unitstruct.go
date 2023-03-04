package eud

import (
	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/object"
)

func unitHasClass(unit ns3.ObjectID, v object.Class) bool {
	if p := ns4.AsObj(unit); p != nil {
		return p.Class().Has(v)
	}
	return false
}

// IsMissileUnit checks if unit is a missile.
func IsMissileUnit(unit ns3.ObjectID) bool {
	return unitHasClass(unit, object.ClassMissile)
}

// IsPlayerUnit checks if unit as a player.
func IsPlayerUnit(unit ns3.ObjectID) bool {
	return unitHasClass(unit, object.ClassPlayer)
}

// IsMonsterUnit checks if unit as a monster.
func IsMonsterUnit(unit ns3.ObjectID) bool {
	return unitHasClass(unit, object.ClassMonster)
}
