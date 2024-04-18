package main

import (
	"fmt"
	"os"
	"time"
)

type Properties struct {
	ID      string   `json:"id"`
	Aliases []string `json:"aliases"`
	Tags    []string `json:"tags"`
	Date    string   `json:"date"`
}

func main() {
	timeStamp := time.Now()
	noteID := "quicknote-" + fmt.Sprint(timeStamp.Unix())

	note := fmt.Sprintf(`---
id: %s
tags:
  - quicknote
date: "%s"
---
`, noteID, timeStamp.Format("2006-01-02"))

	filePath := os.Getenv("HOME") + "/notes/quicknotes/" + noteID + ".md"

	err := os.WriteFile(filePath, []byte(note), 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(filePath)
}
