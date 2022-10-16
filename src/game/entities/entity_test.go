package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntity_TakeDamage(t *testing.T) {
	for name, tc := range map[string]struct {
		damage         int
		expectedHealth int
		expectedIsDead bool
	}{
		"Take 10 damage, should not be dead": {
			damage:         10,
			expectedHealth: 10,
			expectedIsDead: false,
		},
		"Take 20 damage, should be dead": {
			damage:         20,
			expectedHealth: 0,
			expectedIsDead: true,
		},
		"Take 30 damage, should be dead without negative health": {
			damage:         30,
			expectedHealth: 0,
			expectedIsDead: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			entity := Entity{Health: 20}
			isDead := entity.TakeDamage(tc.damage)
			assert.Equal(t, tc.expectedHealth, entity.Health)
			assert.Equal(t, tc.expectedIsDead, isDead)
		})
	}
}
