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

const MONITORING_PER_CYCLE = 2
const DELAY_BETWEEN_CYCLES_IN_SECONDS = 5
const INPUT_FILE_NAME = "sites.txt"
const OUTPUT_FILE_NAME = "logs.txt"
const OUTPUT_FILE_PERMS = 0666

func main() {
	displayIntroduction()
	for {
		displayMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Displaying Logs...")
			displayLogs()
		case 0:
			fmt.Println("Exiting the program")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}

}

func displayIntroduction() {
	name := "Marcos"
	version := 1.2
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is in version", version)
}

func displayMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Display Logs")
	fmt.Println("0- Exit the Program")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("The chosen command was", commandRead)
	fmt.Println("")

	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readURLsFile()

monitoringLoop:
	for i := 0; i < MONITORING_PER_CYCLE; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			err := testURL(site)
			if err != nil {
				break monitoringLoop
			}
		}
		time.Sleep(DELAY_BETWEEN_CYCLES_IN_SECONDS * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testURL(url string) *string {
	if url == "" || !strings.HasPrefix(url, "https") && !strings.HasPrefix(url, "http") {
		err := "Invalid URL"
		fmt.Println(err)
		return &err
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	var isValid bool
	if response.StatusCode == 200 {
		// fmt.Println("Site:", url, "was loaded successfully!")
		isValid = true
	} else {
		// fmt.Println("Site:", url, "has problems. Status Code:", response.StatusCode)
		isValid = false
	}

	writeLog(url, isValid)
	return nil
}

func readURLsFile() []string {
	var sites []string

	file, err := os.Open(INPUT_FILE_NAME)
	defer file.Close()

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	return sites
}

func writeLog(url string, status bool) {
	file, err := os.OpenFile(OUTPUT_FILE_NAME, os.O_RDWR|os.O_CREATE|os.O_APPEND, OUTPUT_FILE_PERMS)
	defer file.Close()

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	statusVerb := "was"
	if !status {
		statusVerb = "wasn't"
	}

	date := time.Now().Local().Format("2006-01-02 15:04:05")

	stringToWrite := fmt.Sprintf("at %s the site %s %s online\n", date, url, statusVerb)
	// url + " - online: " + strconv.FormatBool(status) + "\n"

	file.WriteString(stringToWrite)
}

func displayLogs() {
	file, err := os.ReadFile(OUTPUT_FILE_NAME)

	if err != nil {
		fmt.Println("An error occurred", err)
	}

	fmt.Println(string(file))
}
