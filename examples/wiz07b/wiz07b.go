package wiz07b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	wp7    ns.WaypointID
	wp8    ns.WaypointID
	wp9    ns.WaypointID
	wp10   ns.WaypointID
	wp11   ns.WaypointID
	gvar12 ns.WallGroupID
	gvar13 ns.ObjectGroupID
	gvar14 ns.ObjectGroupID
	gvar15 ns.ObjectGroupID
	gvar16 ns.ObjectGroupID
	wp17   ns.WaypointID
	wp18   ns.WaypointID
	wp19   ns.WaypointID
	wp20   ns.WaypointID
	wp21   ns.WaypointID
	wp22   ns.WaypointID
	wp23   [21]ns.WaypointID
	wp24   [21]ns.WaypointID
	wp25   [21]ns.WaypointID
	gvar26 ns.WallGroupID
	gvar27 ns.ObjectGroupID
	fvar28 float32
	fvar29 float32
	fvar30 float32
	gvar31 int
	ivar32 int
	flag33 bool
)

func init() {
	fvar28 = 8
	fvar29 = 20
	gvar31 = 5
	ivar32 = 600
	flag33 = true
}
func BustEm2() {
	ns.WallGroupBreak(gvar12)
}
func RegenA() {
	if ns.Object("RegenD1") != 0 {
		goto LABEL1
	}
	obj4 = ns.CreateObject("EmberDemon", wp7)
	ns.CreatureHunt(obj4)
LABEL1:
	if ns.Object("RegenD2") != 0 {
		goto LABEL2
	}
	obj5 = ns.CreateObject("EmberDemon", wp8)
	ns.CreatureHunt(obj5)
LABEL2:
	return
}
func RegenB() {
	if ns.Object("RegenD3") != 0 {
		goto LABEL1
	}
	obj6 = ns.CreateObject("EmberDemon", wp9)
	ns.CreatureHunt(obj6)
LABEL1:
	return
}
func AutoSave1() {
	ns.AutoSave()
}
func BustEm() {
	ns.AudioEvent(ns.MagicMissileDetonate, wp11)
	ns.AudioEvent(ns.WallDestroyed, wp10)
	ns.Effect(ns.JIGGLE, 835, 2650, 13, 0)
	ns.ObjectGroupOff(gvar13)
	ns.FrameTimer(7, BustEm2)
}
func Init2() {
	ns.Frozen(ns.GetHost(), false)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp22), ns.GetWaypointY(wp22), 0, 0)
	ns.AudioEvent(ns.TeleportIn, wp22)
	if !ns.HasEnchant(ns.GetHost(), ns.ENCHANT_INVISIBLE) {
		goto LABEL1
	}
	ns.EnchantOff(ns.GetHost(), ns.ENCHANT_INVISIBLE)
LABEL1:
	ns.PrintToAll("Con03C.scr:NewTask")
	ns.JournalEntry(ns.GetHost(), "Wiz07EscapeUnderworld", 2)
}
func Quake() {
	var (
		v0 float32
		v1 float32
		v6 int
	)
	if flag33 {
		goto LABEL1
	}
	return
LABEL1:
	v0 = ns.GetObjectX(ns.GetHost())
	v1 = ns.GetObjectY(ns.GetHost())
	fvar30 = ns.RandomFloat(fvar28, fvar29)
	ns.MoveWaypoint(wp11, v0, v1)
	ns.AudioEvent(ns.BoulderMove, wp11)
	ns.AudioEvent(ns.EarthRumbleMinor, wp11)
	ns.Effect(ns.JIGGLE, v0, v1, fvar30, 0)
	v6 = ns.Random(5, ivar32)
	ns.FrameTimer(v6, Quake)
}
func PlaySubMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(18, 100)
LABEL1:
	return
}
func MapInitialize() {
	gvar16 = ns.ObjectGroup("MDGroup1")
	gvar14 = ns.ObjectGroup("BatGroup1")
	gvar15 = ns.ObjectGroup("ImpGroup1")
	wp17 = ns.Waypoint("Secret01WP")
	wp18 = ns.Waypoint("Secret02WP")
	wp19 = ns.Waypoint("Secret03WP")
	wp20 = ns.Waypoint("Secret04WP")
	wp11 = ns.Waypoint("QuakeAudioOrigin")
	wp21 = ns.Waypoint("CreationWP")
	wp22 = ns.Waypoint("StartOrigin")
	gvar26 = ns.WallGroup("Walls1")
	gvar12 = ns.WallGroup("BustEmWalls")
	obj4 = ns.Object("RegenD1")
	obj5 = ns.Object("RegenD2")
	obj6 = ns.Object("RegenD3")
	wp10 = ns.Waypoint("BustAudioOrigin")
	wp7 = ns.Waypoint("RegenLoc1")
	wp8 = ns.Waypoint("RegenLoc2")
	wp9 = ns.Waypoint("RegenLoc3")
	wp23[0] = ns.Waypoint("Flamer1")
	wp23[1] = ns.Waypoint("Flamer2")
	wp23[2] = ns.Waypoint("Flamer3")
	wp23[3] = ns.Waypoint("Flamer4")
	wp23[4] = ns.Waypoint("Flamer5")
	wp23[5] = ns.Waypoint("Flamer6")
	wp23[6] = ns.Waypoint("Flamer7")
	wp23[7] = ns.Waypoint("Flamer8")
	wp23[8] = ns.Waypoint("Flamer9")
	wp23[9] = ns.Waypoint("Flamer10")
	wp23[10] = ns.Waypoint("Flamer11")
	wp23[11] = ns.Waypoint("Flamer12")
	wp23[12] = ns.Waypoint("Flamer13")
	wp23[13] = ns.Waypoint("Flamer14")
	wp23[14] = ns.Waypoint("Flamer15")
	wp23[15] = ns.Waypoint("Flamer16")
	wp23[16] = ns.Waypoint("Flamer17")
	wp23[17] = ns.Waypoint("Flamer18")
	wp23[18] = ns.Waypoint("Flamer19")
	wp23[19] = ns.Waypoint("Flamer20")
	wp23[20] = ns.Waypoint("Flamer21")
	wp24[0] = ns.Waypoint("FlamerB1")
	wp24[1] = ns.Waypoint("FlamerB2")
	wp24[2] = ns.Waypoint("FlamerB3")
	wp24[3] = ns.Waypoint("FlamerB4")
	wp24[4] = ns.Waypoint("FlamerB5")
	wp24[5] = ns.Waypoint("FlamerB6")
	wp24[6] = ns.Waypoint("FlamerB7")
	wp24[7] = ns.Waypoint("FlamerB8")
	wp24[8] = ns.Waypoint("FlamerB9")
	wp24[9] = ns.Waypoint("FlamerB10")
	wp24[10] = ns.Waypoint("FlamerB11")
	wp24[11] = ns.Waypoint("FlamerB12")
	wp24[12] = ns.Waypoint("FlamerB13")
	wp24[13] = ns.Waypoint("FlamerB14")
	wp24[14] = ns.Waypoint("FlamerB15")
	wp24[15] = ns.Waypoint("FlamerB16")
	wp24[16] = ns.Waypoint("FlamerB17")
	wp24[17] = ns.Waypoint("FlamerB18")
	wp24[18] = ns.Waypoint("FlamerB19")
	wp24[19] = ns.Waypoint("FlamerB20")
	wp24[20] = ns.Waypoint("FlamerB21")
	wp25[0] = ns.Waypoint("FlamerC1")
	wp25[1] = ns.Waypoint("FlamerC2")
	wp25[2] = ns.Waypoint("FlamerC3")
	wp25[3] = ns.Waypoint("FlamerC4")
	wp25[4] = ns.Waypoint("FlamerC5")
	wp25[5] = ns.Waypoint("FlamerC6")
	wp25[6] = ns.Waypoint("FlamerC7")
	wp25[7] = ns.Waypoint("FlamerC8")
	wp25[8] = ns.Waypoint("FlamerC9")
	wp25[9] = ns.Waypoint("FlamerC10")
	wp25[10] = ns.Waypoint("FlamerC11")
	wp25[11] = ns.Waypoint("FlamerC12")
	wp25[12] = ns.Waypoint("FlamerC13")
	wp25[13] = ns.Waypoint("FlamerC14")
	wp25[14] = ns.Waypoint("FlamerC15")
	wp25[15] = ns.Waypoint("FlamerC16")
	wp25[16] = ns.Waypoint("FlamerC17")
	wp25[17] = ns.Waypoint("FlamerC18")
	wp25[18] = ns.Waypoint("FlamerC19")
	wp25[19] = ns.Waypoint("FlamerC20")
	wp25[20] = ns.Waypoint("FlamerC21")
	gvar27 = ns.ObjectGroup("SecretT1")
	gvar13 = ns.ObjectGroup("BustEmT")
	ns.UnBlind()
	Quake()
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(30, Init2)
}
func Bats1Hostilize() {
	ns.GroupAggressionLevel(gvar14, 0.83)
}
func Imps1Hostilize() {
	ns.GroupAggressionLevel(gvar15, 0.83)
}
func MD1Hostilize() {
	ns.GroupAggressionLevel(gvar16, 0.83)
}
func HostilizeMe() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func MapEntry() {
	PlaySubMusic()
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(17, 100)
LABEL1:
	return
}
func PlayTownMusic() {
	ns.Music(15, 100)
}
func WallOpen1() {
	ns.WallGroupOpen(gvar26)
}
func Secret01() {
	ns.AudioEvent(ns.SecretFound, wp17)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret02() {
	ns.ObjectGroupOff(gvar27)
	ns.AudioEvent(ns.SecretFound, wp18)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret03() {
	ns.AudioEvent(ns.SecretFound, wp19)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func Secret04() {
	ns.AudioEvent(ns.SecretFound, wp20)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	case "MapEntry":
		MapEntry()
	}
}
