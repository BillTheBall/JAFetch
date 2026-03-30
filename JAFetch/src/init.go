package src

import (
	"os" //Used to look into the folders
)

func GetInit() string {
	if _, err := os.Stat("/run/systemd/system"); err == nil { //looks for a systemD folder
		return ("systemd")
	} else if _, err := os.Stat("/run/runit"); err == nil { //looks for a runit folder
		return ("runit")
	} else if _, err := os.Stat("/run/openrc"); err == nil { //looks for an openrc folder
		return ("openrc")
	} else {
		return ("Couldn't obtain")
	}
}
