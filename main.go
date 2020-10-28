package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el numero de lineas a escribir")
	scanner.Scan()
	numberOfLinesString := scanner.Text()
	numberOfLines, err := strconv.ParseInt(numberOfLinesString, 10, 32)
	if err != nil {
		fmt.Println("Could not convert string to number", err)
		return
	}

	linesOfData := make([]string, numberOfLines)
	for i := int64(0); i < numberOfLines; i++ {
		fmt.Println("Ingrese la linea ", i+1)
		scanner.Scan()
		linesOfData[i] = scanner.Text()
	}

	sortedData := sort.StringSlice(linesOfData)

	sort.Sort(sortedData)

	file, err := os.OpenFile("./asecendente.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Could not open file with name asecendente", err)
		return
	}

	for _, line := range sortedData {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Could not write line to file", err)
			return
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Could not close file asecendente", err)
		return
	}

	sort.Sort(sort.Reverse(sortedData))

	file, err = os.OpenFile("./descendente.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Could not open file descendente", err)
		return
	}

	for _, line := range sortedData {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Could not write line to file descendente", err)
			return
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Could not close file descendente", err)
		return
	}
}
