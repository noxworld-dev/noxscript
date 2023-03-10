package ns

// This array makes sure we have all NoxScript builtins implemented.
var _ = [211]any{
	Wall,
	WallOpen,
	WallGroupOpen,
	WallClose,
	WallGroupClose,
	WallToggle,
	WallGroupToggle,
	WallBreak,
	WallGroupBreak,
	SecondTimer,
	FrameTimer,
	Move,
	GroupMove,
	LookAtDirection,
	GroupLookAtDirection,
	ObjectOn,
	ObjectGroupOn,
	WayPointOn,
	WayPointGroupOn,
	ObjectOff,
	ObjectGroupOff,
	WayPointOff,
	WayPointGroupOff,
	ObjectToggle,
	ObjectGroupToggle,
	WayPointToggle,
	WayPointGroupToggle,
	Delete,
	GroupDelete,
	Wander,
	GroupWander,
	Unused1f,
	Unused20,
	GoBackHome,
	AudioEvent,
	Print,
	PrintToAll,
	Chat,
	StopScript,
	UnlockDoor,
	LockDoor,
	IsObjectOn,
	IsWaypointOn,
	IsLocked,
	RandomFloat,
	Random,
	SecondTimerWithArg,
	FrameTimerWithArg,
	IntToString,
	FloatToString,
	CreateObject,
	Damage,
	GroupDamage,
	CreateMover,
	GroupCreateMover,
	AwardSpell,
	GroupAwardSpell,
	Enchant,
	GroupEnchant,
	GetHost,
	Object,
	GetObjectX,
	GetWaypointX,
	GetObjectY,
	GetWaypointY,
	GetObjectZ,
	GetDirection,
	MoveObject,
	MoveWaypoint,
	Raise,
	LookWithAngle,
	PushObjectTo,
	PushObject,
	GetLastItem,
	GetPreviousItem,
	HasItem,
	GetHolder,
	Pickup,
	Drop,
	HasClass,
	Unused50,
	HasEnchant,
	EnchantOff,
	CurrentHealth,
	MaxHealth,
	RestoreHealth,
	Distance,
	IsVisibleTo,
	Unused58,
	Unused59,
	Unused5a,
	Unused5b,
	Unused5c,
	Unused5d,
	Unused5e,
	GetCharacterData,
	LookAtObject,
	Walk,
	GroupWalk,
	CancelTimer,
	Effect,
	SetOwner,
	GroupSetOwner,
	SetOwners,
	GroupSetOwners,
	IsOwnedBy,
	GroupIsOwnedBy,
	IsOwnedByAny,
	GroupIsOwnedByAny,
	ClearOwner,
	Waypoint,
	WaypointGroup,
	ObjectGroup,
	WallGroup,
	ChatTimerSeconds,
	ChatTimer,
	Unused74,
	DestroyChat,
	DestroyEveryChat,
	SetQuestStatus,
	SetQuestStatusFloat,
	GetQuestStatus,
	GetQuestStatusFloat,
	ResetQuestStatus,
	IsTrigger,
	IsCaller,
	SetDialog,
	CancelDialog,
	StoryPic,
	TellStory,
	StartDialog,
	CastSpellObjectObject,
	CastSpellObjectLocation,
	CastSpellLocationObject,
	CastSpellLocationLocation,
	UnBlind,
	Blind,
	WideScreen,
	GetElevatorStatus,
	CreatureGuard,
	CreatureGroupGuard,
	CreatureHunt,
	CreatureGroupHunt,
	CreatureIdle,
	CreatureGroupIdle,
	CreatureFollow,
	CreatureGroupFollow,
	AggressionLevel,
	GroupAggressionLevel,
	HitLocation,
	GroupHitLocation,
	HitFarLocation,
	GroupHitFarLocation,
	SetRoamFlag,
	GroupSetRoamFlag,
	Attack,
	GroupAttack,
	JournalEntry,
	JournalDelete,
	JournalEdit,
	RetreatLevel,
	GroupRetreatLevel,
	ResumeLevel,
	GroupResumeLevel,
	RunAway,
	GroupRunAway,
	PauseObject,
	GroupPauseObject,
	IsAttackedBy,
	GetGold,
	ChangeGold,
	GetAnswer,
	GiveXp,
	HasSubclass,
	AutoSave,
	Music,
	StartupScreen,
	IsTalking,
	GetTrigger,
	GetCaller,
	MakeFriendly,
	MakeEnemy,
	BecomePet,
	BecomeEnemy,
	Unknownb8,
	Unknownb9,
	SetHalberd,
	DeathScreen,
	Frozen,
	NoWallSound,
	SetCallback,
	DeleteObjectTimer,
	TrapSpells,
	IsTrading,
	ClearMessages,
	SetShopkeeperText,
	Unknownc4,
	IsSummoned,
	ZombieStayDown,
	ZombieGroupStayDown,
	RaiseZombie,
	RaiseZombieGroup,
	MusicPushEvent,
	MusicPopEvent,
	MusicEvent,
	IsGameBall,
	IsCrown,
	EndGame,
	ImmediateBlind,
	ChangeScore,
	GetScore,
}
