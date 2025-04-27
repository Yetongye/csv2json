# CSV to JSON Lines Converter

##  1. Project Overview

This project provides a command-line Go application that reads housing data from a CSV file and converts it into a JSON Lines (`.jl`) file. Each line in the output file represents one housing record as a valid JSON object.

The program validates user input and ensures smooth user experience by providing clear prompts and usage instructions.



##  2. Building the Executable

To build a standalone executable file:

### For macOS:

```bash
go build -o csv2json main.go
```

This will generate an executable file named `csv2json` in the current directory.

- **Why not use .app?**: The executable is a terminal command-line application, but macOS treats .app as a GUI application, which results in a direct kill because there is no GUI interface.


### For Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o csv2json.exe main.go
```

This will generate an executable file `csv2json.exe` for Windows systems.




## 3. Pre-built Executables Included

This repository already includes pre-built executables:

- `csv2json.exe` for Windows
- `csv2json` for macOS

### Running on Windows:
- You can **double-click** `csv2json.exe` to launch it, which will open a black command-line window.
- Alternatively, you can open Command Prompt (cmd) and run:

```bash
csv2json.exe housesInput.csv housesOutput.jl
```

### Running on macOS:
- Open Terminal.
- Navigate to the project directory.
- Run the program with:

```bash
./csv2json
```

> **Note:** On macOS, you cannot double-click `csv2json` because it is a terminal command-line application, not a graphical `.app` bundle.



##  4. Running the Application

After building the program:

### If running with command-line arguments:

```bash
./csv2json housesInput.csv housesOutput.jl
```

- `housesInput.csv` is the input CSV file.
- `housesOutput.jl` is the output JSON Lines file.

### If no arguments are provided:

If you run:

```bash
./csv2json
```

The program will interactively prompt you to enter:

1. The input CSV file path
2. The output JL file path

Example:

```
No input arguments detected.
Please enter the input CSV file path:
housesInput.csv
Please enter the output JL file path:
housesOutput.jl
Conversion complete. Output saved to housesOutput.jl
```



## 5. Testing the Application

Unit tests are provided to validate critical functionality.

To run all tests:

```bash
go test
```

This will run tests for:
- Parsing a CSV row into a House struct
- Writing House structs into a valid `.jl` file
- Handling invalid input gracefully

Tests are located in `csv2json_test.go`.


##  6. Use of AI Assistants

In completing this assignment, AI assistance (ChatGPT) was used for:
- Code style improvements and function modularization suggestions
- Debugging guidance
- Best practices for Go project structure and executable building (like how to push the project to the github)
- Help generate this markdown template

