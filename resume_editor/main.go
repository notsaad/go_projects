package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var fileName string
	fmt.Print("Enter the LaTeX file name (without extension): ")
	fmt.Scanln(&fileName)

	fmt.Println("Compiling LaTeX file:", fileName+".tex")

	cmd := exec.Command("pdflatex", fileName)
	err := cmd.Run() // Run waits for the command to finish
	// can also use cmd.CombinedOutput() to get the output of the command
	if err != nil {
		fmt.Println("Error compiling LaTeX code:", err)
		return
	}

	// Run another command after pdflatex is done
	secondCmd := exec.Command("rm", fileName+".log", fileName+".aux", fileName+".out")
	err = secondCmd.Run()
	if err != nil {
		fmt.Println("Error running second command:", err)
		return
	}

	// Print the output
	fmt.Println("PDF compiled successfully!")
}
