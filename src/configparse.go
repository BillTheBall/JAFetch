package src

import (
	"JAFetch/src/config"
	"encoding/json" //Decodes the file
	"io" //Used to read the file
	"os" //Used to open the file
	"fmt" //Used to communicate errors (replacement needed)
)

// Contains all the information that can be in the Config
type Config struct {
	Bios          bool `json:"bios"`
	Cpu           bool `json:"cpu"`
	Disk          bool `json:"disk"`
	Gpu           bool `json:"gpu"`
	Init          bool `json:"init"`
	Ip						bool `json:"ip"`
	Memory        bool `json:"memory"`
	Motherboard   bool `json:"motherboard"`
	Os            bool `json:"os"`
	Kernel        bool `json:"kernel"`
	Resolution    bool `json:"resolution"`
	Terminal      bool `json:"terminal"`
	Uptime        bool `json:"uptime"`
	WindowManager bool `json:"windowmanager"`
}

//Gets the config file from src/config/config,json
func GetConfig() []string {
	jsonFile, err := os.Open("/.config/JAFetch/config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close() //Doesn't instantly close it allowing it be read

	byteValue, err := io.ReadAll(jsonFile) //Reads the JSON file
	if err != nil {
		fmt.Println("Error reading file:", err)

	}

	var conf Config
	err = json.Unmarshal(byteValue, &conf) //Decodes the JSON file
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)

	}

info := []string{} //Makes a string used to print the information

//The config.XXXXXX portion says what it is adding
info = append(info, GetHost())
info = append(info, config.HostEnding)

if conf.Bios {
	info = append(info, config.BiosMessage + GetBios())
}
if conf.Cpu {
	info = append(info, config.CpuMessage + GetCpu())
	info = append(info, config.CpuThreadMessage + GetCores())
}
if conf.Disk {
	info = append(info, config.DiskMessage + GetDisk()+GetDiskPercentage())
}
if conf.Gpu {
	info = append(info, config.GpuMessage + GetGpu())
}
if conf.Init {
	info = append(info, config.InitMessage + GetInit())
}
if conf.Memory {
	info = append(info, config.DiskMessage + GetMemory()+GetMemoryPercentage())
}
if conf.Motherboard {
	info = append(info, config.MotherboardMessage + GetMotherboard())
}
if conf.Os {
	info = append(info, config.OsMessage + GetOS())
}
if conf.Kernel {
	info = append(info, config.KernelMessage + GetKernel())
}
if conf.Resolution {
	info = append(info, config.ResolutionMessage+GetResolution())
}
if conf.Terminal {
	info = append(info, config.TerminalMessage + GetTerminal())
}
if conf.Uptime {
	info = append(info, config.UptimeMessage+GetUptime())
}
if conf.WindowManager {
	info = append(info, config.WindowManagerMessage + GetWindowManager())
}
if conf.Ip {
	info = append(info, config.IPMessage + GetIPAddress())

}

return(info)
}
