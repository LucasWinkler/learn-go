package main

import "fmt"

func main() {
	// var intArr [3]int32
	// intArr[0] = 1
	// intArr[1] = 2
	// intArr[2] = 3

	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("Length: %v, Capacity: %v, Value: %v\n", len(intSlice), cap(intSlice), intSlice)

	intSlice = append(intSlice, 7)
	fmt.Printf("Length: %v, Capacity: %v, Value: %v\n", len(intSlice), cap(intSlice), intSlice)

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("Length: %v, Capacity: %v, Value: %v\n", len(intSlice), cap(intSlice), intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("Length: %v, Capacity: %v, Value: %v\n", len(intSlice3), cap(intSlice3), intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Adam"]
	fmt.Println(age, ok)

	// delete(myMap2, "Adam")

	if ok {
		fmt.Printf("Age is %v\n", age)
	} else {
		fmt.Println("Name not found")
	}

	// for name := range myMap2 {
	// 	fmt.Printf("Name: %v\n", name)
	// }

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	for i := range 10 {
		fmt.Println(i)
	}
}
