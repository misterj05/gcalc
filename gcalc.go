package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gookit/color"
)

const ver float32 = 0.63

func clearTerm() {
	white := color.New(color.FgWhite, color.BgBlack).Render
	blue := color.New(color.FgCyan, color.BgBlack).Render
	lightYellow := color.New(color.FgLightYellow, color.BgBlack).Render

	switch runtime.GOOS {
	case "linux":
		_, errLinuxClearTerm := fmt.Println("\033[H\033[2J\033[34mGCalc", ver, "\033[97mby \033[35mWideHardo J\n\033[33mA simple terminal calculator written in \033[34mGo\033[33m.\nFormatting goes {number operation number}. Example:2 + 3\n\033[0m")
		// _, errLinuxClearTerm := fmt.Print(blue("GCalc by "), purple("WideHardo J\n"), yellow("A simple terminal calculator written in "), blue("Go"), yellow(".\nFormatting goes {number operation number}.\n"))
		if errLinuxClearTerm != nil {
			log.Println("linuxPrintErr: ", errLinuxClearTerm)
			time.Sleep(2)
		}
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		_, errWindowsClearTerm := fmt.Print(blue("GCalc ", ver), white(" by "), blue("WideHardo J\n"), lightYellow("A simple terminal calculator written in "), blue("Go"), lightYellow(".\nFormatting goes {number operation number}. Example:2 + 3\n"))
		if errWindowsClearTerm != nil {
			log.Println("defaultPrintErr: ", errWindowsClearTerm)
			time.Sleep(2)
		}
	default:
		_, errDefaultClearTerm := fmt.Println("GCalc ", ver, "by WideHardo J\nA simple terminal calculator written in Go.\nFormatting goes {number operation number}. Example:2 + 3\n")
		if errDefaultClearTerm != nil {
			log.Println("defaultPrintErr: ", errDefaultClearTerm)
			time.Sleep(2)
		}
	}
}

func main() {
	var num1, num2 int64
	var operator string
	clearTerm()
	for {
		_, printExpErr := fmt.Println("Give an expression.\n")
		if printExpErr != nil {
			log.Println("Print Expression Error: ", printExpErr)
			time.Sleep(2)
		}
		_, scanErr := fmt.Scanln(&num1, &operator, &num2)
		if scanErr != nil {
			log.Println("Scan Error: ", scanErr)
			time.Sleep(2)
		}
		clearTerm()
		_, printResultErr := fmt.Printf("%d %s %d = ", num1, operator, num2)
		if printResultErr != nil {
			log.Println("Print Result Error: ", printResultErr)
			time.Sleep(2)
		}
		switch operator {
		case "+":
			fmt.Println(num1 + num2)
		case "p":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "min":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "mul":
			fmt.Println(num1 * num2)
		case "x":
			fmt.Println(num1 * num2)
		case "/":
			fmt.Println(num1 / num2)
		case "d":
			fmt.Println(num1 / num2)
		default:
			fmt.Println("Incorrect Operator, Try Again.")
		}
	}
}
