package some_package

import(
	"slices"
	"errors"
)

type Competition struct{
	Name string
	Players []string
	Prizes []string
}
// methods to be mocked
func (c Competition) HasPlayer(player string)bool{
	return slices.Contains(c.Players, player)
}
func (c Competition) GetPlayerPosition(player string)int{
	return slices.Index(c.Players, player)
}

// method that calls other methods
func (c Competition) PrizeAwarded(player string) (prize string, err error){
	if( len(c.Prizes) == 0 ){ return "", errors.New("don't have prizes") }
	if( !c.HasPlayer(player) ){ return "", errors.New("player is not participating of this competition") }
	position := c.GetPlayerPosition(player)

	if( position >= len(c.Prizes) ){
		return "", errors.New("player doesn't win any prize")
	}

	return c.Prizes[position], nil
}