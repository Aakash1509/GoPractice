package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type person struct {
	name   string
	age    int
	colors []string
	gr     grades
}
type grades struct {
	grade  int
	course string
}

func main() {
	type book struct {
		title  string // can combine different fields of the same type on the same line
		author string //each field must be unique inside a struct
		year   int    //No comma required
	}
	b := Vertex{1, 2}
	p1 := &b
	(*p1).X = 5
	fmt.Println(b)
	v := (Vertex{1, 2})
	p := &v
	(*p).X = 1e9
	fmt.Println(v)
	myBook := book{"Do Epic Shit", "Ankur Warikoo", 2000} //We can't know if values gets swapped i.e order matters

	//Book is struct type while myBook is struct
	fmt.Println(myBook)

	bestBook := book{title: "abc", year: 1945, author: "xyz"} //Order doesn't matter as fields are specified
	fmt.Println(bestBook)

	abook := book{title: "Random"}
	fmt.Printf("%+v\n", abook) //To print both fields and their values // + modifier with %v  printed out both the field names and their values
	fmt.Printf("%#v\n", abook)

	//To retrieve
	fmt.Println(abook.title)

	//To update
	bestBook.title = "Aakash"
	fmt.Println(bestBook)

	//Comparing struct : Can be compared by ==

	anotherBook := book{"Aakash", "xyz", 1945}
	fmt.Println(anotherBook == bestBook)

	//Creating copy of struct (Using =)

	copyBook := anotherBook
	fmt.Println(copyBook)

	//Anonymous structs
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "dwe",
		age:       30,
	}
	fmt.Printf("%+v\n", diana)

	//Anonymous fields
	type Book1 struct {
		string
		float64
		bool
	}
	b1 := Book1{"vwrew", 10.2, false}
	fmt.Printf("%+v\n", b1)
	fmt.Println(b1.string)

	//Mixture is also possible
	type Employee struct {
		name   string
		salary int
		bool
	}
	e := Employee{"John", 4000, false}
	fmt.Printf("%+v\n", e)
	e.bool = true
	fmt.Printf("%+v\n", e)

	//Embedded structs

	type Contact struct {
		email, address string
		phone          int
	}
	type Employee1 struct {
		name        string
		salary      int
		contactInfo Contact
	}
	john := Employee1{
		name:   "John",
		salary: 1000,
		contactInfo: Contact{
			email:   "john.doe@gmail.com",
			address: "wvwvwr",
			phone:   1234567890,
		},
	}
	fmt.Printf("%+v\n", john)
	fmt.Printf("Employee email:%s\n", john.contactInfo.email)
	john.contactInfo.email = "johnnew.doe@gmail.com"
	fmt.Printf("Employee's new dddddemail:%s\n", john.contactInfo.email)

	type Book struct {
		string
		float64
		bool
	}
	exam := Book{"abc", 10.45, false}
	fmt.Printf("%+v\n", exam)

	str := "ţară"
	h := byte(str[0])

	g := "t"
	k := rune(g[0])
	fmt.Println(string(k))
	fmt.Println(h)
	for i, r := range str {
		fmt.Printf("%d:%c\n", i, r)
	}

	tour := person{
		name:   "Aakash",
		age:    21,
		colors: []string{"red", "yellow", "blue"},
		gr:     grades{grade: 85, course: "Python"},
	}
	fmt.Println(tour)

	tour.gr.grade = 100
	fmt.Println(tour)
	fmt.Printf("%+v\n", tour)

}
