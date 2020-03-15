package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	email1 := "aaaaa.bbbb.ccc+test@gmail.com"
	email2 := "aaaaa.bbbbccc+test@gmail.com"
	email3 := "aaaaabbbb.ccc@gmail.com"
	email4 := "a@gmail.com"
	fmt.Println(numUniqueEmails([]string{email1, email2, email3, email4}))
}

func numUniqueEmails(emails []string) int {
	resultMap := make(map[string]string)
	for _, v := range emails {
		email := removeDot(ignoreStringAfterPlus(v))
		resultMap[email] = ""
	}
	return len(resultMap)
}

func removeDot(email string) string {
	splited := strings.Split(email, "@")
	splited[0] = strings.Replace(email, ".", "", -1)
	converted := splited[0] + "@" + splited[1]
	return converted
}

func ignoreStringAfterPlus(email string) string {
	rep := regexp.MustCompile(`\+.*@`)
	converted := rep.ReplaceAllString(email, "@")

	return converted
}
