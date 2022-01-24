package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 5
const delay = 5

func main() {
	wellcomeMessage()

	for {
		menu()
	}
}

func wellcomeMessage() {

	fmt.Println("")
	name, version := returnNameAndAge()
	fmt.Println("Hello Sr.", name)
	fmt.Println("This program is in version", version)
	fmt.Println("")

}

func menu() {
	fmt.Println("1- Start the program")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	option := readOption()

	switch option {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 0:
		fmt.Println("Exit")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
		menu()
	}
}

func readOption() int {
	var optionSelected int
	fmt.Scan(&optionSelected)
	fmt.Println("")
	fmt.Println("Option selected:", optionSelected)
	fmt.Println("")

	return optionSelected
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	fmt.Println("")

	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {

		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")

}

func returnNameAndAge() (string, float64) {
	name := "Gustavo"
	version := 1.1
	return name, version
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "- Site is up!")
		writeLog(site, true)
	} else {
		fmt.Println("Site:", site, "Site is down! code:", resp.StatusCode)
		writeLog(site, false)

	}
}

func readSitesFromFile() []string {
	var sites []string

	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(archive)
	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}
	}
	archive.Close()

	return sites
}

func writeLog(site string, status bool) {

	archive, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}
	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " + strconv.FormatBool(status) + "\n")
	archive.Close()

}

func showLogs() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(archive))
}
