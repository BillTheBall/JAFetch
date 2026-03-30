package src

import (
	"github.com/shirou/gopsutil/v4/host" //Used to access information about uptime
	"strconv"                            //Used to convert the strings
)

func GetUptime() string {
	os, _ := host.Info() //Gets OS information
	uptimeHours := os.Uptime / 60 / 60 //Gets seconds then divides it by 60 twice
	uptimeMinutes := os.Uptime / 60 //Gets seconds and breaks it into minutes
	uptimeMinutesActual := uptimeMinutes % 60 //Gets minutes as long as they're not above 60 to prevent hour issues
	uptimeHoursConv := strconv.Itoa(int(uptimeHours)) //Changes hour into string
	uptimeMinutesConv := strconv.Itoa(int(uptimeMinutesActual)) //Changes Minutes into strings
	if uptimeHours > 0 {
		if uptimeHours == 1 {
			return (uptimeHoursConv + " Hour " + uptimeMinutesConv + " Mins")
		} else {
			return (uptimeHoursConv + " Hours " + uptimeMinutesConv + " Mins")
		}
	} else {
		return (uptimeMinutesConv + " Mins")
	}
}
