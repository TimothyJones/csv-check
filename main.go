package main

import (
	"encoding/csv"
    "log"
	"os"
    "io"
)

func main() {
	r := csv.NewReader(os.Stdin)

	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
