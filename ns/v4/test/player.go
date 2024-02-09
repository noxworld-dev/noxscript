package nstest

import (
	"fmt"
	"log/slog"
	"slices"
	"sync/atomic"

	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/types"
	"golang.org/x/exp/maps"

	"github.com/noxworld-dev/noxscript/ns/v4"
)

type Players struct {
	last uint32
	ByID map[uint32]*Player
	Host *Player
}

func (r *Runtime) HostPlayer() ns.Player {
	return r.PlayersData.Host.asPlayer()
}

func (r *Runtime) Players() []ns.Player {
	ids := maps.Keys(r.PlayersData.ByID)
	slices.Sort(ids)
	var out []ns.Player
	for _, id := range ids {
		pl := r.PlayersData.ByID[id]
		out = append(out, pl)
	}
	return out
}

func (s *Players) New(r *Runtime, name string) *Player {
	id := atomic.AddUint32(&s.last, 1)
	p := &Player{r: r, id: id, name: name}
	if s.ByID == nil {
		s.ByID = make(map[uint32]*Player)
	}
	s.ByID[p.id] = p
	return p
}

func (r *Runtime) NewPlayer(name string) *Player {
	return r.PlayersData.New(r, name)
}

func (r *Runtime) OnPlayerJoin(fnc ns.PlayerJoinFunc) {
	//TODO implement me
}

func (r *Runtime) OnPlayerLeave(fnc ns.PlayerLeaveFunc) {
	//TODO implement me
}

func (r *Runtime) OnPlayerDeath(fnc ns.PlayerDeathFunc) {
	//TODO implement me
}

type Player struct {
	r            *Runtime
	id           uint32
	name         string
	Storage      Storage
	UnitPtr      *Object
	ClassVal     player.Class
	ScoreVal     int
	CursorPosVec types.Pointf
}

func (p *Player) asPlayer() ns.Player {
	if p == nil {
		return nil
	}
	return p
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Unit() ns.Obj {
	return p.UnitPtr.asObj()
}

func (p *Player) GetScore() int {
	return p.ScoreVal
}

func (p *Player) ChangeScore(score int) {
	p.ScoreVal = score
}

func (p *Player) Print(message ns.StringID) {
	slog.Info("print", "player", p.id, "id", message)
}

func (p *Player) PrintStr(message string) {
	slog.Info(message, "player", p.id)
}

func (p *Player) Blind(blind bool) {
	//TODO implement me
}

func (p *Player) HasTeam(t ns.Team) bool {
	//TODO implement me
	panic("implement me")
}

func (p *Player) Team() ns.Team {
	//TODO implement me
	panic("implement me")
}

func (p *Player) CursorPos() types.Pointf {
	return p.CursorPosVec
}

func (p *Player) Store(typ ns.StorageType) ns.Storage {
	switch typ := typ.(type) {
	default:
		panic(fmt.Errorf("unexpected storage type: %T", typ))
	case nil:
		return p.Storage.Session("default")
	case ns.Session:
		return p.Storage.Session(typ.Name)
	case ns.Persistent:
		return p.Storage.Persistent(typ.Name)
	}
}
