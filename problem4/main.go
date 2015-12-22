package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func Mine(secret string, rx string) int {
	regex := regexp.MustCompile(rx)
	var hash string
	for i := 0; i < 1000000000; i++ {
		hash = secret + strconv.Itoa(i)
		if regex.MatchString(getMD5Hash(hash)) {
			return i
		}
	}
	return -1
}

func main() {
	path, _ := filepath.Abs("input")
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println("The first solution is", Mine(scanner.Text(), "^00000"))
		fmt.Println("The second solution is", Mine(scanner.Text(), "^000000"))
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
