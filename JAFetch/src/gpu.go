package src

import (
	"github.com/jaypipes/ghw" //Used to figure out what GPU your using, also really freaking complex
)

func GetGpu() string {
	var cardinfo string   //Used as a generic "" to store the gpu information
	gpu, err := ghw.GPU() //Creates the gpu viewer
	if err != nil {
		return ("Error in obtain GPU")
	}

	for _, card := range gpu.GraphicsCards { //Checks all information about the gpu
		cardinfo = (card.DeviceInfo.Product.Name[6:] + " ") //Grabs just the part of the slice containing product name skipping 6 useless letters
	}
	if cardinfo != "" {
		return (cardinfo)
	} else {
		return ("Error in gathering GPU card")
	}
}
