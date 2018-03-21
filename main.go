package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hpcloud/tail"
)

const TEST_FILE = "./test-logs.txt"
const LOGS = "./logs.txt"

func main () {
	fmt.Println("This is the test mine")

	t, _ := tail.TailFile(
		TEST_FILE, tail.Config{Follow: true})
	for line := range t.Lines {
		fmt.Println(line.Text)
		appendStringToFile(LOGS, line)
		time.Sleep(time.Second)
	}
}

func appendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
