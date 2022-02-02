package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

const (
	NewTime      = 8
	DefaultTime  = 6
	DefaultRound = 18
)

func readNumsInFile(fileName string) ([]int, error) {
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	var nums []int
	for fileScanner.Scan() {
		strNums := strings.Split(fileScanner.Text(), ",")
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("%s, %s", ErrReadFile, err)
	}

	return nums, nil
}

func solve(nums []int) int {
	answer := 9999999999
	len := len(nums)
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	for i := 1; i <= max; i++ {
		cur := i
		sum := 0

		for j := 0; j < len; j++ {
			// log.Print("cur, nums[j]", cur, nums[j])
			curpoint := (int(math.Abs(float64(nums[j] - cur))))
			for point := 1; point <= curpoint; point++ {
				sum += point
			}
			// log.Print("SUM", sum)
		}

		if answer > sum {
			answer = sum
		}
	}

	return answer
}

func main() {
	const (
		fileName = "input2"
	)

	positions, err := readNumsInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	fmt.Print(solve(positions))

}
