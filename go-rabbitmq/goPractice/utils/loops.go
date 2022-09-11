package utils

import "fmt"

func Loops() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and day is %v \n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at lco")

}
