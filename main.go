package main

import (
	"os"
	"fmt"
)
const help = `Tell me what to scrape and where to put the results:
	Example:
		cex dvd dvd.csv 
			Scrapes products in the Film & Tv category and saves the results in a file name dvd.csv
		cex all all.csv
			Scrapes all products and saves the results in a file name all.csv`
func main() {
	if len(os.Args) < 2 || len(os.Args) < 3 {
		fmt.Println(help)
		os.Exit(0)
	}
	if os.Args[1] == "dvd" {
		ScrapeFilmAndTv(os.Args[2])
	}
	
	if os.Args[1] == "all" && len(os.Args) > 2 {
		ScrapeAll(os.Args[2])
	}
}