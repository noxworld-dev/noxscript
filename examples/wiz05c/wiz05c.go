package wiz05c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4   ns.WallGroupID
	gvar5   ns.WallGroupID
	gvar6   ns.WallGroupID
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
	gvar52  ns.ObjectGroupID
	obj53   ns.ObjectID
	obj54   ns.ObjectID
	obj55   ns.ObjectID
	obj56   ns.ObjectID
	gvar57  int
	gvar58  ns.WallGroupID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	obj63   ns.ObjectID
	obj64   ns.ObjectID
	obj65   ns.ObjectID
	obj66   ns.ObjectID
	obj67   ns.ObjectID
	obj68   ns.ObjectID
	obj69   ns.ObjectID
	obj70   ns.ObjectID
	obj71   ns.ObjectID
	obj72   ns.ObjectID
	obj73   ns.ObjectID
	obj74   ns.ObjectID
	obj75   ns.ObjectID
	obj76   ns.ObjectID
	obj77   ns.ObjectID
	obj78   ns.ObjectID
	obj79   ns.ObjectID
	obj80   ns.ObjectID
	obj81   ns.ObjectID
	obj82   ns.ObjectID
	obj83   ns.ObjectID
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
	gvar96  ns.WallGroupID
	wp97    ns.WaypointID
	wp98    ns.WaypointID
	wp99    ns.WaypointID
	wp100   ns.WaypointID
	wp101   ns.WaypointID
	wp102   ns.WaypointID
	wp103   ns.WaypointID
	wp104   ns.WaypointID
	wp105   ns.WaypointID
	wp106   ns.WaypointID
	wp107   ns.WaypointID
	wp108   ns.WaypointID
	wp109   ns.WaypointID
	wp110   ns.WaypointID
	wp111   ns.WaypointID
	wp112   ns.WaypointID
	wp113   ns.WaypointID
	wp114   ns.WaypointID
	wp115   ns.WaypointID
	wp116   ns.WaypointID
	wp117   ns.WaypointID
	wp118   ns.WaypointID
	wp119   ns.WaypointID
	wp120   ns.WaypointID
	wp121   ns.WaypointID
	gvar122 int
	obj123  ns.ObjectID
	obj124  ns.ObjectID
	obj125  ns.ObjectID
	obj126  ns.ObjectID
	wp127   ns.WaypointID
	wp128   ns.WaypointID
	obj129  ns.ObjectID
	obj130  ns.ObjectID
	obj131  ns.ObjectID
	obj132  ns.ObjectID
	wp133   ns.WaypointID
	wp134   ns.WaypointID
)

func DisableElevators() {
	ns.ObjectOff(obj24)
	ns.ObjectOff(obj25)
	ns.ObjectOff(obj26)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	ns.ObjectOff(obj29)
}
func EnableSkeletons() {
	ns.AggressionLevel(obj30, 0.83)
	ns.AggressionLevel(obj31, 0.83)
	ns.AggressionLevel(obj32, 0.83)
	ns.AggressionLevel(obj34, 0.83)
	ns.AggressionLevel(obj35, 0.83)
}
func BreakWalls1() {
	ns.WallGroupOpen(gvar4)
}
func EnableElevators1() {
	ns.ObjectOn(obj24)
	ns.ObjectOn(obj25)
	ns.ObjectOn(obj26)
	ns.ObjectOn(obj27)
	ns.ObjectOn(obj28)
	ns.ObjectOn(obj29)
	ns.FrameTimer(20, SkeletonGoal)
	EnableSkeletons()
}
func SkeletonGoal() {
	ns.Move(obj30, wp51)
	ns.Move(obj31, wp51)
	ns.Move(obj32, wp51)
	ns.Move(obj34, wp51)
	ns.Move(obj35, wp51)
}
func EndTrap1() {
	ns.Move(obj9, wp41)
	ns.Move(obj10, wp42)
	ns.Move(obj11, wp43)
	ns.Move(obj12, wp44)
	ns.ObjectOff(obj36)
	ns.AudioEvent(ns.EarthRumbleMinor, wp37)
	ns.AudioEvent(ns.SecretWallSecret, wp37)
}
func OpenGates() {
	ns.WallOpen(ns.Wall(96, 208))
	ns.WallOpen(ns.Wall(95, 209))
	ns.WallOpen(ns.Wall(94, 210))
}
func CloseBlocks() {
	ns.Move(obj7, wp45)
	ns.Move(obj8, wp46)
	ns.AudioEvent(ns.BoulderMove, wp38)
	ns.AudioEvent(ns.EarthRumbleMinor, wp38)
	ns.AudioEvent(ns.SecretWallOpen, wp38)
	ns.AudioEvent(ns.StoneBlockMove, wp38)
}
func OpenBlocks() {
	ns.ObjectOff(ns.GetTrigger())
	ns.Move(obj7, wp39)
	ns.Move(obj8, wp40)
	ns.AudioEvent(ns.BoulderMove, wp38)
	ns.AudioEvent(ns.EarthRumbleMinor, wp38)
	ns.AudioEvent(ns.SecretWallOpen, wp38)
	ns.AudioEvent(ns.StoneBlockMove, wp38)
}
func Trap1() {
	ns.ObjectOff(obj13)
	ns.ObjectOff(obj14)
	ns.ObjectOff(obj15)
	ns.ObjectOff(obj16)
	ns.ObjectOff(obj17)
	ns.Move(obj9, wp47)
	ns.Move(obj10, wp48)
	ns.Move(obj11, wp49)
	ns.Move(obj12, wp50)
	ns.AudioEvent(ns.EarthRumbleMajor, wp50)
	ns.SecondTimer(1, BreakWalls1)
	ns.SecondTimer(2, EnableElevators1)
	ns.Music(27, 100)
}
func CloseBlocks2() {
	ns.Move(obj7, wp45)
	ns.Move(obj8, wp46)
	ns.Move(obj9, wp47)
	ns.Move(obj10, wp48)
}
func LockCellDoors() {
	ns.LockDoor(obj68)
	ns.LockDoor(obj69)
	ns.LockDoor(obj70)
	ns.LockDoor(obj71)
	ns.LockDoor(obj72)
	ns.LockDoor(obj73)
	ns.LockDoor(obj74)
	ns.LockDoor(obj75)
	ns.LockDoor(obj76)
	ns.LockDoor(obj77)
	ns.LockDoor(obj78)
	ns.LockDoor(obj79)
}
func EnableSpikes() {
	ns.Move(obj81, wp110)
	ns.Move(obj82, wp111)
	ns.Move(obj83, wp112)
	ns.Move(obj84, wp113)
	ns.Move(obj85, wp114)
	ns.Move(obj86, wp115)
	ns.Move(obj87, wp116)
	ns.Move(obj88, wp117)
	ns.Move(obj89, wp118)
	ns.Move(obj90, wp119)
	ns.Move(obj91, wp120)
	ns.Move(obj92, wp121)
	ns.FrameTimer(50, DisableSpikes)
}
func DisableSpikes() {
	ns.Move(obj81, wp98)
	ns.Move(obj82, wp99)
	ns.Move(obj83, wp100)
	ns.Move(obj84, wp101)
	ns.Move(obj85, wp102)
	ns.Move(obj86, wp103)
	ns.Move(obj87, wp104)
	ns.Move(obj88, wp105)
	ns.Move(obj89, wp106)
	ns.Move(obj90, wp107)
	ns.Move(obj91, wp108)
	ns.Move(obj92, wp109)
	ns.FrameTimer(50, EnableSpikes)
}
func EnableGears() {
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj60)
	ns.SecondTimer(2, DisableGears)
}
func DisableGears() {
	ns.ObjectOff(obj59)
	ns.ObjectOff(obj60)
}
func KillSkeleton7() {
	ns.Damage(obj67, 0, 1000, 0)
}
func LetterBoxOff() {
	ns.WideScreen(false)
}
func PlayerDeath() {
	ns.Blind()
	ns.FrameTimer(30, ChapterFail2)
}
func SpikeNoise1() {
	ns.AudioEvent(ns.EarthRumbleMajor, wp119)
	ns.AudioEvent(ns.EarthRumbleMajor, wp113)
	ns.FrameTimer(60, SpikeNoise1)
}
func HorvathDeath() {
	ns.PrintToAll("Wiz05C.scr:ObjectiveFailed")
	ns.FrameTimer(30, PlayerDeath)
}
func MapInitialize() {
	wp133 = ns.Waypoint("Secret01WP")
	wp134 = ns.Waypoint("Secret02WP")
	wp97 = ns.Waypoint("FadeoutExitWP")
	gvar52 = ns.ObjectGroup("BatEntryGroup")
	obj53 = ns.Object("Bear1")
	obj93 = ns.Object("Ogre1")
	obj94 = ns.Object("Ogre2")
	gvar58 = ns.WallGroup("ScorpionWallGroup")
	obj7 = ns.Object("Block1")
	obj8 = ns.Object("Block2")
	obj9 = ns.Object("Block3")
	obj10 = ns.Object("Block4")
	obj11 = ns.Object("Block5")
	obj12 = ns.Object("Block6")
	obj18 = ns.Object("BlockMover1")
	obj19 = ns.Object("BlockMover2")
	obj20 = ns.Object("BlockMover3")
	obj21 = ns.Object("BlockMover4")
	obj22 = ns.Object("BlockMover5")
	obj23 = ns.Object("BlockMover6")
	obj129 = ns.Object("SilverKeyTrigger1")
	obj130 = ns.Object("SilverKeyTrigger2")
	obj131 = ns.Object("SilverKeyTrigger3")
	obj132 = ns.Object("SilverKeyTrigger4")
	gvar96 = ns.WallGroup("SilverKeyWallGroup")
	obj123 = ns.Object("Grunt1")
	obj124 = ns.Object("Grunt2")
	obj125 = ns.Object("Grunt3")
	obj126 = ns.Object("Grunt4")
	gvar4 = ns.WallGroup("AmbushWallGroup")
	obj61 = ns.Object("Pit1")
	obj62 = ns.Object("Pit2")
	obj63 = ns.Object("Pit3")
	obj64 = ns.Object("Pit4")
	obj65 = ns.Object("Pit5")
	obj66 = ns.Object("Pit6")
	obj24 = ns.Object("Elevator1")
	obj25 = ns.Object("Elevator2")
	obj26 = ns.Object("Elevator3")
	obj27 = ns.Object("Elevator4")
	obj28 = ns.Object("Elevator5")
	obj29 = ns.Object("Elevator6")
	obj30 = ns.Object("Skeleton1")
	obj31 = ns.Object("Skeleton2")
	obj32 = ns.Object("Skeleton3")
	obj33 = ns.Object("Skeleton4")
	obj34 = ns.Object("Skeleton5")
	obj35 = ns.Object("Skeleton6")
	obj67 = ns.Object("Skeleton7")
	obj36 = ns.Object("EndTrapTrigger")
	obj59 = ns.Object("Gear1")
	obj60 = ns.Object("Gear2")
	obj13 = ns.Object("Plate1")
	obj14 = ns.Object("Plate2")
	obj15 = ns.Object("Plate3")
	obj16 = ns.Object("Plate4")
	obj17 = ns.Object("Plate5")
	obj81 = ns.Object("Spike1")
	obj82 = ns.Object("Spike2")
	obj83 = ns.Object("Spike3")
	obj84 = ns.Object("Spike4")
	obj85 = ns.Object("Spike5")
	obj86 = ns.Object("Spike6")
	obj87 = ns.Object("Spike7")
	obj88 = ns.Object("Spike8")
	obj89 = ns.Object("Spike9")
	obj90 = ns.Object("Spike10")
	obj91 = ns.Object("Spike11")
	obj92 = ns.Object("Spike12")
	obj68 = ns.Object("CellDoor1")
	obj69 = ns.Object("CellDoor2")
	obj70 = ns.Object("CellDoor3")
	obj71 = ns.Object("CellDoor4")
	obj72 = ns.Object("CellDoor5")
	obj73 = ns.Object("CellDoor6")
	obj74 = ns.Object("CellDoor7")
	obj75 = ns.Object("CellDoor8")
	obj76 = ns.Object("CellDoor9")
	obj77 = ns.Object("CellDoor10")
	obj78 = ns.Object("CellDoor11")
	obj79 = ns.Object("CellDoor12")
	obj80 = ns.Object("CellLever")
	wp39 = ns.Waypoint("BlockOpen1")
	wp45 = ns.Waypoint("BlockClosed1")
	wp40 = ns.Waypoint("BlockOpen2")
	wp46 = ns.Waypoint("BlockClosed2")
	wp41 = ns.Waypoint("BlockOpen3")
	wp47 = ns.Waypoint("BlockClosed3")
	wp42 = ns.Waypoint("BlockOpen4")
	wp48 = ns.Waypoint("BlockClosed4")
	wp43 = ns.Waypoint("BlockOpen5")
	wp49 = ns.Waypoint("BlockClosed5")
	wp44 = ns.Waypoint("BlockOpen6")
	wp50 = ns.Waypoint("BlockClosed6")
	wp127 = ns.Waypoint("Gates1WP")
	wp128 = ns.Waypoint("Gates2WP")
	gvar5 = ns.WallGroup("Gates1")
	gvar6 = ns.WallGroup("Gates2")
	wp38 = ns.Waypoint("BlockAudio1WP")
	wp98 = ns.Waypoint("Spike1Open")
	wp99 = ns.Waypoint("Spike2Open")
	wp100 = ns.Waypoint("Spike3Open")
	wp101 = ns.Waypoint("Spike4Open")
	wp102 = ns.Waypoint("Spike5Open")
	wp103 = ns.Waypoint("Spike6Open")
	wp104 = ns.Waypoint("Spike7Open")
	wp105 = ns.Waypoint("Spike8Open")
	wp106 = ns.Waypoint("Spike9Open")
	wp107 = ns.Waypoint("Spike10Open")
	wp108 = ns.Waypoint("Spike11Open")
	wp109 = ns.Waypoint("Spike12Open")
	wp110 = ns.Waypoint("Spike1Closed")
	wp111 = ns.Waypoint("Spike2Closed")
	wp112 = ns.Waypoint("Spike3Closed")
	wp113 = ns.Waypoint("Spike4Closed")
	wp114 = ns.Waypoint("Spike5Closed")
	wp115 = ns.Waypoint("Spike6Closed")
	wp116 = ns.Waypoint("Spike7Closed")
	wp117 = ns.Waypoint("Spike8Closed")
	wp118 = ns.Waypoint("Spike9Closed")
	wp119 = ns.Waypoint("Spike10Closed")
	wp120 = ns.Waypoint("Spike11Closed")
	wp121 = ns.Waypoint("Spike12Closed")
	wp51 = ns.Waypoint("SkeletonGoal")
	wp37 = ns.Waypoint("EndTrapAudioWP")
	OpenGates()
	CloseBlocks()
	LockCellDoors()
	KillSkeleton7()
	ns.Move(obj11, wp49)
	ns.Move(obj12, wp50)
	EnableSpikes()
}
func BatsHostilize() {
	ns.GroupAggressionLevel(gvar52, 0.83)
}
func BearHostilize() {
	ns.AggressionLevel(obj53, 0.83)
}
func OgresHostilize() {
	ns.AggressionLevel(obj93, 0.83)
	ns.AggressionLevel(obj94, 0.83)
}
func DestroyWalls() {
	ns.WallGroupBreak(gvar58)
	ns.Attack(obj54, ns.GetHost())
	ns.Attack(obj55, ns.GetHost())
	ns.Attack(obj56, ns.GetHost())
}
func OpenCellDoors() {
	ns.UnlockDoor(obj68)
	ns.UnlockDoor(obj69)
	ns.UnlockDoor(obj70)
	ns.UnlockDoor(obj71)
	ns.UnlockDoor(obj72)
	ns.UnlockDoor(obj73)
	ns.UnlockDoor(obj74)
	ns.UnlockDoor(obj75)
	ns.UnlockDoor(obj76)
	ns.UnlockDoor(obj77)
	ns.UnlockDoor(obj78)
	ns.UnlockDoor(obj79)
	ns.ObjectOff(obj80)
	EnableGears()
}
func ToggleGates1() {
	ns.AudioEvent(ns.FireBallExplode, wp127)
	ns.WallGroupOpen(gvar5)
}
func ToggleGates2() {
	ns.AudioEvent(ns.FireBallExplode, wp128)
	ns.WallGroupOpen(gvar6)
}
func ChapterFail2() {
	ns.EnchantOff(ns.GetHost(), ns.ENCHANT_INVULNERABLE)
	ns.DeathScreen(5)
}
func SilverKeyTrap() {
	ns.ObjectOff(obj129)
	ns.ObjectOff(obj130)
	ns.ObjectOff(obj131)
	ns.ObjectOff(obj132)
	ns.WallGroupBreak(gvar96)
	ns.AggressionLevel(obj123, 0.83)
	ns.AggressionLevel(obj124, 0.83)
	ns.AggressionLevel(obj125, 0.83)
	ns.AggressionLevel(obj126, 0.83)
}
func MapEntry() {
	ns.UnBlind()
	LetterBoxOff()
	ns.Frozen(ns.GetHost(), false)
	if ns.GetQuestStatus("Wiz05B:GotHorvath") != 1 {
		goto LABEL1
	}
	obj95 = ns.Object("Wiz05B:Horvath")
	ns.SetCallback(obj95, 5, HorvathDeath)
LABEL1:
	return
}
func PlaySubMusic() {
	ns.Music(20, 100)
}
func MonsterGoHome() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func Secret01() {
	ns.AudioEvent(ns.SecretFound, wp133)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func Secret02() {
	ns.AudioEvent(ns.SecretFound, wp134)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	}
}
