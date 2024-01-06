package war05c

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
	obj57   ns.ObjectID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	gvar62  int
	gvar63  ns.WallGroupID
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
	obj96   ns.ObjectID
	obj97   ns.ObjectID
	obj98   ns.ObjectID
	obj99   ns.ObjectID
	gvar100 ns.WallGroupID
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
	wp122   ns.WaypointID
	wp123   ns.WaypointID
	wp124   ns.WaypointID
	gvar125 int
	obj126  ns.ObjectID
	obj127  ns.ObjectID
	obj128  ns.ObjectID
	obj129  ns.ObjectID
	wp130   ns.WaypointID
	wp131   ns.WaypointID
	obj132  ns.ObjectID
	obj133  ns.ObjectID
	obj134  ns.ObjectID
	obj135  ns.ObjectID
	wp136   ns.WaypointID
	wp137   ns.WaypointID
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
	ns.AggressionLevel(obj33, 0.83)
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
	ns.Move(obj33, wp51)
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
	ns.LockDoor(obj73)
	ns.LockDoor(obj74)
	ns.LockDoor(obj75)
	ns.LockDoor(obj76)
	ns.LockDoor(obj77)
	ns.LockDoor(obj78)
	ns.LockDoor(obj79)
	ns.LockDoor(obj80)
	ns.LockDoor(obj81)
	ns.LockDoor(obj82)
	ns.LockDoor(obj83)
	ns.LockDoor(obj84)
}
func EnableSpikes() {
	ns.Move(obj86, wp113)
	ns.Move(obj87, wp114)
	ns.Move(obj88, wp115)
	ns.Move(obj89, wp116)
	ns.Move(obj90, wp117)
	ns.Move(obj91, wp118)
	ns.Move(obj92, wp119)
	ns.Move(obj93, wp120)
	ns.Move(obj94, wp121)
	ns.Move(obj95, wp122)
	ns.Move(obj96, wp123)
	ns.Move(obj97, wp124)
	ns.FrameTimer(50, DisableSpikes)
}
func DisableSpikes() {
	ns.Move(obj86, wp101)
	ns.Move(obj87, wp102)
	ns.Move(obj88, wp103)
	ns.Move(obj89, wp104)
	ns.Move(obj90, wp105)
	ns.Move(obj91, wp106)
	ns.Move(obj92, wp107)
	ns.Move(obj93, wp108)
	ns.Move(obj94, wp109)
	ns.Move(obj95, wp110)
	ns.Move(obj96, wp111)
	ns.Move(obj97, wp112)
	ns.FrameTimer(50, EnableSpikes)
}
func EnableGears() {
	ns.ObjectOn(obj64)
	ns.ObjectOn(obj65)
	ns.SecondTimer(2, DisableGears)
}
func DisableGears() {
	ns.ObjectOff(obj64)
	ns.ObjectOff(obj65)
}
func KillSkeleton7() {
	ns.Damage(obj72, 0, 1000, 0)
}
func MaidenDeath() {
	ns.PrintToAll("War05B.scr:MaidenDeath")
	ns.DeathScreen(5)
}
func LetterBoxOff() {
	ns.WideScreen(false)
}
func PlayerDeath() {
	ns.PrintToAll("War05A.scr:PlayerDeath")
	ns.Blind()
	ns.FrameTimer(30, ChapterFail2)
}
func SpikeNoise1() {
	ns.AudioEvent(ns.EarthRumbleMajor, wp122)
	ns.AudioEvent(ns.EarthRumbleMajor, wp116)
	ns.FrameTimer(60, SpikeNoise1)
}
func MapInitialize() {
	wp136 = ns.Waypoint("Secret01WP")
	wp137 = ns.Waypoint("Secret02WP")
	gvar52 = ns.ObjectGroup("BatEntryGroup")
	obj58 = ns.Object("Bear1")
	obj98 = ns.Object("Ogre1")
	obj99 = ns.Object("Ogre2")
	gvar63 = ns.WallGroup("ScorpionWallGroup")
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
	obj132 = ns.Object("SilverKeyTrigger1")
	obj133 = ns.Object("SilverKeyTrigger2")
	obj134 = ns.Object("SilverKeyTrigger3")
	obj135 = ns.Object("SilverKeyTrigger4")
	gvar100 = ns.WallGroup("SilverKeyWallGroup")
	obj126 = ns.Object("Grunt1")
	obj127 = ns.Object("Grunt2")
	obj128 = ns.Object("Grunt3")
	obj129 = ns.Object("Grunt4")
	gvar4 = ns.WallGroup("AmbushWallGroup")
	obj66 = ns.Object("Pit1")
	obj67 = ns.Object("Pit2")
	obj68 = ns.Object("Pit3")
	obj69 = ns.Object("Pit4")
	obj70 = ns.Object("Pit5")
	obj71 = ns.Object("Pit6")
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
	obj72 = ns.Object("Skeleton7")
	obj36 = ns.Object("EndTrapTrigger")
	obj64 = ns.Object("Gear1")
	obj65 = ns.Object("Gear2")
	obj13 = ns.Object("Plate1")
	obj14 = ns.Object("Plate2")
	obj15 = ns.Object("Plate3")
	obj16 = ns.Object("Plate4")
	obj17 = ns.Object("Plate5")
	obj86 = ns.Object("Spike1")
	obj87 = ns.Object("Spike2")
	obj88 = ns.Object("Spike3")
	obj89 = ns.Object("Spike4")
	obj90 = ns.Object("Spike5")
	obj91 = ns.Object("Spike6")
	obj92 = ns.Object("Spike7")
	obj93 = ns.Object("Spike8")
	obj94 = ns.Object("Spike9")
	obj95 = ns.Object("Spike10")
	obj96 = ns.Object("Spike11")
	obj97 = ns.Object("Spike12")
	obj73 = ns.Object("CellDoor1")
	obj74 = ns.Object("CellDoor2")
	obj75 = ns.Object("CellDoor3")
	obj76 = ns.Object("CellDoor4")
	obj77 = ns.Object("CellDoor5")
	obj78 = ns.Object("CellDoor6")
	obj79 = ns.Object("CellDoor7")
	obj80 = ns.Object("CellDoor8")
	obj81 = ns.Object("CellDoor9")
	obj82 = ns.Object("CellDoor10")
	obj83 = ns.Object("CellDoor11")
	obj84 = ns.Object("CellDoor12")
	obj85 = ns.Object("CellLever")
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
	wp130 = ns.Waypoint("Gates1WP")
	wp131 = ns.Waypoint("Gates2WP")
	gvar5 = ns.WallGroup("Gates1")
	gvar6 = ns.WallGroup("Gates2")
	wp38 = ns.Waypoint("BlockAudio1WP")
	wp101 = ns.Waypoint("Spike1Open")
	wp102 = ns.Waypoint("Spike2Open")
	wp103 = ns.Waypoint("Spike3Open")
	wp104 = ns.Waypoint("Spike4Open")
	wp105 = ns.Waypoint("Spike5Open")
	wp106 = ns.Waypoint("Spike6Open")
	wp107 = ns.Waypoint("Spike7Open")
	wp108 = ns.Waypoint("Spike8Open")
	wp109 = ns.Waypoint("Spike9Open")
	wp110 = ns.Waypoint("Spike10Open")
	wp111 = ns.Waypoint("Spike11Open")
	wp112 = ns.Waypoint("Spike12Open")
	wp113 = ns.Waypoint("Spike1Closed")
	wp114 = ns.Waypoint("Spike2Closed")
	wp115 = ns.Waypoint("Spike3Closed")
	wp116 = ns.Waypoint("Spike4Closed")
	wp117 = ns.Waypoint("Spike5Closed")
	wp118 = ns.Waypoint("Spike6Closed")
	wp119 = ns.Waypoint("Spike7Closed")
	wp120 = ns.Waypoint("Spike8Closed")
	wp121 = ns.Waypoint("Spike9Closed")
	wp122 = ns.Waypoint("Spike10Closed")
	wp123 = ns.Waypoint("Spike11Closed")
	wp124 = ns.Waypoint("Spike12Closed")
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
	ns.AggressionLevel(obj58, 0.83)
}
func OgresHostilize() {
	ns.AggressionLevel(obj98, 0.83)
	ns.AggressionLevel(obj99, 0.83)
}
func DestroyWalls() {
	ns.WallGroupBreak(gvar63)
	ns.Attack(obj59, ns.GetHost())
	ns.Attack(obj60, ns.GetHost())
	ns.Attack(obj61, ns.GetHost())
}
func OpenCellDoors() {
	ns.UnlockDoor(obj73)
	ns.UnlockDoor(obj74)
	ns.UnlockDoor(obj75)
	ns.UnlockDoor(obj76)
	ns.UnlockDoor(obj77)
	ns.UnlockDoor(obj78)
	ns.UnlockDoor(obj79)
	ns.UnlockDoor(obj80)
	ns.UnlockDoor(obj81)
	ns.UnlockDoor(obj82)
	ns.UnlockDoor(obj83)
	ns.UnlockDoor(obj84)
	ns.ObjectOff(obj85)
	EnableGears()
}
func ToggleGates1() {
	ns.AudioEvent(ns.FireBallExplode, wp130)
	ns.WallGroupOpen(gvar5)
}
func ToggleGates2() {
	ns.AudioEvent(ns.FireBallExplode, wp131)
	ns.WallGroupOpen(gvar6)
}
func ChapterFail2() {
	ns.EnchantOff(ns.GetHost(), ns.ENCHANT_INVULNERABLE)
	ns.DeathScreen(5)
}
func SilverKeyTrap() {
	ns.ObjectOff(obj132)
	ns.ObjectOff(obj133)
	ns.ObjectOff(obj134)
	ns.ObjectOff(obj135)
	ns.WallGroupBreak(gvar100)
	ns.AggressionLevel(obj126, 0.83)
	ns.AggressionLevel(obj127, 0.83)
	ns.AggressionLevel(obj128, 0.83)
	ns.AggressionLevel(obj129, 0.83)
}
func MapEntry() {
	ns.UnBlind()
	LetterBoxOff()
	ns.Frozen(ns.GetHost(), false)
	if ns.GetQuestStatus("MaidensRescued") != 1 {
		goto LABEL1
	}
	obj53 = ns.Object("War05B:W5Maiden1")
	obj54 = ns.Object("War05B:W5Maiden2")
	obj55 = ns.Object("War05B:W5Maiden3")
	obj56 = ns.Object("War05B:W5Maiden4")
	obj57 = ns.Object("War05B:W5Sister")
	ns.SetCallback(obj53, 5, MaidenDeath)
	ns.SetCallback(obj54, 5, MaidenDeath)
	ns.SetCallback(obj55, 5, MaidenDeath)
	ns.SetCallback(obj56, 5, MaidenDeath)
	ns.SetCallback(obj57, 5, MaidenDeath)
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
	ns.AudioEvent(ns.SecretFound, wp136)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func Secret02() {
	ns.AudioEvent(ns.SecretFound, wp137)
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
