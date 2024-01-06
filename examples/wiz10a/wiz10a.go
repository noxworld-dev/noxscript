package wiz10a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4    ns.ObjectID
	obj5    ns.ObjectID
	obj6    ns.ObjectID
	obj7    ns.ObjectID
	obj8    ns.ObjectID
	obj9    ns.ObjectID
	obj10   ns.ObjectID
	obj11   ns.ObjectID
	obj12   ns.ObjectID
	obj13   ns.ObjectID
	obj14   ns.ObjectID
	obj15   ns.ObjectID
	obj16   ns.ObjectID
	obj17   ns.ObjectID
	obj18   ns.ObjectID
	obj19   ns.ObjectID
	obj20   ns.ObjectID
	obj21   ns.ObjectID
	obj22   ns.ObjectID
	obj23   ns.ObjectID
	obj24   ns.ObjectID
	obj25   ns.ObjectID
	obj26   ns.ObjectID
	obj27   ns.ObjectID
	obj28   ns.ObjectID
	obj29   ns.ObjectID
	obj30   ns.ObjectID
	obj31   ns.ObjectID
	obj32   ns.ObjectID
	obj33   ns.ObjectID
	obj34   ns.ObjectID
	obj35   ns.ObjectID
	obj36   ns.ObjectID
	obj37   ns.ObjectID
	obj38   ns.ObjectID
	obj39   ns.ObjectID
	obj40   ns.ObjectID
	obj41   ns.ObjectID
	obj42   ns.ObjectID
	obj43   ns.ObjectID
	obj44   ns.ObjectID
	obj45   ns.ObjectID
	obj46   ns.ObjectID
	obj47   ns.ObjectID
	gvar48  ns.ObjectGroupID
	wp49    ns.WaypointID
	wp50    ns.WaypointID
	wp51    ns.WaypointID
	wp52    ns.WaypointID
	wp53    ns.WaypointID
	wp54    ns.WaypointID
	wp55    ns.WaypointID
	wp56    ns.WaypointID
	wp57    ns.WaypointID
	wp58    ns.WaypointID
	wp59    ns.WaypointID
	wp60    ns.WaypointID
	wp61    ns.WaypointID
	gvar62  ns.WallGroupID
	gvar63  ns.WallGroupID
	gvar64  ns.WallGroupID
	gvar65  ns.ObjectGroupID
	gvar66  ns.ObjectGroupID
	gvar67  ns.ObjectGroupID
	gvar68  ns.ObjectGroupID
	gvar69  ns.ObjectGroupID
	gvar70  ns.ObjectGroupID
	flag71  bool
	flag72  bool
	flag73  bool
	flag74  bool
	gvar75  ns.WallGroupID
	gvar76  ns.WallGroupID
	gvar77  ns.WallGroupID
	gvar78  ns.ObjectGroupID
	obj79   ns.ObjectID
	obj80   ns.ObjectID
	gvar81  ns.WallGroupID
	flag82  bool
	flag83  bool
	wp84    ns.WaypointID
	wp85    ns.WaypointID
	obj86   ns.ObjectID
	obj87   ns.ObjectID
	wp88    ns.WaypointID
	wp89    ns.WaypointID
	obj90   ns.ObjectID
	gvar91  ns.WallGroupID
	gvar92  ns.WallGroupID
	gvar93  ns.ObjectGroupID
	gvar94  ns.ObjectGroupID
	gvar95  ns.ObjectGroupID
	gvar96  ns.ObjectGroupID
	gvar97  ns.ObjectGroupID
	gvar98  ns.ObjectGroupID
	wp99    ns.WaypointID
	wp100   ns.WaypointID
	wp101   ns.WaypointID
	wp102   ns.WaypointID
	wp103   ns.WaypointID
	wp104   ns.WaypointID
	obj105  ns.ObjectID
	obj106  ns.ObjectID
	wp107   ns.WaypointID
	gvar108 ns.ObjectGroupID
	obj109  ns.ObjectID
	obj110  ns.ObjectID
	gvar111 int
	wp112   ns.WaypointID
	wp113   ns.WaypointID
	gvar114 ns.WallGroupID
	gvar115 ns.WallGroupID
	gvar116 ns.ObjectGroupID
	gvar117 ns.WallGroupID
	gvar118 ns.WallGroupID
	wp119   ns.WaypointID
	obj120  ns.ObjectID
	obj121  ns.ObjectID
	obj122  ns.ObjectID
	wp123   ns.WaypointID
	wp124   ns.WaypointID
	fvar125 float32
	fvar126 float32
	flag127 bool
	flag128 bool
)

func init() {
	flag71 = false
	flag72 = false
	flag73 = false
	flag74 = false
	flag83 = false
	flag82 = false
	flag127 = false
	flag128 = false
}
func DisableArrow01() {
	ns.ObjectOff(obj4)
}
func EnableArrow01() {
	ns.ObjectOn(obj4)
	ns.FrameTimer(1, DisableArrow01)
}
func DisableArrow02() {
	ns.ObjectOff(obj5)
}
func EnableArrow02() {
	ns.ObjectOn(obj5)
	ns.FrameTimer(1, DisableArrow02)
}
func DisableArrow03() {
	ns.ObjectOff(obj6)
}
func EnableArrow03() {
	ns.ObjectOn(obj6)
	ns.FrameTimer(1, DisableArrow03)
}
func DisableArrow04() {
	ns.ObjectOff(obj7)
}
func EnableArrow04() {
	ns.ObjectOn(obj7)
	ns.FrameTimer(1, DisableArrow04)
}
func DisableArrow05() {
	ns.ObjectOff(obj8)
}
func EnableArrow05() {
	ns.ObjectOn(obj8)
	ns.FrameTimer(1, DisableArrow05)
}
func DisableArrow06() {
	ns.ObjectOff(obj9)
}
func EnableArrow06() {
	ns.ObjectOn(obj9)
	ns.FrameTimer(1, DisableArrow06)
}
func DisableArrow07() {
	ns.ObjectOff(obj10)
}
func EnableArrow07() {
	ns.ObjectOn(obj10)
	ns.FrameTimer(1, DisableArrow07)
}
func DisableArrow08() {
	ns.ObjectOff(obj11)
}
func EnableArrow08() {
	ns.ObjectOn(obj11)
	ns.FrameTimer(1, DisableArrow08)
}
func DisableArrow09() {
	ns.ObjectOff(obj12)
}
func EnableArrow09() {
	ns.ObjectOn(obj12)
	ns.FrameTimer(1, DisableArrow09)
}
func DisableArrow10() {
	ns.ObjectOff(obj13)
}
func EnableArrow10() {
	ns.ObjectOn(obj13)
	ns.FrameTimer(1, DisableArrow10)
}
func DisableArrow11() {
	ns.ObjectOff(obj14)
}
func EnableArrow11() {
	ns.ObjectOn(obj14)
	ns.FrameTimer(1, DisableArrow11)
}
func DisableArrow12() {
	ns.ObjectOff(obj15)
}
func EnableArrow12() {
	ns.ObjectOn(obj15)
	ns.FrameTimer(1, DisableArrow12)
}
func DisableArrow13() {
	ns.ObjectOff(obj16)
}
func EnableArrow13() {
	ns.ObjectOn(obj16)
	ns.FrameTimer(1, DisableArrow13)
}
func DisableArrow14() {
	ns.ObjectOff(obj17)
}
func EnableArrow14() {
	ns.ObjectOn(obj17)
	ns.FrameTimer(1, DisableArrow14)
}
func DisableArrow15() {
	ns.ObjectOff(obj18)
}
func EnableArrow15() {
	ns.ObjectOn(obj18)
	ns.FrameTimer(1, DisableArrow15)
}
func DisableArrow25() {
	ns.ObjectOff(obj28)
}
func EnableArrow25() {
	ns.ObjectOn(obj28)
	ns.FrameTimer(1, DisableArrow25)
}
func DisableArrow26() {
	ns.ObjectOff(obj29)
}
func EnableArrow26() {
	ns.ObjectOn(obj29)
	ns.FrameTimer(1, DisableArrow26)
}
func DisableArrow27() {
	ns.ObjectOff(obj30)
}
func EnableArrow27() {
	ns.ObjectOn(obj30)
	ns.FrameTimer(1, DisableArrow27)
}
func DisableArrow28() {
	ns.ObjectOff(obj31)
}
func EnableArrow28() {
	ns.ObjectOn(obj31)
	ns.FrameTimer(1, DisableArrow28)
}
func DisableArrow29() {
	ns.ObjectOff(obj32)
}
func EnableArrow29() {
	ns.ObjectOn(obj32)
	ns.FrameTimer(1, DisableArrow29)
}
func DisableArrow30() {
	ns.ObjectOff(obj33)
}
func EnableArrow30() {
	ns.ObjectOn(obj33)
	ns.FrameTimer(1, DisableArrow30)
}
func DisableArrow31() {
	ns.ObjectOff(obj34)
}
func EnableArrow31() {
	ns.ObjectOn(obj34)
	ns.FrameTimer(1, DisableArrow31)
}
func DisableArrow32() {
	ns.ObjectOff(obj35)
}
func EnableArrow32() {
	ns.ObjectOn(obj35)
	ns.FrameTimer(1, DisableArrow32)
}
func DisableArrow33() {
	ns.ObjectOff(obj36)
}
func EnableArrow33() {
	ns.ObjectOn(obj36)
	ns.FrameTimer(1, DisableArrow33)
}
func DisableArrow34() {
	ns.ObjectOff(obj37)
}
func EnableArrow34() {
	ns.ObjectOn(obj37)
	ns.FrameTimer(1, DisableArrow34)
}
func DisableArrow35() {
	ns.ObjectOff(obj38)
}
func EnableArrow35() {
	ns.ObjectOn(obj38)
	ns.FrameTimer(1, DisableArrow35)
}
func DisableArrow36() {
	ns.ObjectOff(obj39)
}
func EnableArrow36() {
	ns.ObjectOn(obj39)
	ns.FrameTimer(1, DisableArrow36)
}
func DisableArrow37() {
	ns.ObjectOff(obj40)
}
func EnableArrow37() {
	ns.ObjectOn(obj40)
	ns.FrameTimer(1, DisableArrow37)
}
func DisableArrow38() {
	ns.ObjectOff(obj41)
}
func EnableArrow38() {
	ns.ObjectOn(obj41)
	ns.FrameTimer(1, DisableArrow38)
}
func DisableArrowGroup1() {
	ns.ObjectOff(obj19)
	ns.ObjectOff(obj20)
	ns.ObjectOff(obj21)
	ns.ObjectOff(obj22)
}
func EnableArrowGroup1() {
	ns.ObjectOn(obj19)
	ns.ObjectOn(obj20)
	ns.ObjectOn(obj21)
	ns.ObjectOn(obj22)
	ns.FrameTimer(2, DisableArrowGroup1)
}
func DisableArrowGroup2() {
	ns.ObjectOff(obj24)
	ns.ObjectOff(obj25)
	ns.ObjectOff(obj26)
	ns.ObjectOff(obj27)
}
func EnableArrowGroup2() {
	ns.ObjectOn(obj24)
	ns.ObjectOn(obj25)
	ns.ObjectOn(obj26)
	ns.ObjectOn(obj27)
	ns.FrameTimer(2, DisableArrowGroup2)
}
func Blocks2MoveBack() {
	ns.Move(obj45, wp55)
	ns.Move(obj46, wp56)
	ns.Move(obj47, wp57)
	ns.AudioEvent(ns.SpikeBlockMove, wp52)
}
func Blocks2Move() {
	ns.ObjectGroupOff(gvar48)
	ns.Move(obj45, wp52)
	ns.Move(obj46, wp53)
	ns.Move(obj47, wp54)
	ns.AudioEvent(ns.SpikeBlockMove, wp55)
}
func OpenBlockWallGroup1() {
	ns.WallGroupOpen(gvar62)
}
func OpenBlockWallGroup2() {
	ns.WallGroupOpen(gvar63)
}
func OpenBlockWallGroup3() {
	ns.WallGroupOpen(gvar64)
}
func FilterBlock5() int {
	if !ns.IsCaller(obj46) {
		goto LABEL1
	}
	return 1
	goto LABEL2
LABEL1:
	return 0
LABEL2:
	return 0
}
func EnableSpikeGroup1() {
	ns.ObjectGroupOn(gvar78)
	ns.AudioEvent(ns.FloorSpikesUp, wp53)
}
func UnlockSpikeDoor() {
	ns.UnlockDoor(obj79)
}
func LockSpikeDoor() {
	ns.LockDoor(obj79)
}
func SkelGroup1Hostilize() {
	ns.GroupAggressionLevel(gvar65, 0.83)
}
func SkelGroup2Hostilize() {
	ns.GroupAggressionLevel(gvar66, 0.83)
}
func SkelGroup3Hostilize() {
	ns.GroupAggressionLevel(gvar67, 0.83)
}
func SkelGroup4Hostilize() {
	ns.GroupAggressionLevel(gvar68, 0.83)
}
func SkelGroup5Hostilize() {
	ns.GroupAggressionLevel(gvar69, 0.83)
}
func SkelGroup6Hostilize() {
	ns.GroupAggressionLevel(gvar70, 0.83)
}
func OpenSecretWallGroup1() {
	ns.WallGroupOpen(gvar75)
}
func OpenSecretWallGroup2() {
	ns.WallGroupOpen(gvar76)
}
func OpenSecretWallGroup3() {
	ns.WallGroupOpen(gvar77)
}
func Blocks1Move() {
	ns.Move(obj42, wp49)
	ns.Move(obj43, wp50)
	ns.Move(obj44, wp51)
	ns.AudioEvent(ns.SpikeBlockMove, wp58)
	ns.AudioEvent(ns.SpikeBlockMove, wp60)
	ns.AudioEvent(ns.SpikeBlockMove, wp61)
}
func Lever1Check() {
	ns.WallGroupOpen(gvar63)
	if flag71 {
		goto LABEL1
	}
	flag71 = true
	ns.AudioEvent(ns.Gear3, wp61)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	if !(flag71 == true && flag72 == true) {
		goto LABEL2
	}
	Blocks1Move()
LABEL2:
	return
}
func Lever2Check() {
	ns.WallGroupOpen(gvar62)
	if flag72 {
		goto LABEL1
	}
	flag72 = true
	ns.AudioEvent(ns.Gear3, wp60)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	if !(flag71 == true && flag72 == true) {
		goto LABEL2
	}
	Blocks1Move()
LABEL2:
	return
}
func Lever3Check() {
	if flag73 {
		goto LABEL1
	}
	flag73 = true
	ns.AudioEvent(ns.Gear3, wp59)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	if !(flag73 == true && flag74 == true) {
		goto LABEL2
	}
	OpenSecretWallGroup3()
	ns.ObjectOn(obj80)
LABEL2:
	return
}
func Lever4Check() {
	if flag74 {
		goto LABEL1
	}
	flag74 = true
	ns.AudioEvent(ns.Gear3, wp59)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	if !(flag73 == true && flag74 == true) {
		goto LABEL2
	}
	OpenSecretWallGroup3()
	ns.ObjectOn(obj80)
LABEL2:
	return
}
func PlayerDeath() {
	ns.DeathScreen(10)
}
func OpenExitWallGroup() {
	ns.WallGroupOpen(gvar81)
}
func PuzzleCheck() {
	if !(flag82 && flag83) {
		goto LABEL1
	}
	OpenExitWallGroup()
LABEL1:
	return
}
func LeftPuzzleActivation() {
	flag83 = true
	PuzzleCheck()
}
func RightPuzzleActivation() {
	flag82 = true
	ns.AudioEvent(ns.CreatureCageAppears, wp84)
	PuzzleCheck()
}
func LeftPuzzleRelease() {
	flag83 = false
	ns.AudioEvent(ns.CreatureCageAppears, wp85)
	PuzzleCheck()
}
func RightPuzzleRelease() {
	flag82 = false
	PuzzleCheck()
}
func OrbLightOff() {
	ns.ObjectOff(obj86)
	ns.MoveObject(obj87, ns.GetWaypointX(wp88), ns.GetWaypointY(wp88))
}
func ForceOrb() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.RandomFloat(-30, 30)
	v1 = ns.RandomFloat(-30, 30)
	v2 = ns.GetObjectX(ns.GetTrigger()) + v0
	v3 = ns.GetObjectY(ns.GetTrigger()) + v1
	ns.MoveObject(obj86, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.MoveObject(obj87, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger())+6)
	ns.MoveWaypoint(wp89, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.ObjectOn(obj86)
	ns.AudioEvent(ns.BallBounce, wp89)
	ns.PushObject(ns.GetCaller(), 60, v2, v3)
	ns.FrameTimer(4, OrbLightOff)
}
func PlaySubMusic() {
	ns.Music(18, 100)
}
func MissionStart() {
	ns.StartupScreen(10)
	ns.JournalEntry(ns.GetHost(), "War10aOrbQuest", 2)
	ns.FrameTimer(10, PlaySubMusic)
}
func InitializeVampireKnights() {
	obj120 = ns.Object("VampireKnight1")
	obj121 = ns.Object("VKBat1")
	wp123 = ns.Waypoint("VampireKnight1WP")
	wp124 = ns.Waypoint("BatCreate")
}
func MapInitialize() {
	obj90 = ns.Object("MissionTrig")
	gvar118 = ns.WallGroup("GoldKeyWallGroup")
	wp119 = ns.Waypoint("GoldKeyExitWP")
	gvar93 = ns.ObjectGroup("PitAmbushGroup")
	gvar91 = ns.WallGroup("SkelPitWallGroup")
	gvar92 = ns.WallGroup("ElevatorWallGroup")
	gvar94 = ns.ObjectGroup("PitSpikesGroup1")
	gvar95 = ns.ObjectGroup("PitSpikesGroup2")
	gvar96 = ns.ObjectGroup("PitSpikesGroup3")
	gvar97 = ns.ObjectGroup("PitSpikesGroup4")
	gvar98 = ns.ObjectGroup("PitSpikesGroup5")
	wp99 = ns.Waypoint("Spikes1AudioWP")
	wp100 = ns.Waypoint("Spikes2AudioWP")
	wp101 = ns.Waypoint("Spikes3AudioWP")
	wp102 = ns.Waypoint("Spikes4AudioWP")
	wp103 = ns.Waypoint("Spikes5AudioWP")
	wp104 = ns.Waypoint("Secret1AudioWP")
	obj4 = ns.Object("Arrow01")
	obj5 = ns.Object("Arrow02")
	obj6 = ns.Object("Arrow03")
	obj7 = ns.Object("Arrow04")
	obj8 = ns.Object("Arrow05")
	obj9 = ns.Object("Arrow06")
	obj10 = ns.Object("Arrow07")
	obj11 = ns.Object("Arrow08")
	obj12 = ns.Object("Arrow09")
	obj13 = ns.Object("Arrow10")
	obj14 = ns.Object("Arrow11")
	obj15 = ns.Object("Arrow12")
	obj16 = ns.Object("Arrow13")
	obj17 = ns.Object("Arrow14")
	obj18 = ns.Object("Arrow15")
	obj19 = ns.Object("Arrow16")
	obj20 = ns.Object("Arrow17")
	obj21 = ns.Object("Arrow18")
	obj22 = ns.Object("Arrow19")
	obj23 = ns.Object("Arrow20")
	obj24 = ns.Object("Arrow21")
	obj25 = ns.Object("Arrow22")
	obj26 = ns.Object("Arrow23")
	obj27 = ns.Object("Arrow24")
	obj28 = ns.Object("Arrow25")
	obj29 = ns.Object("Arrow26")
	obj30 = ns.Object("Arrow27")
	obj31 = ns.Object("Arrow28")
	obj32 = ns.Object("Arrow29")
	obj33 = ns.Object("Arrow30")
	obj34 = ns.Object("Arrow31")
	obj35 = ns.Object("Arrow32")
	obj36 = ns.Object("Arrow33")
	obj37 = ns.Object("Arrow34")
	obj38 = ns.Object("Arrow35")
	obj39 = ns.Object("Arrow36")
	obj40 = ns.Object("Arrow37")
	obj41 = ns.Object("Arrow38")
	obj42 = ns.Object("Block1")
	obj43 = ns.Object("Block2")
	obj44 = ns.Object("Block3")
	obj45 = ns.Object("Block4")
	obj46 = ns.Object("Block5")
	obj47 = ns.Object("Block6")
	gvar48 = ns.ObjectGroup("BlockTriggers")
	wp49 = ns.Waypoint("Block1WP")
	wp50 = ns.Waypoint("Block2WP")
	wp51 = ns.Waypoint("Block3WP")
	wp52 = ns.Waypoint("Block4WP")
	wp53 = ns.Waypoint("Block5WP")
	wp54 = ns.Waypoint("Block6WP")
	wp55 = ns.Waypoint("Block4WPb")
	wp56 = ns.Waypoint("Block5WPb")
	wp57 = ns.Waypoint("Block6WPb")
	wp58 = ns.Waypoint("BlockAudioWP")
	wp59 = ns.Waypoint("BlockAudio2WP")
	wp60 = ns.Waypoint("BlockAudio3WP")
	wp61 = ns.Waypoint("BlockAudio4WP")
	obj105 = ns.Object("Blocks2Trigger")
	gvar62 = ns.WallGroup("BlockWallGroup1")
	gvar63 = ns.WallGroup("BlockWallGroup2")
	gvar64 = ns.WallGroup("BlockWallGroup3")
	gvar78 = ns.ObjectGroup("SpikeGroup1")
	obj79 = ns.Object("SpikeDoor")
	ns.LockDoor(obj79)
	obj106 = ns.Object("StoneBlock1")
	gvar75 = ns.WallGroup("SecretWallGroup1")
	gvar76 = ns.WallGroup("SecretWallGroup2")
	gvar77 = ns.WallGroup("SecretWallGroup3")
	wp107 = ns.Waypoint("Secret2AudioWP")
	gvar108 = ns.ObjectGroup("Secret200Group")
	gvar65 = ns.ObjectGroup("SkelGroup1")
	gvar66 = ns.ObjectGroup("SkelGroup2")
	gvar67 = ns.ObjectGroup("SkelGroup3")
	gvar68 = ns.ObjectGroup("SkelGroup4")
	gvar69 = ns.ObjectGroup("SkelGroup5")
	gvar70 = ns.ObjectGroup("SkelGroup6")
	obj109 = ns.Object("VampireKnight1")
	obj110 = ns.Object("VKbat1")
	wp112 = ns.Waypoint("VampireKnight1WP ")
	wp113 = ns.Waypoint("BatCreate")
	gvar114 = ns.WallGroup("Vampire1WallGroup")
	gvar115 = ns.WallGroup("Vampire1WallGroup2")
	gvar116 = ns.ObjectGroup("Vampire1TrigGroup")
	gvar117 = ns.WallGroup("CherubWallGroup")
	ns.WallGroupOpen(gvar114)
	gvar81 = ns.WallGroup("ExitWallGroup")
	obj86 = ns.Object("OrbLight")
	wp89 = ns.Waypoint("OrbSoundWP")
	obj87 = ns.Object("OrbEffect")
	wp88 = ns.Waypoint("OrbEffectWP")
	wp84 = ns.Waypoint("RightPuzzleWP")
	wp85 = ns.Waypoint("LeftPuzzleWP")
	obj80 = ns.Object("Necro1")
	MissionStart()
	InitializeVampireKnights()
}
func OpenGoldKeyWallGroup() {
	ns.WallGroupOpen(gvar118)
	ns.AudioEvent(ns.FlagDrop, wp119)
}
func OpenCherubWallGroup() {
	ns.WallGroupOpen(gvar117)
}
func Secret100XP() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp104)
}
func Secret200XP() {
	ns.GiveXp(ns.GetHost(), 200)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp107)
	ns.ObjectGroupOff(gvar108)
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(17, 100)
LABEL1:
	return
}
func PlayAction3Music() {
	ns.Music(28, 100)
}
func ToggleSpikes1() {
	ns.ObjectGroupToggle(gvar94)
	ns.AudioEvent(ns.FloorSpikesDown, wp99)
}
func ToggleSpikes2() {
	ns.ObjectGroupToggle(gvar95)
	ns.ObjectGroupToggle(gvar96)
	ns.AudioEvent(ns.FloorSpikesDown, wp100)
	ns.AudioEvent(ns.FloorSpikesDown, wp101)
}
func ToggleSpikes3() {
	ns.ObjectGroupToggle(gvar97)
	ns.ObjectGroupToggle(gvar98)
	ns.AudioEvent(ns.FloorSpikesDown, wp102)
	ns.AudioEvent(ns.FloorSpikesDown, wp103)
}
func OpenSkelPitWallGroup() {
	ns.WallGroupOpen(gvar91)
	ns.GroupAggressionLevel(gvar93, 0.83)
}
func OpenElevatorWallGroup() {
	ns.WallGroupOpen(gvar92)
}
func BatToVampireKnight() {
	ns.Effect(ns.BLUE_SPARKS, fvar125, fvar126, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, fvar125, fvar126, 0, 0)
	ns.Enchant(obj120, ns.ENCHANT_INVISIBLE, 0.25)
	ns.MoveObject(obj120, fvar125, fvar126)
	ns.LookAtObject(obj120, ns.GetHost())
	PlayAction3Music()
}
func SetRetreatBat() {
	ns.AggressionLevel(obj122, 0.83)
	ns.SetCallback(obj122, 7, BatDie)
	ns.SetCallback(obj122, 5, BatDie)
}
func BatDie() {
	var (
		v0 ns.ObjectID
		v1 ns.ObjectID
		v2 ns.ObjectID
	)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj122), ns.GetObjectY(obj122), 0, 0)
	ns.MoveWaypoint(wp124, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.Delete(obj122)
	ns.AudioEvent(ns.BurnCast, wp124)
	v0 = ns.CreateObject("Flame", wp124)
	v1 = ns.CreateObject("MediumFlame", wp124)
	v2 = ns.CreateObject("SmallFlame", wp124)
	ns.DeleteObjectTimer(v0, 80)
	ns.DeleteObjectTimer(v1, 83)
	ns.DeleteObjectTimer(v2, 85)
	ns.WallGroupOpen(gvar114)
	ns.WallGroupOpen(gvar115)
	PlaySubMusic()
	ns.GiveXp(ns.GetHost(), 250)
}
func ChangeOnSight() {
	if !flag128 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.SetCallback(ns.GetTrigger(), 6, BatToVampireKnight)
LABEL1:
	return
}
func InjureVampireKnight() {
	fvar125 = ns.GetObjectX(ns.GetTrigger())
	fvar126 = ns.GetObjectY(ns.GetTrigger())
	ns.ObjectOn(obj120)
	ns.Damage(obj120, 0, 50, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.Delete(ns.GetTrigger())
	ns.FrameTimer(1, BatToVampireKnight)
}
func VKDie() {
	fvar125 = ns.GetObjectX(ns.GetTrigger())
	fvar126 = ns.GetObjectY(ns.GetTrigger())
	ns.Delete(ns.GetTrigger())
	ns.MoveWaypoint(wp124, fvar125, fvar126)
	obj122 = ns.CreateObject("Bat", wp124)
	ns.FrameTimer(1, SetRetreatBat)
}
func OpenVampire1WallGroup() {
	ns.WallGroupOpen(gvar114)
}
func CloseVampire1WallGroup() {
	ns.WallGroupClose(gvar114)
	ns.ObjectGroupOff(gvar116)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
