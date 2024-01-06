package wiz10d

import (
	"github.com/noxworld-dev/noxscript/ns/v3"
	"strconv"
)

var (
	obj4    ns.ObjectID
	obj5    ns.ObjectID
	gvar6   ns.WallGroupID
	gvar7   ns.ObjectGroupID
	gvar8   ns.ObjectGroupID
	gvar9   ns.ObjectGroupID
	gvar10  ns.WallGroupID
	wp11    ns.WaypointID
	wp12    ns.WaypointID
	wp13    ns.WaypointID
	wp14    ns.WaypointID
	gvar15  ns.WallGroupID
	gvar16  ns.ObjectGroupID
	gvar17  ns.ObjectGroupID
	gvar18  ns.ObjectGroupID
	wp19    ns.WaypointID
	wp20    ns.WaypointID
	wp21    ns.WaypointID
	obj22   ns.ObjectID
	wp23    [32]ns.WaypointID
	obj24   ns.ObjectID
	gvar25  ns.WallGroupID
	gvar26  ns.ObjectGroupID
	gvar27  ns.ObjectGroupID
	gvar28  ns.WallGroupID
	gvar29  ns.WallGroupID
	gvar30  ns.ObjectGroupID
	obj31   ns.ObjectID
	obj32   ns.ObjectID
	flag33  bool
	flag34  bool
	wp35    ns.WaypointID
	wp36    ns.WaypointID
	wp37    ns.WaypointID
	wp38    ns.WaypointID
	wp39    ns.WaypointID
	wp40    ns.WaypointID
	wp41    ns.WaypointID
	wp42    ns.WaypointID
	wp43    ns.WaypointID
	wp44    ns.WaypointID
	wp45    ns.WaypointID
	wp46    ns.WaypointID
	wp47    ns.WaypointID
	wp48    ns.WaypointID
	wp49    ns.WaypointID
	wp50    ns.WaypointID
	wp51    ns.WaypointID
	wp52    ns.WaypointID
	obj53   ns.ObjectID
	obj54   ns.ObjectID
	obj55   ns.ObjectID
	obj56   ns.ObjectID
	obj57   ns.ObjectID
	obj58   ns.ObjectID
	gvar59  int
	gvar60  int
	gvar61  int
	gvar62  int
	gvar63  int
	gvar64  int
	gvar65  int
	obj66   ns.ObjectID
	wp67    [6]ns.WaypointID
	gvar68  ns.WallGroupID
	gvar69  int
	gvar70  int
	gvar71  int
	obj72   ns.ObjectID
	obj73   ns.ObjectID
	obj74   ns.ObjectID
	obj75   ns.ObjectID
	wp76    ns.WaypointID
	wp77    ns.WaypointID
	wp78    ns.WaypointID
	wp79    ns.WaypointID
	gvar80  ns.WallGroupID
	gvar81  ns.ObjectGroupID
	gvar82  ns.ObjectGroupID
	gvar83  ns.WallGroupID
	obj84   ns.ObjectID
	obj85   ns.ObjectID
	obj86   ns.ObjectID
	obj87   ns.ObjectID
	obj88   ns.ObjectID
	obj89   ns.ObjectID
	obj90   ns.ObjectID
	obj91   ns.ObjectID
	obj92   ns.ObjectID
	obj93   ns.ObjectID
	obj94   ns.ObjectID
	obj95   ns.ObjectID
	obj96   ns.ObjectID
	obj97   ns.ObjectID
	obj98   ns.ObjectID
	obj99   ns.ObjectID
	obj100  ns.ObjectID
	obj101  ns.ObjectID
	gvar102 ns.ObjectGroupID
	gvar103 ns.ObjectGroupID
	gvar104 ns.ObjectGroupID
	gvar105 ns.ObjectGroupID
	gvar106 ns.ObjectGroupID
	gvar107 ns.ObjectGroupID
	gvar108 ns.WallGroupID
	gvar109 ns.WallGroupID
	gvar110 ns.WallGroupID
	flag111 bool
	flag112 bool
	flag113 bool
	flag114 bool
	wp115   ns.WaypointID
	gvar116 ns.ObjectGroupID
	wp117   ns.WaypointID
	wp118   ns.WaypointID
)

func init() {
	gvar59 = 0
	gvar60 = 1
	gvar61 = 2
	gvar62 = 3
	gvar63 = 4
	gvar64 = 5
	gvar65 = gvar59
	gvar69 = 0
	gvar70 = 1
	gvar71 = gvar69
	flag33 = false
	flag34 = false
	flag111 = false
	flag112 = false
	flag113 = false
	flag114 = false
}
func PlayerDeath() {
	ns.DeathScreen(10)
}
func FilterHecubah() int {
	if !ns.IsCaller(obj5) {
		goto LABEL1
	}
	return 1
	goto LABEL2
LABEL1:
	return 0
LABEL2:
	return 0
}
func ActivateSmokePuffs() {
	var v0 int
	v0 = 0
	for {
		if !(v0 < 32) {
			goto LABEL1
		}
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp23[v0]), ns.GetWaypointY(wp23[v0]), 0, 0)
		v0 += 2
	}
LABEL1:
	return
}
func HecubahGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp13)
	v1 = ns.GetWaypointY(wp13)
	v2 = ns.GetWaypointX(wp14)
	v3 = ns.GetWaypointY(wp14)
	ns.CreatureGuard(obj5, v0, v1, v2, v3, 0)
	ns.NoWallSound(true)
	ns.WallGroupClose(gvar25)
	ns.NoWallSound(false)
}
func EnableOrbMonsterGroup() {
	ns.ObjectGroupOn(gvar7)
	ns.Frozen(ns.GetHost(), false)
	ns.WallGroupOpen(gvar6)
	ns.FrameTimer(2, ActivateSmokePuffs)
	ns.WallGroupClose(gvar10)
}
func HecubahExit() {
	ns.Enchant(obj5, ns.ENCHANT_HASTED, 100)
	ns.Move(obj5, wp13)
	ns.WideScreen(false)
	ns.FrameTimer(60, EnableOrbMonsterGroup)
}
func HecubahDialogEnd() {
	ns.FrameTimer(30, HecubahExit)
	ns.CancelDialog(obj5)
}
func HecubahDialogStart() {
	ns.TellStory(ns.DemonRecognize, "Con07H.scr:HecubahTalk02")
}
func HecubahTrigger() {
	ns.LookAtObject(obj5, ns.GetHost())
	ns.StoryPic(obj5, "HecubahPic")
	ns.SetDialog(obj5, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj5, ns.GetHost())
}
func HecubahEnter() {
	ns.Move(obj5, wp12)
}
func HecubahWallsDown() {
	ns.WallGroupOpen(gvar15)
	ns.FrameTimer(1, HecubahEnter)
}
func GetOrb() {
	ns.ObjectOff(obj4)
	ns.ObjectGroupOff(gvar17)
	ns.SetHalberd(3)
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.FrameTimer(60, HecubahEnter)
	ns.ObjectGroupOff(gvar16)
	ns.JournalEntry(ns.GetHost(), "War11Hecubah", 1)
	ns.JournalEdit(ns.GetHost(), "War10aOrbQuest", 4)
	ns.FrameTimer(60, HecubahWallsDown)
}
func EnableAllMonsters() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar7)
	ns.ObjectGroupOn(gvar8)
	ns.ObjectGroupOn(gvar9)
}
func HecubahVanish1() {
	ns.ObjectOn(obj22)
	ns.ObjectGroupOff(gvar18)
	ns.Move(obj5, wp19)
	ns.AudioEvent(ns.HecubahTaunt, wp13)
	ns.WideScreen(false)
	ns.FrameTimer(60, EnableAllMonsters)
}
func HecubahTakesOff() {
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.FrameTimer(60, HecubahVanish1)
	ns.ObjectGroupOff(gvar7)
	ns.ObjectGroupOff(gvar8)
	ns.ObjectGroupOff(gvar9)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar25)
	ns.NoWallSound(false)
}
func HecubahVanish2() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp20)
	v1 = ns.GetWaypointY(wp20)
	v2 = ns.GetWaypointX(wp19)
	v3 = ns.GetWaypointY(wp19)
	ns.Effect(ns.SMOKE_BLAST, v2, v3, 0, 0)
	ns.AudioEvent(ns.TeleportOut, wp19)
	ns.MoveObject(obj5, v0, v1)
	ns.ObjectOn(obj24)
}
func MapSwitch() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp21)
	v1 = ns.GetWaypointY(wp21)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), v0, v1)
}
func TeleportPlayer() {
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(120, MapSwitch)
}
func EnableMonsterGroup2() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.ObjectGroupOn(gvar8)
}
func EnableMonsterGroup3() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.ObjectGroupOn(gvar9)
}
func DestroyMe() {
	ns.Delete(ns.GetCaller())
}
func LichWallsDown2() {
	var (
		v0  float32
		v1  float32
		v2  float32
		v3  float32
		v4  float32
		v5  float32
		v6  float32
		v7  float32
		v8  float32
		v9  float32
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 float32
	)
	v0 = ns.GetWaypointX(wp43)
	v1 = ns.GetWaypointY(wp43)
	v2 = ns.GetWaypointX(wp44)
	v3 = ns.GetWaypointY(wp44)
	v4 = ns.GetWaypointX(wp45)
	v5 = ns.GetWaypointY(wp45)
	v6 = ns.GetWaypointX(wp46)
	v7 = ns.GetWaypointY(wp46)
	v8 = ns.GetWaypointX(wp47)
	v9 = ns.GetWaypointY(wp47)
	v10 = ns.GetWaypointX(wp48)
	v11 = ns.GetWaypointY(wp48)
	v12 = ns.GetWaypointX(wp49)
	v13 = ns.GetWaypointY(wp49)
	v14 = ns.GetWaypointX(wp50)
	v15 = ns.GetWaypointY(wp50)
	ns.WallGroupOpen(gvar29)
	ns.ObjectGroupOn(gvar30)
	ns.ObjectGroupOn(gvar27)
	ns.FrameTimer(1, LichMove)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v2, v3, v2, v3)
	ns.Effect(ns.SMOKE_BLAST, v4, v5, v4, v5)
	ns.Effect(ns.SMOKE_BLAST, v6, v7, v6, v7)
	ns.Effect(ns.SMOKE_BLAST, v8, v9, v8, v9)
	ns.Effect(ns.SMOKE_BLAST, v10, v11, v10, v11)
	ns.Effect(ns.SMOKE_BLAST, v12, v13, v12, v13)
	ns.Effect(ns.SMOKE_BLAST, v14, v15, v14, v15)
}
func LichDieCheck() {
	if !(flag33 && flag34) {
		goto LABEL1
	}
	LichWallsDown2()
LABEL1:
	return
}
func LichDie1() {
	flag33 = true
	LichDieCheck()
}
func LichDie2() {
	flag34 = true
	LichDieCheck()
}
func LichMove() {
	ns.Move(obj31, wp51)
	ns.Move(obj32, wp52)
}
func LichWallsDown() {
	var (
		v0  float32
		v1  float32
		v2  float32
		v3  float32
		v4  float32
		v5  float32
		v6  float32
		v7  float32
		v8  float32
		v9  float32
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 float32
	)
	v0 = ns.GetWaypointX(wp35)
	v1 = ns.GetWaypointY(wp35)
	v2 = ns.GetWaypointX(wp36)
	v3 = ns.GetWaypointY(wp36)
	v4 = ns.GetWaypointX(wp37)
	v5 = ns.GetWaypointY(wp37)
	v6 = ns.GetWaypointX(wp38)
	v7 = ns.GetWaypointY(wp38)
	v8 = ns.GetWaypointX(wp39)
	v9 = ns.GetWaypointY(wp39)
	v10 = ns.GetWaypointX(wp40)
	v11 = ns.GetWaypointY(wp40)
	v12 = ns.GetWaypointX(wp41)
	v13 = ns.GetWaypointY(wp41)
	v14 = ns.GetWaypointX(wp42)
	v15 = ns.GetWaypointY(wp42)
	ns.WallGroupOpen(gvar28)
	ns.ObjectGroupOn(gvar26)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, v0, v1)
	ns.Effect(ns.SMOKE_BLAST, v2, v3, v2, v3)
	ns.Effect(ns.SMOKE_BLAST, v4, v5, v4, v5)
	ns.Effect(ns.SMOKE_BLAST, v6, v7, v6, v7)
	ns.Effect(ns.SMOKE_BLAST, v8, v9, v8, v9)
	ns.Effect(ns.SMOKE_BLAST, v10, v11, v10, v11)
	ns.Effect(ns.SMOKE_BLAST, v12, v13, v12, v13)
	ns.Effect(ns.SMOKE_BLAST, v14, v15, v14, v15)
}
func SentryRun() {
	var v0 int
	v0 = gvar65
	if v0 == gvar59 {
		goto LABEL1
	}
	if v0 == gvar60 {
		goto LABEL2
	}
	if v0 == gvar61 {
		goto LABEL3
	}
	if v0 == gvar62 {
		goto LABEL4
	}
	if v0 == gvar63 {
		goto LABEL5
	}
	if v0 == gvar64 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	ns.ObjectOff(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOn(obj55)
	ns.ObjectOn(obj56)
	ns.ObjectOn(obj57)
	ns.ObjectOff(obj58)
	ns.FrameTimer(25, SentryRun)
	gvar65 = gvar60
	goto LABEL7
LABEL2:
	ns.ObjectOn(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOn(obj55)
	ns.ObjectOn(obj56)
	ns.ObjectOff(obj57)
	ns.ObjectOn(obj58)
	ns.FrameTimer(25, SentryRun)
	gvar65 = gvar61
	goto LABEL7
LABEL3:
	ns.ObjectOn(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOn(obj55)
	ns.ObjectOff(obj56)
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.FrameTimer(25, SentryRun)
	gvar65 = gvar62
	goto LABEL7
LABEL4:
	ns.ObjectOn(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOff(obj55)
	ns.ObjectOn(obj56)
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.FrameTimer(25, SentryRun)
	gvar65 = gvar63
	goto LABEL7
LABEL5:
	ns.ObjectOn(obj53)
	ns.ObjectOff(obj54)
	ns.ObjectOn(obj55)
	ns.ObjectOn(obj56)
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.FrameTimer(25, SentryRun)
	gvar65 = gvar64
	goto LABEL7
LABEL6:
	ns.ObjectOn(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOn(obj55)
	ns.ObjectOn(obj56)
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.FrameTimer(25, SentryRun)
	gvar65 = gvar59
	goto LABEL7
LABEL7:
	return
}
func LLDies() {
	gvar71 = gvar70
}
func OpenLLWallGroup() {
	ns.WallGroupOpen(gvar68)
}
func Ghost1Face() {
	ns.LookAtObject(obj66, ns.GetHost())
}
func Ghost2Face() {
	ns.LookAtObject(obj66, ns.GetHost())
}
func PlayAction2Music() {
	ns.Music(27, 100)
}
func MapInitialize() {
	var v0 int
	gvar26 = ns.ObjectGroup("LichGroup")
	gvar27 = ns.ObjectGroup("LichGroup2")
	gvar28 = ns.WallGroup("LichWallGroup")
	gvar29 = ns.WallGroup("LichWallGroup2")
	gvar30 = ns.ObjectGroup("LichElevatorGroup")
	obj31 = ns.Object("Lich1")
	obj32 = ns.Object("Lich2")
	wp35 = ns.Waypoint("DestroyWP1")
	wp36 = ns.Waypoint("DestroyWP2")
	wp37 = ns.Waypoint("DestroyWP3")
	wp38 = ns.Waypoint("DestroyWP4")
	wp39 = ns.Waypoint("DestroyWP5")
	wp40 = ns.Waypoint("DestroyWP6")
	wp41 = ns.Waypoint("DestroyWP7")
	wp42 = ns.Waypoint("DestroyWP8")
	wp43 = ns.Waypoint("DestroyWP1a")
	wp44 = ns.Waypoint("DestroyWP2a")
	wp45 = ns.Waypoint("DestroyWP3a")
	wp46 = ns.Waypoint("DestroyWP4a")
	wp47 = ns.Waypoint("DestroyWP5a")
	wp48 = ns.Waypoint("DestroyWP6a")
	wp49 = ns.Waypoint("DestroyWP7a")
	wp50 = ns.Waypoint("DestroyWP8a")
	wp51 = ns.Waypoint("LichWP1")
	wp52 = ns.Waypoint("LichWP2")
	obj53 = ns.Object("Sentry1")
	obj54 = ns.Object("Sentry2")
	obj55 = ns.Object("Sentry3")
	obj56 = ns.Object("Sentry4")
	obj57 = ns.Object("Sentry5")
	obj58 = ns.Object("Sentry6")
	SentryRun()
	obj66 = ns.Object("LLGhost")
	wp67[0] = ns.Waypoint("LLGhostWP1")
	wp67[1] = ns.Waypoint("LLGhostWP2")
	wp67[2] = ns.Waypoint("LLGhostWP3")
	wp67[3] = ns.Waypoint("LLGhostWP4")
	wp67[4] = ns.Waypoint("LLGhostWP5")
	wp67[5] = ns.Waypoint("LLGhostWP6")
	gvar68 = ns.WallGroup("LLWallGroup")
	obj72 = ns.Object("Spike1")
	obj73 = ns.Object("Spike2")
	obj74 = ns.Object("Spike3")
	obj75 = ns.Object("Spike4")
	wp76 = ns.Waypoint("Spike1WP")
	wp77 = ns.Waypoint("Spike2WP")
	wp78 = ns.Waypoint("Spike3WP")
	wp79 = ns.Waypoint("Spike4WP")
	gvar80 = ns.WallGroup("SpikeWallGroup")
	gvar81 = ns.ObjectGroup("SpikeWallTrigGroup")
	gvar82 = ns.ObjectGroup("CherubGroup")
	gvar83 = ns.WallGroup("CherubWallGroup")
	obj84 = ns.Object("Cherub1")
	obj85 = ns.Object("Cherub2")
	obj86 = ns.Object("Cherub3")
	obj87 = ns.Object("Cherub4")
	obj88 = ns.Object("Cherub5")
	obj89 = ns.Object("Cherub6")
	obj90 = ns.Object("Cherub7")
	obj91 = ns.Object("Cherub8")
	obj92 = ns.Object("Cherub9")
	obj93 = ns.Object("Cherub10")
	gvar102 = ns.ObjectGroup("SentryGroup1")
	gvar103 = ns.ObjectGroup("SentryGroup2")
	gvar104 = ns.ObjectGroup("SentryGroup3")
	gvar105 = ns.ObjectGroup("SentryGroup4")
	obj94 = ns.Object("Sentry1L")
	obj95 = ns.Object("Sentry1R")
	obj96 = ns.Object("Sentry2L")
	obj97 = ns.Object("Sentry2R")
	obj98 = ns.Object("Sentry3L")
	obj99 = ns.Object("Sentry3R")
	obj100 = ns.Object("Sentry4L")
	obj101 = ns.Object("Sentry4R")
	gvar107 = ns.ObjectGroup("OrbTriggerGroup")
	gvar106 = ns.ObjectGroup("SentryMonsterGroup")
	gvar108 = ns.WallGroup("OrbWallGroup1")
	gvar109 = ns.WallGroup("OrbWallGroup2")
	gvar110 = ns.WallGroup("OrbWallGroup3")
	gvar10 = ns.WallGroup("OrbWallGroup4")
	gvar116 = ns.ObjectGroup("OrbXPTrigGroup")
	wp115 = ns.Waypoint("OrbXPWP")
	obj4 = ns.Object("OrbPool1")
	obj5 = ns.Object("Hecubah")
	gvar7 = ns.ObjectGroup("OrbMonsterGroup")
	gvar8 = ns.ObjectGroup("OrbMonsterGroup2")
	gvar9 = ns.ObjectGroup("OrbMonsterGroup3")
	gvar6 = ns.WallGroup("OrbWallGroupFinal")
	wp11 = ns.Waypoint("OrbAudioWP")
	gvar15 = ns.WallGroup("HecubahWallGroup")
	wp12 = ns.Waypoint("HecubahEnterWP")
	wp13 = ns.Waypoint("HecubahExitWP")
	wp14 = ns.Waypoint("HecubahLookWP")
	gvar16 = ns.ObjectGroup("OrbSoundGroup")
	gvar17 = ns.ObjectGroup("OrbPoolTrigGroup")
	gvar18 = ns.ObjectGroup("HecubahVanishTrigGroup")
	wp19 = ns.Waypoint("HecubahVanishWP")
	wp20 = ns.Waypoint("HecubahEndWP")
	wp21 = ns.Waypoint("MapSwitchWP")
	obj22 = ns.Object("HecubahEndTrig")
	obj24 = ns.Object("ExitTeleporter")
	wp117 = ns.Waypoint("Secret1AudioWP")
	wp118 = ns.Waypoint("Secret2AudioWP")
	gvar25 = ns.WallGroup("InvisibleWallGroup")
	ns.WallGroupOpen(gvar25)
	PlayAction2Music()
	v0 = 0
	for {
		if !(v0 < 32) {
			goto LABEL1
		}
		wp23[v0] = ns.Waypoint("SmokeWP" + strconv.Itoa(v0+1))
		v0 += 1
	}
LABEL1:
	return
}
func OrbXP() {
	ns.ObjectGroupOff(gvar116)
	ns.AudioEvent(ns.FlagDrop, wp115)
	ns.GiveXp(ns.GetHost(), 2000)
}
func Secret100XP() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.JournalEntryAdd, wp117)
}
func Secret200XP() {
	ns.GiveXp(ns.GetHost(), 200)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.JournalEntryAdd, wp118)
}
func PlayOrbMusic() {
	ns.Music(12, 100)
}
func PlayAction3Music() {
	ns.Music(28, 100)
}
func SpikeMove() {
	ns.ObjectGroupOff(gvar81)
	ns.AudioEvent(ns.SpikeBlockMove, wp79)
	ns.WallGroupOpen(gvar80)
	ns.Move(obj72, wp76)
	ns.Move(obj73, wp77)
	ns.Move(obj74, wp78)
	ns.Move(obj75, wp79)
	PlayAction3Music()
}
func OpenOrbWallGroup1() {
	if !(!ns.IsObjectOn(obj94) && !ns.IsObjectOn(obj95) && !ns.IsObjectOn(obj96) && !ns.IsObjectOn(obj97) && !ns.IsObjectOn(obj98) && !ns.IsObjectOn(obj99) && !ns.IsObjectOn(obj100) && !ns.IsObjectOn(obj101)) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar108)
LABEL1:
	return
}
func ToggleSG1() {
	ns.ObjectGroupToggle(gvar102)
	OpenOrbWallGroup1()
}
func EnableSG1() {
	ns.ObjectGroupOn(gvar102)
}
func ToggleSG2() {
	ns.ObjectGroupToggle(gvar103)
	OpenOrbWallGroup1()
}
func EnableSG2() {
	ns.ObjectGroupOn(gvar103)
}
func ToggleSG3() {
	ns.ObjectGroupToggle(gvar104)
	OpenOrbWallGroup1()
}
func EnableSG3() {
	ns.ObjectGroupOn(gvar104)
}
func ToggleSG4() {
	ns.ObjectGroupToggle(gvar105)
	OpenOrbWallGroup1()
}
func EnableSG4() {
	ns.ObjectGroupOn(gvar105)
}
func OpenOrbWallGroup2() {
	ns.WallGroupOpen(gvar109)
	ns.ObjectGroupOn(gvar106)
	ns.ObjectGroupOn(gvar82)
	ns.WallGroupOpen(gvar83)
}
func OpenOrbWallGroup3() {
	ns.WallGroupOpen(gvar110)
}
func OpenOrbWallGroup4() {
	ns.WallGroupOpen(gvar10)
	ns.ObjectGroupOff(gvar107)
	ns.ObjectGroupOff(gvar102)
	ns.ObjectGroupOff(gvar103)
	ns.ObjectGroupOff(gvar104)
	ns.ObjectGroupOff(gvar105)
}
func BossDieCheck() {
	if !(flag111 && flag112 && flag113 && flag114) {
		goto LABEL1
	}
	OpenOrbWallGroup3()
LABEL1:
	return
}
func BossDie1() {
	flag111 = true
	BossDieCheck()
}
func BossDie2() {
	flag112 = true
	BossDieCheck()
}
func BossDie3() {
	flag113 = true
	BossDieCheck()
}
func BossDie4() {
	flag114 = true
	BossDieCheck()
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
