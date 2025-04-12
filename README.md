# 🎮 Slippi Playtime (Go)

![Go Version](https://img.shields.io/badge/go-1.21+-brightgreen?logo=go)  
**Slippi Playtime (Go)** is a fast and efficient Slippi replay analyzer written in Go. It scans `.slp` replay files to calculate your total playtime per character based on your Slippi code or in-game name. The output is concise, human-readable, and geared toward competitive players looking to analyze their habits, or look to see the matchup spread of opponents.

---

## 🕹️ Features

- ⚡ Fast analysis with parallel processing
- 📈 Breakdown of playtime and matches per character
- 🧠 Smart replay parsing using your Slippi code or username
- 📁 Custom Folder selection
- 🔒 Local-only processing, no network required

---

## ⏱️ Accuracy Notice

> **Note:** Playtime is estimated based on the assumption that all matches run at 60 FPS. Time spent in menus, stage/character select, and during pauses is **not included**.

---

## 🚀 Quick Start (Recommended)

### 📥 Download Prebuilt Release

> 🔗 [**Download the latest release**](https://github.com/9je/slippi-playtime-go/releases)

1. Download the release for your system.
2. Run the executable (`.exe`, `.out`, etc.).
3. Select your replay folder.
4. Enter your Slippi code (`DBC#544`) or username (`dbcooper`).
5. View your character playtime stats!

_This is the easiest and fastest way to use the tool—no setup required._

---

## 🛠️ Running from Source

If you want to build or modify the tool yourself, follow these steps:

### ✅ Prerequisites

- [Go (1.21+)](https://go.dev/dl/) must be installed and in your `PATH`.
- Clone the repository:

```sh
git clone https://github.com/9je/slippi-playtime-go.git
cd slippi-playtime-go
```

### 📦 Install Dependencies

```sh
go mod download
```

### ▶️ Run the Program

```sh
go run main.go
```

### 🏗️ Build Executable

```sh
go build .
```

This will create a binary in your current directory (`slippi-playtime-go` or `slippi-playtime-go.exe` depending on your OS).

---

## 🔍 How It Works

1. Prompts for your Slippi code or tag.
2. Lets you pick a folder containing your `.slp` replays.
3. Scans all files and extracts frame counts per character.
4. Converts frames to minutes/hours/days.
5. Outputs a breakdown of time played per character.

---

## 🤝 Contributing

Found a bug? Have an idea? Open an [issue](https://github.com/9je/slippi-playtime-go/issues) or submit a [pull request](https://github.com/9je/slippi-playtime-go/pulls)!  
All contributions are welcome—whether it's code, design, or feedback.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).