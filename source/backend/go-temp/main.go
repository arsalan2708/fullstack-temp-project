package main

import (
	"fmt"

	"md.cli/md_formatter"
)

func main() {
	heading := md_formatter.CreateHeading("Test Title")
	listItem := md_formatter.CreateListItem("First List Item")

	fmt.Println(heading)
	fmt.Println(listItem)
}
