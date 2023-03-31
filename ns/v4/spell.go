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
// Example:
//
//	ns.CastSpell(spell.DEATH_RAY, ns.Object("CruelDude"), ns.GetHost())
//	ns.CastSpell(spell.DEATH_RAY, ns.Ptf(10, 5), ns.Waypoint("Target"))
func CastSpell(spell spell.Spell, source, target Positioner) {
	if impl == nil {
		return
	}
	impl.CastSpell(spell, source, target)
}

// CastSpellLvl casts a spell from source to target with a given level.
//
// Example:
//
//	ns.CastSpell(spell.DEATH_RAY, 1, ns.Object("CruelDude"), ns.GetHost())
//	ns.CastSpell(spell.DEATH_RAY, 3, ns.Ptf(10, 5), ns.Waypoint("Target"))
func CastSpellLvl(spell spell.Spell, lvl int, source, target Positioner) {
	if impl == nil {
		return
	}
	impl.CastSpellLvl(spell, lvl, source, target)
}

// NewTrap creates a trap at a given point.
//
// Example:
//
//	ns.NewTrap(ns.Ptf(10, 5), spell.DEATH_RAY)
//	ns.NewTrap(ns.GetHost(), spell.DEATH_RAY, spell.SPELL_CONFUSE)
func NewTrap(pos Positioner, spells ...spell.Spell) Obj {
	if impl == nil {
		return nil
	}
	out := make([]TrapSpell, 0, len(spells))
	for _, sp := range spells {
		out = append(out, TrapSpell{Spell: sp})
	}
	return impl.NewTrap(pos, out)
}

type TrapSpell struct {
	Spell spell.Spell
	Level int
}

// NewTrapAdv creates a trap at a given point. It's an advanced version of NewTrap.
//
// Example:
//
//		ns.NewTrapAdv(ns.Ptf(10, 5), ns.TrapSpell{Spell: spell.DEATH_RAY, Level: 3})
//		ns.NewTrapAdv(ns.GetHost(), []ns.TrapSpell{
//	 		{Spell: spell.DEATH_RAY, Level: 1},
//	 		{Spell: spell.SPELL_CONFUSE}, // default level
//		}...)
func NewTrapAdv(pos Positioner, spells []TrapSpell) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewTrap(pos, spells)
}
