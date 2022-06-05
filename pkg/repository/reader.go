package repository

import (
	"bufio"
	"errors"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

const FILEPATH = "config/users.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetUserFromFile(userId int) (string, error) {
	file, err := os.Open(FILEPATH)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sliceSplit := strings.Split(line, ",")
		if len(sliceSplit) < 2 {
			return "", errors.New("Error reading file: the file format is invalid")
		}
		//first element is the id, second the name, third the age
		idFile, _ := strconv.Atoi(sliceSplit[0])
		if userId == idFile && len(sliceSplit) > 1 {
			return sliceSplit[1], nil
		}
	}
	return "", errors.New("User not found")
}
