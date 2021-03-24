package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type numAdderMapping struct {
	Number   int64
	AdderMap map[int64]int64
}

var numberIndex = make(map[int64]*numAdderMapping)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n, n)
	maxElement := int64(0)
	for i := 0; i < len(queries); i++ {
		startIndex := queries[i][0] - 1
		endIndex := queries[i][1] - 1
		adder := int64(queries[i][2])

		for j := startIndex; j <= endIndex; j++ {
			var sum, tempSum int64
			var ok bool
			//arr[j] = arr[j] + adder
			numberToAdderMappingPresent := false
			adderToNumberMappingPresent := false
			numberMapping, adderMapping := getNumberAndAdderFromNumberIndex(arr[j], adder)
			if adderMapping == nil && numberMapping == nil {
				arr[j] = createIndex(arr[j], adder)
				continue
			}
			if numberMapping != nil {
				tempSum, ok = numberMapping.AdderMap[adder]
				if ok {
					//arr[j] = tempSum
					sum = tempSum
					numberToAdderMappingPresent = true
				}
			}

			if adderMapping != nil {
				tempSum, ok = adderMapping.AdderMap[arr[j]]
				if ok {
					adderToNumberMappingPresent = true
					if !numberToAdderMappingPresent {
						sum = tempSum
					}
				}

			}

			if !numberToAdderMappingPresent && adderToNumberMappingPresent {
				addToIndex(arr[j], adder, sum)
			}
			if !adderToNumberMappingPresent && numberToAdderMappingPresent {
				addToIndex(adder, arr[j], sum)
			}
			if !numberToAdderMappingPresent && !adderToNumberMappingPresent {
				sum = arr[j] + adder
				addToIndex(arr[j], adder, sum)
				addToIndex(adder, arr[j], sum)
				//arr[j] = sum
			}
			arr[j] = sum

		}
	}
	for i := int32(0); i < n; i++ {
		if maxElement < arr[i] {
			maxElement = arr[i]
		}
	}
	return maxElement
}

func getNumberAndAdderFromNumberIndex(number, adder int64) (*numAdderMapping, *numAdderMapping) {
	var numberMapping, adderMapping *numAdderMapping
	numberMapping, _ = numberIndex[number]
	adderMapping, _ = numberIndex[adder]
	return numberMapping, adderMapping
}

func addToIndex(number, adder, sum int64) {
	index, numberIndexPresent := numberIndex[number]
	if numberIndexPresent {
		_, ok := index.AdderMap[adder]
		if !ok {
			index.AdderMap[adder] = sum
		}
		return
	}
	numberMapping := &numAdderMapping{
		Number:   number,
		AdderMap: map[int64]int64{adder: sum},
	}
	numberIndex[numberMapping.Number] = numberMapping

}

func createIndex(number, adder int64) int64 {
	sum := number + adder
	numberMapping := &numAdderMapping{
		Number:   number,
		AdderMap: map[int64]int64{adder: sum},
	}
	adderMapping := &numAdderMapping{
		Number:   adder,
		AdderMap: map[int64]int64{number: sum},
	}
	numberIndex[numberMapping.Number] = numberMapping
	numberIndex[adderMapping.Number] = adderMapping
	return sum
}

func main() {
	infutFile, _ := os.Open("input02.txt")
	reader := bufio.NewReaderSize(infutFile, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
