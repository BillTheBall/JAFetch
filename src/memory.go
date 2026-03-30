package src

import (
	"JAFetch/src/libs"                  //Used to round decimals
	"github.com/shirou/gopsutil/v4/mem" //Used to access information about the memory
	"strconv"                           //Converts strings
)

func GetMemory() string {
	memory, err := mem.VirtualMemory() //Gets memory
	if err != nil {
		return ("Error in obtaining memory")
	}

	floatTotal := float64(memory.Total) //Finds total size in bytes
	totalRAM := libs.RoundToDecimals(floatTotal/1024/1024/1024, 2) //Converts the bytes into Gigabytes

	floatUsed := float64(memory.Used) //Gets total size used
	usedRAM := libs.RoundToDecimals(floatUsed/1024/1024/1024-0.01, 2) //Converts the total size used into Gigabytes

	totalRAMconv := strconv.FormatFloat(totalRAM, 'f', -1, 64) //Converts into a string
	usedRAMconv := strconv.FormatFloat(usedRAM, 'f', -1, 64)

	memoryInformation := (usedRAMconv + " GiB / " + totalRAMconv + " GiB")
	return (memoryInformation)
}

func GetMemoryPercentage() string {
	memory, err := mem.VirtualMemory() //Get memory again
	if err != nil {
		return ("Error in obtain memory percentage")
	}
	usedpercentRAM := libs.RoundToDecimals(memory.UsedPercent, 0) //Rounds the percentage
	usedpercentRAMconv := strconv.FormatFloat(usedpercentRAM, 'f', -1, 64) //Converts to string
	ramPercentage := "(" + usedpercentRAMconv + "%)"
	return (ramPercentage)
}
