package main

import (
	"bufio"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
)

var mostVisited map[string]struct{}
var allTabs []string

func init() {
	allTabs = getAllTabs("tab_list.txt")
	mostVisited = make(map[string]struct{})
	setAnalyticsTabs("analytics.xlsx", mostVisited)
}

//TODO: Check if a tab from allTabs exists into mostVisited. If not, mark it to remove later.
func main() {

}

func checkWWW(link string) string {
	link = strings.Replace(link, "/", "", 1) //remove '/' rune on the start of every cell, dunno why that happens
	if strings.HasPrefix(link, "www.") {
		return link
	}
	return "www." + link
}

func setAnalyticsTabs(spreadsheetName string, table map[string]struct{}) {
	f, err := excelize.OpenFile(spreadsheetName)
	if err != nil {
		log.Fatalf("error opening spreadsheet file: %s\n", err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatalf("error closing file: %s\n", err)
		}
	}()

	rows, err := f.GetRows("Dataset1") //GetRows loads all rows into the memory, excluding empty rows
	if err != nil {
		log.Fatalf("error getting rows: %s\n", err)
	}

	for _, row := range rows[1:] { //we skip the spreadsheet header
		table[checkWWW(row[0])] = struct{}{}
	}
}

func getAllTabs(filename string) (lines []string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatalf("error closing file: %s\n", err)
		}
	}()

	bf := bufio.NewScanner(f)
	for bf.Scan() {
		lines = append(lines, bf.Text())
	}

	return lines
}
