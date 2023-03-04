package eud

import (
	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/object"
)

const (
	UNIT_FLAG_BELOW              = 0x1
	UNIT_FLAG_NO_UPDATE          = 0x2
	UNIT_FLAG_ACTIVE             = 0x4
	UNIT_FLAG_ALLOW_OVERLAP      = 0x8
	UNIT_FLAG_SHORT              = 0x10
	UNIT_FLAG_DESTROYED          = 0x20
	UNIT_FLAG_NO_COLLIDE         = 0x40
	UNIT_FLAG_MISSILE_HIT        = 0x80
	UNIT_FLAG_EQUIPPED           = 0x100
	UNIT_FLAG_PARTITIONED        = 0x200
	UNIT_FLAG_NO_COLLIDE_OWNER   = 0x400
	UNIT_FLAG_OWNER_VISIBLE      = 0x800
	UNIT_FLAG_EDIT_VISIBLE       = 0x1000
	UNIT_FLAG_NO_PUSH_CHARACTERS = 0x2000
	UNIT_FLAG_AIRBORNE           = 0x4000
	UNIT_FLAG_DEAD               = 0x8000
	UNIT_FLAG_SHADOW             = 0x10000
	UNIT_FLAG_FALLING            = 0x20000
	UNIT_FLAG_IN_HOLE            = 0x40000
	UNIT_FLAG_RESPAWN            = 0x80000
	UNIT_FLAG_ON_OBJECT          = 0x100000
	UNIT_FLAG_SIGHT_DESTROY      = 0x200000
	UNIT_FLAG_TRANSIENT          = 0x400000
	UNIT_FLAG_BOUNCY             = 0x800000
	UNIT_FLAG_ENABLED            = 0x1000000
	UNIT_FLAG_PENDING            = 0x2000000
	UNIT_FLAG_TRANSLUCENT        = 0x4000000
	UNIT_FLAG_STILL              = 0x8000000
	UNIT_FLAG_NO_AUTO_DROP       = 0x10000000
	UNIT_FLAG_FLICKER            = 0x20000000
	UNIT_FLAG_SELECTED           = 0x40000000
	UNIT_FLAG_MARKED             = 0x80000000
)

const (
	DAMAGE_TYPE_BLADE             = 0
	DAMAGE_TYPE_FLAME             = 1
	DAMAGE_TYPE_CRUSH             = 2
	DAMAGE_TYPE_IMPALE            = 3
	DAMAGE_TYPE_DRAIN             = 4
	DAMAGE_TYPE_POISON            = 5
	DAMAGE_TYPE_DISPEL_UNDEAD     = 6
	DAMAGE_TYPE_EXPLOSION         = 7
	DAMAGE_TYPE_BITE              = 8
	DAMAGE_TYPE_ELECTRIC          = 9
	DAMAGE_TYPE_CLAW              = 10
	DAMAGE_TYPE_IMPACT            = 11
	DAMAGE_TYPE_LAVA              = 12
	DAMAGE_TYPE_DEATH_MAGIC       = 13
	DAMAGE_TYPE_PLASMA            = 14
	DAMAGE_TYPE_MANA_BOMB         = 15
	DAMAGE_TYPE_ZAP_RAY           = 16
	DAMAGE_TYPE_AIRBORNE_ELECTRIC = 17
)

const (
	MON_STATUS_DESTROY_WHEN_DEAD = 1
	MON_STATUS_CHECK             = 2
	MON_STATUS_CAN_BLOCK         = 4
	MON_STATUS_CAN_DODGE         = 8
	MON_STATUS_UNUSED            = 0x10
	MON_STATUS_CAN_CAST_SPELLS   = 0x20
	MON_STATUS_HOLD_YOUR_GROUND  = 0x40
	MON_STATUS_SUMMONED          = 0x80
	MON_STATUS_ALERT             = 0x100
	MON_STATUS_INJURED           = 0x200
	MON_STATUS_CAN_SEE_FRIENDS   = 0x400
	MON_STATUS_CAN_HEAL_SELF     = 0x800
	MON_STATUS_CAN_HEAL_OTHERS   = 0x1000
	MON_STATUS_CAN_RUN           = 0x2000
	MON_STATUS_RUNNING           = 0x4000
	MON_STATUS_ALWAYS_RUN        = 0x8000
	MON_STATUS_NEVER_RUN         = 0x10000
	MON_STATUS_BOT               = 0x20000
	MON_STATUS_MORPHED           = 0x40000
	MON_STATUS_STAY_DEAD         = 0x80000
	MON_STATUS_ON_FIRE           = 0x100000
	MON_STATUS_FRUSTRATED        = 0x200000
)

const (
	UNIT_CLASS_MISSILE          = 0x1
	UNIT_CLASS_MONSTER          = 0x2
	UNIT_CLASS_PLAYER           = 0x4
	UNIT_CLASS_OBSTACLE         = 0x8
	UNIT_CLASS_FOOD             = 0x10
	UNIT_CLASS_EXIT             = 0x20
	UNIT_CLASS_KEY              = 0x40
	UNIT_CLASS_DOOR             = 0x80
	UNIT_CLASS_INFO_BOOK        = 0x100
	UNIT_CLASS_TRIGGER          = 0x200
	UNIT_CLASS_TRANSPORTER      = 0x400
	UNIT_CLASS_HOLE             = 0x800
	UNIT_CLASS_WAND             = 0x1000
	UNIT_CLASS_FIRE             = 0x2000
	UNIT_CLASS_ELEVATOR         = 0x4000
	UNIT_CLASS_ELEVATOR_SHAFT   = 0x8000
	UNIT_CLASS_DANGEROUS        = 0x10000
	UNIT_CLASS_MONSTERGENERATOR = 0x20000
	UNIT_CLASS_READABLE         = 0x40000
	UNIT_CLASS_LIGHT            = 0x80000
	UNIT_CLASS_SIMPLE           = 0x100000
	UNIT_CLASS_COMPLEX          = 0x200000
	UNIT_CLASS_IMMOBILE         = 0x400000
	UNIT_CLASS_VISIBLE_ENABLE   = 0x800000
	UNIT_CLASS_WEAPON           = 0x1000000
	UNIT_CLASS_ARMOR            = 0x2000000
	UNIT_CLASS_NOT_STACKABLE    = 0x4000000
	UNIT_CLASS_TREASURE         = 0x8000000
	UNIT_CLASS_FLAG             = 0x10000000
	UNIT_CLASS_CLIENT_PERSIST   = 0x20000000
	UNIT_CLASS_CLIENT_PREDICT   = 0x40000000
	UNIT_CLASS_PICKUP           = 0x80000000
)

// GetUnitClass returns unit class.
func GetUnitClass(unit ns3.ObjectID) int {
	if p := ns4.AsObj(unit); p != nil {
		return int(p.Class())
	}
	return 0
}

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

// GetUnitThingID returns an index of an object type in game's database.
func GetUnitThingID(unit ns3.ObjectID) int {
	return ns4.AsObj(unit).Type().Index()
}

// SetUnitMaxHealth sets units max health.
func SetUnitMaxHealth(unit ns3.ObjectID, amount int) {
	ns4.AsObj(unit).SetMaxHealth(amount)
}

// DistanceUnitToUnit computes distance between two units.
func DistanceUnitToUnit(unit1, unit2 ns3.ObjectID) float32 {
	obj1, obj2 := ns4.AsObj(unit1), ns4.AsObj(unit2)
	if obj1 == nil || obj2 == nil {
		return 0
	}
	return float32(obj1.Pos().Sub(obj2.Pos()).Len())
}
