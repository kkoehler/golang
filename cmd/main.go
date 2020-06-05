/*Package main shows a simple command-line client implementation in Go.

You can find the corresponding article at https://source-fellows.com/go-commandline/

Run the command with:
	go run main.go <options>

The following options are available:
	-flagOne int
		help message for flagOne (default 1234)
	-flagThree string
		Hier sollte ihr Text stehen. (Required)
	-flagTwo int
		help message for flagTwo (default 4321)
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flagOne := flag.Int("flagOne", 1234, "help message for flagOne")

	var flagTwo int
	flag.IntVar(&flagTwo, "flagTwo", 4321, "help message for flagTwo")

	flagThree := flag.String("flagThree", "", "Hier sollte ihr Text stehen. (Required)")
	flag.Parse()

	if *flagThree == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("FlagOne  : %d\n", *flagOne)
	fmt.Printf("FlagTwo  : %d\n", flagTwo)
	fmt.Printf("flagThree: %s\n", *flagThree)

}
