package ns

import (
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
)

// Effect triggers an effect from point p1 to p2.
// Some effects only have one point, in which case p2 is ignored.
func Effect(effect effect.Effect, p1, p2 Positioner) {
	if impl == nil {
		return
	}
	impl.Effect(effect, p1, p2)
}

// CastSpell casts a spell from source to target.
//
//	Example:
//	  CastSpell(spell.DEATH_RAY, Object("CruelDude"), GetHost())
//	  CastSpell(spell.DEATH_RAY, types.Ptf(10, 5), Waypoint("Target"))
func CastSpell(spell spell.Spell, source, target Positioner) {
	if impl == nil {
		return
	}
	impl.CastSpell(spell, source, target)
}
