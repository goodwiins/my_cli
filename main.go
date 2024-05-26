package main

import (
	"flag"
	"fmt"
	"os"
)

func hellocom(name string, age int) {
	fmt.Printf("Hello, %s! welcome to World! my age: %d ", name, age)

}
func usage() {
	fmt.Println("Usage: hello -name [name] -age [age]")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "-h", "--help":
		usage()
	case "hello":
		// If the command is "hello", call the hello function
		helloCmd := flag.NewFlagSet("hello", flag.ExitOnError)
		name := helloCmd.String("name", "World", "a name to say hello to")
		age := helloCmd.Int("age", 0, "your age")
		if err := helloCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println("Error parsing flags for hello command")
			os.Exit(1)
		}
		hellocom(*name, *age)
	default:
		// If the command is not recognized, display the usage
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		usage()
		os.Exit(1)
	}

}
