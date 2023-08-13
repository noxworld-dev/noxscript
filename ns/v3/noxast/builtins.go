package noxast

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"

	"github.com/noxworld-dev/noxscript/ns/asm"

	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
)

type builtinDef struct {
	Name  string
	Type  reflect.Type
	Idemp bool // no side effects
}

var builtins = []*builtinDef{
	asm.BuiltinWall:                      {Name: "Wall", Type: reflect.TypeOf(ns3.Wall), Idemp: true},
	asm.BuiltinWallOpen:                  {Name: "WallOpen", Type: reflect.TypeOf(ns3.WallOpen)},
	asm.BuiltinWallGroupOpen:             {Name: "WallGroupOpen", Type: reflect.TypeOf(ns3.WallGroupOpen)},
	asm.BuiltinWallClose:                 {Name: "WallClose", Type: reflect.TypeOf(ns3.WallClose)},
	asm.BuiltinWallGroupClose:            {Name: "WallGroupClose", Type: reflect.TypeOf(ns3.WallGroupClose)},
	asm.BuiltinWallToggle:                {Name: "WallToggle", Type: reflect.TypeOf(ns3.WallToggle)},
	asm.BuiltinWallGroupToggle:           {Name: "WallGroupToggle", Type: reflect.TypeOf(ns3.WallGroupToggle)},
	asm.BuiltinWallBreak:                 {Name: "WallBreak", Type: reflect.TypeOf(ns3.WallBreak)},
	asm.BuiltinWallGroupBreak:            {Name: "WallGroupBreak", Type: reflect.TypeOf(ns3.WallGroupBreak)},
	asm.BuiltinSecondTimer:               {Name: "SecondTimer", Type: reflect.TypeOf(ns3.SecondTimer)},
	asm.BuiltinFrameTimer:                {Name: "FrameTimer", Type: reflect.TypeOf(ns3.FrameTimer)},
	asm.BuiltinMove:                      {Name: "Move", Type: reflect.TypeOf(ns3.Move)},
	asm.BuiltinGroupMove:                 {Name: "GroupMove", Type: reflect.TypeOf(ns3.GroupMove)},
	asm.BuiltinLookAtDirection:           {Name: "LookAtDirection", Type: reflect.TypeOf(ns3.LookAtDirection)},
	asm.BuiltinGroupLookAtDirection:      {Name: "GroupLookAtDirection", Type: reflect.TypeOf(ns3.GroupLookAtDirection)},
	asm.BuiltinObjectOn:                  {Name: "ObjectOn", Type: reflect.TypeOf(ns3.ObjectOn)},
	asm.BuiltinObjectGroupOn:             {Name: "ObjectGroupOn", Type: reflect.TypeOf(ns3.ObjectGroupOn)},
	asm.BuiltinWaypointOn:                {Name: "WayPointOn", Type: reflect.TypeOf(ns3.WayPointOn)},
	asm.BuiltinWaypointGroupOn:           {Name: "WayPointGroupOn", Type: reflect.TypeOf(ns3.WayPointGroupOn)},
	asm.BuiltinObjectOff:                 {Name: "ObjectOff", Type: reflect.TypeOf(ns3.ObjectOff)},
	asm.BuiltinObjectGroupOff:            {Name: "ObjectGroupOff", Type: reflect.TypeOf(ns3.ObjectGroupOff)},
	asm.BuiltinWaypointOff:               {Name: "WayPointOff", Type: reflect.TypeOf(ns3.WayPointOff)},
	asm.BuiltinWaypointGroupOff:          {Name: "WayPointGroupOff", Type: reflect.TypeOf(ns3.WayPointGroupOff)},
	asm.BuiltinObjectToggle:              {Name: "ObjectToggle", Type: reflect.TypeOf(ns3.ObjectToggle)},
	asm.BuiltinObjectGroupToggle:         {Name: "ObjectGroupToggle", Type: reflect.TypeOf(ns3.ObjectGroupToggle)},
	asm.BuiltinWaypointToggle:            {Name: "WayPointToggle", Type: reflect.TypeOf(ns3.WayPointToggle)},
	asm.BuiltinWaypointGroupToggle:       {Name: "WayPointGroupToggle", Type: reflect.TypeOf(ns3.WayPointGroupToggle)},
	asm.BuiltinDelete:                    {Name: "Delete", Type: reflect.TypeOf(ns3.Delete)},
	asm.BuiltinGroupDelete:               {Name: "GroupDelete", Type: reflect.TypeOf(ns3.GroupDelete)},
	asm.BuiltinWander:                    {Name: "Wander", Type: reflect.TypeOf(ns3.Wander)},
	asm.BuiltinGroupWander:               {Name: "GroupWander", Type: reflect.TypeOf(ns3.GroupWander)},
	asm.BuiltinUnused1f:                  {Name: "Unused1f", Type: reflect.TypeOf(ns3.Unused1f)},
	asm.BuiltinUnused20:                  {Name: "Unused20", Type: reflect.TypeOf(ns3.Unused20)},
	asm.BuiltinGoBackHome:                {Name: "GoBackHome", Type: reflect.TypeOf(ns3.GoBackHome)},
	asm.BuiltinAudioEvent:                {Name: "AudioEvent", Type: reflect.TypeOf(ns3.AudioEvent)},
	asm.BuiltinPrint:                     {Name: "Print", Type: reflect.TypeOf(ns3.Print)},
	asm.BuiltinPrintToAll:                {Name: "PrintToAll", Type: reflect.TypeOf(ns3.PrintToAll)},
	asm.BuiltinChat:                      {Name: "Chat", Type: reflect.TypeOf(ns3.Chat)},
	asm.BuiltinStopScript:                {Name: "StopScript", Type: reflect.TypeOf(ns3.StopScript)},
	asm.BuiltinUnlockDoor:                {Name: "UnlockDoor", Type: reflect.TypeOf(ns3.UnlockDoor)},
	asm.BuiltinLockDoor:                  {Name: "LockDoor", Type: reflect.TypeOf(ns3.LockDoor)},
	asm.BuiltinIsObjectOn:                {Name: "IsObjectOn", Type: reflect.TypeOf(ns3.IsObjectOn), Idemp: true},
	asm.BuiltinIsWaypointOn:              {Name: "IsWaypointOn", Type: reflect.TypeOf(ns3.IsWaypointOn), Idemp: true},
	asm.BuiltinIsLocked:                  {Name: "IsLocked", Type: reflect.TypeOf(ns3.IsLocked), Idemp: true},
	asm.BuiltinRandomFloat:               {Name: "RandomFloat", Type: reflect.TypeOf(ns3.RandomFloat)},
	asm.BuiltinRandom:                    {Name: "Random", Type: reflect.TypeOf(ns3.Random)},
	asm.BuiltinSecondTimerWithArg:        {Name: "SecondTimerWithArg", Type: reflect.TypeOf(ns3.SecondTimerWithArg)},
	asm.BuiltinFrameTimerWithArg:         {Name: "FrameTimerWithArg", Type: reflect.TypeOf(ns3.FrameTimerWithArg)},
	asm.BuiltinIntToString:               {Name: "IntToString", Type: reflect.TypeOf(ns3.IntToString), Idemp: true},
	asm.BuiltinFloatToString:             {Name: "FloatToString", Type: reflect.TypeOf(ns3.FloatToString), Idemp: true},
	asm.BuiltinCreateObject:              {Name: "CreateObject", Type: reflect.TypeOf(ns3.CreateObject)},
	asm.BuiltinDamage:                    {Name: "Damage", Type: reflect.TypeOf(ns3.Damage)},
	asm.BuiltinGroupDamage:               {Name: "GroupDamage", Type: reflect.TypeOf(ns3.GroupDamage)},
	asm.BuiltinCreateMover:               {Name: "CreateMover", Type: reflect.TypeOf(ns3.CreateMover)},
	asm.BuiltinGroupCreateMover:          {Name: "GroupCreateMover", Type: reflect.TypeOf(ns3.GroupCreateMover)},
	asm.BuiltinAwardSpell:                {Name: "AwardSpell", Type: reflect.TypeOf(ns3.AwardSpell)},
	asm.BuiltinGroupAwardSpell:           {Name: "GroupAwardSpell", Type: reflect.TypeOf(ns3.GroupAwardSpell)},
	asm.BuiltinEnchant:                   {Name: "Enchant", Type: reflect.TypeOf(ns3.Enchant)},
	asm.BuiltinGroupEnchant:              {Name: "GroupEnchant", Type: reflect.TypeOf(ns3.GroupEnchant)},
	asm.BuiltinGetHost:                   {Name: "GetHost", Type: reflect.TypeOf(ns3.GetHost), Idemp: true},
	asm.BuiltinObject:                    {Name: "Object", Type: reflect.TypeOf(ns3.Object), Idemp: true},
	asm.BuiltinGetObjectX:                {Name: "GetObjectX", Type: reflect.TypeOf(ns3.GetObjectX), Idemp: true},
	asm.BuiltinGetWaypointX:              {Name: "GetWaypointX", Type: reflect.TypeOf(ns3.GetWaypointX), Idemp: true},
	asm.BuiltinGetObjectY:                {Name: "GetObjectY", Type: reflect.TypeOf(ns3.GetObjectY), Idemp: true},
	asm.BuiltinGetWaypointY:              {Name: "GetWaypointY", Type: reflect.TypeOf(ns3.GetWaypointY), Idemp: true},
	asm.BuiltinGetObjectZ:                {Name: "GetObjectZ", Type: reflect.TypeOf(ns3.GetObjectZ), Idemp: true},
	asm.BuiltinGetDirection:              {Name: "GetDirection", Type: reflect.TypeOf(ns3.GetDirection), Idemp: true},
	asm.BuiltinMoveObject:                {Name: "MoveObject", Type: reflect.TypeOf(ns3.MoveObject)},
	asm.BuiltinMoveWaypoint:              {Name: "MoveWaypoint", Type: reflect.TypeOf(ns3.MoveWaypoint)},
	asm.BuiltinRaise:                     {Name: "Raise", Type: reflect.TypeOf(ns3.Raise)},
	asm.BuiltinLookWithAngle:             {Name: "LookWithAngle", Type: reflect.TypeOf(ns3.LookWithAngle)},
	asm.BuiltinPushObjectTo:              {Name: "PushObjectTo", Type: reflect.TypeOf(ns3.PushObjectTo)},
	asm.BuiltinPushObject:                {Name: "PushObject", Type: reflect.TypeOf(ns3.PushObject)},
	asm.BuiltinGetLastItem:               {Name: "GetLastItem", Type: reflect.TypeOf(ns3.GetLastItem), Idemp: true},
	asm.BuiltinGetPreviousItem:           {Name: "GetPreviousItem", Type: reflect.TypeOf(ns3.GetPreviousItem), Idemp: true},
	asm.BuiltinHasItem:                   {Name: "HasItem", Type: reflect.TypeOf(ns3.HasItem), Idemp: true},
	asm.BuiltinGetHolder:                 {Name: "GetHolder", Type: reflect.TypeOf(ns3.GetHolder), Idemp: true},
	asm.BuiltinPickup:                    {Name: "Pickup", Type: reflect.TypeOf(ns3.Pickup)},
	asm.BuiltinDrop:                      {Name: "Drop", Type: reflect.TypeOf(ns3.Drop)},
	asm.BuiltinHasClass:                  {Name: "HasClass", Type: reflect.TypeOf(ns3.HasClass), Idemp: true},
	asm.BuiltinUnused50:                  {Name: "Unused50", Type: reflect.TypeOf(ns3.Unused50)},
	asm.BuiltinHasEnchant:                {Name: "HasEnchant", Type: reflect.TypeOf(ns3.HasEnchant), Idemp: true},
	asm.BuiltinEnchantOff:                {Name: "EnchantOff", Type: reflect.TypeOf(ns3.EnchantOff)},
	asm.BuiltinCurrentHealth:             {Name: "CurrentHealth", Type: reflect.TypeOf(ns3.CurrentHealth), Idemp: true},
	asm.BuiltinMaxHealth:                 {Name: "MaxHealth", Type: reflect.TypeOf(ns3.MaxHealth), Idemp: true},
	asm.BuiltinRestoreHealth:             {Name: "RestoreHealth", Type: reflect.TypeOf(ns3.RestoreHealth)},
	asm.BuiltinDistance:                  {Name: "Distance", Type: reflect.TypeOf(ns3.Distance), Idemp: true},
	asm.BuiltinIsVisibleTo:               {Name: "IsVisibleTo", Type: reflect.TypeOf(ns3.IsVisibleTo), Idemp: true},
	asm.BuiltinUnused58:                  {Name: "Unused58", Type: reflect.TypeOf(ns3.Unused58)},
	asm.BuiltinUnused59:                  {Name: "Unused59", Type: reflect.TypeOf(ns3.Unused59)},
	asm.BuiltinUnused5a:                  {Name: "Unused5a", Type: reflect.TypeOf(ns3.Unused5a)},
	asm.BuiltinUnused5b:                  {Name: "Unused5b", Type: reflect.TypeOf(ns3.Unused5b)},
	asm.BuiltinUnused5c:                  {Name: "Unused5c", Type: reflect.TypeOf(ns3.Unused5c)},
	asm.BuiltinUnused5d:                  {Name: "Unused5d", Type: reflect.TypeOf(ns3.Unused5d)},
	asm.BuiltinUnused5e:                  {Name: "Unused5e", Type: reflect.TypeOf(ns3.Unused5e)},
	asm.BuiltinGetCharacterData:          {Name: "GetCharacterData", Type: reflect.TypeOf(ns3.GetCharacterData)},
	asm.BuiltinLookAtObject:              {Name: "LookAtObject", Type: reflect.TypeOf(ns3.LookAtObject)},
	asm.BuiltinWalk:                      {Name: "Walk", Type: reflect.TypeOf(ns3.Walk)},
	asm.BuiltinGroupWalk:                 {Name: "GroupWalk", Type: reflect.TypeOf(ns3.GroupWalk)},
	asm.BuiltinCancelTimer:               {Name: "CancelTimer", Type: reflect.TypeOf(ns3.CancelTimer)},
	asm.BuiltinEffect:                    {Name: "Effect", Type: reflect.TypeOf(ns3.Effect)},
	asm.BuiltinSetOwner:                  {Name: "SetOwner", Type: reflect.TypeOf(ns3.SetOwner)},
	asm.BuiltinGroupSetOwner:             {Name: "GroupSetOwner", Type: reflect.TypeOf(ns3.GroupSetOwner)},
	asm.BuiltinSetOwners:                 {Name: "SetOwners", Type: reflect.TypeOf(ns3.SetOwners)},
	asm.BuiltinGroupSetOwners:            {Name: "GroupSetOwners", Type: reflect.TypeOf(ns3.GroupSetOwners)},
	asm.BuiltinIsOwnedBy:                 {Name: "IsOwnedBy", Type: reflect.TypeOf(ns3.IsOwnedBy), Idemp: true},
	asm.BuiltinGroupIsOwnedBy:            {Name: "GroupIsOwnedBy", Type: reflect.TypeOf(ns3.GroupIsOwnedBy), Idemp: true},
	asm.BuiltinIsOwnedByAny:              {Name: "IsOwnedByAny", Type: reflect.TypeOf(ns3.IsOwnedByAny), Idemp: true},
	asm.BuiltinGroupIsOwnedByAny:         {Name: "GroupIsOwnedByAny", Type: reflect.TypeOf(ns3.GroupIsOwnedByAny), Idemp: true},
	asm.BuiltinClearOwner:                {Name: "ClearOwner", Type: reflect.TypeOf(ns3.ClearOwner)},
	asm.BuiltinWaypoint:                  {Name: "Waypoint", Type: reflect.TypeOf(ns3.Waypoint), Idemp: true},
	asm.BuiltinWaypointGroup:             {Name: "WaypointGroup", Type: reflect.TypeOf(ns3.WaypointGroup), Idemp: true},
	asm.BuiltinObjectGroup:               {Name: "ObjectGroup", Type: reflect.TypeOf(ns3.ObjectGroup), Idemp: true},
	asm.BuiltinWallGroup:                 {Name: "WallGroup", Type: reflect.TypeOf(ns3.WallGroup), Idemp: true},
	asm.BuiltinChatTimerSeconds:          {Name: "ChatTimerSeconds", Type: reflect.TypeOf(ns3.ChatTimerSeconds)},
	asm.BuiltinChatTimer:                 {Name: "ChatTimer", Type: reflect.TypeOf(ns3.ChatTimer)},
	asm.BuiltinUnused74:                  {Name: "Unused74", Type: reflect.TypeOf(ns3.Unused74)},
	asm.BuiltinDestroyChat:               {Name: "DestroyChat", Type: reflect.TypeOf(ns3.DestroyChat)},
	asm.BuiltinDestroyEveryChat:          {Name: "DestroyEveryChat", Type: reflect.TypeOf(ns3.DestroyEveryChat)},
	asm.BuiltinSetQuestStatus:            {Name: "SetQuestStatus", Type: reflect.TypeOf(ns3.SetQuestStatus)},
	asm.BuiltinSetQuestStatusFloat:       {Name: "SetQuestStatusFloat", Type: reflect.TypeOf(ns3.SetQuestStatusFloat)},
	asm.BuiltinGetQuestStatus:            {Name: "GetQuestStatus", Type: reflect.TypeOf(ns3.GetQuestStatus), Idemp: true},
	asm.BuiltinGetQuestStatusFloat:       {Name: "GetQuestStatusFloat", Type: reflect.TypeOf(ns3.GetQuestStatusFloat), Idemp: true},
	asm.BuiltinResetQuestStatus:          {Name: "ResetQuestStatus", Type: reflect.TypeOf(ns3.ResetQuestStatus)},
	asm.BuiltinIsTrigger:                 {Name: "IsTrigger", Type: reflect.TypeOf(ns3.IsTrigger), Idemp: true},
	asm.BuiltinIsCaller:                  {Name: "IsCaller", Type: reflect.TypeOf(ns3.IsCaller), Idemp: true},
	asm.BuiltinSetDialog:                 {Name: "SetDialog", Type: reflect.TypeOf(ns3.SetDialog)},
	asm.BuiltinCancelDialog:              {Name: "CancelDialog", Type: reflect.TypeOf(ns3.CancelDialog)},
	asm.BuiltinStoryPic:                  {Name: "StoryPic", Type: reflect.TypeOf(ns3.StoryPic)},
	asm.BuiltinTellStory:                 {Name: "TellStory", Type: reflect.TypeOf(ns3.TellStory)},
	asm.BuiltinStartDialog:               {Name: "StartDialog", Type: reflect.TypeOf(ns3.StartDialog)},
	asm.BuiltinCastSpellObjectObject:     {Name: "CastSpellObjectObject", Type: reflect.TypeOf(ns3.CastSpellObjectObject)},
	asm.BuiltinCastSpellObjectLocation:   {Name: "CastSpellObjectLocation", Type: reflect.TypeOf(ns3.CastSpellObjectLocation)},
	asm.BuiltinCastSpellLocationObject:   {Name: "CastSpellLocationObject", Type: reflect.TypeOf(ns3.CastSpellLocationObject)},
	asm.BuiltinCastSpellLocationLocation: {Name: "CastSpellLocationLocation", Type: reflect.TypeOf(ns3.CastSpellLocationLocation)},
	asm.BuiltinUnBlind:                   {Name: "UnBlind", Type: reflect.TypeOf(ns3.UnBlind)},
	asm.BuiltinBlind:                     {Name: "Blind", Type: reflect.TypeOf(ns3.Blind)},
	asm.BuiltinWideScreen:                {Name: "WideScreen", Type: reflect.TypeOf(ns3.WideScreen)},
	asm.BuiltinGetElevatorStatus:         {Name: "GetElevatorStatus", Type: reflect.TypeOf(ns3.GetElevatorStatus), Idemp: true},
	asm.BuiltinCreatureGuard:             {Name: "CreatureGuard", Type: reflect.TypeOf(ns3.CreatureGuard)},
	asm.BuiltinCreatureGroupGuard:        {Name: "CreatureGroupGuard", Type: reflect.TypeOf(ns3.CreatureGroupGuard)},
	asm.BuiltinCreatureHunt:              {Name: "CreatureHunt", Type: reflect.TypeOf(ns3.CreatureHunt)},
	asm.BuiltinCreatureGroupHunt:         {Name: "CreatureGroupHunt", Type: reflect.TypeOf(ns3.CreatureGroupHunt)},
	asm.BuiltinCreatureIdle:              {Name: "CreatureIdle", Type: reflect.TypeOf(ns3.CreatureIdle)},
	asm.BuiltinCreatureGroupIdle:         {Name: "CreatureGroupIdle", Type: reflect.TypeOf(ns3.CreatureGroupIdle)},
	asm.BuiltinCreatureFollow:            {Name: "CreatureFollow", Type: reflect.TypeOf(ns3.CreatureFollow)},
	asm.BuiltinCreatureGroupFollow:       {Name: "CreatureGroupFollow", Type: reflect.TypeOf(ns3.CreatureGroupFollow)},
	asm.BuiltinAggressionLevel:           {Name: "AggressionLevel", Type: reflect.TypeOf(ns3.AggressionLevel)},
	asm.BuiltinGroupAggressionLevel:      {Name: "GroupAggressionLevel", Type: reflect.TypeOf(ns3.GroupAggressionLevel)},
	asm.BuiltinHitLocation:               {Name: "HitLocation", Type: reflect.TypeOf(ns3.HitLocation)},
	asm.BuiltinGroupHitLocation:          {Name: "GroupHitLocation", Type: reflect.TypeOf(ns3.GroupHitLocation)},
	asm.BuiltinHitFarLocation:            {Name: "HitFarLocation", Type: reflect.TypeOf(ns3.HitFarLocation)},
	asm.BuiltinGroupHitFarLocation:       {Name: "GroupHitFarLocation", Type: reflect.TypeOf(ns3.GroupHitFarLocation)},
	asm.BuiltinSetRoamFlag:               {Name: "SetRoamFlag", Type: reflect.TypeOf(ns3.SetRoamFlag)},
	asm.BuiltinGroupSetRoamFlag:          {Name: "GroupSetRoamFlag", Type: reflect.TypeOf(ns3.GroupSetRoamFlag)},
	asm.BuiltinAttack:                    {Name: "Attack", Type: reflect.TypeOf(ns3.Attack)},
	asm.BuiltinGroupAttack:               {Name: "GroupAttack", Type: reflect.TypeOf(ns3.GroupAttack)},
	asm.BuiltinJournalEntry:              {Name: "JournalEntry", Type: reflect.TypeOf(ns3.JournalEntry)},
	asm.BuiltinJournalDelete:             {Name: "JournalDelete", Type: reflect.TypeOf(ns3.JournalDelete)},
	asm.BuiltinJournalEdit:               {Name: "JournalEdit", Type: reflect.TypeOf(ns3.JournalEdit)},
	asm.BuiltinRetreatLevel:              {Name: "RetreatLevel", Type: reflect.TypeOf(ns3.RetreatLevel)},
	asm.BuiltinGroupRetreatLevel:         {Name: "GroupRetreatLevel", Type: reflect.TypeOf(ns3.GroupRetreatLevel)},
	asm.BuiltinResumeLevel:               {Name: "ResumeLevel", Type: reflect.TypeOf(ns3.ResumeLevel)},
	asm.BuiltinGroupResumeLevel:          {Name: "GroupResumeLevel", Type: reflect.TypeOf(ns3.GroupResumeLevel)},
	asm.BuiltinRunAway:                   {Name: "RunAway", Type: reflect.TypeOf(ns3.RunAway)},
	asm.BuiltinGroupRunAway:              {Name: "GroupRunAway", Type: reflect.TypeOf(ns3.GroupRunAway)},
	asm.BuiltinPauseObject:               {Name: "PauseObject", Type: reflect.TypeOf(ns3.PauseObject)},
	asm.BuiltinGroupPauseObject:          {Name: "GroupPauseObject", Type: reflect.TypeOf(ns3.GroupPauseObject)},
	asm.BuiltinIsAttackedBy:              {Name: "IsAttackedBy", Type: reflect.TypeOf(ns3.IsAttackedBy)},
	asm.BuiltinGetGold:                   {Name: "GetGold", Type: reflect.TypeOf(ns3.GetGold), Idemp: true},
	asm.BuiltinChangeGold:                {Name: "ChangeGold", Type: reflect.TypeOf(ns3.ChangeGold)},
	asm.BuiltinGetAnswer:                 {Name: "GetAnswer", Type: reflect.TypeOf(ns3.GetAnswer), Idemp: true},
	asm.BuiltinGiveXp:                    {Name: "GiveXp", Type: reflect.TypeOf(ns3.GiveXp)},
	asm.BuiltinHasSubclass:               {Name: "HasSubclass", Type: reflect.TypeOf(ns3.HasSubclass), Idemp: true},
	asm.BuiltinAutoSave:                  {Name: "AutoSave", Type: reflect.TypeOf(ns3.AutoSave)},
	asm.BuiltinMusic:                     {Name: "Music", Type: reflect.TypeOf(ns3.Music)},
	asm.BuiltinStartupScreen:             {Name: "StartupScreen", Type: reflect.TypeOf(ns3.StartupScreen)},
	asm.BuiltinIsTalking:                 {Name: "IsTalking", Type: reflect.TypeOf(ns3.IsTalking), Idemp: true},
	asm.BuiltinGetTrigger:                {Name: "GetTrigger", Type: reflect.TypeOf(ns3.GetTrigger), Idemp: true},
	asm.BuiltinGetCaller:                 {Name: "GetCaller", Type: reflect.TypeOf(ns3.GetCaller), Idemp: true},
	asm.BuiltinMakeFriendly:              {Name: "MakeFriendly", Type: reflect.TypeOf(ns3.MakeFriendly)},
	asm.BuiltinMakeEnemy:                 {Name: "MakeEnemy", Type: reflect.TypeOf(ns3.MakeEnemy)},
	asm.BuiltinBecomePet:                 {Name: "BecomePet", Type: reflect.TypeOf(ns3.BecomePet)},
	asm.BuiltinBecomeEnemy:               {Name: "BecomeEnemy", Type: reflect.TypeOf(ns3.BecomeEnemy)},
	asm.BuiltinUnknownb8:                 {Name: "Unknownb8", Type: reflect.TypeOf(ns3.Unknownb8)},
	asm.BuiltinUnknownb9:                 {Name: "Unknownb9", Type: reflect.TypeOf(ns3.Unknownb9)},
	asm.BuiltinSetHalberd:                {Name: "SetHalberd", Type: reflect.TypeOf(ns3.SetHalberd)},
	asm.BuiltinDeathScreen:               {Name: "DeathScreen", Type: reflect.TypeOf(ns3.DeathScreen)},
	asm.BuiltinFrozen:                    {Name: "Frozen", Type: reflect.TypeOf(ns3.Frozen)},
	asm.BuiltinNoWallSound:               {Name: "NoWallSound", Type: reflect.TypeOf(ns3.NoWallSound)},
	asm.BuiltinSetCallback:               {Name: "SetCallback", Type: reflect.TypeOf(ns3.SetCallback)},
	asm.BuiltinDeleteObjectTimer:         {Name: "DeleteObjectTimer", Type: reflect.TypeOf(ns3.DeleteObjectTimer)},
	asm.BuiltinTrapSpells:                {Name: "TrapSpells", Type: reflect.TypeOf(ns3.TrapSpells)},
	asm.BuiltinIsTrading:                 {Name: "IsTrading", Type: reflect.TypeOf(ns3.IsTrading)},
	asm.BuiltinClearMessages:             {Name: "ClearMessages", Type: reflect.TypeOf(ns3.ClearMessages)},
	asm.BuiltinSetShopkeeperText:         {Name: "SetShopkeeperText", Type: reflect.TypeOf(ns3.SetShopkeeperText)},
	asm.BuiltinUnknownc4:                 {Name: "Unknownc4", Type: reflect.TypeOf(ns3.Unknownc4)},
	asm.BuiltinIsSummoned:                {Name: "IsSummoned", Type: reflect.TypeOf(ns3.IsSummoned), Idemp: true},
	asm.BuiltinZombieStayDown:            {Name: "ZombieStayDown", Type: reflect.TypeOf(ns3.ZombieStayDown)},
	asm.BuiltinZombieGroupStayDown:       {Name: "ZombieGroupStayDown", Type: reflect.TypeOf(ns3.ZombieGroupStayDown)},
	asm.BuiltinRaiseZombie:               {Name: "RaiseZombie", Type: reflect.TypeOf(ns3.RaiseZombie)},
	asm.BuiltinRaiseZombieGroup:          {Name: "RaiseZombieGroup", Type: reflect.TypeOf(ns3.RaiseZombieGroup)},
	asm.BuiltinMusicPushEvent:            {Name: "MusicPushEvent", Type: reflect.TypeOf(ns3.MusicPushEvent)},
	asm.BuiltinMusicPopEvent:             {Name: "MusicPopEvent", Type: reflect.TypeOf(ns3.MusicPopEvent)},
	asm.BuiltinMusicEvent:                {Name: "MusicEvent", Type: reflect.TypeOf(ns3.MusicEvent)},
	asm.BuiltinIsGameBall:                {Name: "IsGameBall", Type: reflect.TypeOf(ns3.IsGameBall), Idemp: true},
	asm.BuiltinIsCrown:                   {Name: "IsCrown", Type: reflect.TypeOf(ns3.IsCrown), Idemp: true},
	asm.BuiltinEndGame:                   {Name: "EndGame", Type: reflect.TypeOf(ns3.EndGame)},
	asm.BuiltinImmediateBlind:            {Name: "ImmediateBlind", Type: reflect.TypeOf(ns3.ImmediateBlind)},
	asm.BuiltinChangeScore:               {Name: "ChangeScore", Type: reflect.TypeOf(ns3.ChangeScore)},
	asm.BuiltinGetScore:                  {Name: "GetScore", Type: reflect.TypeOf(ns3.GetScore), Idemp: true},
}

type callBuiltinFunc func(args []ast.Expr) ast.Expr

func (t *translator) callBuiltin(ind asm.Builtin) (rt reflect.Type, callExp callBuiltinFunc) {
	switch ind {
	case asm.BuiltinIntToString:
		rt = reflect.TypeOf((*func(int) string)(nil)).Elem()
		callExp = func(args []ast.Expr) ast.Expr {
			return callSel(t.imports.strconv, "Itoa", args[0])
		}
	case asm.BuiltinFloatToString:
		rt = reflect.TypeOf((*func(float32) string)(nil)).Elem()
		callExp = func(args []ast.Expr) ast.Expr {
			return callSel(t.imports.strconv, "FormatFloat",
				call(ast.NewIdent("float64"), args[0]),
				&ast.BasicLit{Kind: token.CHAR, Value: "'g'"},
				intLit(-1), intLit(32),
			)
		}
	default:
		var fnc ast.Expr
		if ind >= 0 && int(ind) < len(builtins) {
			fnc = t.builtins[ind]
			if fnc == nil {
				panic(fmt.Errorf("no builtin: %v", ind))
			}
		} else {
			fnc = ast.NewIdent(fmt.Sprintf("builtin_overflow_%d", uint32(ind)))
		}
		rt, _ = getType(fnc).(reflect.Type)
		callExp = func(args []ast.Expr) ast.Expr {
			switch ind {
			case asm.BuiltinSetDialog:
				if s, ok := asStr(args[1]); ok {
					switch s {
					case "NORMAL", "NEXT", "YESNO", "YESNONEXT":
						args[1] = sel(t.imports.ns3, s)
					}
				}
				if fp, ok := asInt(args[2]); ok && fp >= 0 && fp < len(t.funcs) {
					args[2] = t.funcs[fp]
				}
				if fp, ok := asInt(args[3]); ok && fp >= 0 && fp < len(t.funcs) {
					args[3] = t.funcs[fp]
				}
			case asm.BuiltinAudioEvent, asm.BuiltinTellStory:
				if s, ok := asStr(args[0]); ok {
					switch s {
					case "RestoreHealth":
						s = "RestoreHealthName"
					}
					args[0] = sel(t.imports.ns3, s)
				}
			case asm.BuiltinEnchant, asm.BuiltinGroupEnchant, asm.BuiltinEnchantOff, asm.BuiltinHasEnchant:
				if s, ok := asStr(args[1]); ok && strings.HasPrefix(s, "ENCHANT_") {
					args[1] = sel(t.imports.ns3, s)
				}
			case asm.BuiltinEffect:
				if s, ok := asStr(args[0]); ok {
					args[0] = sel(t.imports.ns3, s)
				}
			case asm.BuiltinSetCallback:
				if fp, ok := asInt(args[2]); ok && fp >= 0 && fp < len(t.funcs) {
					args[2] = t.funcs[fp]
				}
			case asm.BuiltinSecondTimer, asm.BuiltinFrameTimer:
				if fp, ok := asInt(args[1]); ok && fp >= 0 && fp < len(t.funcs) {
					args[1] = t.funcs[fp]
				}
			case asm.BuiltinSecondTimerWithArg, asm.BuiltinFrameTimerWithArg:
				if fp, ok := asInt(args[2]); ok && fp >= 0 && fp < len(t.funcs) {
					args[2] = t.funcs[fp]
				}
			case asm.BuiltinHasClass, asm.BuiltinHasSubclass:
				if s, ok := asStr(args[1]); ok {
					args[1] = sel(t.imports.ns3, s)
				}
			case asm.BuiltinAwardSpell, asm.BuiltinGroupAwardSpell:
				if s, ok := asStr(args[1]); ok && strings.HasPrefix(s, "SPELL_") {
					args[1] = sel(t.imports.ns3, s)
				}
			case asm.BuiltinCastSpellLocationObject, asm.BuiltinCastSpellObjectLocation,
				asm.BuiltinCastSpellObjectObject, asm.BuiltinCastSpellLocationLocation:
				if s, ok := asStr(args[0]); ok && strings.HasPrefix(s, "SPELL_") {
					args[0] = sel(t.imports.ns3, s)
				}
			case asm.BuiltinTrapSpells:
				if s, ok := asStr(args[1]); ok && strings.HasPrefix(s, "SPELL_") {
					args[1] = sel(t.imports.ns3, s)
				} else if s == "NULL" {
					args[1] = stringLit("")
				}
				if s, ok := asStr(args[2]); ok && strings.HasPrefix(s, "SPELL_") {
					args[2] = sel(t.imports.ns3, s)
				} else if s == "NULL" {
					args[2] = stringLit("")
				}
				if s, ok := asStr(args[3]); ok && strings.HasPrefix(s, "SPELL_") {
					args[3] = sel(t.imports.ns3, s)
				} else if s == "NULL" {
					args[3] = stringLit("")
				}
			case asm.BuiltinLookAtDirection, asm.BuiltinGroupLookAtDirection:
				if d, ok := asInt(args[1]); ok {
					var name string
					switch d {
					case 0:
						name = "NW"
					case 1:
						name = "N"
					case 2:
						name = "NE"
					case 3:
						name = "W"
					case 5:
						name = "E"
					case 6:
						name = "SW"
					case 7:
						name = "S"
					case 8:
						name = "SE"
					}
					if name != "" {
						args[1] = sel(t.imports.ns3, name)
					}
				}
			}
			return &ast.CallExpr{Fun: fnc, Args: args}
		}
	}
	return
}
