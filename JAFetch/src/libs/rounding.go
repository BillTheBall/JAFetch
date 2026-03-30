package libs

import (
	"math" //Used to do the Math portion of the code
)

func RoundToDecimals(num float64, decimals int) float64 {
	shift := math.Pow(10, float64(decimals)) //Puts everything to the 10th power, and then makes it decimals
	return math.Round(num*shift) / shift //Corrects and rounds it
}
