package src

import (
	"github.com/jaypipes/ghw" //Used to figure out what Motherboard your using
)

func GetMotherboard() string {
	motherboard, err := ghw.Baseboard(ghw.WithDisableWarnings()) //Checks for motherboards
	if err != nil {
		return ("Error getting motherboard")
	}
	return (motherboard.Product) //Shows just the product name
}
