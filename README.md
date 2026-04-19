# 🚀 Go Mastery - A Comprehensive Learning Journey

A structured, hands-on learning project covering Go fundamentals from zero to intermediate proficiency. This repository contains 28 progressive modules designed to build a solid foundation in the Go programming language.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Learning Path](#learning-path)
- [Module Guide](#module-guide)
- [How to Run](#how-to-run)
- [Project Structure](#project-structure)
- [Tips for Learning](#tips-for-learning)

---

## 🎯 Overview

**Go Mastery** is a self-paced learning curriculum designed for developers who want to master Go from the ground up. Each module focuses on a specific concept, combining theory with practical code examples.

### Key Features
- ✅ **Progressive Difficulty**: Modules build upon each other
- ✅ **Hands-on Learning**: Every concept includes working code examples
- ✅ **Well-Organized**: Logical grouping of related topics
- ✅ **Self-Contained**: Each module is independent and runnable
- ✅ **Best Practices**: Covers Go idioms and conventions

---

## 📦 Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.25.6 or later** - [Download Go](https://golang.org/dl/)
- A text editor or IDE (VS Code, GoLand, etc.)
- Basic command-line familiarity
- Understanding of general programming concepts

### Verify Installation

```bash
go version
```

---

## 🏃 Getting Started

1. **Clone or download this repository**
   ```bash
   git clone <repository-url>
   cd go-mastery
   ```

2. **Navigate to any module**
   ```bash
   cd 01_setup_first-program
   ```

3. **Run the code**
   ```bash
   go run main.go
   ```

---

## 📚 Learning Path

The modules are structured to follow a natural learning progression:

| Phase | Modules | Topics |
|-------|---------|--------|
| **Foundations** | 01-02 | Setup, variables, types |
| **Syntax Basics** | 03-09 | Packages, operators, control flow |
| **Collections** | 10-18 | Arrays, slices, maps |
| **Functions** | 19-23 | Functions, return values, error handling, defer |
| **Advanced Types** | 24-27 | Pointers, structs, methods |

---

## 🔍 Module Guide

### Phase 1: Foundations (Modules 01-02)

| # | Module | Topics |
|---|--------|--------|
| 01 | **Setup & First Program** | Go project structure, `main` package, `func main()`, Hello World |
| 02 | **Variables & Types** | Variable declaration with `var`, basic types (string, int, float64) |

### Phase 2: Syntax Basics (Modules 03-09)

| # | Module | Topics |
|---|--------|--------|
| 03 | **Packages & Imports** | Package organization, import statements, visibility |
| 04 | **var vs Short Declaration** | `var` vs `:=` operator, scope, type inference |
| 05 | **Basic Types: String** | String literals, string operations |
| 06 | **Basic Types: Integer** | Integer types, signed/unsigned integers |
| 07 | **Basic Types: Boolean** | Boolean values, logical operators |
| 08 | **Constants** | Constant declaration, immutability, type specification |
| 09 | **If and Else** | Conditional statements, if-else chains |

### Phase 3: Control Flow (Modules 10-12)

| # | Module | Topics |
|---|--------|--------|
| 10 | **If with Short Statement** | Inline variable declaration in if conditions |
| 11 | **For Classic Loop** | Traditional for loops, loop patterns |
| 12 | **Switch Statement** | Switch cases, fall-through, default |

### Phase 4: Collections (Modules 13-18)

| # | Module | Topics |
|---|--------|--------|
| 13 | **Arrays** | Array declaration, fixed-size collections |
| 14 | **Slice Literals** | Slice basics, slice vs array |
| 15 | **Len and Capacity of Slice** | Slice length, capacity, growth |
| 16 | **For Range Over Slice** | Iterating slices with range keyword |
| 17 | **Map** | Map declaration, key-value pairs, iteration |
| 18 | **Read Value & OK** | Map value checking with the "ok" idiom |

### Phase 5: Functions (Modules 19-23)

| # | Module | Topics |
|---|--------|--------|
| 19 | **Functions** | Function declaration, parameters, return values |
| 20 | **Named Return Values** | Functions with named return values |
| 21 | **Variadic Functions** | Variable number of arguments with `...` |
| 22 | **Error Return Pattern** | Error handling, returning errors |
| 23 | **Defer Basics** | Defer statements, deferred function calls |

### Phase 6: Advanced Types (Modules 24-27)

| # | Module | Topics |
|---|--------|--------|
| 24 | **Pointers** | Pointer syntax, dereferencing, memory addresses |
| 25 | **Structs** | Struct declaration, fields, struct literals |
| 26 | **Methods: Value Receiver** | Methods on types with value receivers |
| 27 | **Methods: Pointer Receiver** | Methods on types with pointer receivers |

---

## 🖥️ How to Run

### Run a Specific Module

```bash
# Navigate to the module
cd 019_functions

# Run the code
go run main.go
```

### Run All Modules (Sequential)

Create a script or use your shell to run all modules:

**On Linux/macOS:**
```bash
for dir in */; do
  echo "Running $dir..."
  (cd "$dir" && go run main.go)
  echo "---"
done
```

**On Windows (PowerShell):**
```powershell
Get-ChildItem -Directory | ForEach-Object {
  Write-Host "Running $_..."
  & go run "./$_/main.go"
  Write-Host "---"
}
```

### Modify and Experiment

Each module encourages hands-on learning:

1. Read the code and comments
2. Run the program
3. Modify values or logic
4. Run again to see the changes
5. Experiment with new ideas

---

## 📁 Project Structure

```
go-mastery/
├── README.md                          # This file
├── go.mod                             # Go module definition
├── 01_setup_first-program/
│   └── main.go
├── 02_variables_and_types/
│   └── main.go
├── 03_packages_imports/
│   └── main.go
│
├── ... [modules 04-09] ...
│
├── 010_if_with_shortstatement/
│   └── main.go
├── 011_for_classic_loop/
│   └── main.go
├── 012_switch_statement/
│   └── main.go
│
├── ... [modules 13-27] ...
│
└── .git/                              # Version control
```

**Note:** Each module is self-contained with its own `main.go` file.

---

## 🎓 Tips for Learning

### Best Practices

1. **Read the Code First** - Before running, read through the code and comments
2. **Run Each Module** - Execute the code to see it in action
3. **Experiment** - Modify values, try different inputs, break things intentionally
4. **Take Notes** - Document your understanding of each concept
5. **Progress Sequentially** - Don't skip modules; each builds on previous knowledge

### Common Patterns to Watch For

- `package main` and `func main()` - Entry point pattern
- `:=` - Short variable declaration (Go's type inference)
- `defer` - Deferred execution (cleanup operations)
- `_, err` - Error handling pattern
- Value vs. Pointer receivers - When to use which

### Additional Resources

- **Official Docs**: [golang.org](https://golang.org)
- **Go by Example**: [gobyexample.com](https://gobyexample.com)
- **Effective Go**: [Go best practices guide](https://golang.org/doc/effective_go)

---

## 📊 Learning Milestones

Track your progress:

- [ ] **Module 01-02**: Understand Go setup and basic variables
- [ ] **Module 03-09**: Master syntax fundamentals
- [ ] **Module 10-12**: Control flow mastery
- [ ] **Module 13-18**: Collections and data structures
- [ ] **Module 19-23**: Functions and error handling
- [ ] **Module 24-27**: Pointers, structs, and methods

Once all modules are completed, you'll have a solid foundation in Go and be ready to build real-world applications!

---

## 🚀 Next Steps After Mastery

After completing all 27 modules, consider:

- Building small CLI tools
- Creating web services with Go
- Contributing to open-source Go projects
- Exploring advanced topics (interfaces, concurrency, reflection)
- Reading source code from popular Go projects

---

## 📝 License

This project is provided for educational purposes. Feel free to use, modify, and share!

---

## 💡 Contributing

Have suggestions or improvements? Feel free to:
- Create issues for unclear concepts
- Submit improvements to code examples
- Add additional learning resources

---

**Happy Learning! 🎉**

*Start with Module 01 and progress through the modules at your own pace. Happy coding!*
