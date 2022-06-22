package main

import (
	"bufio"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
)

var mostVisited []string
var allTabs []string

func init() {

	allTabs = getAllTabs("tab_list.txt")
}

func main() {

}

func checkWWW(link string) string {
	link = strings.Replace(link, "/", "", 1) //remove '/' rune on the start of every cell, dunno why that happens
	if strings.HasPrefix(link, "www.") {
		return link
	}
	return "www." + link
}

func getAnalyticsTabs(spreadsheetName string) (tabs []string) {
	f, err := excelize.OpenFile("analytics.xlsx")
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
		tabs = append(tabs, checkWWW(row[0]))
	}

	return tabs
}

func getAllTabs(filename string) (lines []string) {
	f, err := os.Open("tab_list.txt")
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
