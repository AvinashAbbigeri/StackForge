package main

import (
	"encoding/json"
	"fmt"
	"os"

	"StackForge/engine"
)

func printHelp() {
	fmt.Print(`StackForge – cross-platform dev environment generator

Usage:
  stackforge init <project> <modules|presets> [--dry-run]
  stackforge add <modules> [--dry-run]
  stackforge list
  stackforge presets
  stackforge version
  stackforge --help

Examples:
  stackforge init api python-api
  stackforge init web nextjs-app --dry-run
  stackforge add pytest ruff
  stackforge list
  stackforge presets

Flags:
  --dry-run   Show what will happen without making changes
`)
}

func main() {
	// Enable embedded assets
	engine.SetEmbeddedFS(embeddedFS)

	// ---------------- Parse flags ----------------
	dryRun := false
	var args []string

	for _, a := range os.Args[1:] {
		if a == "--dry-run" {
			dryRun = true
		} else {
			args = append(args, a)
		}
	}

	if len(args) == 0 {
		printHelp()
		return
	}

	cmd := args[0]

	// ---------------- HELP / VERSION ----------------
	switch cmd {
	case "--help", "-h":
		printHelp()
		return
	case "version", "--version", "-v":
		printVersion()
		return
	}

	// ---------------- LIST ----------------
	if cmd == "list" {
		mods, err := engine.LoadModules("modules")
		if err != nil {
			fmt.Println("Error loading modules:", err)
			return
		}
		for _, m := range engine.ListModules(mods) {
			fmt.Println(m)
		}
		return
	}

	// ---------------- PRESETS ----------------
	if cmd == "presets" {
		presets, err := engine.LoadPresets("presets.json")
		if err != nil {
			fmt.Println("Error loading presets:", err)
			return
		}
		for _, p := range engine.PresetNames(presets) {
			fmt.Println(p)
		}
		return
	}

	// ---------------- ADD ----------------
	if cmd == "add" {
		if len(args) < 2 {
			fmt.Println("Usage: stackforge add <module> [module...] [--dry-run]")
			return
		}

		manifest, err := engine.LoadManifest(".stackforge/manifest.json")
		if err != nil {
			fmt.Println("Not inside a StackForge project")
			return
		}

		manifest.Modules = engine.Unique(append(manifest.Modules, args[1:]...))

		modulesMap, err := engine.LoadModules("modules")
		if err != nil {
			panic(err)
		}

		resolved, err := engine.Resolve(manifest.Modules, modulesMap)
		if err != nil {
			panic(err)
		}

		base, err := engine.ReadTemplate("templates/base.sh")
		if err != nil {
			panic(err)
		}

		osinfo := engine.DetectOS()
		plan, err := engine.Assemble(resolved, string(base), osinfo)
		if err != nil {
			panic(err)
		}

		if dryRun {
			fmt.Println("StackForge dry run")
			fmt.Println("\nWill add:")
			for _, m := range args[1:] {
				fmt.Println(" ", m)
			}
			fmt.Println("\nCommands:")
			for _, c := range plan.Installs {
				fmt.Println(" ", c)
			}
			return
		}

		_ = os.WriteFile(".stackforge/setup.sh", []byte(plan.Script), 0755)
		engine.SaveManifest(".stackforge/manifest.json", manifest)

		fmt.Println("Modules added. Run:")
		fmt.Println("  bash .stackforge/setup.sh")
		return
	}

	// ---------------- INIT ----------------
	if cmd != "init" {
		fmt.Println("Unknown command:", cmd)
		fmt.Println()
		printHelp()
		return
	}

	if len(args) < 3 {
		fmt.Println("Usage: stackforge init <project> <modules|presets> [--dry-run]")
		return
	}

	projectName := args[1]
	input := args[2:]

	presets, err := engine.LoadPresets("presets.json")
	if err != nil {
		panic(err)
	}

	var modules []string
	for _, name := range input {
		if preset, ok := presets[name]; ok {
			modules = append(modules, preset...)
		} else {
			modules = append(modules, name)
		}
	}

	modules = engine.Unique(modules)

	modulesMap, err := engine.LoadModules("modules")
	if err != nil {
		panic(err)
	}

	resolved, err := engine.Resolve(modules, modulesMap)
	if err != nil {
		panic(err)
	}

	base, err := engine.ReadTemplate("templates/base.sh")
	if err != nil {
		panic(err)
	}

	osinfo := engine.DetectOS()
	plan, err := engine.Assemble(resolved, string(base), osinfo)
	if err != nil {
		panic(err)
	}

	if dryRun {
		fmt.Println("StackForge dry run")
		fmt.Println("\nWill install:")
		for _, m := range modules {
			fmt.Println(" ", m)
		}
		fmt.Println("\nCommands:")
		for _, c := range plan.Installs {
			fmt.Println(" ", c)
		}
		fmt.Println("\nFiles:")
		for _, f := range plan.Files {
			fmt.Println(" ", f)
		}
		return
	}

	_ = os.Mkdir(projectName, 0755)
	_ = os.Mkdir(projectName+"/.stackforge", 0755)

	_ = os.WriteFile(projectName+"/.stackforge/setup.sh", []byte(plan.Script), 0755)

	manifest := engine.Manifest{
		Modules: modules,
		OS:      osinfo.PackageManager,
	}

	data, _ := json.MarshalIndent(manifest, "", "  ")
	_ = os.WriteFile(projectName+"/.stackforge/manifest.json", data, 0644)

	fmt.Println("Project created:", projectName)
	fmt.Println("Next:")
	fmt.Println("  cd", projectName)
	fmt.Println("  bash .stackforge/setup.sh")
}
