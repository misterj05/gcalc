package main

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const ver float32 = 0.61

func clearTerm() {
	white := color.New(color.FgWhite, color.BgBlack).Render
	blue := color.New(color.FgLightBlue, color.BgBlack).Render
	lightYellow := color.New(color.FgLightYellow, color.BgBlack).Render

	switch runtime.GOOS {
	case "linux":
		_, errLinuxClearTerm := fmt.Println("\033[H\033[2J\033[34mGCalc", ver, "\033[97mby \033[35mWideHardo J\n\033[33mA simple terminal calculator written in \033[34mGo\033[33m.\nFormatting goes {number operation number}.\n\033[0m")
		// _, errLinuxClearTerm := fmt.Print(blue("GCalc by "), purple("WideHardo J\n"), yellow("A simple terminal calculator written in "), blue("Go"), yellow(".\nFormatting goes {number operation number}.\n"))
		if errLinuxClearTerm != nil {
			log.Fatal("linuxPrintErr: ", errLinuxClearTerm)
		}
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		_, errWindowsClearTerm := fmt.Print(blue("GCalc "), white(ver, "by "), blue("WideHardo J\n"), lightYellow("A simple terminal calculator written in "), blue("Go"), lightYellow(".\nFormatting goes {number operation number}.\n"))
		if errWindowsClearTerm != nil {
			log.Fatal("defaultPrintErr: ", errWindowsClearTerm)
		}
	default:
		_, errDefaultPrint := fmt.Println("GCalc ", ver, "by WideHardo J\nA simple terminal calculator written in Go.\nFormatting goes {number operation number}.\n")
		if errDefaultPrint != nil {
			log.Fatal("defaultPrintErr: ", errDefaultPrint)
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
			log.Fatal("Print Expression Error: ", printExpErr)
		}
		_, scanErr := fmt.Scanln(&num1, &operator, &num2)
		if scanErr != nil {
			log.Println("Scan Error: ", scanErr)
		}
		clearTerm()
		_, printResultErr := fmt.Printf("%d %s %d = ", num1, operator, num2)
		if printResultErr != nil {
			log.Fatal("Print Result Error: ", printResultErr)
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
