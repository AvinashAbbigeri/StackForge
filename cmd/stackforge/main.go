package main

import (
	"fmt"
	"os"

	"StackForge/engine"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "init" {
		fmt.Println("Usage: stackforge init <project-name> <module> [module...]")
		return
	}

	projectName := os.Args[2]
	modules := os.Args[3:]

	// Create project directory
	if err := os.Mkdir(projectName, 0755); err != nil {
		panic(err)
	}

	modulesMap, err := engine.LoadModules("modules")
	if err != nil {
		panic(err)
	}

	resolved, err := engine.Resolve(modules, modulesMap)
	if err != nil {
		panic(err)
	}

	osinfo := engine.DetectOS()

	base, err := os.ReadFile("templates/base.sh")
	if err != nil {
		panic(err)
	}

	script, err := engine.Assemble(resolved, string(base), osinfo)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(projectName+"/setup.sh", []byte(script), 0755)
	if err != nil {
		panic(err)
	}

	fmt.Println("Project created:", projectName)
	fmt.Println("Next:")
	fmt.Println("  cd", projectName)
	fmt.Println("  bash setup.sh")
}
