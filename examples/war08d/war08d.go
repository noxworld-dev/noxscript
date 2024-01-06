package war08d

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	gvar5  ns.ObjectGroupID
	gvar6  ns.ObjectGroupID
	gvar7  ns.ObjectGroupID
	gvar8  ns.ObjectGroupID
	wp9    ns.WaypointID
	ivar10 int
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	obj14  ns.ObjectID
	obj15  ns.ObjectID
	obj16  ns.ObjectID
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	wp25   ns.WaypointID
	wp26   ns.WaypointID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
	obj30  ns.ObjectID
	obj31  ns.ObjectID
	obj32  ns.ObjectID
	obj33  ns.ObjectID
	obj34  ns.ObjectID
	obj35  ns.ObjectID
	obj36  ns.ObjectID
	obj37  ns.ObjectID
	obj38  ns.ObjectID
	obj39  ns.ObjectID
	obj40  ns.ObjectID
	fvar41 float32
	fvar42 float32
	ivar43 int
	str44  [3]string
	obj45  ns.ObjectID
	obj46  ns.ObjectID
	obj47  ns.ObjectID
	wp48   ns.WaypointID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	flag53 bool
	obj54  ns.ObjectID
	obj55  ns.ObjectID
	gvar56 ns.WallGroupID
	obj57  ns.ObjectID
	obj58  ns.ObjectID
	obj59  ns.ObjectID
	obj60  ns.ObjectID
	obj61  ns.ObjectID
	obj62  ns.ObjectID
	obj63  ns.ObjectID
	obj64  ns.ObjectID
	obj65  ns.ObjectID
	obj66  ns.ObjectID
	obj67  ns.ObjectID
	gvar68 ns.ObjectGroupID
	wp69   ns.WaypointID
	wp70   ns.WaypointID
	wp71   ns.WaypointID
	wp72   ns.WaypointID
	wp73   ns.WaypointID
	wp74   ns.WaypointID
	wp75   ns.WaypointID
	obj76  ns.ObjectID
	obj77  [15]ns.ObjectID
	gvar78 [21]ns.ObjectGroupID
	gvar79 int
	ivar80 int
	ivar81 int
	ivar82 int
	ivar83 int
	ivar84 int
	ivar85 [21]int
	flag86 [21]bool
	wp87   [21]ns.WaypointID
	str88  [6]string
)

func init() {
	ivar10 = 0
	fvar41 = 0
	fvar42 = 0
	ivar43 = 0
	flag53 = false
	gvar79 = 0
	ivar80 = 1
	ivar81 = 0
	ivar82 = 10
	ivar83 = 180
	ivar84 = 0
}
func ArrowTrapAudio() {
	ns.MoveWaypoint(wp9, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BowShoot, wp9)
	ns.AudioEvent(ns.CrossBowShoot, wp9)
}
func DisableArrowTrap01() {
	ns.ObjectGroupOff(gvar5)
}
func ActivateArrowTrap01() {
	ns.ObjectGroupOn(gvar5)
	ns.FrameTimer(1, DisableArrowTrap01)
}
func DisableArrowTrap02() {
	ns.ObjectGroupOff(gvar6)
}
func ActivateArrowTrap02() {
	ns.ObjectGroupOn(gvar6)
	ns.FrameTimer(1, DisableArrowTrap02)
}
func DisableArrowTrap03() {
	ns.ObjectGroupOff(gvar7)
}
func ActivateArrowTrap03() {
	ns.ObjectGroupOn(gvar7)
	ns.FrameTimer(1, DisableArrowTrap03)
}
func DisableUltimateArrowTrapGuns() {
	ns.ObjectGroupOff(gvar8)
}
func ResetUltimateArrowTrap() {
	ns.ObjectOn(obj4)
}
func UltimateArrowTrapLoop() {
	if !(ivar10 < 8) {
		goto LABEL1
	}
	ivar10 += 1
	ns.ObjectGroupOn(gvar8)
	ns.FrameTimer(1, DisableUltimateArrowTrapGuns)
	ns.FrameTimer(10, UltimateArrowTrapLoop)
LABEL1:
	if ivar10 != 8 {
		goto LABEL2
	}
	ivar10 = 0
	ns.ObjectGroupOn(gvar8)
	ns.FrameTimer(1, DisableUltimateArrowTrapGuns)
	ns.FrameTimer(10, ResetUltimateArrowTrap)
LABEL2:
	return
}
func ActivateUltimateArrowTrap() {
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectGroupOn(gvar8)
	ns.FrameTimer(1, DisableUltimateArrowTrapGuns)
	ns.FrameTimer(10, UltimateArrowTrapLoop)
}
func InitializeBeholders() {
	str44[0] = "SPELL_STUN"
	str44[1] = "SPELL_CONFUSE"
	str44[2] = "SPELL_FUMBLE"
}
func BeholderIdle() {
	ns.Wander(ns.GetTrigger())
}
func Beholder01Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj11) > 0 && ns.CurrentHealth(obj27) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj11)
	v1 = ns.GetObjectY(obj11)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj27 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj27, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj27 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj27, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj27 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj27, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj11, obj27)
	ns.CreatureFollow(obj27, obj11)
LABEL1:
	return
}
func Beholder01Recognize() {
	ns.FrameTimer(90, Beholder01Bombers)
}
func Beholder01Die() {
	ns.Damage(obj27, 0, 500, 0)
}
func Beholder02Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj12) > 0 && ns.CurrentHealth(obj28) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj12)
	v1 = ns.GetObjectY(obj12)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj28 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj28, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj28 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj28, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj28 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj28, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj12, obj28)
	ns.CreatureFollow(obj28, obj12)
LABEL1:
	return
}
func Beholder02Recognize() {
	ns.FrameTimer(90, Beholder02Bombers)
}
func Beholder02Die() {
	ns.Damage(obj28, 0, 500, 0)
}
func Beholder03Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj13) > 0 && ns.CurrentHealth(obj29) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj13)
	v1 = ns.GetObjectY(obj13)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj29 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj29, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj29 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj29, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj29 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj29, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj13, obj29)
	ns.CreatureFollow(obj29, obj13)
LABEL1:
	return
}
func Beholder03Recognize() {
	ns.FrameTimer(90, Beholder03Bombers)
}
func Beholder03Die() {
	ns.Damage(obj29, 0, 500, 0)
}
func Beholder04Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj14) > 0 && ns.CurrentHealth(obj30) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj14)
	v1 = ns.GetObjectY(obj14)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj30 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj30, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj30 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj30, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj30 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj30, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj14, obj30)
	ns.CreatureFollow(obj30, obj14)
LABEL1:
	return
}
func Beholder04Recognize() {
	ns.FrameTimer(90, Beholder04Bombers)
}
func Beholder04Die() {
	ns.Damage(obj30, 0, 500, 0)
}
func Beholder05Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj15) > 0 && ns.CurrentHealth(obj31) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj15)
	v1 = ns.GetObjectY(obj15)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj31 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj31, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj31 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj31, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj31 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj31, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj15, obj31)
	ns.CreatureFollow(obj31, obj15)
LABEL1:
	return
}
func Beholder05Recognize() {
	ns.FrameTimer(90, Beholder05Bombers)
}
func Beholder05Die() {
	ns.Damage(obj31, 0, 500, 0)
}
func Beholder06Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj16) > 0 && ns.CurrentHealth(obj32) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj16)
	v1 = ns.GetObjectY(obj16)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj32 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj32, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj32 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj32, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj32 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj32, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj16, obj32)
	ns.CreatureFollow(obj32, obj16)
LABEL1:
	return
}
func Beholder06Recognize() {
	ns.FrameTimer(90, Beholder06Bombers)
}
func Beholder06Die() {
	ns.Damage(obj32, 0, 500, 0)
}
func Beholder07Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj17) > 0 && ns.CurrentHealth(obj33) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj17)
	v1 = ns.GetObjectY(obj17)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj33 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj33, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj33 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj33, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj33 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj33, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj17, obj33)
	ns.CreatureFollow(obj33, obj17)
LABEL1:
	return
}
func Beholder07Recognize() {
	ns.FrameTimer(90, Beholder07Bombers)
}
func Beholder07Die() {
	ns.Damage(obj33, 0, 500, 0)
}
func Beholder08Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj18) > 0 && ns.CurrentHealth(obj34) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj18)
	v1 = ns.GetObjectY(obj18)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj34 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj34, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj34 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj34, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj34 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj34, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj18, obj34)
	ns.CreatureFollow(obj34, obj18)
LABEL1:
	return
}
func Beholder08Recognize() {
	ns.FrameTimer(90, Beholder08Bombers)
}
func Beholder08Die() {
	ns.Damage(obj34, 0, 500, 0)
}
func Beholder09Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj19) > 0 && ns.CurrentHealth(obj35) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj19)
	v1 = ns.GetObjectY(obj19)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj35 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj35, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj35 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj35, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj35 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj35, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj19, obj35)
	ns.CreatureFollow(obj35, obj19)
LABEL1:
	return
}
func Beholder09Recognize() {
	ns.FrameTimer(90, Beholder09Bombers)
}
func Beholder09Die() {
	ns.Damage(obj35, 0, 500, 0)
}
func Beholder10Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj20) > 0 && ns.CurrentHealth(obj36) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj20)
	v1 = ns.GetObjectY(obj20)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj36 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj36, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj36 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj36, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj36 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj36, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj20, obj36)
	ns.CreatureFollow(obj36, obj20)
LABEL1:
	return
}
func Beholder10Recognize() {
	ns.FrameTimer(90, Beholder10Bombers)
}
func Beholder10Die() {
	ns.Damage(obj36, 0, 500, 0)
}
func Beholder11Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj21) > 0 && ns.CurrentHealth(obj37) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj21)
	v1 = ns.GetObjectY(obj21)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj37 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj37, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj37 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj37, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj37 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj37, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj21, obj37)
	ns.CreatureFollow(obj37, obj21)
LABEL1:
	return
}
func Beholder11Recognize() {
	ns.FrameTimer(90, Beholder11Bombers)
}
func Beholder11Die() {
	ns.Damage(obj37, 0, 500, 0)
}
func Beholder12Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj22) > 0 && ns.CurrentHealth(obj38) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj22)
	v1 = ns.GetObjectY(obj22)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj38 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj38, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj38 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj38, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj38 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj38, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj22, obj38)
	ns.CreatureFollow(obj38, obj22)
LABEL1:
	return
}
func Beholder12Recognize() {
	ns.FrameTimer(90, Beholder12Bombers)
}
func Beholder12Die() {
	ns.Damage(obj38, 0, 500, 0)
}
func Beholder13Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj23) > 0 && ns.CurrentHealth(obj39) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj23)
	v1 = ns.GetObjectY(obj23)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj39 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj39, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj39 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj39, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj39 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj39, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj23, obj39)
	ns.CreatureFollow(obj39, obj23)
LABEL1:
	return
}
func Beholder13Recognize() {
	ns.FrameTimer(90, Beholder13Bombers)
}
func Beholder13Die() {
	ns.Damage(obj39, 0, 500, 0)
}
func Beholder14Bombers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	_ = v3
	_ = v2
	if !(ns.CurrentHealth(obj24) > 0 && ns.CurrentHealth(obj40) <= 0) {
		goto LABEL1
	}
	v0 = ns.GetObjectX(obj24)
	v1 = ns.GetObjectY(obj24)
	ns.MoveWaypoint(wp26, v0, v1)
	ns.AudioEvent(ns.SummonCast, wp26)
	v2 = ns.RandomFloat(fvar41, fvar42)
	v3 = ns.RandomFloat(fvar41, fvar42)
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ivar43 = ns.Random(0, 2)
	v4 = ivar43
	if v4 == 0 {
		goto LABEL2
	}
	if v4 == 1 {
		goto LABEL3
	}
	if v4 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj40 = ns.CreateObject("BomberBlue", wp25)
	ns.TrapSpells(obj40, str44[ivar43], "", "")
	goto LABEL5
LABEL3:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj40 = ns.CreateObject("BomberGreen", wp25)
	ns.TrapSpells(obj40, str44[ivar43], "", "")
	goto LABEL5
LABEL4:
	ns.MoveWaypoint(wp25, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	obj40 = ns.CreateObject("BomberYellow", wp25)
	ns.TrapSpells(obj40, str44[ivar43], "", "")
	goto LABEL5
LABEL5:
	ns.SetOwner(obj24, obj40)
	ns.CreatureFollow(obj40, obj24)
LABEL1:
	return
}
func Beholder14Recognize() {
	ns.FrameTimer(90, Beholder14Bombers)
}
func Beholder14Die() {
	ns.Damage(obj40, 0, 500, 0)
}
func OpenBlockPassage() {
	ns.ObjectOff(ns.GetTrigger())
	ns.AudioEvent(ns.SpikeBlockMove, wp52)
	ns.Move(obj45, wp48)
	ns.Move(obj46, wp49)
	ns.Move(obj47, wp51)
}
func SealBlockRoom() {
	ns.ObjectOff(ns.GetTrigger())
	ns.AudioEvent(ns.SpikeBlockMove, wp50)
	ns.Move(obj46, wp50)
}
func CagedAnimalDie() {
	if flag53 {
		goto LABEL1
	}
	flag53 = true
	ns.WallGroupBreak(gvar56)
LABEL1:
	return
}
func UnlockBearCage() {
	ns.ObjectOff(ns.GetTrigger())
	ns.UnlockDoor(obj54)
	ns.UnlockDoor(obj55)
}
func PlayerDeath() {
	ns.DeathScreen(8)
}
func FireFONTrap01() {
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj57, ns.GetWaypointX(wp69), ns.GetWaypointY(wp69))
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj58, ns.GetWaypointX(wp70), ns.GetWaypointY(wp70))
	ns.SecondTimer(5, FireFONTrap01)
}
func EndFONSetpiece() {
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
}
func BearAttack() {
	ns.EnchantOff(obj62, ns.ENCHANT_INVULNERABLE)
	ns.AggressionLevel(obj62, 0.83)
	ns.Attack(obj62, ns.GetHost())
	ns.FrameTimer(45, FireFONTrap01)
}
func InitializeFONtraps() {
	obj57 = ns.Object("FON_Origin01")
	obj58 = ns.Object("FON_Origin02")
	obj59 = ns.Object("FON_Origin03")
	obj60 = ns.Object("FON_Origin04")
	obj61 = ns.Object("FON_Origin05")
	obj62 = ns.Object("Bear01")
	obj63 = ns.Object("SpikeTrapDoor01")
	obj64 = ns.Object("SpikeTrapDoor02")
	obj65 = ns.Object("SpikeTrapDoor03")
	obj66 = ns.Object("SpikeTrapDoor04")
	gvar68 = ns.ObjectGroup("SpikeGroup01")
	wp69 = ns.Waypoint("FON_Target01")
	wp70 = ns.Waypoint("FON_Target02")
	wp71 = ns.Waypoint("FON_Target03")
	wp72 = ns.Waypoint("FON_Target04")
	wp73 = ns.Waypoint("FON_Target05")
	wp74 = ns.Waypoint("SpikeTrapAudioOrigin")
	ns.Enchant(obj62, ns.ENCHANT_INVULNERABLE, 0)
}
func StartFONSetpiece() {
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.CreatureIdle(ns.GetHost())
	BearAttack()
}
func BearDie() {
	ns.FrameTimer(60, EndFONSetpiece)
}
func BearInjured() {
	ns.Damage(ns.GetTrigger(), 0, 300, 0)
}
func FireFONTrap02() {
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj59, ns.GetWaypointX(wp71), ns.GetWaypointY(wp71))
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj60, ns.GetWaypointX(wp72), ns.GetWaypointY(wp72))
}
func EnableSpikeTrap() {
	ns.ObjectGroupOn(gvar68)
	ns.AudioEvent(ns.FloorSpikesUp, wp74)
	ns.LockDoor(obj63)
	ns.LockDoor(obj64)
	ns.LockDoor(obj65)
	ns.LockDoor(obj66)
	ns.ObjectOn(obj67)
}
func DisableSpikeTrap() {
	ns.ObjectGroupOff(gvar68)
	ns.AudioEvent(ns.FloorSpikesDown, wp74)
	ns.UnlockDoor(obj63)
	ns.UnlockDoor(obj64)
	ns.UnlockDoor(obj65)
	ns.UnlockDoor(obj66)
	ns.ObjectOff(obj67)
}
func InitializeSummoningCircles() {
	var v0 int
	v0 = 0
	for {
		if !(v0 < 21) {
			goto LABEL1
		}
		flag86[v0] = true
		ivar85[v0] = 0
		v0 += 1
	}
LABEL1:
	str88[0] = "UrchinShaman"
	str88[1] = "Mimic"
	str88[2] = "Bear"
	str88[3] = "Urchin"
	str88[4] = "Imp"
	str88[5] = "WhiteWolf"
	obj76 = ns.Object("SummoningCircleTreasureLight")
	obj77[0] = ns.Object("SummoningCircleTreasure01")
	obj77[1] = ns.Object("SummoningCircleTreasure02")
	obj77[2] = ns.Object("SummoningCircleTreasure03")
	obj77[3] = ns.Object("SummoningCircleTreasure04")
	obj77[4] = ns.Object("SummoningCircleTreasure05")
	obj77[5] = ns.Object("SummoningCircleTreasure06")
	obj77[6] = ns.Object("SummoningCircleTreasure07")
	obj77[7] = ns.Object("SummoningCircleTreasure08")
	obj77[8] = ns.Object("SummoningCircleTreasure09")
	obj77[9] = ns.Object("SummoningCircleTreasure10")
	obj77[10] = ns.Object("SummoningCircleTreasure11")
	obj77[11] = ns.Object("SummoningCircleTreasure12")
	obj77[12] = ns.Object("SummoningCircleTreasure13")
	obj77[13] = ns.Object("SummoningCircleTreasure14")
	obj77[14] = ns.Object("SummoningCircleTreasure15")
	wp87[0] = ns.Waypoint("SummoningCircle01WP")
	wp87[1] = ns.Waypoint("SummoningCircle02WP")
	wp87[2] = ns.Waypoint("SummoningCircle03WP")
	wp87[3] = ns.Waypoint("SummoningCircle04WP")
	wp87[4] = ns.Waypoint("SummoningCircle05WP")
	wp87[5] = ns.Waypoint("SummoningCircle06WP")
	wp87[6] = ns.Waypoint("SummoningCircle07WP")
	wp87[7] = ns.Waypoint("SummoningCircle08WP")
	wp87[8] = ns.Waypoint("SummoningCircle09WP")
	wp87[9] = ns.Waypoint("SummoningCircle10WP")
	wp87[10] = ns.Waypoint("SummoningCircle11WP")
	wp87[11] = ns.Waypoint("SummoningCircle12WP")
	wp87[12] = ns.Waypoint("SummoningCircle13WP")
	wp87[13] = ns.Waypoint("SummoningCircle14WP")
	wp87[14] = ns.Waypoint("SummoningCircle15WP")
	wp87[15] = ns.Waypoint("SummoningCircle16WP")
	wp87[16] = ns.Waypoint("SummoningCircle17WP")
	wp87[17] = ns.Waypoint("SummoningCircle18WP")
	wp87[18] = ns.Waypoint("SummoningCircle19WP")
	wp87[19] = ns.Waypoint("SummoningCircle20WP")
	wp87[20] = ns.Waypoint("SummoningCircle21WP")
	gvar78[0] = ns.ObjectGroup("SummonCircle01Lights")
	gvar78[1] = ns.ObjectGroup("SummonCircle02Lights")
	gvar78[2] = ns.ObjectGroup("SummonCircle03Lights")
	gvar78[3] = ns.ObjectGroup("SummonCircle04Lights")
	gvar78[4] = ns.ObjectGroup("SummonCircle05Lights")
	gvar78[5] = ns.ObjectGroup("SummonCircle06Lights")
	gvar78[6] = ns.ObjectGroup("SummonCircle07Lights")
	gvar78[7] = ns.ObjectGroup("SummonCircle08Lights")
	gvar78[8] = ns.ObjectGroup("SummonCircle09Lights")
	gvar78[9] = ns.ObjectGroup("SummonCircle10Lights")
	gvar78[10] = ns.ObjectGroup("SummonCircle11Lights")
	gvar78[11] = ns.ObjectGroup("SummonCircle12Lights")
	gvar78[12] = ns.ObjectGroup("SummonCircle13Lights")
	gvar78[13] = ns.ObjectGroup("SummonCircle14Lights")
	gvar78[14] = ns.ObjectGroup("SummonCircle15Lights")
	gvar78[15] = ns.ObjectGroup("SummonCircle16Lights")
	gvar78[16] = ns.ObjectGroup("SummonCircle17Lights")
	gvar78[17] = ns.ObjectGroup("SummonCircle18Lights")
	gvar78[18] = ns.ObjectGroup("SummonCircle19Lights")
	gvar78[19] = ns.ObjectGroup("SummonCircle20Lights")
	gvar78[20] = ns.ObjectGroup("SummonCircle21Lights")
}
func MapInitialize() {
	obj45 = ns.Object("StoneBlock01")
	obj46 = ns.Object("StoneBlock02")
	obj47 = ns.Object("StoneBlock03")
	obj11 = ns.Object("Beholder01")
	obj12 = ns.Object("Beholder02")
	obj13 = ns.Object("Beholder03")
	obj14 = ns.Object("Beholder04")
	obj15 = ns.Object("Beholder05")
	obj16 = ns.Object("Beholder06")
	obj17 = ns.Object("Beholder07")
	obj18 = ns.Object("Beholder08")
	obj19 = ns.Object("Beholder09")
	obj20 = ns.Object("Beholder10")
	obj21 = ns.Object("Beholder11")
	obj22 = ns.Object("Beholder12")
	obj23 = ns.Object("Beholder13")
	obj24 = ns.Object("Beholder14")
	obj54 = ns.Object("CageGateL")
	obj55 = ns.Object("CageGateR")
	obj4 = ns.Object("UltimateArrowTrapSwitch")
	wp9 = ns.Waypoint("ArrowTrapAudioOrigin")
	obj67 = ns.Object("SpikeSFX")
	wp48 = ns.Waypoint("StoneBlock01WP")
	wp49 = ns.Waypoint("StoneBlock02WP")
	wp50 = ns.Waypoint("StoneBlock02WP2")
	wp51 = ns.Waypoint("StoneBlock03WP")
	wp52 = ns.Waypoint("StoneBlockAudioOrigin")
	wp25 = ns.Waypoint("BomberCreationWP")
	wp26 = ns.Waypoint("BomberAudioOrigin")
	wp75 = ns.Waypoint("SecretAudioOrigin")
	gvar5 = ns.ObjectGroup("ArrowTrapGroup01")
	gvar6 = ns.ObjectGroup("ArrowTrapGroup02")
	gvar7 = ns.ObjectGroup("ArrowTrapGroup03")
	gvar8 = ns.ObjectGroup("ArrowTrapGroup04")
	gvar56 = ns.WallGroup("CageWalls")
	ns.Wander(obj11)
	ns.Wander(obj12)
	ns.Wander(obj13)
	ns.AggressionLevel(obj13, 0.83)
	ns.Wander(obj14)
	ns.AggressionLevel(obj14, 0.83)
	ns.Wander(obj15)
	ns.AggressionLevel(obj15, 0.83)
	ns.Wander(obj16)
	ns.AggressionLevel(obj16, 0.83)
	ns.Wander(obj17)
	ns.AggressionLevel(obj17, 0.83)
	ns.Wander(obj18)
	ns.AggressionLevel(obj18, 0.83)
	ns.Wander(obj19)
	ns.AggressionLevel(obj19, 0.83)
	ns.Wander(obj20)
	ns.AggressionLevel(obj20, 0.83)
	ns.Wander(obj21)
	ns.AggressionLevel(obj21, 0.83)
	ns.Wander(obj22)
	ns.AggressionLevel(obj22, 0.83)
	ns.LockDoor(obj54)
	ns.LockDoor(obj55)
	InitializeBeholders()
	InitializeSummoningCircles()
	InitializeFONtraps()
}
func OpenPitElevatorWall() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallOpen(ns.Wall(56, 134))
}
func MonstersGoHome() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func SecretSFX() {
	ns.MoveWaypoint(wp75, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp75)
}
func BoulderSecretFound() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 200)
	SecretSFX()
}
func OpenBoulderSecret() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallOpen(ns.Wall(104, 40))
	ns.AudioEvent(ns.SecretWallOpen, ns.Waypoint("BoulderSecretAudioOrigin"))
	SecretSFX()
}
func SecretFound() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
	SecretSFX()
}
func ResetSummonCircle(a1 int) {
	flag86[a1] = true
	ns.ObjectGroupOn(gvar78[a1])
}
func SummoningCircleCleared(a1 int) {
	ivar84 += 1
	if ivar84 != 15 {
		goto LABEL1
	}
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 500)
	ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(wp87[a1]), ns.GetWaypointY(wp87[a1]), 0, 0)
	ns.MoveObject(obj76, ns.GetWaypointX(wp87[a1]), ns.GetWaypointY(wp87[a1]))
	ns.MoveObject(obj77[ivar84-1], ns.GetWaypointX(wp87[a1]), ns.GetWaypointY(wp87[a1]))
	ns.AudioEvent(ns.SmallGong, wp87[a1])
	goto LABEL2
LABEL1:
	ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(wp87[a1]), ns.GetWaypointY(wp87[a1]), 0, 0)
	ns.MoveObject(obj76, ns.GetWaypointX(wp87[a1]), ns.GetWaypointY(wp87[a1]))
	ns.MoveObject(obj77[ivar84-1], ns.GetWaypointX(wp87[a1]), ns.GetWaypointY(wp87[a1]))
	ns.AudioEvent(ns.BigGong, wp87[a1])
LABEL2:
	return
}
func SummonCircle01Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 0
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle02Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 1
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle03Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 2
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle04Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 3
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle05Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 4
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle06Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 5
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle07Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 6
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle08Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 7
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle09Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 8
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle10Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 9
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle11Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 10
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle12Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 11
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle13Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 12
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle14Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 13
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle15Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 14
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle16Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 15
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle17Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 16
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle18Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 17
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle19Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 18
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle20Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 19
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle21Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar79 = 20
	if !(ns.IsCaller(ns.GetHost()) && flag86[gvar79] == true) {
		goto LABEL1
	}
	flag86[gvar79] = false
	ns.ObjectGroupOff(gvar78[gvar79])
	ns.AudioEvent(ns.SummonCast, wp87[gvar79])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp87[gvar79]), ns.GetWaypointY(wp87[gvar79]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar82) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp87[gvar79])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp87[gvar79])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar81 = ns.Random(0, 5)
	ns.CreateObject(str88[ivar81], wp87[gvar79])
LABEL6:
	ivar85[gvar79] += 1
	if !(ivar85[gvar79] > ivar80) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar79, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar83, gvar79, ResetSummonCircle)
LABEL1:
	return
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
