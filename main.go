package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	var buf bytes.Buffer
	for scanner.Scan() {
		text := scanner.Text()
		cells := strings.Split(text, "|")
		fmt.Fprintln(&buf, text)
		if len(cells) < 2 {
			continue
		}
		scanner.Scan()
		text = scanner.Text()
		fmt.Fprintln(&buf, text)
		lines := [][]string{}
		for scanner.Scan() {
			text = scanner.Text()
			cells = strings.Split(text, "|")
			if len(cells) < 2 {
				break
			}
			if strings.HasPrefix(cells[4], "http") {
				cells[4] = "<" + cells[4] + ">"
			}
			lines = append(lines, cells)
			text = ""
		}
		sort.Slice(lines, func(i, j int) bool {
			if lines[i][2] != "Etc" && lines[j][2] == "Etc" {
				return true
			}
			if lines[i][2] == "Etc" && lines[j][2] != "Etc" {
				return false
			}
			if lines[i][2] < lines[j][2] {
				return true
			}
			if lines[i][2] > lines[j][2] {
				return false
			}
			return lines[i][1] < lines[j][1]
		})
		for _, line := range lines {
			fmt.Fprintln(&buf, strings.Join(line, "|"))
		}
		if text != "" {
			fmt.Fprintln(&buf, text)
		}
	}
	err = ioutil.WriteFile("README.md", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
