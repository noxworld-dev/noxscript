package war10d

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
	gvar26  int
	gvar27  int
	gvar28  int
	gvar29  int
	gvar30  ns.ObjectGroupID
	gvar31  ns.ObjectGroupID
	gvar32  ns.WallGroupID
	gvar33  ns.WallGroupID
	gvar34  ns.ObjectGroupID
	obj35   ns.ObjectID
	obj36   ns.ObjectID
	flag37  bool
	flag38  bool
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
	wp53    ns.WaypointID
	wp54    ns.WaypointID
	wp55    ns.WaypointID
	wp56    ns.WaypointID
	obj57   ns.ObjectID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	gvar63  int
	gvar64  int
	gvar65  int
	gvar66  int
	gvar67  int
	gvar68  int
	gvar69  int
	obj70   ns.ObjectID
	wp71    [6]ns.WaypointID
	gvar72  ns.WallGroupID
	gvar73  int
	gvar74  int
	gvar75  int
	obj76   ns.ObjectID
	obj77   ns.ObjectID
	obj78   ns.ObjectID
	obj79   ns.ObjectID
	wp80    ns.WaypointID
	wp81    ns.WaypointID
	wp82    ns.WaypointID
	wp83    ns.WaypointID
	gvar84  ns.WallGroupID
	gvar85  ns.ObjectGroupID
	gvar86  ns.ObjectGroupID
	gvar87  ns.WallGroupID
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
	obj102  ns.ObjectID
	obj103  ns.ObjectID
	obj104  ns.ObjectID
	obj105  ns.ObjectID
	gvar106 ns.ObjectGroupID
	gvar107 ns.ObjectGroupID
	gvar108 ns.ObjectGroupID
	gvar109 ns.ObjectGroupID
	gvar110 ns.ObjectGroupID
	gvar111 ns.ObjectGroupID
	gvar112 ns.WallGroupID
	gvar113 ns.WallGroupID
	gvar114 ns.WallGroupID
	flag115 bool
	flag116 bool
	flag117 bool
	flag118 bool
	wp119   ns.WaypointID
	gvar120 ns.ObjectGroupID
	wp121   ns.WaypointID
	wp122   ns.WaypointID
)

func init() {
	gvar63 = 0
	gvar64 = 1
	gvar65 = 2
	gvar66 = 3
	gvar67 = 4
	gvar68 = 5
	gvar69 = gvar63
	gvar73 = 0
	gvar74 = 1
	gvar75 = gvar73
	flag37 = false
	flag38 = false
	flag115 = false
	flag116 = false
	flag117 = false
	flag118 = false
}
func PlayerDeath() {
	ns.DeathScreen(10)
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
	ns.FrameTimer(60, HecubahWallsDown)
	ns.ObjectGroupOff(gvar16)
	ns.JournalEntry(ns.GetHost(), "War11Hecubah", 1)
	ns.JournalEdit(ns.GetHost(), "War10aOrbQuest", 4)
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
func DestroyMe() {
	ns.Delete(ns.GetCaller())
}
func PlayerFadeOut() {
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(60, MapSwitch)
}
func TeleportPlayer() {
	ns.Blind()
	ns.Frozen(ns.GetHost(), true)
}
func EnableMonsterGroup2() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.ObjectGroupOn(gvar8)
}
func EnableMonsterGroup3() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.ObjectGroupOn(gvar9)
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
	v0 = ns.GetWaypointX(wp47)
	v1 = ns.GetWaypointY(wp47)
	v2 = ns.GetWaypointX(wp48)
	v3 = ns.GetWaypointY(wp48)
	v4 = ns.GetWaypointX(wp49)
	v5 = ns.GetWaypointY(wp49)
	v6 = ns.GetWaypointX(wp50)
	v7 = ns.GetWaypointY(wp50)
	v8 = ns.GetWaypointX(wp51)
	v9 = ns.GetWaypointY(wp51)
	v10 = ns.GetWaypointX(wp52)
	v11 = ns.GetWaypointY(wp52)
	v12 = ns.GetWaypointX(wp53)
	v13 = ns.GetWaypointY(wp53)
	v14 = ns.GetWaypointX(wp54)
	v15 = ns.GetWaypointY(wp54)
	ns.WallGroupOpen(gvar33)
	ns.ObjectGroupOn(gvar34)
	ns.ObjectGroupOn(gvar31)
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
	if !(flag37 && flag38) {
		goto LABEL1
	}
	LichWallsDown2()
LABEL1:
	return
}
func LichDie1() {
	flag37 = true
	LichDieCheck()
}
func LichDie2() {
	flag38 = true
	LichDieCheck()
}
func LichMove() {
	ns.Move(obj35, wp55)
	ns.Move(obj36, wp56)
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
	v0 = ns.GetWaypointX(wp39)
	v1 = ns.GetWaypointY(wp39)
	v2 = ns.GetWaypointX(wp40)
	v3 = ns.GetWaypointY(wp40)
	v4 = ns.GetWaypointX(wp41)
	v5 = ns.GetWaypointY(wp41)
	v6 = ns.GetWaypointX(wp42)
	v7 = ns.GetWaypointY(wp42)
	v8 = ns.GetWaypointX(wp43)
	v9 = ns.GetWaypointY(wp43)
	v10 = ns.GetWaypointX(wp44)
	v11 = ns.GetWaypointY(wp44)
	v12 = ns.GetWaypointX(wp45)
	v13 = ns.GetWaypointY(wp45)
	v14 = ns.GetWaypointX(wp46)
	v15 = ns.GetWaypointY(wp46)
	ns.WallGroupOpen(gvar32)
	ns.ObjectGroupOn(gvar30)
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
	v0 = gvar69
	if v0 == gvar63 {
		goto LABEL1
	}
	if v0 == gvar64 {
		goto LABEL2
	}
	if v0 == gvar65 {
		goto LABEL3
	}
	if v0 == gvar66 {
		goto LABEL4
	}
	if v0 == gvar67 {
		goto LABEL5
	}
	if v0 == gvar68 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	ns.ObjectOff(obj57)
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOff(obj62)
	ns.FrameTimer(25, SentryRun)
	gvar69 = gvar64
	goto LABEL7
LABEL2:
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOff(obj61)
	ns.ObjectOn(obj62)
	ns.FrameTimer(25, SentryRun)
	gvar69 = gvar65
	goto LABEL7
LABEL3:
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOff(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.FrameTimer(25, SentryRun)
	gvar69 = gvar66
	goto LABEL7
LABEL4:
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.ObjectOff(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.FrameTimer(25, SentryRun)
	gvar69 = gvar67
	goto LABEL7
LABEL5:
	ns.ObjectOn(obj57)
	ns.ObjectOff(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.FrameTimer(25, SentryRun)
	gvar69 = gvar68
	goto LABEL7
LABEL6:
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.FrameTimer(25, SentryRun)
	gvar69 = gvar63
	goto LABEL7
LABEL7:
	return
}
func LLDies() {
	gvar75 = gvar74
}
func OpenLLWallGroup() {
	ns.WallGroupOpen(gvar72)
}
func Ghost1Face() {
	ns.LookAtObject(obj70, ns.GetHost())
}
func Ghost2Face() {
	ns.LookAtObject(obj70, ns.GetHost())
}
func PlayAction2Music() {
	ns.Music(27, 100)
}
func MapInitialize() {
	var v0 int
	gvar30 = ns.ObjectGroup("LichGroup")
	gvar31 = ns.ObjectGroup("LichGroup2")
	gvar32 = ns.WallGroup("LichWallGroup")
	gvar33 = ns.WallGroup("LichWallGroup2")
	gvar34 = ns.ObjectGroup("LichElevatorGroup")
	obj35 = ns.Object("Lich1")
	obj36 = ns.Object("Lich2")
	wp39 = ns.Waypoint("DestroyWP1")
	wp40 = ns.Waypoint("DestroyWP2")
	wp41 = ns.Waypoint("DestroyWP3")
	wp42 = ns.Waypoint("DestroyWP4")
	wp43 = ns.Waypoint("DestroyWP5")
	wp44 = ns.Waypoint("DestroyWP6")
	wp45 = ns.Waypoint("DestroyWP7")
	wp46 = ns.Waypoint("DestroyWP8")
	wp47 = ns.Waypoint("DestroyWP1a")
	wp48 = ns.Waypoint("DestroyWP2a")
	wp49 = ns.Waypoint("DestroyWP3a")
	wp50 = ns.Waypoint("DestroyWP4a")
	wp51 = ns.Waypoint("DestroyWP5a")
	wp52 = ns.Waypoint("DestroyWP6a")
	wp53 = ns.Waypoint("DestroyWP7a")
	wp54 = ns.Waypoint("DestroyWP8a")
	wp55 = ns.Waypoint("LichWP1")
	wp56 = ns.Waypoint("LichWP2")
	obj57 = ns.Object("Sentry1")
	obj58 = ns.Object("Sentry2")
	obj59 = ns.Object("Sentry3")
	obj60 = ns.Object("Sentry4")
	obj61 = ns.Object("Sentry5")
	obj62 = ns.Object("Sentry6")
	SentryRun()
	obj70 = ns.Object("LLGhost")
	wp71[0] = ns.Waypoint("LLGhostWP1")
	wp71[1] = ns.Waypoint("LLGhostWP2")
	wp71[2] = ns.Waypoint("LLGhostWP3")
	wp71[3] = ns.Waypoint("LLGhostWP4")
	wp71[4] = ns.Waypoint("LLGhostWP5")
	wp71[5] = ns.Waypoint("LLGhostWP6")
	gvar72 = ns.WallGroup("LLWallGroup")
	obj76 = ns.Object("Spike1")
	obj77 = ns.Object("Spike2")
	obj78 = ns.Object("Spike3")
	obj79 = ns.Object("Spike4")
	wp80 = ns.Waypoint("Spike1WP")
	wp81 = ns.Waypoint("Spike2WP")
	wp82 = ns.Waypoint("Spike3WP")
	wp83 = ns.Waypoint("Spike4WP")
	gvar84 = ns.WallGroup("SpikeWallGroup")
	gvar85 = ns.ObjectGroup("SpikeWallTrigGroup")
	gvar86 = ns.ObjectGroup("CherubGroup")
	gvar87 = ns.WallGroup("CherubWallGroup")
	obj88 = ns.Object("Cherub1")
	obj89 = ns.Object("Cherub2")
	obj90 = ns.Object("Cherub3")
	obj91 = ns.Object("Cherub4")
	obj92 = ns.Object("Cherub5")
	obj93 = ns.Object("Cherub6")
	obj94 = ns.Object("Cherub7")
	obj95 = ns.Object("Cherub8")
	obj96 = ns.Object("Cherub9")
	obj97 = ns.Object("Cherub10")
	gvar106 = ns.ObjectGroup("SentryGroup1")
	gvar107 = ns.ObjectGroup("SentryGroup2")
	gvar108 = ns.ObjectGroup("SentryGroup3")
	gvar109 = ns.ObjectGroup("SentryGroup4")
	obj98 = ns.Object("Sentry1L")
	obj99 = ns.Object("Sentry1R")
	obj100 = ns.Object("Sentry2L")
	obj101 = ns.Object("Sentry2R")
	obj102 = ns.Object("Sentry3L")
	obj103 = ns.Object("Sentry3R")
	obj104 = ns.Object("Sentry4L")
	obj105 = ns.Object("Sentry4R")
	gvar111 = ns.ObjectGroup("OrbTriggerGroup")
	gvar110 = ns.ObjectGroup("SentryMonsterGroup")
	gvar112 = ns.WallGroup("OrbWallGroup1")
	gvar113 = ns.WallGroup("OrbWallGroup2")
	gvar114 = ns.WallGroup("OrbWallGroup3")
	gvar10 = ns.WallGroup("OrbWallGroup4")
	wp119 = ns.Waypoint("OrbXPWP")
	gvar120 = ns.ObjectGroup("OrbXPTrigGroup")
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
	wp121 = ns.Waypoint("Secret1AudioWP")
	wp122 = ns.Waypoint("Secret2AudioWP")
	PlayAction2Music()
	gvar25 = ns.WallGroup("InvisibleWallGroup")
	ns.WallGroupOpen(gvar25)
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
	ns.ObjectGroupOff(gvar120)
	ns.AudioEvent(ns.FlagDrop, wp119)
	ns.GiveXp(ns.GetHost(), 2000)
}
func Secret100XP() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp121)
}
func Secret200XP() {
	ns.GiveXp(ns.GetHost(), 200)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp122)
}
func HostilizeMe() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func PlayAction3Music() {
	ns.Music(28, 100)
}
func PlayOrbMusic() {
	ns.Music(12, 100)
}
func SpikeMove() {
	ns.ObjectGroupOff(gvar85)
	ns.AudioEvent(ns.SpikeBlockMove, wp83)
	ns.WallGroupOpen(gvar84)
	ns.Move(obj76, wp80)
	ns.Move(obj77, wp81)
	ns.Move(obj78, wp82)
	ns.Move(obj79, wp83)
	PlayAction3Music()
}
func OpenOrbWallGroup1() {
	if !(!ns.IsObjectOn(obj98) && !ns.IsObjectOn(obj99) && !ns.IsObjectOn(obj100) && !ns.IsObjectOn(obj101) && !ns.IsObjectOn(obj102) && !ns.IsObjectOn(obj103) && !ns.IsObjectOn(obj104) && !ns.IsObjectOn(obj105)) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar112)
LABEL1:
	return
}
func ToggleSG1() {
	ns.ObjectGroupToggle(gvar106)
	OpenOrbWallGroup1()
}
func EnableSG1() {
	ns.ObjectGroupOn(gvar106)
}
func ToggleSG2() {
	ns.ObjectGroupToggle(gvar107)
	OpenOrbWallGroup1()
}
func EnableSG2() {
	ns.ObjectGroupOn(gvar107)
}
func ToggleSG3() {
	ns.ObjectGroupToggle(gvar108)
	OpenOrbWallGroup1()
}
func EnableSG3() {
	ns.ObjectGroupOn(gvar108)
}
func ToggleSG4() {
	ns.ObjectGroupToggle(gvar109)
	OpenOrbWallGroup1()
}
func EnableSG4() {
	ns.ObjectGroupOn(gvar109)
}
func OpenOrbWallGroup2() {
	ns.WallGroupOpen(gvar113)
	ns.ObjectGroupOn(gvar110)
	ns.ObjectGroupOn(gvar86)
	ns.WallGroupOpen(gvar87)
}
func OpenOrbWallGroup3() {
	ns.WallGroupOpen(gvar114)
}
func OpenOrbWallGroup4() {
	ns.WallGroupOpen(gvar10)
	ns.ObjectGroupOff(gvar111)
	ns.ObjectGroupOff(gvar106)
	ns.ObjectGroupOff(gvar107)
	ns.ObjectGroupOff(gvar108)
	ns.ObjectGroupOff(gvar109)
}
func BossDieCheck() {
	if !(flag115 && flag116 && flag117 && flag118) {
		goto LABEL1
	}
	OpenOrbWallGroup3()
LABEL1:
	return
}
func BossDie1() {
	flag115 = true
	BossDieCheck()
}
func BossDie2() {
	flag116 = true
	BossDieCheck()
}
func BossDie3() {
	flag117 = true
	BossDieCheck()
}
func BossDie4() {
	flag118 = true
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
