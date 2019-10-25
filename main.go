package main

import (
	"fmt"
	"strings"
)

func add(i, j int) int {
	return i + j
}

func getUserListSQL(username, email string) string {
	sql := "select * from user"
	where := []string{}

	if username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", username))
	}

	if email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

func main() {
	fmt.Println(getUserListSQL("lunwaac", ""))

	fmt.Println(getUserListSQL("lunwaac", "lunwaac@gmail.com"))

}
