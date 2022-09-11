package utils

import (
	"fmt"
	"sort"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func Pointers() {
	fmt.Println("Studying about pointers")

	// var ptr *int
	// fmt.Println("value of pointer is", ptr)

	myNumber := 23
	var myNumberPointer = &myNumber
	fmt.Println(myNumberPointer)
	fmt.Println(*myNumberPointer)

	*myNumberPointer = *myNumberPointer * 2
	fmt.Println(*myNumberPointer)
}

func Arrays() {
	fmt.Println("Classes for arrays")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "kiwi"
	// fruitList[2] = "tomato"
	fruitList[3] = "mango"

	fmt.Println("List of fruits are", fruitList)
	fmt.Println("Length of fruit", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie List", vegList)
	fmt.Println("Length of veggies", len(vegList))
}

func Slices() {
	fmt.Println("Studying slices")

	var fruitList = []string{"mango", "banana", "apple", "orange"}
	fmt.Printf("Type of fruitlist is %T \n", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList, "peaches", "kiwi")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 344
	highScores[2] = 123
	highScores[3] = 350

	highScores = append(highScores, 450, 340)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// remove value from slices based on index

	var courses = []string{"reactjs", "go", "javascript", "swift", "python"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

func Maps() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	languages["ts"] = "typescript"

	fmt.Println("List of languages", languages)
	fmt.Println("JS Short for", languages["js"])

	delete(languages, "rb")
	fmt.Println(languages)

	// Looping in maps
	for key, value := range languages {
		fmt.Printf("For key %v value %v\n", key, value)
	}
}

func Structs() {
	jane := User{"jane", "jane@test.dev", true, 23}
	fmt.Printf("Jane details are: %+v\n", jane)
	fmt.Printf("Jane's name is: %v\n", jane.Name)
}
