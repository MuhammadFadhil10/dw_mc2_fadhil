package main

import "fmt"

type Profile struct {
	Name               string
	Health, Power, Exp int
}

func main() {
	// make a profile
	profile := MakeProfile("Naruto")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("===Heroes POWER UP===")

	// power up profile
	powerUp := PowerUp(profile, 3)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)

}

func MakeProfile(name string) Profile {
	var profile Profile
	if name == "Sasuke" || name == "sasuke" {
		profile.Name = "Sasuke"
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
	} else if name == "Goku" || name == "goku" {
		profile.Name = "Goku"
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
	} else if name == "Naruto" || name == "naruto" {
		profile.Name = "Naruto"
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
	}

	return profile
}

func PowerUp(p Profile, multiplier int) Profile {

	p.Health = p.Health + (p.Health * multiplier)
	p.Power = p.Power + (p.Power * multiplier)
	p.Exp = p.Exp + (p.Exp * multiplier)

	return p
}
