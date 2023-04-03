/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadFile(f string) []int {
	r, err := os.Open(f)
	if err != nil {
		fmt.Println("error reading the file", err)
	}
	defer r.Close()

	reader := csv.NewReader(r)
	var input []string
	var final []int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		input = record
	}

	for _, v := range input {
		i, _ := strconv.Atoi(v)
		final = append(final, i)
	}
	return final
}
