package wiz09c

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
	obj48   ns.ObjectID
	obj49   ns.ObjectID
	obj50   ns.ObjectID
	obj51   ns.ObjectID
	obj52   ns.ObjectID
	obj53   ns.ObjectID
	obj54   ns.ObjectID
	gvar55  ns.ObjectGroupID
	gvar56  ns.ObjectGroupID
	gvar57  ns.ObjectGroupID
	gvar58  ns.ObjectGroupID
	gvar59  ns.ObjectGroupID
	gvar60  ns.ObjectGroupID
	gvar61  ns.ObjectGroupID
	gvar62  ns.ObjectGroupID
	gvar63  ns.ObjectGroupID
	gvar64  ns.ObjectGroupID
	gvar65  ns.ObjectGroupID
	gvar66  [6]ns.ObjectGroupID
	gvar67  ns.ObjectGroupID
	gvar68  ns.ObjectGroupID
	gvar69  ns.WallGroupID
	gvar70  ns.WallGroupID
	gvar71  ns.WallGroupID
	gvar72  ns.WallGroupID
	gvar73  ns.WallGroupID
	gvar74  ns.WallGroupID
	wp75    ns.WaypointID
	wp76    ns.WaypointID
	wp77    ns.WaypointID
	wp78    ns.WaypointID
	wp79    ns.WaypointID
	wp80    ns.WaypointID
	wp81    ns.WaypointID
	wp82    ns.WaypointID
	wp83    ns.WaypointID
	wp84    ns.WaypointID
	wp85    ns.WaypointID
	wp86    ns.WaypointID
	wp87    ns.WaypointID
	wp88    ns.WaypointID
	wp89    ns.WaypointID
	wp90    ns.WaypointID
	wp91    ns.WaypointID
	wp92    ns.WaypointID
	wp93    ns.WaypointID
	wp94    ns.WaypointID
	wp95    ns.WaypointID
	wp96    ns.WaypointID
	wp97    ns.WaypointID
	flag98  bool
	flag99  bool
	ivar100 int
	ivar101 int
	gvar102 int
	gvar103 int
	ivar104 int
	gvar105 int
	obj106  ns.ObjectID
)

func init() {
	flag99 = true
	ivar100 = 22
	ivar101 = ivar100
	ivar104 = 0
	gvar105 = 0
	flag98 = true
}
func ToggleSpikes1() {
	ns.ObjectGroupToggle(gvar61)
	ns.AudioEvent(ns.FloorSpikesUp, ns.Waypoint("SpikeWP01"))
	ns.AudioEvent(ns.FloorSpikesDown, ns.Waypoint("SpikeWP06"))
	ns.AudioEvent(ns.FloorSpikesUp, ns.Waypoint("SpikeWP09"))
	ns.FrameTimer(60, ToggleSpikes2)
}
func ToggleSpikes2() {
	ns.ObjectGroupToggle(gvar62)
	ns.AudioEvent(ns.FloorSpikesDown, ns.Waypoint("SpikeWP04"))
	ns.AudioEvent(ns.FloorSpikesUp, ns.Waypoint("SpikeWP07"))
	ns.AudioEvent(ns.FloorSpikesDown, ns.Waypoint("SpikeWP12"))
	ns.FrameTimer(60, ToggleSpikes1)
}
func GoToPost() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp94)
	v1 = ns.GetWaypointY(wp94)
	v2 = ns.GetWaypointX(wp95)
	v3 = ns.GetWaypointY(wp95)
	ns.CreatureGuard(obj28, v0, v1, v2, v3, 100)
}
func SwitchSpikeState() {
	ns.ObjectGroupOff(gvar66[ivar104])
	if !flag99 {
		goto LABEL1
	}
	ivar104 += 1
	if !(ivar104 > 5) {
		goto LABEL2
	}
	ivar104 = 0
LABEL2:
	ns.ObjectGroupOn(gvar66[ivar104])
	ns.AudioEvent(ns.FloorSpikesUp, wp97)
	ns.FrameTimer(30, SwitchSpikeState)
	goto LABEL3
LABEL1:
	return
LABEL3:
	return
}
func CloseWalls() {
	ns.MusicPopEvent()
	ns.WallGroupClose(gvar69)
	ns.WallGroupClose(gvar70)
}
func PlayAction() {
	ns.MusicPushEvent()
	ns.Music(26, 100)
}
func PlayCaves() {
	ns.Music(18, 100)
}
func MapEntry() {
	PlayCaves()
	ns.NoWallSound(false)
}
func Secret01Declare() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp96, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp96)
	ns.GiveXp(ns.GetHost(), 500)
}
func Secret01Exit() {
	ns.WallGroupOpen(gvar71)
}
func Secret02Declare() {
	ns.ObjectGroupOff(gvar68)
	ns.MoveWaypoint(wp96, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp96)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 250)
}
func Secret03Declare() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp96, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp96)
	ns.GiveXp(ns.GetHost(), 50)
}
func Secret04Declare() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp96, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp96)
	ns.GiveXp(ns.GetHost(), 100)
}
func SpikeTimerAudioLoop() {
	if flag98 {
		goto LABEL1
	}
	ns.AudioEvent(ns.Gear1, wp75)
	ns.FrameTimer(ivar101, SpikeTimerAudioLoop)
	ivar101 -= 1
	goto LABEL2
LABEL1:
	ns.AudioEvent(ns.TriggerPressed, wp75)
LABEL2:
	return
}
func ResetOgreSpikeDoor01() {
	ns.ObjectOn(obj4)
	ns.ObjectOn(obj5)
	ns.ObjectGroupOn(gvar55)
	ns.LockDoor(obj6)
	flag98 = true
	ivar101 = ivar100
}
func ToggleOgreTrapSpikes01() {
	ns.ObjectOff(obj4)
	ns.ObjectOff(obj5)
	ns.ObjectGroupOff(gvar55)
	if !flag98 {
		goto LABEL1
	}
	ns.UnlockDoor(obj6)
	flag98 = false
	SpikeTimerAudioLoop()
	ns.FrameTimer(150, ResetOgreSpikeDoor01)
	goto LABEL2
LABEL1:
	ns.LockDoor(obj6)
	flag98 = true
LABEL2:
	return
}
func MoveBlocks01() {
	ns.ObjectGroupOff(gvar57)
	ns.ObjectGroupOff(gvar58)
	PlayAction()
	ns.Move(obj7, wp76)
	ns.Move(obj8, wp77)
	ns.Move(obj9, wp78)
	ns.Move(obj10, wp79)
	ns.Move(obj11, wp80)
	ns.Move(obj12, wp81)
	ns.Move(obj13, wp82)
	ns.Move(obj14, wp83)
	ns.Move(obj15, wp84)
	ns.Move(obj16, wp85)
	ns.Move(obj17, wp86)
	ns.Move(obj18, wp87)
	ns.WallGroupOpen(gvar69)
	ns.ObjectOn(obj21)
	ns.Wander(obj21)
}
func CallNextOgre() {
	ns.WallGroupOpen(gvar70)
	ns.ObjectOn(obj22)
	ns.Wander(obj22)
}
func MoveBlocks02() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp90)
	v1 = ns.GetWaypointY(wp90)
	ns.MoveObject(obj19, v0, v1)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	ns.AudioEvent(ns.KeyDrop, wp90)
	ns.ObjectOn(obj20)
	ns.Move(obj9, wp88)
	ns.Move(obj10, wp89)
	ns.Move(obj7, wp78)
	ns.Move(obj8, wp79)
	ns.FrameTimer(45, CloseWalls)
}
func OpenWayOut() {
	ns.WallGroupOpen(gvar74)
}
func GetBuddy() {
	ns.RetreatLevel(obj23, 0)
	ns.Move(obj23, wp91)
}
func FollowOgrePatrol01() {
	if !ns.IsCaller(obj23) {
		goto LABEL1
	}
	ns.CreatureFollow(ns.GetTrigger(), obj23)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.Walk(obj23, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL1:
	return
}
func WaitASec() {
	ns.PauseObject(ns.GetTrigger(), 30)
}
func Patrol() {
	ns.Wander(ns.GetTrigger())
}
func GoMedieval() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
LABEL1:
	return
}
func GoNormal() {
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
}
func StopAndListen() {
	var v0 int
	v0 = ns.Random(1, 3)
	if !(ns.IsCaller(ns.GetHost()) && v0 == 1) {
		goto LABEL1
	}
	ns.PauseObject(ns.GetTrigger(), 45)
	ns.LookAtObject(ns.GetTrigger(), ns.GetCaller())
	if !ns.IsVisibleTo(ns.GetTrigger(), ns.GetCaller()) {
		goto LABEL2
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
	goto LABEL1
LABEL2:
	ns.CreatureIdle(ns.GetTrigger())
LABEL1:
	return
}
func ReturnHome() {
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
	ns.GoBackHome(ns.GetTrigger())
}
func SpikeHall() {
	ns.ObjectGroupOff(gvar59)
	ns.ObjectGroupOn(gvar60)
	ToggleSpikes1()
}
func PressButton() {
	ns.AudioEvent(ns.GruntRecognize, wp93)
	ns.Move(obj28, wp92)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.AggressionLevel(obj31, 0.83)
	ns.FrameTimer(60, GoToPost)
}
func OgrePressButton() {
	ns.FrameTimer(10, PressButton)
}
func LockCageDoor() {
	if ns.IsLocked(obj32) {
		goto LABEL1
	}
	ns.LockDoor(obj32)
	goto LABEL2
LABEL1:
	ns.UnlockDoor(obj32)
LABEL2:
	return
}
func FoundTreasureTrove() {
	ns.PrintToAll("War09c:FoundTrove")
	ns.GiveXp(ns.GetHost(), 1000)
}
func ToggleElevatorGroup3() {
	ns.ObjectGroupToggle(gvar65)
}
func ToggleElevatorGroup1() {
	ns.ObjectGroupToggle(gvar63)
}
func StartSpikes() {
	ns.ObjectGroupOff(gvar67)
	ns.LockDoor(obj33)
	ns.FrameTimer(60, SwitchSpikeState)
}
func UnlockSpikeSegRoomDoor() {
	ns.UnlockDoor(obj33)
	flag99 = false
}
func CallSpitter1() {
	ns.Wander(obj35)
}
func SetSpitter1PathColor() {
	ns.SetRoamFlag(obj35, 1|128)
}
func LockPitDoor() {
	ns.LockDoor(obj34)
}
func UnlockPitDoor() {
	ns.UnlockDoor(obj34)
}
func ReleaseA() {
	ns.UnlockDoor(obj36)
	ns.UnlockDoor(obj37)
	ns.CreatureHunt(obj48)
}
func ReleaseB() {
	ns.UnlockDoor(obj38)
	ns.UnlockDoor(obj39)
	ns.CreatureHunt(obj49)
}
func ReleaseC() {
	ns.UnlockDoor(obj40)
	ns.UnlockDoor(obj41)
	ns.CreatureHunt(obj50)
}
func ReleaseD() {
	ns.UnlockDoor(obj42)
	ns.UnlockDoor(obj43)
	ns.CreatureHunt(obj51)
}
func ReleaseE() {
	ns.UnlockDoor(obj44)
	ns.UnlockDoor(obj45)
	ns.CreatureHunt(obj52)
}
func ReleaseF() {
	ns.UnlockDoor(obj46)
	ns.UnlockDoor(obj47)
	ns.CreatureHunt(obj53)
}
func ToggleElevatorGroup2() {
	ns.ObjectOff(obj54)
	ns.ObjectGroupToggle(gvar64)
	ns.GiveXp(ns.GetHost(), 2500)
}
func PlayerDeath() {
	ns.DeathScreen(9)
}
func MapInitialize() {
	obj6 = ns.Object("OgreSpikeDoor01")
	obj4 = ns.Object("OgreTrapSpikePlate01")
	obj5 = ns.Object("OgreTrapSpikeSwitch01")
	obj7 = ns.Object("Block01")
	obj8 = ns.Object("Block02")
	obj9 = ns.Object("Block03")
	obj10 = ns.Object("Block04")
	obj11 = ns.Object("Block05")
	obj12 = ns.Object("Block06")
	obj13 = ns.Object("Block07")
	obj14 = ns.Object("Block08")
	obj15 = ns.Object("Block09")
	obj16 = ns.Object("Block10")
	obj17 = ns.Object("Block11")
	obj18 = ns.Object("Block12")
	obj19 = ns.Object("SilverKey")
	obj20 = ns.Object("SilverKeyLight")
	obj21 = ns.Object("OgreFighter01")
	obj22 = ns.Object("OgreFighter02")
	obj23 = ns.Object("OgrePatrol01")
	obj24 = ns.Object("OgrePatrol02")
	obj25 = ns.Object("OgrePatrol03")
	obj26 = ns.Object("OgrePatrolBuddy01")
	obj27 = ns.Object("OgrePatrolBuddy02")
	obj28 = ns.Object("ButtonPresser")
	obj29 = ns.Object("OgrePsycho01")
	obj30 = ns.Object("OgrePsycho02")
	obj31 = ns.Object("OgrePsycho03")
	obj32 = ns.Object("CageDoor")
	obj33 = ns.Object("SpikeSegRoomDoor")
	obj34 = ns.Object("CrumblingPitDoor")
	obj36 = ns.Object("CreaturePitDoorA1")
	obj37 = ns.Object("CreaturePitDoorA2")
	obj38 = ns.Object("CreaturePitDoorB1")
	obj39 = ns.Object("CreaturePitDoorB2")
	obj40 = ns.Object("CreaturePitDoorC1")
	obj41 = ns.Object("CreaturePitDoorC2")
	obj42 = ns.Object("CreaturePitDoorD1")
	obj43 = ns.Object("CreaturePitDoorD2")
	obj44 = ns.Object("CreaturePitDoorE1")
	obj45 = ns.Object("CreaturePitDoorE2")
	obj46 = ns.Object("CreaturePitDoorF1")
	obj47 = ns.Object("CreaturePitDoorF2")
	obj48 = ns.Object("CreatureA")
	obj49 = ns.Object("CreatureB")
	obj50 = ns.Object("CreatureC")
	obj51 = ns.Object("CreatureD")
	obj52 = ns.Object("CreatureE")
	obj53 = ns.Object("CreatureF")
	obj106 = ns.Object("CreaturePitElev")
	obj54 = ns.Object("CreaturePitSwitch")
	gvar55 = ns.ObjectGroup("OgreTrapSpikes01")
	gvar57 = ns.ObjectGroup("BlockTrigger01")
	gvar56 = ns.ObjectGroup("BlockMoverGroup01")
	gvar58 = ns.ObjectGroup("SecretLight01")
	gvar59 = ns.ObjectGroup("SpikeHallTriggers")
	gvar60 = ns.ObjectGroup("SpikeMoverGroup")
	gvar61 = ns.ObjectGroup("HallSpikes1")
	gvar62 = ns.ObjectGroup("HallSpikes2")
	gvar66[0] = ns.ObjectGroup("SpikeSeg01")
	gvar66[1] = ns.ObjectGroup("SpikeSeg02")
	gvar66[2] = ns.ObjectGroup("SpikeSeg03")
	gvar66[3] = ns.ObjectGroup("SpikeSeg04")
	gvar66[4] = ns.ObjectGroup("SpikeSeg05")
	gvar66[5] = ns.ObjectGroup("SpikeSeg06")
	gvar67 = ns.ObjectGroup("SpikeSegTriggers")
	gvar68 = ns.ObjectGroup("Secret02Triggers")
	gvar63 = ns.ObjectGroup("ElevatorGroup1")
	gvar64 = ns.ObjectGroup("ElevatorGroup2")
	gvar65 = ns.ObjectGroup("ElevatorGroup3")
	gvar69 = ns.WallGroup("SecretWallGroup01")
	gvar70 = ns.WallGroup("SecretWallGroup02")
	gvar71 = ns.WallGroup("Secret01Wall")
	gvar72 = ns.WallGroup("SecretHallWalls01")
	gvar73 = ns.WallGroup("CreaturePitWalls")
	gvar74 = ns.WallGroup("WayOut")
	wp75 = ns.Waypoint("SpikeTimerAudioOrigin")
	wp76 = ns.Waypoint("BlockWP01")
	wp77 = ns.Waypoint("BlockWP02")
	wp78 = ns.Waypoint("BlockWP03")
	wp79 = ns.Waypoint("BlockWP04")
	wp80 = ns.Waypoint("BlockWP05")
	wp81 = ns.Waypoint("BlockWP06")
	wp82 = ns.Waypoint("BlockWP07")
	wp83 = ns.Waypoint("BlockWP08")
	wp84 = ns.Waypoint("BlockWP09")
	wp85 = ns.Waypoint("BlockWP10")
	wp86 = ns.Waypoint("BlockWP11")
	wp87 = ns.Waypoint("BlockWP12")
	wp88 = ns.Waypoint("BlockWP13")
	wp89 = ns.Waypoint("BlockWP14")
	wp90 = ns.Waypoint("SilverKeySpot")
	wp91 = ns.Waypoint("BuddyWP")
	wp92 = ns.Waypoint("ButtonLocationWP")
	wp93 = ns.Waypoint("OgreAudioOrigin")
	wp94 = ns.Waypoint("PostWP")
	wp95 = ns.Waypoint("FacingWP")
	wp96 = ns.Waypoint("PlayerSounds")
	wp97 = ns.Waypoint("RoomSounds")
	ns.LockDoor(obj36)
	ns.LockDoor(obj37)
	ns.LockDoor(obj38)
	ns.LockDoor(obj39)
	ns.LockDoor(obj40)
	ns.LockDoor(obj41)
	ns.LockDoor(obj42)
	ns.LockDoor(obj43)
	ns.LockDoor(obj44)
	ns.LockDoor(obj45)
	ns.LockDoor(obj46)
	ns.LockDoor(obj47)
	ns.UnBlind()
}
func OnEvent(typ string) {
	switch typ {
	case "MapEntry":
		MapEntry()
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
