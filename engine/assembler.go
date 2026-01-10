package engine

import "strings"

func Assemble(mods []Module, baseTemplate string, osinfo OSInfo) (string, error) {
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

	out := strings.ReplaceAll(baseTemplate, "{{install}}", strings.Join(installs, "\n"))
	out = strings.ReplaceAll(out, "{{files}}", strings.Join(files, "\n"))
	out = strings.ReplaceAll(out, "{{tests}}", strings.Join(tests, "\n"))

	return out, nil
}
