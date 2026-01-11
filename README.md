# StackForge

**StackForge** is a lightweight, cross-platform CLI tool that bootstraps development environments and project structures using **declarative modules and presets**.

It generates a reproducible setup script instead of performing magic behind your back — so you always know **what will be installed, what files will be created, and how your project is set up**.

> One binary. No runtime dependencies. ~3 MB.

---

## Why StackForge?

Modern project bootstrapping tools often:
- install global dependencies implicitly
- hide logic inside frameworks
- lock you into opinionated workflows
- require heavy runtimes (Node, Docker, etc.)

**StackForge takes a different approach:**

- 🧱 Composable modules (language, framework, tools)
- 📦 Presets for common stacks
- 📜 Transparent shell scripts
- ♻️ Reproducible setups
- ⚡ Fast, tiny, standalone binary

---

## Features

- Multi-language support (Python, Node.js, Go, Java)
- Dependency resolution between modules
- Conflict detection
- Project presets (API, CLI, frontend, backend)
- `--dry-run` mode (preview before install)
- Add modules to existing projects
- Cross-platform (Linux, macOS, Windows*)
- No background daemons, no lock-in

---

## CLI Usage

```
stackforge init <project> <modules|presets> [--dry-run]
stackforge add <modules> [--dry-run]
stackforge list
stackforge presets
stackforge --help
stackforge --version
```

### Examples

```
stackforge init api python-api
stackforge init web nextjs-app --dry-run
stackforge add pytest ruff
stackforge list
stackforge presets
```

---

## Version Information

```
stackforge --version
```

Displays the StackForge ASCII logo, tool description, and current version.

---

## Quick Start

### Initialize a project using a preset

```
stackforge init my-api python-api
cd my-api
bash .stackforge/setup.sh
```

### Dry run (recommended)

```
stackforge init my-api python-api --dry-run
```

---

## Add modules to an existing project

```
stackforge add black ruff
bash .stackforge/setup.sh
```

---

## Supported Presets

### Python
- python-basic
- python-api
- flask-app
- django-app

### Node.js
- node-basic
- node-api
- react-app
- nextjs-app

### Go
- go-basic
- go-cli
- go-api

### Java
- java-basic
- java-api

List presets:
```
stackforge presets
```

---

## Supported Modules

Modules are small, declarative JSON definitions that describe:
- install commands
- generated files
- basic tests

List all modules:
```
stackforge list
```

---

## Project Structure

```
my-project/
├── .stackforge/
│   ├── setup.sh
│   └── manifest.json
├── generated project files
```

- `setup.sh` → reproducible setup script
- `manifest.json` → enabled modules and detected OS

---

## Design Philosophy

- Explicit over implicit
- Shell scripts over magic
- Composition over frameworks
- Readable configs over DSLs
- Small binaries over heavy runtimes

StackForge coordinates existing tools — it does not replace them.

---

## Roadmap

- Rebuild from manifest & lockfile
- Package manager helpers (dnf / apt / brew)
- Official website & documentation
- Community modules
- Optional version locking

---

## License

MIT License

---

## Contributing

Contributions are welcome:
- new modules
- new presets
- template improvements
- bug fixes

Open an issue or pull request to get started.

---

Built with care.
Designed to stay simple, fast, and honest.
