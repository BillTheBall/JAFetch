package src

import (
	"github.com/shirou/gopsutil/v4/host" //Used to access information about uptime
	"os/exec" //Executes commands
	"os/user" //Checks users
	"strings" //Clarifies command output
)

// Works!
func GetOS() string {
	operatingsystem, err := host.Info() //gets OS info
	if err != nil {
		return ("Error in finding host.Info")
	}
	return (operatingsystem.Platform + " " + operatingsystem.OS + " " + operatingsystem.KernelArch) //Finds the name of disto (void, arch). then if it's linux or Windows, then it will find if it's Arm, or X86-64
}

// Fails
func GetHost() string {
	operatingsystem, err := host.Info()
	if err != nil {
		return ("Error in finding host.Info for computer name")
	}

	currentUser, err := user.Current()
	if err != nil {
		return ("Error getting current user")
	}

	return (currentUser.Name + " " + operatingsystem.Hostname) //Finds the name of the computer, then finds the name of the user
}

func GetKernel() string {
	cmd := exec.Command("uname", "-r") //Uses uname -r to find Kernel version i,e 4.20, 6.12, 2.6
	output, err := cmd.Output()
	if err != nil {
		return ("Error executing uname -r")
	}
	return (strings.TrimSpace(string(output))) //Clarifies and outputs it
}
