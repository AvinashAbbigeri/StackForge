package main

import (
	"fmt"
	"os"

	"StackForge/engine"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: stackforge <module> [module...]")
		return
	}

	modules, err := engine.LoadModules("modules")
	if err != nil {
		panic(err)
	}

	resolved, err := engine.Resolve(os.Args[1:], modules)
	if err != nil {
		panic(err)
	}

	base, err := os.ReadFile("templates/base.sh")
	if err != nil {
		panic(err)
	}

	script, err := engine.Assemble(resolved, string(base))
	if err != nil {
		panic(err)
	}

	os.WriteFile("setup.sh", []byte(script), 0755)
	fmt.Println("Generated setup.sh")
}
