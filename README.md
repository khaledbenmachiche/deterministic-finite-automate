# Deterministic Finite Automaton (DFA) Implementation in Golang

This project contains a deterministic finite automaton (DFA) implemented in the Go programming language.

## Overview

A Deterministic Finite Automaton (DFA) is a mathematical model that recognizes a set of strings based on a defined set of states and transition rules. This implementation showcases the basic structure of a DFA.

## Files

- `main.go`: Contains the main code implementing the DFA, its states, transitions, and string recognition.
- `README.md`: The file you're currently reading, providing information about the project.

## Prerequisites

To run this DFA implementation, ensure you have Go installed. If not, you can download and install it from the [official Go website](https://golang.org/).

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/khaledbenmachiche/deterministic-finite-automate/
   cd deterministic-finite-automate
2. Compile
   ```bash
   go build main.go
3. execute:
   ```bash
    ./main

# Example
Consider an example DFA that recognizes strings of 0s and 1s where the number of 0s is divisible by 3.

"000" is accepted.\n
"110" is rejected.\n
Modify the transition table and rules in the code to create and test different DFAs based on your requirements.

# License
This project is licensed under the MIT License - you're free to use, modify, and distribute the code as per the license terms.

Feel free to contribute, improve, or use this DFA implementation for your projects. Happy coding!
