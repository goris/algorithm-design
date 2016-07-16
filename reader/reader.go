package reader

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts(dir string) ([]int, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, num)
        //fmt.Println("Read: ", num)
	}

	return lines, scanner.Err()

}
