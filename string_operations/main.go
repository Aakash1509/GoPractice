package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println

	result := strings.Contains("Hello World", "Hello")
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'x')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("dwdwergAAA"))
	p(strings.ToUpper("Go Hello"))

	p("Go" == "go")

	//Use EqualFold when letter case is not taken into account
	p(strings.EqualFold("Go", "go"))

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2) //Will replace first 2 occurrences
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1) //Will replace all occurrences
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a,b,c,", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "golang"}
	myStr = strings.Join(s, "-")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t Goodby Windows, Welcome Linux! \n")
	fmt.Println(s1)

	s2 := strings.Trim("...Hello World !!!??", ".!?")
	fmt.Println(s2)
}
