package con03b

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
	ivar15  int
	gvar16  int
	gvar17  int
	gvar18  int
	gvar19  int
	gvar20  int
	ivar21  int
	ivar22  int
	wp23    ns.WaypointID
	wp24    ns.WaypointID
	wp25    ns.WaypointID
	wp26    ns.WaypointID
	wp27    ns.WaypointID
	wp28    ns.WaypointID
	obj29   ns.ObjectID
	gvar30  ns.WallGroupID
	gvar31  ns.WallGroupID
	gvar32  ns.WallGroupID
	gvar33  ns.WallID
	gvar34  ns.WallID
	gvar35  ns.WallID
	gvar36  ns.WallID
	gvar37  ns.WallID
	gvar38  ns.WallID
	gvar39  ns.WallID
	gvar40  ns.WallID
	gvar41  ns.WallID
	gvar42  ns.WallID
	gvar43  ns.WallID
	gvar44  ns.WallID
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
	obj55   ns.ObjectID
	obj56   ns.ObjectID
	obj57   ns.ObjectID
	obj58   ns.ObjectID
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
	gvar74  int
	obj75   ns.ObjectID
	gvar76  int
	obj77   ns.ObjectID
	obj78   ns.ObjectID
	obj79   ns.ObjectID
	obj80   ns.ObjectID
	gvar81  int
	gvar82  int
	gvar83  int
	wp84    ns.WaypointID
	wp85    [24]ns.WaypointID
	wp86    [7]ns.WaypointID
	wp87    [4]ns.WaypointID
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
	wp98    ns.WaypointID
	obj99   ns.ObjectID
	obj100  ns.ObjectID
	obj101  ns.ObjectID
	obj102  ns.ObjectID
	obj103  ns.ObjectID
	obj104  ns.ObjectID
	obj105  ns.ObjectID
	obj106  ns.ObjectID
	obj107  ns.ObjectID
	obj108  ns.ObjectID
	obj109  ns.ObjectID
	obj110  ns.ObjectID
	obj111  ns.ObjectID
	gvar112 ns.ObjectGroupID
	gvar113 ns.ObjectGroupID
	gvar114 ns.ObjectGroupID
	gvar115 ns.ObjectGroupID
	gvar116 ns.ObjectGroupID
	obj117  ns.ObjectID
	obj118  ns.ObjectID
	obj119  ns.ObjectID
	obj120  ns.ObjectID
	wp121   ns.WaypointID
	obj122  ns.ObjectID
	obj123  ns.ObjectID
	obj124  ns.ObjectID
	obj125  ns.ObjectID
	gvar126 int
	gvar127 int
	gvar128 int
	gvar129 int
	gvar130 int
)

func init() {
	ivar22 = 0
	gvar81 = 1
	ivar21 = 0
	ivar15 = 1
	gvar16 = 0
	gvar17 = 0
	gvar18 = 0
	gvar19 = 0
	gvar20 = 0
	gvar83 = 0
	gvar127 = 0
	gvar128 = 0
	gvar129 = 0
	gvar130 = 0
}
func Worker1Start() {
	ns.ObjectOn(obj99)
	ns.Move(obj4, wp94)
	ns.Chat(obj4, "Con03B.scr:Worker1ChatA")
}
func TopFloorCleared() {
	ns.ObjectOn(ns.Object("Worker1Unlock"))
	ns.Chat(obj4, "Con03B.scr:Worker1ChatE")
	ns.Move(obj4, wp97)
	ivar15 = 6
}
func GiveHint1() {
	ns.PrintToAll("Con03.scr:Hint1")
	ns.JournalEntry(ns.GetHost(), "Con03Hint1", 8)
}
func Worker1DialogStart() {
	var v0 int
	ns.LookAtObject(obj4, ns.GetHost())
	ns.DestroyEveryChat()
	v0 = ivar15
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	if v0 == 4 {
		goto LABEL4
	}
	if v0 == 5 {
		goto LABEL5
	}
	if v0 == 6 {
		goto LABEL6
	}
	if v0 == 7 {
		goto LABEL7
	}
	goto LABEL8
LABEL1:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1B")
	goto LABEL9
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1C")
	goto LABEL9
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1D")
	goto LABEL9
LABEL4:
	if ivar21 != 0 {
		goto LABEL10
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1Gi")
LABEL10:
	if ivar21 != 1 {
		goto LABEL11
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1Gii")
LABEL11:
	if ivar21 != 2 {
		goto LABEL12
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1Giii")
LABEL12:
	goto LABEL9
LABEL5:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1Hi")
	goto LABEL9
LABEL6:
	if ivar21 != 3 {
		goto LABEL13
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1Ii")
LABEL13:
	if ivar21 != 4 {
		goto LABEL14
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1Iii")
LABEL14:
	goto LABEL9
LABEL7:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:Worker1Ji")
	goto LABEL9
LABEL8:
	goto LABEL9
LABEL9:
	return
}
func Worker1DialogEnd() {
	var v0 int
	v0 = ivar15
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 5 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(2, Worker1Start)
	ns.Pickup(ns.GetHost(), obj10)
	ns.PrintToAll("Con03B.scr:GotKey")
	ivar22 += 1
	ns.CancelDialog(obj4)
	goto LABEL4
LABEL2:
	ns.FrameTimer(2, TopFloorCleared)
	ns.Pickup(ns.GetHost(), obj14)
	ns.PrintToAll("Con03B.scr:GotKey")
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func Worker2DialogStart() {
	var v0 int
	ns.LookAtObject(obj5, ns.GetHost())
	ns.DestroyEveryChat()
	v0 = gvar16
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	if !(ivar22 <= 2) {
		goto LABEL6
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkAKey")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
LABEL7:
	goto LABEL8
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
	goto LABEL8
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkB")
	goto LABEL8
LABEL4:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkC")
	goto LABEL8
LABEL5:
	goto LABEL8
LABEL8:
	return
}
func Worker2DialogEnd() {
	var v0 int
	v0 = gvar16
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	ns.CreatureFollow(obj5, ns.GetHost())
	if !(ivar22 <= 2) {
		goto LABEL4
	}
	ns.Pickup(ns.GetHost(), obj11)
	ns.PrintToAll("Con03B.scr:GotKey")
	ivar22 += 1
LABEL4:
	gvar16 = 1
	if ivar21 != 0 {
		goto LABEL5
	}
	ns.FrameTimer(300, GiveHint1)
LABEL5:
	ns.CancelDialog(obj5)
	goto LABEL6
LABEL2:
	ns.CreatureFollow(obj5, ns.GetHost())
	goto LABEL6
LABEL3:
	goto LABEL6
LABEL6:
	return
}
func Worker3DialogStart() {
	var v0 int
	ns.LookAtObject(obj6, ns.GetHost())
	ns.DestroyEveryChat()
	v0 = gvar17
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	if !(ivar22 <= 2) {
		goto LABEL6
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkAKey")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
LABEL7:
	goto LABEL8
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
	goto LABEL8
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkB")
	goto LABEL8
LABEL4:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkC")
	goto LABEL8
LABEL5:
	goto LABEL8
LABEL8:
	return
}
func Worker3DialogEnd() {
	var v0 int
	v0 = gvar17
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	ns.CreatureFollow(obj6, ns.GetHost())
	if !(ivar22 <= 2) {
		goto LABEL4
	}
	ns.Pickup(ns.GetHost(), obj12)
	ns.PrintToAll("Con03B.scr:GotKey")
	ivar22 += 1
LABEL4:
	gvar17 = 1
	if ivar21 != 0 {
		goto LABEL5
	}
	ns.FrameTimer(300, GiveHint1)
LABEL5:
	ns.CancelDialog(obj6)
	goto LABEL6
LABEL2:
	ns.CreatureFollow(obj6, ns.GetHost())
	goto LABEL6
LABEL3:
	goto LABEL6
LABEL6:
	return
}
func Worker4DialogStart() {
	var v0 int
	ns.LookAtObject(obj7, ns.GetHost())
	ns.DestroyEveryChat()
	v0 = gvar18
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
	goto LABEL6
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
	goto LABEL6
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkB")
	goto LABEL6
LABEL4:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkC")
	goto LABEL6
LABEL5:
	goto LABEL6
LABEL6:
	return
}
func Worker4DialogEnd() {
	var v0 int
	v0 = gvar18
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	ns.CreatureFollow(obj7, ns.GetHost())
	gvar18 = 1
	ns.CancelDialog(obj7)
	goto LABEL4
LABEL2:
	ns.CreatureFollow(obj7, ns.GetHost())
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func Worker5DialogStart() {
	var v0 int
	ns.LookAtObject(obj8, ns.GetHost())
	ns.DestroyEveryChat()
	v0 = gvar19
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
	goto LABEL6
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
	goto LABEL6
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkB")
	goto LABEL6
LABEL4:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkC")
	goto LABEL6
LABEL5:
	goto LABEL6
LABEL6:
	return
}
func Worker5DialogEnd() {
	var v0 int
	v0 = gvar19
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	ns.CreatureFollow(obj8, ns.GetHost())
	gvar19 = 1
	ns.CancelDialog(obj8)
	goto LABEL4
LABEL2:
	ns.CreatureFollow(obj8, ns.GetHost())
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func Worker6DialogStart() {
	var v0 int
	ns.LookAtObject(obj9, ns.GetHost())
	ns.DestroyEveryChat()
	v0 = gvar20
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	if !(ivar22 <= 2) {
		goto LABEL6
	}
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkAKey")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
LABEL7:
	goto LABEL8
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkA")
	goto LABEL8
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkB")
	goto LABEL8
LABEL4:
	ns.TellStory(ns.DemonRecognize, "Con03B.scr:WorkersTalkC")
	goto LABEL8
LABEL5:
	goto LABEL8
LABEL8:
	return
}
func Worker6DialogEnd() {
	var v0 int
	v0 = gvar20
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	ns.CreatureFollow(obj9, ns.GetHost())
	if !(ivar22 <= 2) {
		goto LABEL4
	}
	ns.Pickup(ns.GetHost(), obj13)
	ns.PrintToAll("Con03B.scr:GotKey")
	ivar22 += 1
LABEL4:
	gvar20 = 1
	if ivar21 != 0 {
		goto LABEL5
	}
	ns.FrameTimer(300, GiveHint1)
LABEL5:
	ns.CancelDialog(obj9)
	goto LABEL6
LABEL2:
	ns.CreatureFollow(obj9, ns.GetHost())
	goto LABEL6
LABEL3:
	goto LABEL6
LABEL6:
	return
}
func WallOpen2() {
	ns.ObjectGroupOn(ns.ObjectGroup("Walls2Spiders"))
	ns.ObjectGroupOff(ns.ObjectGroup("Walls2Triggers"))
	ns.WallGroupOpen(gvar31)
	ns.WallBreak(gvar35)
	ns.WallBreak(gvar36)
	ns.WallBreak(gvar37)
	ns.WallBreak(gvar38)
	ns.WallBreak(gvar39)
	ns.WallBreak(gvar40)
	ns.WallBreak(gvar41)
	ns.WallBreak(gvar42)
	ns.WallBreak(gvar43)
	ns.WallBreak(gvar44)
}
func WallOpen3() {
	ns.WallGroupOpen(gvar32)
}
func WallClose3() {
	ns.WallGroupClose(gvar32)
}
func ToggleElevator1() {
	ns.ObjectToggle(obj46)
	ns.ObjectToggle(obj47)
	ns.ObjectToggle(obj48)
	ns.ObjectToggle(obj49)
	ns.ObjectToggle(obj50)
	ns.ObjectToggle(obj51)
	ns.ObjectToggle(obj52)
	ns.ObjectToggle(obj53)
	ns.ObjectToggle(obj54)
	if !ns.IsObjectOn(obj46) {
		goto LABEL1
	}
	ns.AudioEvent(ns.Gear3, ns.Waypoint("ElevatorSFX"))
LABEL1:
	return
}
func ToggleElevator2() {
	ns.ObjectToggle(obj55)
	ns.ObjectToggle(obj56)
	ns.ObjectToggle(obj57)
	ns.ObjectToggle(obj58)
	ns.ObjectToggle(obj59)
	ns.ObjectToggle(obj60)
	ns.ObjectToggle(obj61)
}
func ToggleElevator3() {
	ns.ObjectToggle(obj62)
	ns.ObjectToggle(obj63)
	ns.ObjectToggle(obj64)
	ns.ObjectToggle(obj65)
	ns.ObjectToggle(obj66)
	ns.ObjectToggle(obj67)
	ns.ObjectToggle(obj68)
}
func SetupElevators() {
	if ns.GetElevatorStatus(obj47) != 2 {
		goto LABEL1
	}
	ns.ObjectOff(obj46)
	ns.ObjectOff(obj47)
	ns.ObjectOff(obj55)
	ns.ObjectOff(obj56)
	ns.ObjectOff(obj62)
	ns.ObjectOff(obj63)
	goto LABEL2
LABEL1:
	ns.FrameTimer(5, SetupElevators)
LABEL2:
	return
}
func RunAway1() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	if !ns.IsAttackedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.Move(ns.GetCaller(), wp23)
LABEL1:
	return
}
func RunAway2() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	if ns.IsOwnedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.Move(ns.GetCaller(), wp24)
LABEL1:
	return
}
func RunAway3() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	if ns.IsOwnedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.Move(ns.GetCaller(), wp25)
LABEL1:
	return
}
func RunAway4() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	if ns.IsOwnedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.Move(ns.GetCaller(), wp26)
LABEL1:
	return
}
func RunAway5() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	if ns.IsOwnedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.Move(ns.GetCaller(), wp27)
LABEL1:
	return
}
func RunAway6() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	if ns.IsOwnedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.Move(ns.GetCaller(), wp28)
LABEL1:
	return
}
func ForemanDialogStart() {
	var v0 int
	v0 = gvar81
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	if v0 == 4 {
		goto LABEL4
	}
	if v0 == 5 {
		goto LABEL5
	}
	if v0 == 6 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	ns.TellStory(ns.DemonRecognize, "Con03C.scr:ForemanA")
	goto LABEL8
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03C.scr:ForemanB")
	goto LABEL8
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03C.scr:ForemanC")
	goto LABEL8
LABEL4:
	ns.TellStory(ns.DemonRecognize, "Con03C.scr:ForemanD")
	goto LABEL8
LABEL5:
	ns.TellStory(ns.DemonRecognize, "Con03C.scr:ForemanE")
	goto LABEL8
LABEL6:
	ns.TellStory(ns.DemonRecognize, "Con03C.scr:ForemanF")
	goto LABEL8
LABEL7:
	goto LABEL8
LABEL8:
	return
}
func ForemanDialogEnd() {
	var v0 int
	v0 = gvar81
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	if v0 == 4 {
		goto LABEL4
	}
	if v0 == 5 {
		goto LABEL5
	}
	if v0 == 6 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	ns.JournalEdit(ns.GetHost(), "LocateMineForeman", 4)
	ns.PrintToAll("Con03C.scr:TaskComplete")
	ns.JournalEntry(ns.GetHost(), "Con03RescueMineWorkers", 2)
	ns.PrintToAll("Con03C.scr:NewTask")
	ns.Pickup(ns.GetHost(), obj78)
	ns.Pickup(ns.GetHost(), obj79)
	ns.Pickup(ns.GetHost(), obj80)
	ns.PrintToAll("GeneralPrint:GainedKey")
	gvar81 = 2
	goto LABEL8
LABEL2:
	gvar81 = 3
	goto LABEL8
LABEL3:
	gvar81 = 2
	goto LABEL8
LABEL4:
	ns.JournalEntry(ns.GetHost(), "Con03FindAirshipCap", 2)
	ns.PrintToAll("Con03C.scr:NewTask")
	ns.Pickup(ns.GetHost(), obj77)
	gvar81 = 5
	goto LABEL8
LABEL5:
	gvar81 = 6
	goto LABEL8
LABEL6:
	gvar81 = 5
	goto LABEL8
LABEL7:
	goto LABEL8
LABEL8:
	return
}
func InitializeForeman() {
	obj75 = ns.Object("Foreman")
	obj77 = ns.Object("Gold500")
	obj78 = ns.Object("BlueKey")
	obj79 = ns.Object("EntranceKey")
	obj80 = ns.Object("ImpBeast")
	ns.SetOwner(ns.GetHost(), obj75)
	ns.StoryPic(obj75, "Townsman4Pic")
	ns.SetDialog(obj75, ns.NORMAL, ForemanDialogStart, ForemanDialogEnd)
}
func PlayOutdoorMusic() {
	ns.Music(21, 100)
}
func SetupWorkers() {
	ns.SetOwner(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.SetOwner(ns.GetHost(), obj6)
	ns.SetOwner(ns.GetHost(), obj7)
	ns.SetOwner(ns.GetHost(), obj8)
	ns.SetOwner(ns.GetHost(), obj9)
	ns.LockDoor(obj100)
	ns.LockDoor(obj101)
	ns.Damage(obj102, 0, 1000, 0)
	ns.Damage(obj103, 0, 1000, 0)
	ns.Damage(obj104, 0, 1000, 0)
}
func MapInitialize() {
	wp84 = ns.Waypoint("SecretAudioOrigin")
	wp23 = ns.Waypoint("RunAwayTo1")
	wp24 = ns.Waypoint("RunAwayTo2")
	wp25 = ns.Waypoint("RunAwayTo3")
	wp26 = ns.Waypoint("RunAwayTo4")
	wp27 = ns.Waypoint("RunAwayTo5")
	wp28 = ns.Waypoint("RunAwayTo6")
	wp86[0] = ns.Waypoint("CreateBat1")
	wp86[1] = ns.Waypoint("CreateBat2")
	wp86[2] = ns.Waypoint("CreateBat3")
	wp86[3] = ns.Waypoint("CreateBat4")
	wp86[4] = ns.Waypoint("CreateBat5")
	wp86[5] = ns.Waypoint("CreateBat6")
	wp86[6] = ns.Waypoint("CreateBat7")
	wp87[0] = ns.Waypoint("CreateImp1")
	wp87[1] = ns.Waypoint("CreateImp2")
	wp87[2] = ns.Waypoint("CreateImp3")
	wp87[3] = ns.Waypoint("CreateImp4")
	wp85[0] = ns.Waypoint("CreateSmSpider1")
	wp85[1] = ns.Waypoint("CreateSmSpider2")
	wp85[2] = ns.Waypoint("CreateSmSpider3")
	wp85[3] = ns.Waypoint("CreateSmSpider4")
	wp85[4] = ns.Waypoint("CreateSmSpider5")
	wp85[5] = ns.Waypoint("CreateSmSpider6")
	wp85[6] = ns.Waypoint("CreateSmSpider7")
	wp85[7] = ns.Waypoint("CreateSmSpider8")
	wp85[8] = ns.Waypoint("CreateSmSpider9")
	wp85[9] = ns.Waypoint("CreateSmSpider10")
	wp85[10] = ns.Waypoint("CreateSmSpider11")
	wp85[11] = ns.Waypoint("CreateSmSpider12")
	wp85[12] = ns.Waypoint("CreateSmSpider13")
	wp85[13] = ns.Waypoint("CreateSmSpider14")
	wp85[14] = ns.Waypoint("CreateSmSpider15")
	wp85[15] = ns.Waypoint("CreateSmSpider16")
	wp85[16] = ns.Waypoint("CreateSmSpider17")
	wp85[17] = ns.Waypoint("CreateSmSpider18")
	wp85[18] = ns.Waypoint("CreateSmSpider19")
	wp85[19] = ns.Waypoint("CreateSmSpider20")
	wp85[20] = ns.Waypoint("CreateSmSpider21")
	wp85[21] = ns.Waypoint("CreateSmSpider22")
	wp85[22] = ns.Waypoint("CreateSmSpider23")
	wp85[23] = ns.Waypoint("CreateSmSpider24")
	wp88 = ns.Waypoint("WorkerEnd1")
	wp89 = ns.Waypoint("WorkerEnd2")
	wp90 = ns.Waypoint("WorkerEnd3")
	wp91 = ns.Waypoint("WorkerEnd4")
	wp92 = ns.Waypoint("WorkerEnd5")
	wp93 = ns.Waypoint("WorkerEnd6")
	obj102 = ns.Object("DeadGuy1")
	obj103 = ns.Object("DeadGuy2")
	obj104 = ns.Object("DeadGuy3")
	obj105 = ns.Object("StubbornDeadGuy")
	ns.Damage(obj105, 0, 1000, 0)
	wp94 = ns.Waypoint("Button2Push")
	wp95 = ns.Waypoint("Elevator2Down")
	wp96 = ns.Waypoint("Worker1Wait")
	wp97 = ns.Waypoint("LevelGateUnlock")
	wp98 = ns.Waypoint("Worker1Elevator3")
	obj99 = ns.Object("ToggleElevator2Button")
	obj100 = ns.Object("LevelGateA")
	obj101 = ns.Object("LevelGateB")
	obj106 = ns.Object("Worker1Down")
	gvar30 = ns.WallGroup("Walls1")
	gvar31 = ns.WallGroup("Walls2")
	gvar32 = ns.WallGroup("Walls3")
	gvar33 = ns.Wall(164, 92)
	gvar34 = ns.Wall(165, 93)
	gvar35 = ns.Wall(19, 111)
	gvar36 = ns.Wall(20, 112)
	gvar37 = ns.Wall(28, 110)
	gvar38 = ns.Wall(29, 111)
	gvar39 = ns.Wall(31, 111)
	gvar40 = ns.Wall(33, 109)
	gvar41 = ns.Wall(34, 108)
	gvar42 = ns.Wall(36, 106)
	gvar43 = ns.Wall(37, 105)
	gvar44 = ns.Wall(38, 104)
	obj45 = ns.Object("WallOpen1")
	obj110 = ns.Object("WallOpen2A")
	obj111 = ns.Object("WallOpen2B")
	obj4 = ns.Object("Alex")
	obj5 = ns.Object("Logan")
	obj6 = ns.Object("Naldo")
	obj7 = ns.Object("Garrit")
	obj8 = ns.Object("Dudley")
	obj9 = ns.Object("Claude")
	obj10 = ns.Object("RedKey1")
	obj11 = ns.Object("RedKey2")
	obj12 = ns.Object("RedKey3")
	obj13 = ns.Object("RedKey4")
	obj14 = ns.Object("GoldKey")
	gvar112 = ns.ObjectGroup("Worker1Rec")
	gvar113 = ns.ObjectGroup("Worker2Rec")
	gvar114 = ns.ObjectGroup("Worker3Rec")
	gvar115 = ns.ObjectGroup("Worker5Rec")
	gvar116 = ns.ObjectGroup("Worker6Rec")
	obj117 = ns.Object("InitialDoor")
	obj118 = ns.Object("FinalDoor")
	obj119 = ns.Object("FinalDoor2A")
	obj120 = ns.Object("FinalDoor2B")
	wp121 = ns.Waypoint("FinalPosition")
	obj122 = ns.Object("HuntImp1")
	obj123 = ns.Object("HuntImp2")
	obj124 = ns.Object("HuntImp3")
	obj125 = ns.Object("HuntImp4")
	obj29 = ns.Object("Scorpion1")
	obj46 = ns.Object("Elevator1a")
	obj47 = ns.Object("Elevator1b")
	obj48 = ns.Object("Gear1a")
	obj49 = ns.Object("Gear1b")
	obj50 = ns.Object("Gear1c")
	obj51 = ns.Object("Gear1d")
	obj52 = ns.Object("Gear1e")
	obj53 = ns.Object("Gear1f")
	obj54 = ns.Object("Gear1g")
	obj55 = ns.Object("Elevator2a")
	obj56 = ns.Object("Elevator2b")
	obj57 = ns.Object("Gear2a")
	obj58 = ns.Object("Gear2b")
	obj59 = ns.Object("Gear2c")
	obj60 = ns.Object("Gear2d")
	obj61 = ns.Object("Gear2e")
	obj62 = ns.Object("Elevator3a")
	obj63 = ns.Object("Elevator3b")
	obj64 = ns.Object("Gear3a")
	obj65 = ns.Object("Gear3b")
	obj66 = ns.Object("Gear3c")
	obj67 = ns.Object("Gear3d")
	obj68 = ns.Object("Gear3e")
	obj107 = ns.Object("Secret02Spider01")
	obj108 = ns.Object("Secret02Spider02")
	obj109 = ns.Object("Secret02Spider03")
	obj69 = ns.Object("Urch1")
	obj70 = ns.Object("Urch2")
	obj71 = ns.Object("Urch3")
	obj72 = ns.Object("Urch4")
	obj73 = ns.Object("Urch5")
	ns.StoryPic(obj4, "Miner1Pic")
	ns.StoryPic(obj5, "Miner2Pic")
	ns.StoryPic(obj6, "Miner2Pic")
	ns.StoryPic(obj7, "Miner2Pic")
	ns.StoryPic(obj8, "Miner2Pic")
	ns.StoryPic(obj9, "Miner2Pic")
	ns.SetDialog(obj4, ns.NORMAL, Worker1DialogStart, Worker1DialogEnd)
	PlayOutdoorMusic()
	SetupWorkers()
	InitializeForeman()
	ns.FrameTimer(40, SetupElevators)
}
func MonsterGoHome() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func PlayUndergroundMusic() {
	ns.Music(18, 100)
}
func SecretSFX() {
	ns.MoveWaypoint(wp84, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp84)
}
func Secret01Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
	SecretSFX()
}
func Secret02Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 25)
	ns.AggressionLevel(obj107, 0.83)
	ns.AggressionLevel(obj108, 0.83)
	ns.AggressionLevel(obj109, 0.83)
	SecretSFX()
}
func MapShutdown() {
	ns.NoWallSound(false)
}
func MapEntry() {
	ns.NoWallSound(true)
}
func PlayerDeath() {
	ns.DeathScreen(3)
}
func Worker1Rec() {
	ns.ObjectGroupOff(gvar112)
	ns.PrintToAll("Con03B.scr:Rescued1")
	ns.GiveXp(ns.GetHost(), 250)
	ns.ObjectOn(obj4)
	ns.BecomePet(obj4)
	ns.CreatureFollow(obj4, ns.GetHost())
	ns.Frozen(ns.GetHost(), true)
	ns.StartDialog(obj4, ns.GetHost())
}
func WorkerRescued() {
	var v0 int
	v0 = ivar21
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	if v0 == 4 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.PrintToAll("Con03B.scr:Rescued2")
	goto LABEL7
LABEL2:
	ns.PrintToAll("Con03B.scr:Rescued3")
	goto LABEL7
LABEL3:
	ns.PrintToAll("Con03B.scr:Rescued4")
	goto LABEL7
LABEL4:
	ns.PrintToAll("Con03D.scr:Rescued5")
	goto LABEL7
LABEL5:
	ns.PrintToAll("Con03D.scr:Rescued6")
	goto LABEL7
LABEL6:
	goto LABEL7
LABEL7:
	return
}
func Worker2MoveIt() {
	if gvar128 != 0 {
		goto LABEL1
	}
	ns.Move(obj5, wp89)
	ns.FrameTimer(30, Worker2MoveIt)
LABEL1:
	return
}
func Worker3MoveIt() {
	if gvar129 != 0 {
		goto LABEL1
	}
	ns.Move(obj6, wp90)
	ns.FrameTimer(30, Worker3MoveIt)
LABEL1:
	return
}
func Worker6MoveIt() {
	if gvar130 != 0 {
		goto LABEL1
	}
	ns.Move(obj9, wp93)
	ns.FrameTimer(30, Worker6MoveIt)
LABEL1:
	return
}
func AllFound() {
	ns.JournalEdit(ns.GetHost(), "Con03RescueMineWorkers", 4)
	ns.MoveWaypoint(wp84, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagDrop, wp84)
	ns.PrintToAll("Con03C.scr:TaskComplete")
	ns.GiveXp(ns.GetHost(), 500)
	ns.SetQuestStatus(1, "BeatMine")
	gvar81 = 4
	ivar15 = 7
	gvar16 = 3
	gvar17 = 3
	gvar18 = 3
	gvar19 = 3
	gvar20 = 3
}
func Worker1Reminder() {
	if gvar127 != 1 {
		goto LABEL1
	}
	ns.Move(obj4, wp95)
LABEL1:
	return
}
func Worker1Movement() {
	var v0 int
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	v0 = gvar127
	if v0 == 0 {
		goto LABEL2
	}
	if v0 == 1 {
		goto LABEL3
	}
	if v0 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	ns.Move(obj4, wp95)
	ns.Chat(obj4, "Con03B.scr:Worker1ChatB")
	gvar127 = 1
	goto LABEL1
LABEL3:
	if gvar83 != 0 {
		goto LABEL6
	}
	ns.Move(obj4, wp96)
	ns.ObjectOff(obj106)
	gvar127 = 2
	goto LABEL7
LABEL6:
	ns.Chat(obj4, "Con03B.scr:Worker1ChatC")
	ivar15 = 2
	ns.CreatureGuard(obj4, ns.GetWaypointX(wp88), ns.GetWaypointY(wp88), ns.GetWaypointX(wp92), ns.GetWaypointY(wp92), 0)
LABEL7:
	goto LABEL1
LABEL4:
	goto LABEL1
LABEL5:
	goto LABEL1
LABEL1:
	return
}
func BobDown() {
	gvar83 = 1
	if gvar127 != 2 {
		goto LABEL1
	}
	ns.Chat(obj4, "Con03B.scr:Worker1ChatC")
	ivar15 = 2
	ns.CreatureGuard(obj4, ns.GetWaypointX(wp88), ns.GetWaypointY(wp88), ns.GetWaypointX(wp92), ns.GetWaypointY(wp92), 0)
LABEL1:
	return
}
func BobUp() {
	gvar83 = 0
}
func Worker1Stop() {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Chat(obj4, "Con03B.scr:Worker1ChatD")
	if !(ivar15 < 5) {
		goto LABEL2
	}
	ivar15 = 3
LABEL2:
	ns.SetDialog(obj4, ns.NORMAL, Worker1DialogStart, Worker1DialogEnd)
LABEL1:
	return
}
func Worker1Unlock() {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	ns.UnlockDoor(obj100)
	ns.UnlockDoor(obj101)
	ns.Move(obj4, wp98)
LABEL1:
	return
}
func Worker1Elevator3() {
	ns.CreatureGuard(obj4, ns.GetWaypointX(wp88), ns.GetWaypointY(wp88), ns.GetWaypointX(wp92), ns.GetWaypointY(wp92), 0)
}
func Worker1Filter() int {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	return 1
	goto LABEL2
LABEL1:
	return 0
LABEL2:
	return 0
}
func Worker2Rec() {
	ns.ObjectGroupOff(gvar113)
	ns.GiveXp(ns.GetHost(), 100)
	ns.ObjectOn(obj5)
	ns.BecomePet(obj5)
	ns.CreatureFollow(obj5, ns.GetHost())
	ns.CreateObject("SmallAlbinoSpider", wp85[0])
	ns.CreateObject("SmallAlbinoSpider", wp85[1])
	ns.CreateObject("SmallAlbinoSpider", wp85[2])
	ns.CreateObject("SmallAlbinoSpider", wp85[3])
	ns.CreateObject("SmallAlbinoSpider", wp85[4])
	ns.CreateObject("SmallAlbinoSpider", wp85[5])
	ns.CreateObject("SmallAlbinoSpider", wp85[6])
	ns.CreateObject("SmallAlbinoSpider", wp85[7])
	ns.CreateObject("SmallAlbinoSpider", wp85[8])
	ns.CreateObject("SmallAlbinoSpider", wp85[9])
	ns.CreateObject("SmallAlbinoSpider", wp85[10])
	ns.CreateObject("SmallAlbinoSpider", wp85[11])
	ns.CreateObject("SmallAlbinoSpider", wp85[12])
	ns.CreateObject("Imp", wp87[3])
	ns.Frozen(ns.GetHost(), true)
	ns.SetDialog(obj5, ns.NORMAL, Worker2DialogStart, Worker2DialogEnd)
	ns.StartDialog(obj5, ns.GetHost())
}
func Worker2Return() {
	if !ns.IsCaller(obj5) {
		goto LABEL1
	}
	ns.CreatureGuard(obj5, ns.GetWaypointX(wp89), ns.GetWaypointY(wp89), ns.GetWaypointX(wp90), ns.GetWaypointY(wp90), 0)
	if gvar16 != 1 {
		goto LABEL2
	}
	ns.SetDialog(obj5, ns.NORMAL, Worker2DialogStart, Worker2DialogEnd)
	WorkerRescued()
	ns.GiveXp(ns.GetHost(), 250)
	ns.Chat(obj5, "Con03B.scr:WorkersReturn")
	ivar15 = 4
	ivar21 += 1
	if ivar21 != 3 {
		goto LABEL2
	}
	ivar15 = 5
LABEL2:
	gvar16 = 2
LABEL1:
	return
}
func Worker2Stop() {
	if !ns.IsCaller(obj5) {
		goto LABEL1
	}
	gvar128 = 1
	ns.Chat(obj5, "Con03B.scr:WorkersStop")
	ns.CreatureGuard(obj5, 1779, 2501, 1763, 2508, 60)
LABEL1:
	return
}
func Worker2Filter() int {
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
func Worker3Rec() {
	ns.ObjectGroupOff(gvar114)
	ns.GiveXp(ns.GetHost(), 100)
	ns.ObjectOn(obj6)
	ns.BecomePet(obj6)
	ns.CreatureFollow(obj6, ns.GetHost())
	ns.ObjectOn(obj110)
	ns.ObjectOn(obj111)
	ns.CreateObject("SmallAlbinoSpider", wp85[13])
	ns.CreateObject("SmallAlbinoSpider", wp85[14])
	ns.CreateObject("SmallAlbinoSpider", wp85[15])
	ns.CreateObject("SmallAlbinoSpider", wp85[16])
	ns.CreateObject("SmallAlbinoSpider", wp85[17])
	ns.CreateObject("SmallAlbinoSpider", wp85[18])
	ns.CreateObject("SmallAlbinoSpider", wp85[19])
	ns.CreateObject("SmallAlbinoSpider", wp85[20])
	ns.CreateObject("SmallAlbinoSpider", wp85[21])
	ns.CreateObject("SmallAlbinoSpider", wp85[22])
	ns.CreateObject("SmallAlbinoSpider", wp85[23])
	ns.CreateObject("Imp", wp87[0])
	ns.CreateObject("Imp", wp87[1])
	ns.CreateObject("Imp", wp87[2])
	ns.Frozen(ns.GetHost(), true)
	ns.SetDialog(obj6, ns.NORMAL, Worker3DialogStart, Worker3DialogEnd)
	ns.StartDialog(obj6, ns.GetHost())
}
func Worker3Return() {
	if !ns.IsCaller(obj6) {
		goto LABEL1
	}
	ns.CreatureGuard(obj6, ns.GetWaypointX(wp90), ns.GetWaypointY(wp90), ns.GetWaypointX(wp92), ns.GetWaypointY(wp92), 0)
	if gvar17 != 1 {
		goto LABEL2
	}
	ns.SetDialog(obj6, ns.NORMAL, Worker3DialogStart, Worker3DialogEnd)
	WorkerRescued()
	ns.Chat(obj6, "Con03B.scr:WorkersReturn")
	ns.GiveXp(ns.GetHost(), 250)
	ivar15 = 4
	ivar21 += 1
	if ivar21 != 3 {
		goto LABEL2
	}
	ivar15 = 5
LABEL2:
	gvar17 = 2
LABEL1:
	return
}
func Worker3Stop() {
	if !ns.IsCaller(obj6) {
		goto LABEL1
	}
	gvar129 = 1
	ns.Chat(obj6, "Con03B.scr:WorkersStop")
	ns.CreatureGuard(obj6, 1790, 2347, 1788, 2330, 60)
LABEL1:
	return
}
func Worker3Filter() int {
	if !ns.IsCaller(obj6) {
		goto LABEL1
	}
	return 1
	goto LABEL2
LABEL1:
	return 0
LABEL2:
	return 0
}
func Worker4Return() {
	if !ns.IsCaller(obj7) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	if gvar18 != 1 {
		goto LABEL1
	}
	ns.SetDialog(obj7, ns.NORMAL, Worker4DialogStart, Worker4DialogEnd)
	ns.CreatureGuard(obj7, ns.GetWaypointX(wp91), ns.GetWaypointY(wp91), ns.GetWaypointX(wp90), ns.GetWaypointY(wp90), 0)
	WorkerRescued()
	ns.GiveXp(ns.GetHost(), 250)
	ns.Chat(obj7, "Con03B.scr:WorkersReturn")
	ivar21 += 1
	if ivar21 != 5 {
		goto LABEL2
	}
	AllFound()
	goto LABEL1
LABEL2:
	gvar18 = 2
LABEL1:
	return
}
func Worker5Return() {
	if !ns.IsCaller(obj8) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	if gvar19 != 1 {
		goto LABEL1
	}
	ns.SetDialog(obj8, ns.NORMAL, Worker5DialogStart, Worker5DialogEnd)
	ns.CreatureGuard(obj8, ns.GetWaypointX(wp92), ns.GetWaypointY(wp92), ns.GetWaypointX(wp90), ns.GetWaypointY(wp90), 0)
	WorkerRescued()
	ns.GiveXp(ns.GetHost(), 250)
	ns.Chat(obj8, "Con03B.scr:WorkersReturn")
	ivar21 += 1
	if ivar21 != 5 {
		goto LABEL2
	}
	AllFound()
	goto LABEL1
LABEL2:
	gvar19 = 2
LABEL1:
	return
}
func Worker5Rec() {
	var v0 ns.ObjectID
	ns.ObjectGroupOff(gvar115)
	ns.GiveXp(ns.GetHost(), 100)
	ns.ObjectOn(obj8)
	ns.BecomePet(obj8)
	ns.CreatureFollow(obj8, ns.GetHost())
	ns.ObjectOn(obj7)
	ns.BecomePet(obj7)
	ns.CreatureFollow(obj7, ns.GetHost())
	v0 = ns.CreateObject("Imp", ns.Waypoint("CreateImp01"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Imp01Face"))) // FIXME
	v0 = ns.CreateObject("Imp", ns.Waypoint("CreateImp02"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Imp02Face"))) // FIXME
	v0 = ns.CreateObject("Imp", ns.Waypoint("CreateImp03"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Imp03Face"))) // FIXME
	v0 = ns.CreateObject("Imp", ns.Waypoint("CreateImp04"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Imp04Face"))) // FIXME
	v0 = ns.CreateObject("Imp", ns.Waypoint("CreateImp05"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Imp05Face"))) // FIXME
	v0 = ns.CreateObject("Imp", ns.Waypoint("CreateImp06"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Imp06Face"))) // FIXME
	v0 = ns.CreateObject("Imp", ns.Waypoint("CreateImp07"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Imp07Face"))) // FIXME
	v0 = ns.CreateObject("Scorpion", ns.Waypoint("CreateScorpion01"))
	ns.LookAtObject(v0, ns.ObjectID(ns.Waypoint("Scorpion01Face"))) // FIXME
	ns.Frozen(ns.GetHost(), true)
	ns.SetDialog(obj8, ns.NORMAL, Worker5DialogStart, Worker5DialogEnd)
	gvar18 = 1
	ns.StartDialog(obj8, ns.GetHost())
}
func Worker6Rec() {
	ns.ObjectGroupOff(gvar116)
	ns.GiveXp(ns.GetHost(), 100)
	ns.ObjectOn(obj9)
	ns.BecomePet(obj9)
	ns.CreatureFollow(obj9, ns.GetHost())
	ns.CreateObject("Bat", wp86[0])
	ns.CreateObject("Bat", wp86[1])
	ns.CreateObject("Bat", wp86[2])
	ns.CreateObject("Bat", wp86[3])
	ns.CreateObject("Bat", wp86[4])
	ns.CreateObject("Bat", wp86[5])
	ns.CreateObject("Bat", wp86[6])
	ns.Frozen(ns.GetHost(), true)
	ns.SetDialog(obj9, ns.NORMAL, Worker6DialogStart, Worker6DialogEnd)
	ns.StartDialog(obj9, ns.GetHost())
}
func Worker6Return() {
	if !ns.IsCaller(obj9) {
		goto LABEL1
	}
	ns.CreatureGuard(obj9, ns.GetWaypointX(wp93), ns.GetWaypointY(wp93), ns.GetWaypointX(wp89), ns.GetWaypointY(wp89), 0)
	if gvar20 != 1 {
		goto LABEL2
	}
	ns.SetDialog(obj9, ns.NORMAL, Worker6DialogStart, Worker6DialogEnd)
	WorkerRescued()
	ns.Chat(obj9, "Con03B.scr:WorkersReturn")
	ns.GiveXp(ns.GetHost(), 250)
	ivar21 += 1
	ivar15 = 4
	if ivar21 != 3 {
		goto LABEL2
	}
	ivar15 = 5
LABEL2:
	gvar20 = 2
LABEL1:
	return
}
func Worker6Stop() {
	if !ns.IsCaller(obj9) {
		goto LABEL1
	}
	gvar130 = 1
	ns.Chat(obj9, "Con03B.scr:WorkersStop")
	ns.CreatureGuard(obj9, 1906, 2465, 1907, 2447, 60)
LABEL1:
	return
}
func Worker6Filter() int {
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
func WorkerDie() {
	ns.DeathScreen(3)
}
func WorkerHurt() {
	r0 := ns.Random(1, 2)
	if r0 != 1 {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con03B.scr:WorkerHurt")
LABEL1:
	return
}
func WorkerEngage() {
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "MapShutdown":
		MapShutdown()
	case "MapEntry":
		MapEntry()
	case "PlayerDeath":
		PlayerDeath()
	}
}
