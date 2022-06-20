package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.FieldsPerRecord = -1
	var indexColumn = flag.String("i", "", "indexColumn column")
	flag.Parse()

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var indexColumnIndex = 0
	var header = records[0]
	for headerIndex, headerItem := range header {
		if *indexColumn == headerItem {
			indexColumnIndex = headerIndex
		}
	}
	records = records[1:]
	for lineIndex, line := range records {
		for headerIndex, word := range line {
			if *indexColumn != "" {
				var indexKey = line[indexColumnIndex]
				fmt.Printf("/%s/%s = %s\n", indexKey, header[headerIndex], word)
			} else {
				fmt.Printf("/%d/%s = %s\n", lineIndex, header[headerIndex], word)
			}
		}
	}
}
