package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)

}

func openFile(filename string) []string {
	fmt.Printf("Going to read %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		// err is printable
		fmt.Println("Error:", err)
		return nil
	}

	// automatically call Close() at the end of the current method
	defer file.Close()

	reader := csv.NewReader(file)
	lineCount := 0
	for {
		// read just one record
		record, err := reader.Read()
		// end of file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		// record is an array of string so is directly printable
		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		// and we can iterate on top of that
		for i := 0; i < len(record); i++ {
			fmt.Print(" ", record[i])
		}
		fmt.Println()
		lineCount++
	}

	l := []string{
		"location1",
		"location2",
	}
	return l
}

/**
 * Example from https://blog.tutum.co/2015/01/27/getting-started-with-golang-on-docker/
 */
func main() {
	args := os.Args
	if len(args) > 1 {
		locations := openFile(args[1])
		fmt.Println(locations)
	} else {
		fmt.Println("Sorry, you forgot to give a file name as input.")
	}

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
