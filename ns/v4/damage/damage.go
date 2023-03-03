package damage

import "strconv"

type Type int

func (t Type) String() string {
	switch t {
	case BLADE:
		return "BLADE"
	case FLAME:
		return "FLAME"
	case CRUSH:
		return "CRUSH"
	case IMPALE:
		return "IMPALE"
	case DRAIN:
		return "DRAIN"
	case POISON:
		return "POISON"
	case DISPEL_UNDEAD:
		return "DISPEL_UNDEAD"
	case EXPLOSION:
		return "EXPLOSION"
	case BITE:
		return "BITE"
	case ELECTRIC:
		return "ELECTRIC"
	case CLAW:
		return "CLAW"
	case IMPACT:
		return "IMPACT"
	case LAVA:
		return "LAVA"
	case DEATH_MAGIC:
		return "DEATH_MAGIC"
	case PLASMA:
		return "PLASMA"
	case MANA_BOMB:
		return "MANA_BOMB"
	case ZAP_RAY:
		return "ZAP_RAY"
	case AIRBORNE_ELECTRIC:
		return "AIRBORNE_ELECTRIC"
	default:
		return strconv.Itoa(int(t))
	}
}

const (
	BLADE             = Type(0)
	FLAME             = Type(1)
	CRUSH             = Type(2)
	IMPALE            = Type(3)
	DRAIN             = Type(4)
	POISON            = Type(5)
	DISPEL_UNDEAD     = Type(6)
	EXPLOSION         = Type(7)
	BITE              = Type(8)
	ELECTRIC          = Type(9)
	CLAW              = Type(10)
	IMPACT            = Type(11)
	LAVA              = Type(12)
	DEATH_MAGIC       = Type(13)
	PLASMA            = Type(14)
	MANA_BOMB         = Type(15)
	ZAP_RAY           = Type(16)
	AIRBORNE_ELECTRIC = Type(17)
)
