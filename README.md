
# Or use the full binary path
./ASCII-Converter-Gen-Z-Edition/brainrot-ascii
```

**The installer puts everything exactly where you need it.** Just restart your shell and go wild! 🎉# ⚠️ Things You Need to Know Before Installing

Prepare to be *mindfucked* by the brainrot outputs — handcrafted chaos, just for you the user.

You **must** have a Unix-based system (Linux, macOS, BSD, etc.)
**Windows users**: try WSL

---

# 📦 Dependencies

* **Python 3.6+** (for the installer)
* **Go 1.19 or newer** (auto-installed by the Python installer)
* **Rust** (auto-installed by the Python installer)
* **Git** (auto-installed by the Python installer)

**Literally just run the Python installer and forget about it.** No manual dependency hunting. No drama.

---

# 🛠️ How to Install

## Quick Install (Recommended)

Use our universal Python installer that handles **EVERYTHING** for you:

```bash
# Download the installer
curl -sSL https://raw.githubusercontent.com/Zsombyy/ASCII-Converter-Gen-Z-Edition/main/install_deps.py -o install_deps.py

# Run it
python3 install_deps.py
```

**What it does:**
- 🔍 Auto-detects your Linux distro (Arch, Debian, Fedora, openSUSE)
- 📦 Installs system dependencies (git, tar, curl, build tools, Go)
- 🦀 Installs Rust toolchain via rustup
- 📂 Downloads the entire ASCII Converter repository
- 🔨 Builds the `brainrot-ascii` binary automatically
- ⚙️ Configures your shell environment with aliases
- 🎯 Installs required Rust crates (tar, flate2)

**After installation, you get:**
- ✅ Ready-to-use `brainrot-ascii` binary
- ✅ Shell aliases: `br`, `brmax`, `brquiet`
- ✅ Complete development environment
- ✅ All dependencies sorted

**Supported distros:**
- **Arch family**: Arch, Manjaro, EndeavourOS, Artix, Garuda
- **Debian family**: Debian, Ubuntu, Mint, Pop!_OS, Elementary, Kali
- **Fedora family**: Fedora, CentOS, RHEL, Rocky Linux, AlmaLinux
- **openSUSE family**: openSUSE Leap, Tumbleweed, SLED, SLES
- **Other**: macOS with Homebrew

## Manual Build (if you hate convenience)

Only do this if you enjoy pain and already have all dependencies:

```bash
git clone https://github.com/Zsombyy/ASCII-Converter-Gen-Z-Edition.git
cd ASCII-Converter-Gen-Z-Edition
go build -o brainrot-ascii main.go
```

Then run it:

```bash
./brainrot-ascii
```

> ❗ If you see *`Permission denied`*, just:

```bash
chmod +x brainrot-ascii
```

**Pro tip:** Just use the Python installer instead. It's literally faster and handles everything.

## Pre-Built Binary (old school)

**Note:** The Python installer is way better, but if you insist...

[Click here to download](https://github.com/Zsombyy/ASCII-Converter-Gen-Z-Edition/releases/download/release/brainrot-ascii)

After download, navigate to the download folder and run:

```bash
chmod +x brainrot-ascii
./brainrot-ascii
```

## One-Liner Install (for the ultra-lazy)

```bash
curl -sSL https://raw.githubusercontent.com/Zsombyy/ASCII-Converter-Gen-Z-Edition/main/install_deps.py | python3
```

This single command will:
1. ⚡ Run the complete dependency installer
2. 📥 Download the entire repository
3. 🔧 Build the binary automatically
4. 🎪 Set up all shell aliases
5. 🚀 Leave you ready to generate brainrot

**That's it. One command. Maximum chaos enabled.** 🧠💥

---

# 🎉 Main Features

* Multiple output formats
* Certified brainrot output
* Random Gen-Z brainrot quotes
* Multiple brainrot output styles
* Adjustable brainrot level (casual to catastrophic)
* Cross-platform Rust integration for advanced processing


# 🚀 Shell Integration

After installation, you can use these aliases (automatically added to your shell config):

```bash
# Quick brainrot
alias br='brainrot-ascii'

# Maximum chaos mode
alias brmax='BRAINROT_LEVEL=10 brainrot-ascii'

# Silent mode (for scripting)
alias brquiet='brainrot-ascii --quiet'
```

---

# 🐛 Found Bugs?

Hit me up:

* [GitHub Issues](https://github.com/Zsombyy/ASCII-Converter-Gen-Z-Edition/issues)
* Email: [info@cubxy.dev](mailto:info@cubxy.dev)

---

# 🤝 Contributing

Want to add more brainrot? Here's how:

1. **Easy way:** Run `python3 install_deps.py` (sets up everything)
2. Make your chaotic changes
3. Test with: `go test ./...`
4. Submit a PR with maximum brainrot energy

**Or the manual way** (if you hate efficiency):
1. Fork the repo  
2. Install Go, Rust, dependencies manually
3. Clone your fork
4. Suffer through manual setup
5. Make changes and test
6. Submit PR

**Seriously, just use the installer.** 🤷‍♂️

---

# 📄 License

This project is licensed under the "Do Whatever The Fuck You Want" license.
Seriously, go wild. 🎉

---

# 🙏 Acknowledgments

* The Gen-Z community for inspiring this beautiful chaos
* Rust community for the blazingly fast dependencies
* Go community for keeping it simple
* Everyone who contributes to the brainrot

---

**Remember**: This tool is designed to break your brain in the best possible way. Use responsibly (or don't, we're not your parents). 🧠💥
