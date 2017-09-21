package main

import (
	"fmt"
)

type attacker struct {
	attackpower int
	dmgbonus    int
}

type sword struct {
	attacker
	twohanded bool
}

type gun struct {
	attacker
	bulletsremaining int
}

type power struct {
	attack  int
	defense int
}

type location struct {
	x float32
	y float32
	z float32
}

type nonPlayerCharacter struct {
	name  string
	speed int
	hp    int
	power power
	loc   location
}

func (s sword) Wield() bool {
	fmt.Println("You`ve wielded a sword!")
	return true
}

func (g gun) Wield() bool {
	fmt.Println("You`ve wielded a gun!")
	return true
}

func main() {
	fmt.Println("Stucts...")

	demon := nonPlayerCharacter{
		name:  "Alfred",
		speed: 21,
		hp:    100,
		power: power{attack: 75, defense: 50},
		loc:   location{x: 1075.123, y: 521.123, z: 211.231},
	}

	fmt.Println(demon)

	anotherDemon := nonPlayerCharacter{name: "Beelzebub", speed: 30, hp: 5000, power: power{attack: 10, defense: 10}, loc: location{x: 32.03, y: 72.45, z: 65.231}}

	fmt.Println(anotherDemon)

	sword1 := sword{attacker: attacker{attackpower: 1, dmgbonus: 5}, twohanded: true}
	gun1 := gun{attacker: attacker{attackpower: 10, dmgbonus: 20}, bulletsremaining: 11}
	fmt.Printf("Weapons: sword: %v, gun: %v\n", sword1, gun1)
	sword1.Wield()
	gun1.Wield()

}
