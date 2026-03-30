package src

import (
	"JAFetch/src/libs"                   //Used for rounding decimals
	"github.com/shirou/gopsutil/v4/disk" //Gathers disk information
	"strconv"                            //To convert strings
)

func GetDisk() string {
	disk, err := disk.Usage("/") //Checks all information in the disk starting from / known as home directory
	if err != nil {
		return ("Error in disk information")
	}

	floatDISKUsed := float64(disk.Used)                       //Converts used disk into float
	floatDISKUsedactual := floatDISKUsed / 1024 / 1024 / 1024 //Divides it into the GiB from bytes
	usedDISK := libs.RoundToDecimals(floatDISKUsedactual, 2)  //Rounds it to only the hundruths

	floatDISKTotal := float64(disk.Total)
	floatDISKTotalactual := floatDISKTotal / 1024 / 1024 / 1024
	totalDISK := libs.RoundToDecimals(floatDISKTotalactual, 2)

	usedDISKconv := strconv.FormatFloat(usedDISK, 'f', -1, 64) //Converts the float back into a string
	totalDISKconv := strconv.FormatFloat(totalDISK, 'f', -1, 64)

	diskInformation := (usedDISKconv + "GiB / " + totalDISKconv + "GiB")
	if diskInformation != "" {
		return (diskInformation)
	} else {
		return ("Error in disk printing")
	}
}

func GetDiskPercentage() string {
	disk, err := disk.Usage("/") //Checks all information in the disk starting from / known as home directory
	if err != nil {
		return ("Error in disk information percentage")
	}
	DISKusedpercentage := libs.RoundToDecimals(disk.UsedPercent, 0)
	DISKusedpercentageconv := strconv.FormatFloat(DISKusedpercentage, 'f', -1, 64)
	return ("(" + DISKusedpercentageconv + "%)")
}
