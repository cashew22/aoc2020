package utils

import (
	"bufio"
	"os"
	"strconv"
)

func GetFileContentStr(path string) []string {
	var content []string
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}

func GetFileContentByteArray(path string) [][]byte {
	var content [][]byte
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, []byte(scanner.Text()))
	}

	return content
}

func GetFileContentInt(path string) []int {
	var content []int
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		content = append(content, val)
	}

	return content
}

func GetUniqueValueMapByte(b []byte) map[byte]int {
	dict := make(map[byte]int)
	for _, num := range b {
		dict[num] = dict[num] + 1
	}
	return dict
}

func FindStringInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
