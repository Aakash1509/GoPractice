package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 6

	f := 4.56

	str := fmt.Sprintf("%f", f)

	fmt.Println(str)

	s := strconv.Itoa(i)

	fmt.Printf("%T\n", s)
	fmt.Println(s)

	astring := "-17.6"

	ans, err := strconv.ParseFloat(astring, 64)

	switch v := i.(type) {
	case int:
		fmt.Println(v)
	}

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)

	var r rune = 'f'

	fmt.Printf("%T\n", r)
	fmt.Printf("%c", r)

	map1 := map[string]interface{}{
		"Aakash": 34,
		"rrr":    45.6,
	}

	clone := map[string]interface{}{}

	for k, v := range map1 {
		clone[k] = v
	}

	fmt.Println(clone)

	map1["rrr"] = 44

	fmt.Println(map1)
	fmt.Println(clone)

	sl := []int{1, 2, 3}
	fmt.Printf("%T\n", sl)

	slic := make([]int, len(sl), len(sl)+1)
	fmt.Printf("%#v\n", slic)

	de := 4133535.35135353315

	ans1 := strconv.FormatFloat(de, 'f', 4, 64)

	fmt.Println(ans1)
}
