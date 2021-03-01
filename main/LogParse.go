package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Log struct {
	Timestamp time.Time
	Username  string
	Operation string
	Size      int
}

func main() {
	logs := csvReader()
	processLogs(logs)
}

func csvReader() []Log {
	logs := make([]Log, 0)
	recordFile, err := os.Open("./main/server_log.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	defer recordFile.Close()

	reader := csv.NewReader(recordFile)
	records, _ := reader.ReadAll()

	for _, record := range records {
		var data Log
		tstamp, err := time.Parse(time.UnixDate, record[0])
		if err != nil {
			fmt.Println("An error encountered parsing timestamp::", err)
		}

		size, err := strconv.Atoi(record[3])
		if err != nil {
			fmt.Println("An error encountered parsing size::", err)
		}

		data.Timestamp = tstamp
		data.Username = record[1]
		data.Operation = record[2]
		data.Size = size

		logs = append(logs, data)
	}
	return logs
}

func processLogs(logs []Log) {

	q1 := make(map[string]int)
	q2 := 0
	q3 := 0

	check1, _ := time.Parse(time.UnixDate, "Tue Apr 14 23:59:59 UTC 2020")
	check2, _ := time.Parse(time.UnixDate, "Thu Apr 16 00:00:00 UTC 2020")

	for _, log := range logs {
		if val, ok := q1[log.Username]; ok {
			q1[log.Username] += val
		} else {
			q1[log.Username] = 1
		}

		if log.Operation == "upload" && log.Size > 50 {
			q2 += 1
		}

		if log.Username == "jeff22" && check1.Before(log.Timestamp) && check2.After(log.Timestamp) {
			q3 += 1
		}
	}

	fmt.Println("Unique Users : " + strconv.Itoa(len(q1)))
	fmt.Println("Uploads larger than 50kb : " + strconv.Itoa(q2))
	fmt.Println("User 'jeff22' on Apr 15th : " + strconv.Itoa(q3))
}
