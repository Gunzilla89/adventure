package object

type IPlayer interface {
	setName(name string)
	getName() string
	setCoin(coin int)
	getCoin() int
}
type Player struct {
	Name string
	Coin int
}

//set and get player Name
func (p *Player) setName(name string) {
	p.Name = name
}

func (p *Player) getName() string {
	return p.Name
}

//set and get player Coins
func (p *Player) setCoin(coin int) {
	p.Coin = coin
}

func (p *Player) getCoin() int {
	return p.Coin
}
