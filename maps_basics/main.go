package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No. of pairs:%d\n", len(employees))

	//Retrieving value of a key not present in key will give zero value
	fmt.Printf("The value for key %q is %q\n", "abc", employees["abc"])

	device := map[string]float64{}

	device1 := make(map[string]float64, 5)

	device1["device2"] = 42.0

	device["my"] = 0.5

	fmt.Println(device)

	//employees["abc"] = "xyz" //Will produce an error as map is uninitialized

	people := map[string]float64{} //Initialized map
	people["abc"] = 10.5
	people["xyz"] = 10
	fmt.Println(people)

	map1 := make(map[string]string, 5)
	map1["abc"] = "xyz"
	fmt.Println(map1)

	fmt.Println(people["efg"])

	// "comma ok" idiom is used to distinguish between a missing key:value pair and an existing key with value zero

	v, ok := people["abc"]

	if ok {
		fmt.Println("The key value is ", v)
	} else {
		fmt.Println("The key doesn't exist")
	}

	for i, j := range people {
		fmt.Printf("Key: %#v, Value: %#v\n", i, j)
	}
	delete(people, "abc")
	fmt.Println(people)

	//Comparing maps : Using Sprintf() only if both key and value is string
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Printf("Maps are equal")
	} else {
		fmt.Printf("Maps are not equal")
	}

	//** CLONING A MAP **//

	// When creating a map variable Go creates a pointer to a map header value in memory.
	// The key: value pairs of the map are not stored directly into the map.
	// They are stored in memory at the address referenced by the map header.

	friends := map[string]int{"Dan": 40, "Maria": 35}

	// neighbors is not a copy of friends.
	// both maps reference the same data structure in memory
	neighbors := friends

	// modifying friends AND neighbors
	friends["Dan"] = 30

	fmt.Println(neighbors) // -> map[Dan:30 Maria:35]

	// How to clone a map?
	// 1. initialize a new map
	colleagues := make(map[string]int, 5)

	// colleagues = friends // -> ERROR, illegal with maps!

	// 2. use a for loop to copy each element into the new map
	for k, v := range friends {
		colleagues[k] = v
	}

	// colleagues and friends are different maps in memory!

	var employees2 map[string]string
	fmt.Printf("%#v\n", employees2)

	employees1 := map[string]string{}
	fmt.Printf("%#v\n", employees1)

	var blank map[string]int
	fmt.Printf("%T %#v\n", blank, blank)

	m2 := map[int]string{
		1: "Aakash",
		2: "Dhruv",
	}
	fmt.Println(m2)

	m2[1] = "Sky"
	fmt.Println(m2)

	tour := map[int]bool{1: true, 2: true}
	fmt.Println(tour)

	delete(tour, 1)
	fmt.Println(tour)

	tour[1] = true

	for key, value := range tour {
		fmt.Println(key, value)
	}
}
