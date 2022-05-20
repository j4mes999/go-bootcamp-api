package fileReader

import (
	_ "bufio"
	"fmt"
	"os"
)

const FILEPATH = "config/users.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetUserFromFile(userId int) string {
	dat, err := os.ReadFile(FILEPATH)
	check(err)
	fmt.Printf(string(dat))
	return "holi"
}
