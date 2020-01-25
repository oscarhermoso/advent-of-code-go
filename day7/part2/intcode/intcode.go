package intcode

import "fmt"

// Intcode runs the intcode computer on an inputted slice
// func Intcode(puzzleInput []int) {
// 	// default input slice, index 1 and 2 will be replaced
// 	// puzzleInput := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1, 192, 154, 224, 101, -161, 224, 224, 4, 224, 102, 8, 223, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1001, 157, 48, 224, 1001, 224, -61, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1102, 15, 28, 225, 1002, 162, 75, 224, 1001, 224, -600, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 102, 32, 57, 224, 1001, 224, -480, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 224, 223, 223, 1101, 6, 23, 225, 1102, 15, 70, 224, 1001, 224, -1050, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 101, 53, 196, 224, 1001, 224, -63, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 1101, 64, 94, 225, 1102, 13, 23, 225, 1101, 41, 8, 225, 2, 105, 187, 224, 1001, 224, -60, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1101, 10, 23, 225, 1101, 16, 67, 225, 1101, 58, 10, 225, 1101, 25, 34, 224, 1001, 224, -59, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 329, 101, 1, 223, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 344, 1001, 223, 1, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 359, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 404, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 419, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 101, 1, 223, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 464, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 494, 101, 1, 223, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 1001, 223, 1, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 101, 1, 223, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 101, 1, 223, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 644, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 659, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

// 	// Part 1: ID of System is 1 (use as the input)
// 	// runDiagnostics(puzzleInput, 1)

// 	// Part 2: ID of System is 5 (use as the input)
// 	RunDiagnostics(puzzleInput, 0)
// }

// RunDiagnostics runs the entire computer on a puzzleInput and an inputValue
// three returned ints will be
// 1. the index that the intcode compiler stopped on
// 2. the code that triggered the "exit"
// 3. the output
func RunDiagnostics(puzzleInput []int, inputValue int, index int) (int, int, int) {
	// store the lastOutputValue to be returned upon a 99
	var lastOutputValue int

	for i := index; i < len(puzzleInput); {
		// find op code (last 2 digits of number), a 1, 2, 3, 4, or 99
		// find param1 and param2 which are the 100's and 1000's digit in the i-th element
		opCode, param1, param2, _ := returnCodes(puzzleInput[i])
		var firstVal, secondVal, thirdVal int

		// set value variables based on what the opCodes are (some of them will throw errors if I attempt to set them...)
		if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
			firstVal, secondVal = getValue(puzzleInput, param1, puzzleInput[i+1]), getValue(puzzleInput, param2, puzzleInput[i+2])
			thirdVal = puzzleInput[i+3]
		} else if opCode == 3 || opCode == 4 {
			firstVal = puzzleInput[i+1]
		} else if opCode == 5 || opCode == 6 {
			firstVal, secondVal = getValue(puzzleInput, param1, puzzleInput[i+1]), getValue(puzzleInput, param2, puzzleInput[i+2])
		}

		// switch statement to handle the opcode value
		switch opCode {
		case 99:
			// fmt.Println("99 halts IntCode instance")
			return i, 99, lastOutputValue
		case 1:
			// add and place (by position)
			// firstToAdd, secondToAdd := getValue(puzzleInput, param1, puzzleInput[i+1]), getValue(puzzleInput, param2, puzzleInput[i+2])
			puzzleInput[thirdVal] = firstVal + secondVal
			i += 4
		case 2:
			// multiply and place (by position)
			puzzleInput[thirdVal] = firstVal * secondVal
			i += 4
		case 3:
			// fmt.Println("--INPUT OPCODE 3")
			// write inputValue to puzzle input by position
			puzzleInput[firstVal] = inputValue

			// adding functionality for a second input value (to get the above lines to use it, overwrite the inputValue variable with the next argument that was passed in)
			// inputValue = secondInputValue
			i += 2
		case 4:
			// output the value to the console, always by position
			// fmt.Println("output: ", puzzleInput[firstVal])
			lastOutputValue = puzzleInput[firstVal]
			i += 2
			// return the index (to pass back in, a 4 to signal that it was not a halt, and the outputted value)
			return i, 4, lastOutputValue
		case 5:
			// jump-if-true
			// if first param is != 0 set instruction pointer to value from second parameter
			if firstVal != 0 {
				// fmt.Println("5 moves i to", secondVal)
				i = secondVal
			} else {
				// otherwise do nothing except increment i
				// increment by 3 (like opCode 6)
				i += 3
			}
		case 6:
			// jump-if-false
			// if first param is != 0 set instruction pointer to value from second parameter
			if firstVal == 0 {
				i = secondVal
			} else {
				// otherwise increment i to the next instruction
				// this instruction has 3 values, so i increments by 3
				i += 3
			}
		case 7:
			// less than - if firstVal < secondVal, store a 1 @ position of thirdVal, otherwise store a 0
			if firstVal < secondVal {
				puzzleInput[thirdVal] = 1
			} else {
				puzzleInput[thirdVal] = 0
			}
			i += 4
		case 8:
			// equals, same as opCode 7 but conditional is firstVal is equal to secondVal
			if firstVal == secondVal {
				puzzleInput[thirdVal] = 1
			} else {
				puzzleInput[thirdVal] = 0
			}
			i += 4
		default:
			fmt.Println("bad opCode!!! returning -999")
			return 0, 0, -999
		}
	}

	fmt.Println("shouldn't have reached end of function, returning -1000")
	return 0, 0, -1000
}

func returnCodes(number int) (int, int, int, int) {
	return number % 100, (number % 1000) / 100, (number % 10000) / 1000, (number % 100000) / 10000
}

func getValue(puzzleInput []int, positionOrImmediateCode int, nextOrTwoPastValue int) int {
	// read if the param is for position or immediate mode and return the value (to be added or multiplied)
	if positionOrImmediateCode == 0 {
		// 0 is position mode
		return puzzleInput[nextOrTwoPastValue]
	}
	// otherwise it's immediate mode (no-else-return)
	return nextOrTwoPastValue
}

// note: a solution to part 1 that I looked up to get some insight into golang
// package main

// import (
// 	"fmt"
// 	"math"
// 	"os"
// )

// func main() {

// 	// var instructionStrings []string
// 	var results []int
// 	// input := 1

// 	// file, err := os.Open("./input.txt")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer file.Close()

// 	// scanner := bufio.NewScanner(file)
// 	// for scanner.Scan() {
// 	// 	line := scanner.Text()
// 	// 	instructionStrings = strings.Split(line, ",")
// 	// }

// 	// instructions := make([]int, len(instructionStrings))

// 	// for i, v := range instructionStrings {
// 	// 	instructions[i], _ = strconv.Atoi(v)
// 	// }

// 	// note pasted in instructions slice
// 	instructions := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1, 192, 154, 224, 101, -161, 224, 224, 4, 224, 102, 8, 223, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1001, 157, 48, 224, 1001, 224, -61, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1102, 15, 28, 225, 1002, 162, 75, 224, 1001, 224, -600, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 102, 32, 57, 224, 1001, 224, -480, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 224, 223, 223, 1101, 6, 23, 225, 1102, 15, 70, 224, 1001, 224, -1050, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 101, 53, 196, 224, 1001, 224, -63, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 1101, 64, 94, 225, 1102, 13, 23, 225, 1101, 41, 8, 225, 2, 105, 187, 224, 1001, 224, -60, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1101, 10, 23, 225, 1101, 16, 67, 225, 1101, 58, 10, 225, 1101, 25, 34, 224, 1001, 224, -59, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 329, 101, 1, 223, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 344, 1001, 223, 1, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 359, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 404, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 419, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 101, 1, 223, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 464, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 494, 101, 1, 223, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 1001, 223, 1, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 101, 1, 223, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 101, 1, 223, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 644, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 659, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}
// 	input := 1

// 	index := 0
// 	for {

// 		opcode, arg1, arg2, arg3 := getValues(instructions, index)
// 		switch opcode {
// 		case 1:
// 			instructions[arg3] = arg1 + arg2
// 			index += 4
// 		case 2:
// 			instructions[arg3] = arg1 * arg2
// 			index += 4
// 		case 3:
// 			instructions[arg1] = input
// 			index += 2
// 		case 4:
// 			results = append(results, instructions[arg1])
// 			index += 2
// 		case 99:
// 			fmt.Println(results)
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Invalid opcode")
// 			os.Exit(1)
// 		}
// 	}
// }

// func getValues(instructions []int, index int) (int, int, int, int) {

// 	instruction := instructions[index]
// 	if instruction == 99 {
// 		return instruction, 0, 0, 0
// 	}

// 	opcode := instruction % 100

// 	if opcode == 3 || opcode == 4 {
// 		return opcode, instructions[index+1], 0, 0
// 	}

// 	// tests if opcode is 1 or 2...
// 	// if inSlice([]int{1, 2}, opcode) {
// 	if opcode == 1 || opcode == 2 {
// 		// variables to store the values that the op code should use
// 		var arg1, arg2, arg3 int
// 		// arg 3 is where to write the next value to, (it's always by position)
// 		arg3 = instructions[index+3]

// 		if math.Floor(float64((instruction%1000)/100)) == 1 {
// 			// 1 is immediate mode
// 			arg1 = instructions[index+1]
// 		} else {
// 			// otherwise 0 is position mode
// 			arg1 = instructions[instructions[index+1]]
// 		}
// 		if math.Floor(float64((instruction%10000)/1000)) == 1 {
// 			arg2 = instructions[index+2] // immediate
// 		} else {
// 			arg2 = instructions[instructions[index+2]] // position
// 		}
// 		// return the list of arrays
// 		return opcode, arg1, arg2, arg3
// 	}

// 	return opcode, 0, 0, 0
// }