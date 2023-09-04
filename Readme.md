# Brainfuck Interpreter in Go

This is a simple Brainfuck interpreter implemented in Go. The interpreter reads Brainfuck code, executes it step by step, and displays the memory tape and pointer's state after each step. It also shows the action executed in each step and the current output of the program.

## How to Run

1. **Install Go:** If you haven't already, you can install Go by following the official installation guide: [https://golang.org/doc/install](https://golang.org/doc/install).

2. **Clone or Create a Go File:**

   - Clone this repository to your local machine or create a new Go file and paste the provided code.

3. **Run the Interpreter:**

   - Open your terminal and navigate to the directory where your Go file is located.

   - Execute the following command

     ```shell
     go run main.go
     ```

4. **Interprete Brainfuck Code:**

## Features

- The interpreter can displays the memory tape and pointer's position after each step of execution.
- It shows the action executed in each step, such as incrementing the pointer, decrementing the value, outputting a character, and more.
- The interpreter collects the output of the Brainfuck program and displays it as the code executes.

## Customization

- You can modify the `brainfuckCode` variable in the `main` function to run any Brainfuck code you'd like.
- To adjust the delay between each step (for debugging purposes), you can modify the value in the `time.Sleep` function within the `executeStep` function.

## Flags and Piping

- **Debug Flag:** You can enable the debug output using the `-debug` flag. For example:

   ```shell
   go run main.go -debug
   ```

- **Piping Input:** You can pipe Brainfuck code as input to the interpreter. For example:

   ```shell
   echo ">+++++[<++++>-]<.>++++++[<-------->-]<." | go run main.go
   ```

- **Capture Output:** To capture the output in a file, you can use the standard shell pipe operator (`|`). For example, to execute Brainfuck code and save the output to a file:

   ```shell
   echo ">+++++[<++++>-]<.>++++++[<-------->-]<." | go run main.go > output.txt
   ```

Feel free to explore and enhance this Brainfuck interpreter to suit your needs or to use it as a learning tool for Go programming.

## Dependencies

- This program uses the standard Go library for most of its functionality and does not require additional external libraries.
