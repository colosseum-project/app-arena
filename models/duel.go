package models

import (
	"arena/utils"
	"errors"
	"fmt"
)

type Duel struct {
	atk      *Gladiator // attacker
	def      *Gladiator // defender
	turn     int
	resolved bool
}

func (d *Duel) shouldChangeAttacker() bool {
	var rand = new(utils.Rand)
	if d.turn == 1 {
		return rand.OneinXChance(2)
	}
	return rand.OneinXChance(4)
}

func (d *Duel) swapGladiators() {
	d.atk, d.def = d.def, d.atk
}

func (d Duel) Resolve(gs []Gladiator) (DuelResult, error) {
	var dr DuelResult

	if len(gs) != 2 {
		return dr, errors.New("number of gladiators must be stricly equal to 2")
	}
	d.atk = &gs[0]
	d.def = &gs[1]

	for d.turn, d.resolved = 1, false; !d.resolved && d.turn <= 100; d.turn++ {
		if d.shouldChangeAttacker() {
			d.swapGladiators()
		}

		var bodyPart string = "none"
		var damage int = 0
		var evasion, parade, criticalHit bool

		evasion = d.def.Evasion()
		if !evasion {
			damage = d.atk.RandDamage()
			bodyPart, damage = d.def.BodyPartHit(damage)
			parade, damage = d.def.Parade(damage)
			if !parade {
				criticalHit, damage = d.atk.CriticalHit(damage)
			}
			d.def.TakeDamage(damage)
		}

		dt := duelTurn{
			d.atk.Name,     // AttackerName
			d.def.Name,     // DefenderName
			d.def.HitPoint, // DefenderRemainingHP
			false,          // DefenderDead
			evasion,        // Evasion
			parade,         // Parade
			criticalHit,    // CriticalHit
			bodyPart,       // BodyPart
			damage,         // Damage
		}

		if d.def.IsDead() {
			d.resolved = true
			dt.DefenderDead = true
			dr.WinnerId = d.def.Id
		}

		dr.CombatLogs = append(dr.CombatLogs, dt.toString())
	}

	return dr, nil
}

type duelTurn struct {
	AttackerName        string
	DefenderName        string
	DefenderRemainingHP int
	DefenderDead        bool
	Evasion             bool
	Parade              bool
	CriticalHit         bool
	BodyPart            string
	Damage              int
}

func (dt *duelTurn) toString() string {
	var str string
	if dt.Evasion {
		str = fmt.Sprintf(
			"%s missed %s.",
			dt.AttackerName, dt.DefenderName)
	} else {
		if dt.Parade {
			str = fmt.Sprintf(
				"%s parried %s and protected his %s.",
				dt.DefenderName, dt.AttackerName, dt.BodyPart)
		} else {
			str = fmt.Sprintf(
				"%s hit %s's %s",
				dt.AttackerName, dt.DefenderName, dt.BodyPart)
			if dt.CriticalHit {
				str += " critically"
			}
			str += "."
		}
		str += fmt.Sprintf(
			" %s dealt %d damage and reduced %s's HP to %d.",
			dt.AttackerName, dt.Damage, dt.DefenderName, dt.DefenderRemainingHP)
		if dt.DefenderDead {
			str += fmt.Sprintf(
				" %s died.",
				dt.DefenderName)
		}
	}
	return str
}

type DuelResult struct {
	WinnerId   int      `json:"winnerId"`
	CombatLogs []string `json:"combatLogs"`
}
