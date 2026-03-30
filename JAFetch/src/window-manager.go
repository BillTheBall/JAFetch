package src

import (
	"os/exec" //Executes the commands
	"strings" //Clarifies the commands
)

func GetWindowManager() string {
	// Using xprop to get window manager name
	cmd := exec.Command("xprop", "-root", "_NET_WM_NAME") //Checks the xprop under root for Window Manager name
	output, err := cmd.Output()
	if err != nil {
		return ("Error executing xprop")
	}
	return (strings.TrimSpace(string(output)[28:])) //Cleans it up to start only at the 28th letter to give only important information
}
