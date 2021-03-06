package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

func ReadLinesInFile(fileName string) ([]string, error) {
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("%s, %s", ErrReadFile, err)
	}

	return lines, nil
}

func ReadCharsInFile(fileName string, opts []string) ([][]string, error) {
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	buf := make([]byte, 1)

	var chars [][]string
	var line []string
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New(ErrReadFile)
		}

		bufData := buf[:n]
		data := string(bufData)
		if Exists(opts, data) {
			line = append(line, data)
		} else {
			os := runtime.GOOS
			switch os {
			case "linux":
				if !(bytes.Compare(bufData, []byte{13}) == 0 ||
					bytes.Compare(bufData, []byte{10}) == 0) {
					continue
				}

			case "windows":
				if !(bytes.Compare(bufData, []byte{13}) == 0) {
					continue
				}
			}
			chars = append(chars, line)
			line = []string{}
		}
	}
	chars = append(chars, line)

	return chars, nil
}

func ReadNumsInFile(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(ErrOpenFile)
	}
	defer file.Close()

	buf := make([]byte, 1)

	var nums [][]int
	var line []int
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New(ErrReadFile)
		}

		bufData := buf[:n]
		if data, err := strconv.Atoi(string(bufData)); err == nil {
			line = append(line, data)
		} else {
			os := runtime.GOOS
			switch os {
			case "linux":
				if !(bytes.Compare(bufData, []byte{13}) == 0 ||
					bytes.Compare(bufData, []byte{10}) == 0) {
					continue
				}

			case "windows":
				if !(bytes.Compare(bufData, []byte{13}) == 0) {
					continue
				}
			}
			nums = append(nums, line)
			line = []int{}
		}
	}
	nums = append(nums, line)

	return nums, nil
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("[ERROR] %s", err)
	}

	return i
}

func GreaterInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func PrintMatrix(nums [][]int, rowLen, colLen int) {
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			fmt.Printf("%5d ", nums[i][j])
		}
		fmt.Println()
	}
}

func PrintStringMatrix(array [][]string, rowLen, colLen int) {
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			fmt.Printf("%s ", array[i][j])
		}
		fmt.Println()
	}
}

func PrintIntMap(m map[string]int) {
	log.Println("======= printm START ========")
	sortKeys := make([]string, 0, len(m))
	for k, _ := range m {
		sortKeys = append(sortKeys, k)
	}
	sort.Strings(sortKeys)
	for _, k := range sortKeys {
		fmt.Println(k, m[k])
	}
	log.Println("======= printm END ========")
}

func MaxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func MinInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Exists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if !(arr.Kind() == reflect.Array || arr.Kind() == reflect.Slice) {
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
