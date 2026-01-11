# Installing StackForge

This document explains how to install **StackForge** on supported platforms.

StackForge is distributed as a **single static binary**.
There is no installer, package manager, or runtime dependency.

---

## Supported Platforms

- Linux (x86_64)
- macOS (Intel & Apple Silicon)
- Windows (native binary + WSL for setup scripts)

---

## Linux

### Download
Download `stackforge-linux-amd64` from the releases page.

### Install
```bash
chmod +x stackforge-linux-amd64
sudo mv stackforge-linux-amd64 /usr/local/bin/stackforge
```

### Verify
```bash
stackforge --version
```

---

## macOS

### Download
- Apple Silicon: `stackforge-darwin-arm64`
- Intel: `stackforge-darwin-intel-amd64`

### Install
```bash
chmod +x stackforge-darwin-*
sudo mv stackforge-darwin-* /usr/local/bin/stackforge
xattr -d com.apple.quarantine /usr/local/bin/stackforge
```

### Verify
```bash
stackforge --version
```

---

## Windows (Native Binary + WSL)

StackForge itself runs **natively on Windows**.  
However, generated project setup scripts (`.sh`) require **WSL**.

---

### 1. Install StackForge (Windows binary)

1. Download `stackforge-windows-amd64.exe`
2. Create a tools directory:
```powershell
mkdir C:\Tools
```

3. Move and rename the binary:
```text
C:\Tools\stackforge.exe
```

---

### 2. Add StackForge to PATH

1. Open **Environment Variables**
2. Edit **User variables → Path**
3. Add:
```text
C:\Tools
```

4. Open a new terminal

Verify:
```powershell
stackforge --version
```

---

### 3. Install WSL (Required for setup scripts)

Run in **PowerShell (Administrator)**:
```powershell
wsl --install
```

Reboot if prompted.

---

### 4. Running project setup scripts

Setup scripts must be executed inside WSL:
```bash
bash .stackforge/setup.sh
```

Running `.sh` files directly from cmd.exe or PowerShell will not work.

---

## Troubleshooting

- `command not found`: ensure StackForge is in PATH
- `permission denied`: run `chmod +x stackforge`
- Windows script errors: confirm you are inside WSL

---

Happy forging ⚒️
