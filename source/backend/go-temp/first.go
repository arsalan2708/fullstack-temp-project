package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

// func main() {
// 	var sum int
// 	fmt.Print("First test message:: ")
// 	fmt.Println(quote.Go())
// 	fmt.Println(math.Abs(5 - 9))

// 	name := "Arsalan"
// 	a := 5
// 	b := 34
// 	sum = a + b

// 	fmt.Println(greeting(name))
// 	fmt.Printf("Sum total: %d\n", sum)

// 	arrayInput := [5]int{2, 5, 4, 10, 13}
// 	sum2 := sumArray(arrayInput[:])

// 	fmt.Printf("Sum total of array %v: %d\n", arrayInput, sum2[len(sum2)-1])

// 	write_a_file("new_file2.md", "# This is a title of 2nd File\n\n- [ ] first step\n\n- [ ] second step")
// 	read_file("new_file2.md")

// 	list_files()
// 	readUserInput("This is the first Line.")
// }

func greeting(name string) string {
	message := fmt.Sprintf("Hello %v. Welcome to new go project!", name)
	return message
}

func sumArray(input []int) []int {
	total := 0
	fmt.Printf("sumArray:: input len: %v ; input cap: %v\n",
		len(input), cap(input))

	for _, val := range input {
		total = total + val
	}

	input = append(input, total)
	return input
}

func write_a_file(filename string, inputData string) {
	data := []byte(inputData)
	filePath := fmt.Sprintf("./storage/%v", filename)

	os.Mkdir("./storage", 0755)
	f := os.WriteFile(filePath, data, 0644)

	fmt.Println(f)
}

func read_file(filename string) {
	storage := "./storage"
	path := fmt.Sprintf("%v/%v", storage, filename)

	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	fields := strings.Split(string(data), "\n\n")

	fmt.Printf("number of lines: %v\n", len(fields))

}

func list_files() {
	storage := "./storage"
	files, err := os.ReadDir(storage)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name()[:len(file.Name())-3])
	}
}

// look into "huh" TUI for user editable fields
func readUserInput(def_val string) {

	huh.NewInput().
		Title("Edit input: ").
		Value(&def_val).
		Run()
}
