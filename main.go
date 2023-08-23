package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Token int

const (
	IncrementPointer Token = iota
	DecrementPointer
	IncrementValue
	DecrementValue
	Output
	Input
	LoopStart
	LoopEnd
	Unknown
)

const timeoutDuration = time.Millisecond * 200
const debugPrint = true

var tokenLookup = map[byte]Token{
	'>': IncrementPointer,
	'<': DecrementPointer,
	'+': IncrementValue,
	'-': DecrementValue,
	'.': Output,
	',': Input,
	'[': LoopStart,
	']': LoopEnd,
}

type Tokenizer struct {
	code        string
	codePointer int
	tokens      []Token
}

func NewTokenizer(code string) *Tokenizer {
	return &Tokenizer{
		code:        code,
		codePointer: 0,
		tokens:      []Token{},
	}
}

func (t *Tokenizer) tokenize() {
	for t.codePointer < len(t.code) {
		if token, ok := tokenLookup[t.code[t.codePointer]]; ok {
			t.tokens = append(t.tokens, token)
		}
		t.codePointer++
	}
}

type BrainfuckInterpreter struct {
	memory        [300]int
	pointer       int
	tokens        []Token
	tokenIndex    int
	output        strings.Builder
	showTapeState bool
	deadLoops     map[int]bool
}

func NewBrainfuckInterpreter() *BrainfuckInterpreter {
	return &BrainfuckInterpreter{
		pointer:   0,
		deadLoops: make(map[int]bool),
	}
}

func (bf *BrainfuckInterpreter) incrementPointer() {
	bf.pointer = (bf.pointer + 1) % len(bf.memory)
}

func (bf *BrainfuckInterpreter) decrementPointer() {
	bf.pointer = (bf.pointer - 1 + len(bf.memory)) % len(bf.memory)
}

func (bf *BrainfuckInterpreter) incrementValue() {
	bf.memory[bf.pointer] = (bf.memory[bf.pointer] + 1) % 256
}

func (bf *BrainfuckInterpreter) decrementValue() {
	bf.memory[bf.pointer] = (bf.memory[bf.pointer] - 1 + 256) % 256
}

func (bf *BrainfuckInterpreter) loadCode(code string) {
	tokenizer := NewTokenizer(code)
	tokenizer.tokenize()
	bf.tokens = tokenizer.tokens
	bf.output.Reset()
}

func (bf *BrainfuckInterpreter) displayTape() {
	fmt.Printf("Tape: %s\n", bf.formatTape())
	fmt.Printf("Output: %s\n", bf.output.String())
}

func (bf *BrainfuckInterpreter) formatTape() string {
	nonZeroCells := make([]string, 0, len(bf.memory))
	for i, cell := range bf.memory {
		if i == bf.pointer {
			nonZeroCells = append(nonZeroCells, fmt.Sprintf("[%5d]", cell))
		} else {
			nonZeroCells = append(nonZeroCells, fmt.Sprintf(" %5d ", cell))
		}
	}
	return strings.Join(nonZeroCells, "")
}

func (bf *BrainfuckInterpreter) inputInteger() int {
	fmt.Print("Input value (integer): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer.")
		return bf.inputInteger()
	}
	return intInput
}

func (bf *BrainfuckInterpreter) executeStep() {
	if bf.tokenIndex >= len(bf.tokens) {
		return
	}

	token := bf.tokens[bf.tokenIndex]
	var action string

	switch token {
	case IncrementPointer:
		bf.incrementPointer()
		action = "Increment pointer"
	case DecrementPointer:
		bf.decrementPointer()
		action = "Decrement pointer"
	case IncrementValue:
		bf.incrementValue()
		action = "Increment value"
	case DecrementValue:
		bf.decrementValue()
		action = "Decrement value"
	case Output:
		bf.output.WriteByte(byte(bf.memory[bf.pointer]))
		action = "Output value"
	case Input:
		bf.memory[bf.pointer] = bf.inputInteger()
		action = fmt.Sprintf("Input value: %d", bf.memory[bf.pointer])
	case LoopStart:
		if bf.memory[bf.pointer] == 0 {
			bf.tokenIndex = bf.findMatchingLoopEnd(bf.tokenIndex)
		}
		action = "Start loop"
	case LoopEnd:
		if bf.memory[bf.pointer] != 0 {
			bf.tokenIndex = bf.findMatchingLoopStart(bf.tokenIndex)
		}
		action = "End loop"
	default:
		action = "Unknown action"
	}

	if bf.showTapeState {
		time.Sleep(timeoutDuration)
		fmt.Printf("Action: %s\n", action)
		bf.displayTape()
	}

	bf.tokenIndex++
}

func (bf *BrainfuckInterpreter) execute() {
	for bf.tokenIndex < len(bf.tokens) {
		if bf.deadLoops[bf.tokenIndex] {
			bf.tokenIndex = bf.findMatchingLoopEnd(bf.tokenIndex) + 1
		} else {
			bf.executeStep()
		}
	}

	fmt.Printf("Final Output:\n%s\n", bf.output.String())
}

func (bf *BrainfuckInterpreter) findMatchingLoopStart(endIndex int) int {
	loopLevel := 1
	for i := endIndex - 1; i >= 0; i-- {
		switch bfToken := bf.tokens[i]; bfToken {
		case LoopEnd:
			loopLevel++
		case LoopStart:
			loopLevel--
			if loopLevel == 0 {
				return i
			}
		}
	}
	return endIndex // Unbalanced loop
}

func (bf *BrainfuckInterpreter) findMatchingLoopEnd(startIndex int) int {
	loopLevel := 1
	for i := startIndex + 1; i < len(bf.tokens); i++ {
		switch bfToken := bf.tokens[i]; bfToken {
		case LoopStart:
			loopLevel++
		case LoopEnd:
			loopLevel--
			if loopLevel == 0 {
				return i
			}
		}
	}
	return startIndex // Unbalanced loop
}

func main() {
	brainfuckCode := `
	>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<+
	+.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-
	]<+.
	`

	interpreter := NewBrainfuckInterpreter()
	interpreter.loadCode(brainfuckCode)
	interpreter.showTapeState = debugPrint
	interpreter.execute()
}
