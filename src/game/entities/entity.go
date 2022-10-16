package entities

type Entity struct {
	Health int
}

func (entity *Entity) TakeDamage(damage int) bool {
	entity.Health -= damage
	
	if entity.Health <= 0 {
		entity.Health = 0
		return true
	}

	return false
}
