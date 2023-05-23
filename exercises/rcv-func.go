//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	health, maxHealth float64
	energy, maxEnergy float64
	name              string
}

func (p *Player) changeHealth(amt float64) {
	if p.health+amt >= p.maxHealth {
		p.health = 100
		return
	}
	p.health += amt
}
func (p *Player) changeEnergy(amt float64) {
	if p.energy+amt >= p.maxEnergy {
		p.energy = 100
		return
	}
	p.energy += amt
}
func main() {
	p := Player{
		health:    50,
		maxHealth: 100,
		energy:    60,
		maxEnergy: 200,
		name:      "BEAN",
	}
	p.changeEnergy(100)
	fmt.Println(p)
	p.changeHealth(55)
	fmt.Println(p)
}
