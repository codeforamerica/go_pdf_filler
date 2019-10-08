package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"github.com/jessevdk/go-flags"
)

type runOpts struct {
	InputCSV           string `long:"input-csv" description:"The files containing criminal histories from CA DOJ"`
}

var opts struct {
	Run       runOpts           `command:"run" description:"Process a CSV and output a PDF"`
}

func ReadCSV(filepath string) (string, error) {
	dojFile, err := os.Open(filepath)
	if err != nil {
		return "I'm borked", err
	}

	bufferedReader := bufio.NewReader(dojFile)
	sourceCSV := csv.NewReader(bufferedReader)
	rows, err := sourceCSV.ReadAll()
	stringToWrite := rows[0][0]
	return stringToWrite, nil
	// dump into PDF
}

func (o runOpts) Execute(args []string) error {
	string, err := ReadCSV(o.InputCSV)
	fmt.Print(string)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
