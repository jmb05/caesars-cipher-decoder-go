package main

import (
	"fmt"
	"os"
	"jmb19905.net/styling"
)

func main() {	
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments supplied")
		return
	} else if len(os.Args) < 2 {
		fmt.Println("Please provide a ciphered text as argument")
		return
	}

	infoCopyrightAndWarranty := styling.ColorBold("white", `
Caesar's Cipher cracker

Command Line Tool that can decode text that is encoded with the caesar's cipher

Copyright (C) 2021-2022, Jared M. Bennett

This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.
	`);

	fmt.Println(infoCopyrightAndWarranty)
	fmt.Println("Ciphered text is", styling.Color("red", os.Args[1]))

	for !decipher([]byte(os.Args[1])) {
		fmt.Println("Cycled through whole alphabet")
		fmt.Println("Retry?")
		if !ReadYesNo() {
			break
		}
		fmt.Println("Retrying...")
	}
}

func decipher(input []byte) bool {	
	for i := 0; i < 25; i++ {
		for j := 0; j < len(input); j++ {
			currentCh := CharToUppercase(input[j])
			if currentCh >= 65 && currentCh <= 90 {
				if currentCh == 65 {
					input[j] = byte(currentCh + 25)
				} else {
					input[j] = byte(currentCh - 1)
				}
			}	
		}
		fmt.Println("Is this readable?\n", string(input))

		if(ReadYesNo()) {
			fmt.Println("Decoded:", styling.Color("green", string(input)))
			fmt.Println("The alphabet was shifted", i + 1 , "places")
			return true
		}
	}
	return false
}

func CharToUppercase(char byte) byte {
	if(char >= 'a' && char <= 'z') {
		char = char - 32
	}
	return char
}

func ReadYesNo() bool {
	fmt.Print("[y/N] ")
	var buf string
	fmt.Scanln(&buf)
	if(buf == "y" || buf == "Y") {
		return true
	}
	return false
}
