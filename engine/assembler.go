package engine

import "strings"

func Assemble(mods []Module, baseTemplate string, osinfo OSInfo) (Plan, error) {
	var installs []string
	var tests []string
	var files []string

	for _, m := range mods {
		cmds, ok := m.Install[osinfo.PackageManager]
		if ok {
			installs = append(installs, cmds...)
		}

		tests = append(tests, m.Test...)

		for path, content := range m.Files {
			files = append(files,
				"mkdir -p $(dirname "+path+")",
				"cat << 'EOF' > "+path+"\n"+content+"\nEOF",
			)
		}
	}

	script := strings.ReplaceAll(baseTemplate, "{{install}}", strings.Join(installs, "\n"))
	script = strings.ReplaceAll(script, "{{files}}", strings.Join(files, "\n"))
	script = strings.ReplaceAll(script, "{{tests}}", strings.Join(tests, "\n"))

	return Plan{
		Installs: installs,
		Files:    files,
		Tests:    tests,
		Script:   script,
	}, nil
}
