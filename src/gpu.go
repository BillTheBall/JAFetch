package src

import (
	"github.com/jaypipes/ghw" //Used to figure out what GPU your using, also really freaking complex
	"os/exec"
	"strings"
)

func GetGpu() string {
	var gpuUsed string
 cmd := exec.Command("nvidia-smi", "--query-gpu=name,memory.used,memory.total", "--format=csv") 	
output, err := cmd.Output()
	if err != nil {
					gpu, err := ghw.GPU()
			if err != nil {
				return("Error getting GPU info")
			}
			gpuin := 0
			if len(gpu.GraphicsCards) > 1 {
				for _, c := range gpu.GraphicsCards {
					gpuUsed = c.DeviceInfo.Product.Name
					gpuin++
				}
			} else {
				gpuUsed = gpu.GraphicsCards[0].DeviceInfo.Product.Name
			}
			return(gpuUsed)
	} else {
		return(strings.TrimSpace(string(output)[44:]))
	}
}
