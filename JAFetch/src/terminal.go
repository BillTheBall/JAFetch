// PROBLEM
// PROBLEM
// PROBLEM
package src

import (
	"os" //Looks for the Shell your using
)

func GetTerminal() string {
	if termProgram, confirm := os.LookupEnv("TERM"); confirm { //Checks TERM environment to find the shell being used
		return (termProgram) //DO NOT USE, it's far too general and will only provide very basic information, not advised!
	} else {
		return ("Error getting term program")
	}
}
