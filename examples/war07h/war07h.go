package war07h

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4 ns.ObjectID
	obj5 ns.ObjectID
	obj6 ns.ObjectID
)

func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.Music(22, 100)
	ns.FrameTimer(10, BriefingBegin)
}
func BriefingBegin() {
	ns.StartDialog(obj4, ns.GetHost())
}
func JandorBriefStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:JandorTalk01")
}
func JandorBriefEnd() {
	ns.JournalEntry(ns.GetHost(), "GetHON", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.CancelDialog(obj4)
}
func LanceDialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "War01A.scr:Guard1Talk01")
}
func LanceDialogEnd() {
}
func MapInitialize() {
	obj4 = ns.Object("Jandor")
	obj5 = ns.Object("Lance")
	obj6 = ns.Object("FacadeGate")
	ns.StoryPic(obj4, "AirshipCaptainPic")
	ns.SetDialog(obj4, ns.NORMAL, JandorBriefStart, JandorBriefEnd)
	ns.LockDoor(obj6)
	ns.StartupScreen(7)
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
