package ns

import ns4 "github.com/noxworld-dev/noxscript/ns/v4"

type WallID int

func (id WallID) ScriptID() int {
	return int(id)
}

func (id WallID) WallScriptID() int {
	return int(id)
}

type WallGroupID int

func (id WallGroupID) ScriptID() int {
	return int(id)
}

func (id WallGroupID) WallGroupScriptID() int {
	return int(id)
}

// NoWallSound sets no wall sound flag.
func NoWallSound(noWallSound bool) {
	ns4.NoWallSound(noWallSound)
}

// Wall gets a pointer to a wall by its wall coordinates.
func Wall(x, y int) WallID {
	return asWall(ns4.Wall(x, y))
}

// WallOpen opens a wall.
func WallOpen(wall WallID) {
	ns4.AsWall(wall).Enable(false)
}

// WallClose closes a wall.
func WallClose(wall WallID) {
	ns4.AsWall(wall).Enable(true)
}

// WallToggle toggles a wall.
//
// This will toggle wall between opened and closed.
func WallToggle(wall WallID) {
	ns4.AsWall(wall).Toggle()
}

// WallBreak breaks a wall.
//
// This will break a wall. The wall must be breakable.
func WallBreak(wall WallID) {
	ns4.AsWall(wall).Destroy()
}

// WallGroup looks up wall group by name.
func WallGroup(name string) WallGroupID {
	return asWallGroup(ns4.WallGroup(name))
}

// WallGroupOpen opens walls in a group.
func WallGroupOpen(g WallGroupID) {
	ns4.AsWallGroup(g).Enable(false)
}

// WallGroupClose closes walls in a group.
func WallGroupClose(g WallGroupID) {
	ns4.AsWallGroup(g).Enable(true)
}

// WallGroupToggle toggles walls in a group.
//
// This will toggle a wall group between opened and closed.
func WallGroupToggle(g WallGroupID) {
	ns4.AsWallGroup(g).Toggle()
}

// WallGroupBreak breaks walls in a group.
//
// This will break walls in a group. The walls must be breakable.
func WallGroupBreak(g WallGroupID) {
	ns4.AsWallGroup(g).Destroy()
}
