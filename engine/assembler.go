package engine

import "strings"

func Assemble(mods []Module, baseTemplate string, osinfo OSInfo) (Plan, error) {
	var installs []string
	var commands []string
	var tests []string
	var files []string

	for _, m := range mods {

		// 1️⃣ OS-level installs
		if m.Install != nil {
			if cmds, ok := m.Install[osinfo.PackageManager]; ok {
				installs = append(installs, cmds...)
			}
		}

		// 2️⃣ Project-level commands
		if len(m.Commands) > 0 {
			commands = append(commands, m.Commands...)
		}

		// 3️⃣ Files
		for path, content := range m.Files {
			files = append(files,
				"mkdir -p $(dirname "+path+")",
				"cat << 'EOF' > "+path+"\n"+content+"\nEOF",
			)
		}

		// 4️⃣ Tests
		if len(m.Tests) > 0 {
			tests = append(tests, m.Tests...)
		}
	}

	// Inject into base.sh
	script := strings.ReplaceAll(baseTemplate, "{{install}}", strings.Join(installs, "\n"))
	script = strings.ReplaceAll(script, "{{commands}}", strings.Join(commands, "\n"))
	script = strings.ReplaceAll(script, "{{files}}", strings.Join(files, "\n"))
	script = strings.ReplaceAll(script, "{{tests}}", strings.Join(tests, "\n"))

	return Plan{
		Installs: installs,
		Commands: commands,
		Files:    files,
		Tests:    tests,
		Script:   script,
	}, nil
}
