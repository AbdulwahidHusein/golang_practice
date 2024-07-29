package main

import (
	"fmt"
	"fundamentals/utils"
)

func main() {

	var word string
	fmt.Println("enter a word to check if it is palingromic and count the frequency of each letter: ")
	fmt.Scan(&word)

	if utils.CheckPalindrome(word) {
		fmt.Println("the word is palingromic")
	}
	if !utils.CheckPalindrome(word) {
		fmt.Println("the word is not palingromic")
	}

	frequency := utils.CountFrequency(word)
	fmt.Println("frequency of words : ", frequency)

}
