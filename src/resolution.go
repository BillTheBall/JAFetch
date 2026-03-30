package src

import (
	"github.com/fstanis/screenresolution" //Used for screen resolution
)

func GetResolution() string {
	resolution := screenresolution.GetPrimary().String() //Gets main monitor screen resolution
	if resolution != "" {
		return (resolution)
	} else {
		return ("Error in getting screen resolution")
	}
}
