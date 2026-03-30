package src

import (
	"github.com/jaypipes/ghw" //Used to figure out what BIOS your using
)

func GetBios() string {
	bios, err := ghw.BIOS() //Obtains BIOS
	if err != nil {
		return ("Error getting bios")
	}
	return (bios.Vendor) //Returns the BIOS information
}
