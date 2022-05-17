package models

import (
	"arena/utils"
)

type damage struct {
	Maximum int `json:"maximum"`
	Minimum int `json:"minimum"`
}

type criticalHit struct {
	ChancePercentage int     `json:"chancePercentage"`
	Multiplier       float32 `json:"multiplier"`
}

type attack struct {
	Damage      damage      `json:"damage"`
	CriticalHit criticalHit `json:"criticalHit"`
}

type resistance struct {
	HeadReductionPercentage      int `json:"headReductionPercentage"`
	UpperBodyReductionPercentage int `json:"upperBodyReductionPercentage"`
	LowerBodyReductionPercentage int `json:"lowerBodyReductionPercentage"`
}

type defense struct {
	EvasionChancePercentage int        `json:"evasionChancePercentage"`
	ParadeChancePercentage  int        `json:"paradeChancePercentage"`
	Resistance              resistance `json:"resistance"`
}

type Gladiator struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	HitPoint int     `json:"hitPoint"`
	Attack   attack  `json:"attack"`
	Defense  defense `json:"defense"`
}

// Evasion randomizes the evasion
// It returns whether the evasion is successful or not.
func (g *Gladiator) Evasion() bool {
	var rand = new(utils.Rand)
	return rand.XPercentChance(g.Defense.EvasionChancePercentage)
}

// RandDamage randomizes the damage.
// It returns the generated damage value.
func (g *Gladiator) RandDamage() int {
	var rand = new(utils.Rand)
	return rand.Range(g.Attack.Damage.Minimum, g.Attack.Damage.Maximum)
}

// RandDamage randomizes the critical hit.
// It returns whether the critical hit is successful or not and the altered damage value.
func (g *Gladiator) CriticalHit(dmg int) (bool, int) {
	var rand = new(utils.Rand)
	if rand.XPercentChance(g.Attack.CriticalHit.ChancePercentage) {
		return true, int(float32(dmg) * g.Attack.CriticalHit.Multiplier)
	}
	return false, dmg
}

// BodyPartHit randomizes the body part to hit.
// It returns the name of the body part and the altered damage value.
func (g *Gladiator) BodyPartHit(dmg int) (string, int) {
	var rand = new(utils.Rand)
	var name string
	var reduc int
	switch rand.Range(1, 3) {
	case 1:
		name = "head"
		reduc = g.Defense.Resistance.HeadReductionPercentage
		dmg = int(float32(dmg) * 1.2) // mutiplied if head is hit
	case 2:
		name = "upper body"
		reduc = g.Defense.Resistance.UpperBodyReductionPercentage
	case 3:
		name = "lower body"
		reduc = g.Defense.Resistance.LowerBodyReductionPercentage
	}
	return name, dmg - (dmg * reduc / 100)
}

// Parade randomizes the parade.
// It returns whether the parade is successful or not and the altered damage value.
func (g *Gladiator) Parade(dmg int) (bool, int) {
	var rand = new(utils.Rand)
	if rand.XPercentChance(g.Defense.ParadeChancePercentage) {
		return true, dmg / 3
	}
	return false, dmg
}

// TakeDamage reduces hit point of the gladiator from damage value.
func (g *Gladiator) TakeDamage(dmg int) {
	if g.HitPoint < dmg {
		g.HitPoint = 0
	} else {
		g.HitPoint -= dmg
	}
}

// IsDead returns whether the gladiator is dead or not.
func (g *Gladiator) IsDead() bool {
	return g.HitPoint <= 0
}
