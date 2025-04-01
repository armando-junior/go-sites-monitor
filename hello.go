package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorCount = 3
const monitorInterval = 10 * time.Second

func main() {

	introduction()
	for {
		showMenu()
		option := readInput()
		selectOption(option)
	}

	// if option == 1 {
	// 	fmt.Println("Monitoring started")
	// } else if option == 2 {
	// 	fmt.Println("Logs displayed")
	// } else if option == 0 {
	// 	fmt.Println("Exiting program")
	// } else {
	// 	fmt.Println("Invalid option")
	// }
}

func introduction() {
	name, version := returnNameAndVersion()

	fmt.Println("Hello, Mr.", name, "!")

	fmt.Println("This program is version:", version)
}

func returnNameAndVersion() (string, float32) {
	return "Armando Junior", 1.0
}

func readInput() int {
	var option int
	fmt.Scan(&option)

	return option
}

func selectOption(option int) {
	switch option {
	case 1:
		monitor()
	case 2:
		fmt.Println("Logs displayed")
	case 0:
		fmt.Println("Exiting program")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
		os.Exit(-1)
	}
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Display Logs")
	fmt.Println("0- Exit Program")
}

func monitor() {
	fmt.Println("Monitoring...")
	sites := readFromFile()

	for i := 0; i < monitorCount; i++ {
		validateStatus(sites)
		fmt.Println("Waiting", monitorInterval, "seconds to check again...")
		fmt.Println("")
		time.Sleep(monitorInterval)
	}

}

func validateStatus(sites []string) {
	for _, site := range sites {
		response, err := http.Get(site)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}

		if response.StatusCode == 200 {
			fmt.Println("Site:", site, "was loaded successfully")
		} else {
			fmt.Println("Site:", site, "is not running")
		}
	}
}

func readFromFile() []string {
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	reader := bufio.NewReader(file)

	var sites []string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return nil
		}
		sites = append(sites, strings.TrimSpace(line))
	}

	file.Close()
	return sites
}
