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
		_, err := fmt.Println("\033[H\033[2J\033[34mGCalc", ver, "\033[97mby \033[35mMister J\n\033[33mA simple terminal calculator written in \033[34mGo\033[33m.\nFormatting goes {number operation number}. Example:2 + 3\033[0m")
		if err != nil {
			log.Println("linuxClearTermPrintErr: ", err)
			time.Sleep(2 * time.Second)
		}
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		_, err := fmt.Print(blue("GCalc ", ver), white(" by "), blue("Mister J\n"), lightYellow("A simple terminal calculator written in "), blue("Go"), lightYellow(".\nFormatting goes {number operation number}. Example:2 + 3"))
		if err != nil {
			log.Println("windowsClearTermPrintErr: ", err)
			time.Sleep(2 * time.Second)
		}
	default:
		_, err := fmt.Println("GCalc ", ver, "by Mister J\nA simple terminal calculator written in Go.\nFormatting goes {number operation number}. Example:2 + 3")
		if err != nil {
			log.Println("defaultClearTermPrintErr: ", err)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	var num1, num2 int64
	var operator string
	clearTerm()
	for {
		_, err := fmt.Println("Give an expression.")
		if err != nil {
			log.Println("Print Expression Error: ", err)
			time.Sleep(2 * time.Second)
		}
		_, err = fmt.Scanln(&num1, &operator, &num2)
		if err != nil {
			log.Println("Scan Error: ", err)
			time.Sleep(2 * time.Second)
		}
		clearTerm()
		_, err = fmt.Printf("%d %s %d = ", num1, operator, num2)
		if err != nil {
			log.Println("Print Result Error: ", err)
			time.Sleep(2 * time.Second)
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
