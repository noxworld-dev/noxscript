package ns

// Team is an interface for game teams.
type Team interface {
	// Name of the team.
	Name() string
	// Players returns a list of players belonging to the team.
	Players() []Player
	// GetScore returns the current team score.
	GetScore() int
	// ChangeScore adds or removes team's score(s).
	ChangeScore(score int)
	// Color returns team color.
	Color() Color
}

// Teams return current teams.
func Teams() []Team {
	if impl == nil {
		return nil
	}
	return impl.Teams()
}
