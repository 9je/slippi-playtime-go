# Slippi Playtime (Go)

## Overview

Slippi Playtime (Go) is a fast and efficient Slippi replay analyzer written in Go. It scans Slippi `.slp` replay files to calculate total playtime per character, based on the provided Slippi code or in-game name. The tool outputs detailed statistics, including time spent per character and total gameplay duration.

## In-Game Time Accuracy Notice

Playtime is not 100% accurate as it assumes your gameplay runs at a constant rate of 60 FPS. Menu, character/stage select, and pause time is not accounted for.

## Features

- Parses Slippi replay files to extract gameplay statistics.
- Calculates total playtime and game count per character.
- Supports parallel processing for fast analysis.
- Outputs results in a clear, human-readable format.

## Installation & Usage

### Prerequisites

Ensure you have [Go](https://go.dev/) installed on your system.

### Running the Program

To test the program without building an executable, use:

```sh
 go run main.go
```

### Building the Executable

To build the program as an executable:

```sh
 go build .
```

This will generate a binary that you can run directly.

## How It Works

1. Prompts the user for their Slippi code or in-game name.
2. Allows the user to select their Slippi replay folder.
3. Scans all `.slp` files in the directory, extracting relevant metadata.
4. Calculates the total frames played per character and converts them into minutes, hours, or days.
5. Outputs a detailed breakdown of playtime per character.

## Contributing

Feel free to open issues or submit pull requests to enhance functionality or improve performance.

## License

This project is open-source under the MIT License.

