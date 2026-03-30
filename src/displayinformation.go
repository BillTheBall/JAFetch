package src 

import (
	"fmt" //Used to print everything
	"JAFetch/src/config" //Used to get the information, contains all the modules
)

func GetInformation() {
maxWidth := 20 //Just sets Max width to a value, only matters at large numbers since gets changes
for _, line := range config.Logo { //Checks every line in config
	if len(line) > maxWidth { //If line too large for max width change max width
		maxWidth = len(line)
	}
}

for i := 0; i < len(GetConfig()); i++ { //Checks everything in getconfig
	var left, right string

	if i < len(config.Logo) { //Gets all the information for the logo to be printed
		left = config.Logo[i]
	}
	if i < len(GetConfig()) { //Gets all the information for the actual important info to be printed
		right = GetConfig()[i]
	}

	fmt.Printf("%-*s %s\n", maxWidth+2, left, right) //Prints it at the right size
}	
}
