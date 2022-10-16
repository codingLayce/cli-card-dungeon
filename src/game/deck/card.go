package deck

type Card struct {
	ID           int
	Name         string
	AttackValue  int
	DefenseValue int
}

func NewCard(ID int, name string, attack, defense int) Card {
	return Card{
		ID:           ID,
		Name:         name,
		AttackValue:  attack,
		DefenseValue: defense,
	}
}
