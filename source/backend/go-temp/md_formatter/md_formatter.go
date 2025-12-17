package md_formatter

import (
	"fmt"
)

func CreateHeading(heading string) string {
	out := fmt.Sprintf("# %v", heading)
	return out
}

func CreateListItem(listText string) string {
	out := fmt.Sprintf("- [ ] %v", listText)
	return out
}

func ToggleListItem(index int) error {
}
