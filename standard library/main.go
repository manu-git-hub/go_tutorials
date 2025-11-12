package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello, World!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "World", "Everyone"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "World"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(strings.Join(names, ", "))

}