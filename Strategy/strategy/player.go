package strategy

import "fmt"

// Player is struct
type Player struct {
	Name                           string
	PlayerStrategy                 Strategy
	wincount, losecount, gamecount int
}

// NextHand func can handle result of preHand
func (p *Player) NextHand() *Hand {
	return p.PlayerStrategy.NextHand()
}

// Win func can judge result as game win
func (p *Player) Win() {
	p.PlayerStrategy.study(true)
	p.wincount++
	p.gamecount++
}

// Lose func can judge result as game lose
func (p *Player) Lose() {
	p.PlayerStrategy.study(false)
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
