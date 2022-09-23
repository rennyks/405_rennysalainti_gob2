package main

import (
	"fmt"
	"strings"
)

func main() {

	var studentLists = []string{"Airell", "nanda", "mailo", "schannel", "marco"}
	
	var find = findStudent{studentLists}

	fmt.Println(find("airell"))
}

func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = 1
				break
			}
		}
		if student == "" {
		return fmt.Sprintf("%s does'nt exist!!!", s)
	}
	return fmt.Sprintf("we found %s at position %d", s, position+1)
}
}