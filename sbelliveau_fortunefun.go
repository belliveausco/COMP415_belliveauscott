package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	name := make(chan bool)
	go fortuneFunction(name)
	//var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like a fortune?")
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\r\n", "", -1)
		if input == "yes" {
			name <- true
		} else if input == "Yes" {
			name <- true
		} else if input == "YEs" {
			name <- true
		} else if input == "YES" {
			name <- true
		} else if input == "yEs" {
			name <- true
		} else if input == "YeS" {
			name <- true
		} else if input == "yES" {
			name <- true
		} else if input == "yeS" {
			name <- true
		} else if input == "no" {
			close(name)
			break
		} else if input == "No" {
			close(name)
			break
		} else if input == "NO" {
			close(name)
			break
		} else if input == "nO" {
			close(name)
			break
		} else {
			fmt.Println("Try again")
			fmt.Println("Would you like a fortune?")
		}
	}
}
func fortuneFunction(name <-chan bool) {
	fileAsBytes, err := ioutil.ReadFile("fortunes.txt")
	if err != nil {
		log.Fatalln("Error reading file", err)
	}
	fileContents := string(fileAsBytes)
	s := strings.Split(fileContents, "%%")

	for {
		<-name
		var a [1]int
		rand.Seed(time.Now().UnixNano())
		//for i := 0; i < 1; i++ {
		a[0] = rand.Intn(1689)
		fmt.Println(s[a[0]])
		fmt.Println("Would you like a fortune?")
	}
}

//The main function will
//1. create a channel and then spawn the fortune function as a go routine taking the channel as a
//parameter.
//2. Enter a forever loop and ask the user if they want another fortune. If the user says yes (with any
//capitalization) then send a message down the channel to the for the fortune function to select a
//fortune and display it.
//3. If the user answers no (with any capitalization), then end the program
//4. if the user answers anything else, then ask again
//The fortune function will
//1. Open the fortunes.txt file that I sent last class in your email and read it in.
//2. split the contents of the fortunes file on the %% string to get a slice of strings
//3. loop forever
//1. wait for a message on the channel
//2. when a message is received, randomly select one of the fortunes from your slice
//3. and print it to the screen.
