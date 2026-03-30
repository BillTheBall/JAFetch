package src

import (
	"github.com/shirou/gopsutil/v4/cpu" //Used to figure out what CPU your using
			"os/exec" //Executes the commands
	"strings" //Clarifies the commands
)

func GetCpu() string {
	processor, err := cpu.Info() //Creates the variable that access information about CPU, so cpu,Info,x() isn't required
	if err != nil {
		return ("Error in CPU gathering")
	}
	if len(processor) > 0 {
		cpuInformation := ("[" + processor[0].ModelName + "]") //Gets the CPU information starting from the very start of the cpu.Information.Modelname Slice
		return (cpuInformation)
	} else {
		return ("Error in CPU parsing")
	}
}

func GetCores() string {
	// Using nproc to get thread count
	cmd := exec.Command("nproc")
	output, err := cmd.Output()
	if err != nil {
		return ("Error executing xrandr")
	}
	return(strings.TrimSpace(string(output)))
}
