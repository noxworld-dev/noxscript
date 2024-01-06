package con10d

import (
	"github.com/noxworld-dev/noxscript/ns/v3"
	"strconv"
)

var (
	obj4    ns.ObjectID
	obj5    ns.ObjectID
	obj6    ns.ObjectID
	obj7    ns.ObjectID
	obj8    ns.ObjectID
	obj9    ns.ObjectID
	gvar10  ns.WallGroupID
	gvar11  ns.ObjectGroupID
	gvar12  ns.ObjectGroupID
	gvar13  ns.ObjectGroupID
	gvar14  ns.WallGroupID
	wp15    ns.WaypointID
	wp16    ns.WaypointID
	wp17    ns.WaypointID
	wp18    ns.WaypointID
	gvar19  ns.WallGroupID
	gvar20  ns.ObjectGroupID
	gvar21  ns.ObjectGroupID
	gvar22  ns.ObjectGroupID
	wp23    ns.WaypointID
	wp24    ns.WaypointID
	wp25    ns.WaypointID
	obj26   ns.ObjectID
	wp27    [32]ns.WaypointID
	obj28   ns.ObjectID
	gvar29  ns.WallGroupID
	gvar30  [85]int
	gvar31  ns.ObjectGroupID
	gvar32  ns.ObjectGroupID
	gvar33  ns.WallGroupID
	gvar34  ns.WallGroupID
	gvar35  ns.ObjectGroupID
	obj36   ns.ObjectID
	obj37   ns.ObjectID
	flag38  bool
	flag39  bool
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
	wp57    ns.WaypointID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	obj63   ns.ObjectID
	gvar64  int
	gvar65  int
	gvar66  int
	gvar67  int
	gvar68  int
	gvar69  int
	gvar70  int
	obj71   ns.ObjectID
	wp72    [6]ns.WaypointID
	gvar73  ns.WallGroupID
	gvar74  int
	gvar75  int
	gvar76  int
	obj77   ns.ObjectID
	obj78   ns.ObjectID
	obj79   ns.ObjectID
	obj80   ns.ObjectID
	wp81    ns.WaypointID
	wp82    ns.WaypointID
	wp83    ns.WaypointID
	wp84    ns.WaypointID
	gvar85  ns.WallGroupID
	gvar86  ns.ObjectGroupID
	gvar87  ns.ObjectGroupID
	gvar88  ns.WallGroupID
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
	obj106  ns.ObjectID
	gvar107 ns.ObjectGroupID
	gvar108 ns.ObjectGroupID
	gvar109 ns.ObjectGroupID
	gvar110 ns.ObjectGroupID
	gvar111 ns.ObjectGroupID
	gvar112 ns.ObjectGroupID
	gvar113 ns.WallGroupID
	gvar114 ns.WallGroupID
	gvar115 ns.WallGroupID
	flag116 bool
	flag117 bool
	flag118 bool
	flag119 bool
	gvar120 ns.ObjectGroupID
	wp121   ns.WaypointID
	wp122   ns.WaypointID
	wp123   ns.WaypointID
)

func init() {
	gvar64 = 0
	gvar65 = 1
	gvar66 = 2
	gvar67 = 3
	gvar68 = 4
	gvar69 = 5
	gvar70 = gvar64
	gvar74 = 0
	gvar75 = 1
	gvar76 = gvar74
	flag38 = false
	flag39 = false
	flag116 = false
	flag117 = false
	flag118 = false
	flag119 = false
}
func PlayerDeath() {
	ns.DeathScreen(10)
}
func EnableAllMonsters() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar11)
	ns.ObjectGroupOn(gvar12)
	ns.ObjectGroupOn(gvar13)
}
func FilterHecubah() int {
	if !ns.IsCaller(obj9) {
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
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp27[v0]), ns.GetWaypointY(wp27[v0]), 0, 0)
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
	v0 = ns.GetWaypointX(wp17)
	v1 = ns.GetWaypointY(wp17)
	v2 = ns.GetWaypointX(wp18)
	v3 = ns.GetWaypointY(wp18)
	ns.CreatureGuard(obj9, v0, v1, v2, v3, 0)
	ns.NoWallSound(true)
	ns.WallGroupClose(gvar29)
	ns.NoWallSound(false)
}
func EnableOrbMonsterGroup() {
	ns.ObjectGroupOn(gvar11)
	ns.ObjectGroupOn(gvar12)
	ns.ObjectGroupOn(gvar13)
	ns.Frozen(ns.GetHost(), false)
	ns.WallGroupOpen(gvar10)
	ns.FrameTimer(2, ActivateSmokePuffs)
	ns.WallGroupClose(gvar14)
	ns.Frozen(ns.GetHost(), false)
}
func HecubahExit() {
	ns.Enchant(obj9, ns.ENCHANT_HASTED, 100)
	ns.Move(obj9, wp17)
	ns.WideScreen(false)
	ns.FrameTimer(60, EnableOrbMonsterGroup)
}
func HecubahDialogEnd() {
	ns.FrameTimer(30, HecubahExit)
	ns.CancelDialog(obj9)
}
func HecubahDialogStart() {
	ns.TellStory(ns.DemonRecognize, "Con07H.scr:HecubahTalk02")
}
func HecubahTrigger() {
	ns.LookAtObject(obj9, ns.GetHost())
	ns.StoryPic(obj9, "HecubahPic")
	ns.SetDialog(obj9, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj9, ns.GetHost())
}
func HecubahEnter() {
	ns.Move(obj9, wp16)
}
func HecubahWallsDown() {
	ns.WallGroupOpen(gvar19)
	ns.FrameTimer(1, HecubahEnter)
}
func GetOrb() {
	ns.ObjectOff(obj8)
	ns.ObjectGroupOff(gvar21)
	ns.SetHalberd(3)
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.FrameTimer(60, HecubahWallsDown)
	ns.ObjectGroupOff(gvar20)
	ns.JournalEntry(ns.GetHost(), "War11Hecubah", 1)
	ns.JournalEdit(ns.GetHost(), "War10aOrbQuest", 4)
}
func HecubahVanish1() {
	ns.ObjectOn(obj26)
	ns.ObjectGroupOff(gvar22)
	ns.Move(obj9, wp23)
	ns.AudioEvent(ns.HecubahTaunt, wp17)
	ns.WideScreen(false)
	ns.FrameTimer(60, EnableAllMonsters)
}
func HecubahTakesOff() {
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.FrameTimer(60, HecubahVanish1)
	ns.ObjectGroupOff(gvar11)
	ns.ObjectGroupOff(gvar12)
	ns.ObjectGroupOff(gvar13)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar29)
	ns.NoWallSound(false)
}
func HecubahVanish2() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp24)
	v1 = ns.GetWaypointY(wp24)
	v2 = ns.GetWaypointX(wp23)
	v3 = ns.GetWaypointY(wp23)
	ns.Effect(ns.SMOKE_BLAST, v2, v3, 0, 0)
	ns.AudioEvent(ns.TeleportOut, wp23)
	ns.MoveObject(obj9, v0, v1)
	ns.ObjectOn(obj28)
}
func MapSwitch() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp25)
	v1 = ns.GetWaypointY(wp25)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), v0, v1)
}
func EnableMonsterGroup2() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectGroupOn(gvar12)
LABEL1:
	return
}
func EnableMonsterGroup3() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectGroupOn(gvar13)
LABEL1:
	return
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
	v0 = ns.GetWaypointX(wp48)
	v1 = ns.GetWaypointY(wp48)
	v2 = ns.GetWaypointX(wp49)
	v3 = ns.GetWaypointY(wp49)
	v4 = ns.GetWaypointX(wp50)
	v5 = ns.GetWaypointY(wp50)
	v6 = ns.GetWaypointX(wp51)
	v7 = ns.GetWaypointY(wp51)
	v8 = ns.GetWaypointX(wp52)
	v9 = ns.GetWaypointY(wp52)
	v10 = ns.GetWaypointX(wp53)
	v11 = ns.GetWaypointY(wp53)
	v12 = ns.GetWaypointX(wp54)
	v13 = ns.GetWaypointY(wp54)
	v14 = ns.GetWaypointX(wp55)
	v15 = ns.GetWaypointY(wp55)
	ns.WallGroupOpen(gvar34)
	ns.ObjectGroupOn(gvar35)
	ns.ObjectGroupOn(gvar32)
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
	if !(flag38 && flag39) {
		goto LABEL1
	}
	LichWallsDown2()
LABEL1:
	return
}
func LichDie1() {
	flag38 = true
	LichDieCheck()
}
func LichDie2() {
	flag39 = true
	LichDieCheck()
}
func LichMove() {
	ns.Move(obj36, wp56)
	ns.Move(obj37, wp57)
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
	v0 = ns.GetWaypointX(wp40)
	v1 = ns.GetWaypointY(wp40)
	v2 = ns.GetWaypointX(wp41)
	v3 = ns.GetWaypointY(wp41)
	v4 = ns.GetWaypointX(wp42)
	v5 = ns.GetWaypointY(wp42)
	v6 = ns.GetWaypointX(wp43)
	v7 = ns.GetWaypointY(wp43)
	v8 = ns.GetWaypointX(wp44)
	v9 = ns.GetWaypointY(wp44)
	v10 = ns.GetWaypointX(wp45)
	v11 = ns.GetWaypointY(wp45)
	v12 = ns.GetWaypointX(wp46)
	v13 = ns.GetWaypointY(wp46)
	v14 = ns.GetWaypointX(wp47)
	v15 = ns.GetWaypointY(wp47)
	ns.WallGroupOpen(gvar33)
	ns.ObjectGroupOn(gvar31)
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
	v0 = gvar70
	if v0 == gvar64 {
		goto LABEL1
	}
	if v0 == gvar65 {
		goto LABEL2
	}
	if v0 == gvar66 {
		goto LABEL3
	}
	if v0 == gvar67 {
		goto LABEL4
	}
	if v0 == gvar68 {
		goto LABEL5
	}
	if v0 == gvar69 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	ns.ObjectOff(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.ObjectOff(obj63)
	ns.FrameTimer(25, SentryRun)
	gvar70 = gvar65
	goto LABEL7
LABEL2:
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOff(obj62)
	ns.ObjectOn(obj63)
	ns.FrameTimer(25, SentryRun)
	gvar70 = gvar66
	goto LABEL7
LABEL3:
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOff(obj61)
	ns.ObjectOn(obj62)
	ns.ObjectOn(obj63)
	ns.FrameTimer(25, SentryRun)
	gvar70 = gvar67
	goto LABEL7
LABEL4:
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOff(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.ObjectOn(obj63)
	ns.FrameTimer(25, SentryRun)
	gvar70 = gvar68
	goto LABEL7
LABEL5:
	ns.ObjectOn(obj58)
	ns.ObjectOff(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.ObjectOn(obj63)
	ns.FrameTimer(25, SentryRun)
	gvar70 = gvar69
	goto LABEL7
LABEL6:
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.ObjectOn(obj63)
	ns.FrameTimer(25, SentryRun)
	gvar70 = gvar64
	goto LABEL7
LABEL7:
	return
}
func LLDies() {
	gvar76 = gvar75
}
func OpenLLWallGroup() {
	ns.WallGroupOpen(gvar73)
}
func Ghost1Face() {
	ns.LookAtObject(obj71, ns.GetHost())
}
func Ghost2Face() {
	ns.LookAtObject(obj71, ns.GetHost())
}
func PlayAction2Music() {
	ns.Music(27, 100)
}
func MapInitialize() {
	var v0 int
	gvar31 = ns.ObjectGroup("LichGroup")
	gvar32 = ns.ObjectGroup("LichGroup2")
	gvar33 = ns.WallGroup("LichWallGroup")
	gvar34 = ns.WallGroup("LichWallGroup2")
	gvar35 = ns.ObjectGroup("LichElevatorGroup")
	obj36 = ns.Object("Lich1")
	obj37 = ns.Object("Lich2")
	wp40 = ns.Waypoint("DestroyWP1")
	wp41 = ns.Waypoint("DestroyWP2")
	wp42 = ns.Waypoint("DestroyWP3")
	wp43 = ns.Waypoint("DestroyWP4")
	wp44 = ns.Waypoint("DestroyWP5")
	wp45 = ns.Waypoint("DestroyWP6")
	wp46 = ns.Waypoint("DestroyWP7")
	wp47 = ns.Waypoint("DestroyWP8")
	wp48 = ns.Waypoint("DestroyWP1a")
	wp49 = ns.Waypoint("DestroyWP2a")
	wp50 = ns.Waypoint("DestroyWP3a")
	wp51 = ns.Waypoint("DestroyWP4a")
	wp52 = ns.Waypoint("DestroyWP5a")
	wp53 = ns.Waypoint("DestroyWP6a")
	wp54 = ns.Waypoint("DestroyWP7a")
	wp55 = ns.Waypoint("DestroyWP8a")
	wp56 = ns.Waypoint("LichWP1")
	wp57 = ns.Waypoint("LichWP2")
	obj58 = ns.Object("Sentry1")
	obj59 = ns.Object("Sentry2")
	obj60 = ns.Object("Sentry3")
	obj61 = ns.Object("Sentry4")
	obj62 = ns.Object("Sentry5")
	obj63 = ns.Object("Sentry6")
	SentryRun()
	obj71 = ns.Object("LLGhost")
	wp72[0] = ns.Waypoint("LLGhostWP1")
	wp72[1] = ns.Waypoint("LLGhostWP2")
	wp72[2] = ns.Waypoint("LLGhostWP3")
	wp72[3] = ns.Waypoint("LLGhostWP4")
	wp72[4] = ns.Waypoint("LLGhostWP5")
	wp72[5] = ns.Waypoint("LLGhostWP6")
	gvar73 = ns.WallGroup("LLWallGroup")
	obj77 = ns.Object("Spike1")
	obj78 = ns.Object("Spike2")
	obj79 = ns.Object("Spike3")
	obj80 = ns.Object("Spike4")
	wp81 = ns.Waypoint("Spike1WP")
	wp82 = ns.Waypoint("Spike2WP")
	wp83 = ns.Waypoint("Spike3WP")
	wp84 = ns.Waypoint("Spike4WP")
	gvar85 = ns.WallGroup("SpikeWallGroup")
	gvar86 = ns.ObjectGroup("SpikeWallTrigGroup")
	gvar87 = ns.ObjectGroup("CherubGroup")
	gvar88 = ns.WallGroup("CherubWallGroup")
	obj89 = ns.Object("Cherub1")
	obj90 = ns.Object("Cherub2")
	obj91 = ns.Object("Cherub3")
	obj92 = ns.Object("Cherub4")
	obj93 = ns.Object("Cherub5")
	obj94 = ns.Object("Cherub6")
	obj95 = ns.Object("Cherub7")
	obj96 = ns.Object("Cherub8")
	obj97 = ns.Object("Cherub9")
	obj98 = ns.Object("Cherub10")
	gvar107 = ns.ObjectGroup("SentryGroup1")
	gvar108 = ns.ObjectGroup("SentryGroup2")
	gvar109 = ns.ObjectGroup("SentryGroup3")
	gvar110 = ns.ObjectGroup("SentryGroup4")
	obj99 = ns.Object("Sentry1L")
	obj100 = ns.Object("Sentry1R")
	obj101 = ns.Object("Sentry2L")
	obj102 = ns.Object("Sentry2R")
	obj103 = ns.Object("Sentry3L")
	obj104 = ns.Object("Sentry3R")
	obj105 = ns.Object("Sentry4L")
	obj106 = ns.Object("Sentry4R")
	gvar112 = ns.ObjectGroup("OrbTriggerGroup")
	gvar111 = ns.ObjectGroup("SentryMonsterGroup")
	gvar113 = ns.WallGroup("OrbWallGroup1")
	gvar114 = ns.WallGroup("OrbWallGroup2")
	gvar115 = ns.WallGroup("OrbWallGroup3")
	gvar14 = ns.WallGroup("OrbWallGroup4")
	wp121 = ns.Waypoint("OrbXPWP")
	gvar120 = ns.ObjectGroup("OrbXPTrigGroup")
	obj4 = ns.Object("Group2AlertA")
	obj5 = ns.Object("Group2AlertB")
	obj6 = ns.Object("Group3AlertA")
	obj7 = ns.Object("Group3AlertB")
	obj8 = ns.Object("OrbPool1")
	obj9 = ns.Object("Hecubah")
	gvar11 = ns.ObjectGroup("OrbMonsterGroup")
	gvar12 = ns.ObjectGroup("OrbMonsterGroup2")
	gvar13 = ns.ObjectGroup("OrbMonsterGroup3")
	gvar10 = ns.WallGroup("OrbWallGroupFinal")
	wp15 = ns.Waypoint("OrbAudioWP")
	gvar19 = ns.WallGroup("HecubahWallGroup")
	wp16 = ns.Waypoint("HecubahEnterWP")
	wp17 = ns.Waypoint("HecubahExitWP")
	wp18 = ns.Waypoint("HecubahLookWP")
	gvar20 = ns.ObjectGroup("OrbSoundGroup")
	gvar21 = ns.ObjectGroup("OrbPoolTrigGroup")
	gvar22 = ns.ObjectGroup("HecubahVanishTrigGroup")
	wp23 = ns.Waypoint("HecubahVanishWP")
	wp24 = ns.Waypoint("HecubahEndWP")
	wp25 = ns.Waypoint("MapSwitchWP")
	obj26 = ns.Object("HecubahEndTrig")
	obj28 = ns.Object("ExitTeleporter")
	wp122 = ns.Waypoint("Secret1AudioWP")
	wp123 = ns.Waypoint("Secret2AudioWP")
	gvar29 = ns.WallGroup("InvisibleWallGroup")
	ns.WallGroupOpen(gvar29)
	PlayAction2Music()
	v0 = 0
	for {
		if !(v0 < 32) {
			goto LABEL1
		}
		wp27[v0] = ns.Waypoint("SmokeWP" + strconv.Itoa(v0+1))
		v0 += 1
	}
LABEL1:
	return
}
func OrbXP() {
	ns.ObjectGroupOff(gvar120)
	ns.AudioEvent(ns.FlagDrop, wp121)
	ns.GiveXp(ns.GetHost(), 2000)
}
func Secret100XP() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp122)
}
func Secret200XP() {
	ns.GiveXp(ns.GetHost(), 200)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp123)
}
func PlaySubMusic() {
	ns.Music(18, 100)
}
func PlayAction3Music() {
	ns.Music(28, 100)
}
func PlayOrbMusic() {
	ns.Music(12, 100)
}
func SpikeMove() {
	ns.ObjectGroupOff(gvar86)
	ns.AudioEvent(ns.SpikeBlockMove, wp83)
	ns.WallGroupOpen(gvar85)
	ns.Move(obj77, wp81)
	ns.Move(obj78, wp82)
	ns.Move(obj79, wp83)
	ns.Move(obj80, wp84)
	PlayAction3Music()
}
func OpenOrbWallGroup1() {
	if !(!ns.IsObjectOn(obj99) && !ns.IsObjectOn(obj100) && !ns.IsObjectOn(obj101) && !ns.IsObjectOn(obj102) && !ns.IsObjectOn(obj103) && !ns.IsObjectOn(obj104) && !ns.IsObjectOn(obj105) && !ns.IsObjectOn(obj106)) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar113)
LABEL1:
	return
}
func ToggleSG1() {
	ns.ObjectGroupToggle(gvar107)
	OpenOrbWallGroup1()
}
func EnableSG1() {
	ns.ObjectGroupOn(gvar107)
}
func ToggleSG2() {
	ns.ObjectGroupToggle(gvar108)
	OpenOrbWallGroup1()
}
func EnableSG2() {
	ns.ObjectGroupOn(gvar108)
}
func ToggleSG3() {
	ns.ObjectGroupToggle(gvar109)
	OpenOrbWallGroup1()
}
func EnableSG3() {
	ns.ObjectGroupOn(gvar109)
}
func ToggleSG4() {
	ns.ObjectGroupToggle(gvar110)
	OpenOrbWallGroup1()
}
func EnableSG4() {
	ns.ObjectGroupOn(gvar110)
}
func OpenOrbWallGroup2() {
	ns.WallGroupOpen(gvar114)
	ns.ObjectGroupOn(gvar111)
	ns.ObjectGroupOn(gvar87)
	ns.WallGroupOpen(gvar88)
}
func OpenOrbWallGroup3() {
	ns.WallGroupOpen(gvar115)
}
func OpenOrbWallGroup4() {
	ns.WallGroupOpen(gvar14)
	ns.ObjectGroupOff(gvar112)
	ns.ObjectGroupOff(gvar107)
	ns.ObjectGroupOff(gvar108)
	ns.ObjectGroupOff(gvar109)
	ns.ObjectGroupOff(gvar110)
}
func BossDieCheck() {
	if !(flag116 && flag117 && flag118 && flag119) {
		goto LABEL1
	}
	OpenOrbWallGroup3()
LABEL1:
	return
}
func BossDie1() {
	flag116 = true
	BossDieCheck()
}
func BossDie2() {
	flag117 = true
	BossDieCheck()
}
func BossDie3() {
	flag118 = true
	BossDieCheck()
}
func BossDie4() {
	flag119 = true
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
