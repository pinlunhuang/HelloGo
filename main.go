package nonmain

import (
	"fmt"
	"strings"
)

func add(i, j int) int {
	return i + j
}

type searchOpts struct {
	username string
	email    string
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}

	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}

	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

func main() {
	// fmt.Println(getUserListOptsSQL("lunwaac", ""))

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "lunwaac",
	}))

	// fmt.Println(getUserListOptsSQL("lunwaac", "lunwaac@gmail.com"))
	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "lunwaac",
		email:    "lunwaac@gmail.com",
	}))

}
