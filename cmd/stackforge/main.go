package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
	// Locate StackForge root (where modules/, templates/, presets.json live)
	exe, _ := os.Executable()
	root := filepath.Dir(exe)

	// ---------------- Parse --dry-run ----------------
	dryRun := false
	var cleanArgs []string

	for _, a := range os.Args[1:] {
		if a == "--dry-run" {
			dryRun = true
		} else {
			cleanArgs = append(cleanArgs, a)
		}
	}

	os.Args = append([]string{os.Args[0]}, cleanArgs...)

	// ---------------- HELP / VERSION ----------------
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]

	if cmd == "--help" || cmd == "-h" {
		printHelp()
		return
	}

	if cmd == "version" || cmd == "--version" || cmd == "-v" {
		printVersion()
		return
	}

	// ---------------- LIST ----------------
	if cmd == "list" {
		mods, err := engine.LoadModules(root + "/modules")
		if err != nil {
			panic(err)
		}
		for _, m := range engine.ListModules(mods) {
			fmt.Println(m)
		}
		return
	}

	// ---------------- PRESETS ----------------
	if cmd == "presets" {
		presets, _ := engine.LoadPresets(root + "/presets.json")
		for _, p := range engine.PresetNames(presets) {
			fmt.Println(p)
		}
		return
	}

	// ---------------- ADD ----------------
	if cmd == "add" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: stackforge add <module> [module...] [--dry-run]")
			return
		}

		manifest, err := engine.LoadManifest(".stackforge/manifest.json")
		if err != nil {
			panic(err)
		}

		manifest.Modules = engine.Unique(append(manifest.Modules, os.Args[2:]...))

		modulesMap, err := engine.LoadModules(root + "/modules")
		if err != nil {
			panic(err)
		}

		resolved, err := engine.Resolve(manifest.Modules, modulesMap)
		if err != nil {
			panic(err)
		}

		osinfo := engine.DetectOS()
		base, _ := os.ReadFile(root + "/templates/base.sh")

		plan, err := engine.Assemble(resolved, string(base), osinfo)
		if err != nil {
			panic(err)
		}

		if dryRun {
			fmt.Println("StackForge dry run")
			fmt.Println()
			fmt.Println("Will add:")
			for _, m := range os.Args[2:] {
				fmt.Println(" ", m)
			}
			fmt.Println("\nCommands:")
			for _, c := range plan.Installs {
				fmt.Println(" ", c)
			}
			return
		}

		os.WriteFile(".stackforge/setup.sh", []byte(plan.Script), 0755)
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

	projectName := os.Args[2]
	input := os.Args[3:]

	presets, _ := engine.LoadPresets(root + "/presets.json")

	var modules []string
	for _, name := range input {
		if preset, ok := presets[name]; ok {
			modules = append(modules, preset...)
		} else {
			modules = append(modules, name)
		}
	}

	modules = engine.Unique(modules)

	modulesMap, err := engine.LoadModules(root + "/modules")
	if err != nil {
		panic(err)
	}

	resolved, err := engine.Resolve(modules, modulesMap)
	if err != nil {
		panic(err)
	}

	osinfo := engine.DetectOS()
	base, _ := os.ReadFile(root + "/templates/base.sh")

	plan, _ := engine.Assemble(resolved, string(base), osinfo)

	if dryRun {
		fmt.Println("StackForge dry run")
		fmt.Println()
		fmt.Println("Will install:")
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

	os.Mkdir(projectName, 0755)
	os.Mkdir(projectName+"/.stackforge", 0755)

	os.WriteFile(projectName+"/.stackforge/setup.sh", []byte(plan.Script), 0755)

	manifest := engine.Manifest{
		Modules: modules,
		OS:      osinfo.PackageManager,
	}

	data, _ := json.MarshalIndent(manifest, "", "  ")
	os.WriteFile(projectName+"/.stackforge/manifest.json", data, 0644)

	fmt.Println("Project created:", projectName)
	fmt.Println("Next:")
	fmt.Println("  cd", projectName)
	fmt.Println("  bash .stackforge/setup.sh")
}
