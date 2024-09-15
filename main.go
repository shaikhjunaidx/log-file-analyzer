package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	logFilePath := flag.String("file", "", "Path to log file")
	urlEndPoint := flag.String("url", "", "URL endpoint to search for")
	flag.Parse()

	if *logFilePath == "" || *urlEndPoint == "" {
		log.Fatal("Both file path and URL endpoint must be provided")
	}

	processLogFile(*logFilePath, *urlEndPoint)

}

func processLogFile(filePath, endpoint string) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ipSet := make(map[string]struct{})

	re := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s\w+\s` + regexp.QuoteMeta(endpoint) + `\s`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if len(matches) > 1 {
			ip := matches[1]
			ipSet[ip] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unique IP addresses that accessed the URL endpoint:")

	for ip := range ipSet {
		fmt.Println(ip)
	}
}
