package game

type Game struct {
	UserData *UserData
}

func Create() *Game {

	return new(Game)
}
