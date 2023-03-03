package ns

import (
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
	"github.com/noxworld-dev/noxscript/ns/v4/enchant"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
)

// AwardSpell award spell level to object.
//
// This will raise the spell level of the object. If the object can not cast
// this spell then it will have no effect.
func AwardSpell(id ObjectID, spellID Spell) bool {
	return ns4.AsObj(id).AwardSpell(spell.Spell(spellID))
}

// GroupAwardSpell awards spell level to objects in a group.
//
// This will raise the spell level of the objects in the group. If an object can
// not cast this spell then it will have no effect on that object.
func GroupAwardSpell(objectGroup ObjectGroupID, spellID Spell) {
	ns4.AsObjGroup(objectGroup).AwardSpell(spell.Spell(spellID))
}

// HasEnchant gets whether object has an enchant.
func HasEnchant(id ObjectID, enchantID EnchantID) bool {
	return ns4.AsObj(id).HasEnchant(enchant.Enchant(enchantID))
}

// Enchant grants object an enchantment.
//
// This will grant an object an enchantment of a specified duration.
func Enchant(id ObjectID, enchantID EnchantID, duration float32) {
	ns4.AsObj(id).Enchant(enchant.Enchant(enchantID), timeSeconds(duration))
}

// EnchantOff removes enchant from an object.
func EnchantOff(id ObjectID, enchantID EnchantID) {
	ns4.AsObj(id).EnchantOff(enchant.Enchant(enchantID))
}

// GroupEnchant grants objects in a group an enchantment.
//
// This will grant the objects in a group an enchantment of a specified
// duration.
func GroupEnchant(id ObjectGroupID, enchantID EnchantID, duration float32) {
	ns4.AsObjGroup(id).Enchant(enchant.Enchant(enchantID), timeSeconds(duration))
}

// CastSpellObjectObject casts a spell from source to target.
//
// Example usage:
//
//	ns.CastSpellOn("SPELL_DEATH_RAY", ns.Object("CruelDude"), ns.GetHost())
func CastSpellObjectObject(spellID Spell, source ObjectID, target ObjectID) {
	ns4.CastSpell(spell.Spell(spellID), ns4.AsObj(source), ns4.AsObj(target))
}

// CastSpellObjectLocation casts a spell from source to target.
func CastSpellObjectLocation(spellID Spell, source ObjectID, x, y float32) {
	ns4.CastSpell(spell.Spell(spellID), ns4.AsObj(source), ns4.Ptf(x, y))
}

// CastSpellLocationObject casts a spell from source to target.
func CastSpellLocationObject(spellID Spell, x, y float32, target ObjectID) {
	ns4.CastSpell(spell.Spell(spellID), ns4.Ptf(x, y), ns4.AsObj(target))
}

// CastSpellLocationLocation casts a spell from source to target.
func CastSpellLocationLocation(spellID Spell, x1, y1, x2, y2 float32) {
	ns4.CastSpell(spell.Spell(spellID), ns4.Ptf(x1, y1), ns4.Ptf(x2, y2))
}

// Effect trigger an effect
//
// This will trigger an effect from point (x1,y1) to (x2,y2). Some effects only
// have one point, in which case (x2,y2) is ignored.
func Effect(effectID EffectID, x1, y1, x2, y2 float32) {
	ns4.Effect(effect.Effect(effectID), ns4.Ptf(x1, y1), ns4.Ptf(x2, y2))
}
