package strategy

import "fmt"

// Player is struct
type Player struct {
	Name                           string
	strategy                       Strategy
	wincount, losecount, gamecount int
}

// NewPlayer func for initializing Player
func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{
		Name:     name,
		strategy: strategy,
	}
}

// NextHand func can handle result of preHand
func (p *Player) NextHand() *Hand {
	return p.strategy.NextHand()
}

// Win func can judge result as game win
func (p *Player) Win() {
	p.strategy.study(true)
	p.wincount++
	p.gamecount++
}

// Lose func can judge result as game lose
func (p *Player) Lose() {
	p.strategy.study(false)
	p.losecount++
	p.gamecount++
}

// Even func can judge result as game even
func (p *Player) Even() {
	p.gamecount++
}

// ToString func can display results of all games
func (p *Player) ToString() string {
	str := fmt.Sprintf("[%s: %dgames,  %dwin,  %dlose]", p.Name,
		p.gamecount,
		p.wincount,
		p.losecount)
	return str
}
