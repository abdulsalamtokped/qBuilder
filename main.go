package main

import (
	"fmt"

	"github.com/abdulsalam/qBuilder/db"
)

func main() {
	qHelper := db.GetInstance()
	query := qHelper.Select("id, name, avatar").From("users").Where("id", "=", 1).Where("name", "!=", "Abdul").Build()

	fmt.Println(query)
}
