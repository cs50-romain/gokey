package colors

import (
	"fmt"
)

const (
	Reset   = "\033[0m"
	Bold 	= "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"	
)

// Custom
const (
	FrostWhite  = "\033[38;2;252;252;252m"  // #fcfcfc
	ElectricPink= "\033[38;2;236;11;228m"   // #ec0be4
	RoyalPurple = "\033[38;2;163;85;185m"   // #a355b9
	TealGreen   = "\033[38;2;90;159;142m"   // #5a9f8e
	LimeGreen   = "\033[38;2;17;233;100m"   // #11e964
)

func Print(input string, color string) {
	fmt.Printf(color + input + "\n" + Reset)
}

func PrintBold(input string, color string) {
	fmt.Printf(Bold + color + input + Reset)
}
