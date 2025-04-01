package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
		readLog()
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

		if response.StatusCode >= 200 && response.StatusCode < 300 {
			fmt.Println("Site:", site, "was loaded successfully")
		} else {
			fmt.Println("Site:", site, "is not running")
			writeLog(site, response.StatusCode)
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

func writeLog(site string, status int) {
	// os.O_CREATE: create the file if it doesn't exist
	// os.O_RDWR: open the file for reading and writing
	// os.O_APPEND: append to the file
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " +
		strconv.Itoa(status) + "\n")
	writer.Flush()

	file.Close()
}

func readLog() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(strings.TrimSpace(line))
	}

	file.Close()
}
