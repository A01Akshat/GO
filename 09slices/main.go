package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{"Grapes", "Pappaya", "Avocado"}
	fmt.Println("List:", fruits)

	fruits = append(fruits, "Watermelon", "Kiwi")
	fmt.Println("Updated List:", fruits)

	fruits = append(fruits[2:4])
	fmt.Println((fruits))

	// Something Different

	score:=make([]int,4)
	score[0]=654
	score[1]=211
	score[2]=476
	score[3]=195

	fmt.Println(score)

	// Now if I do this
	// score[4]=789
	// fmt will give error obvio out of range
	// But if I do this:

	score=append(score, 321,541,90)
	fmt.Println(score) // No error, append changes mem allocation

	sort.Ints(score) // sort the slice in ascending order

	var a bool=sort.IntsAreSorted(score) // gives boolean value whether slices is sorted or not
	fmt.Println(a)

    // REMOVE VALUE BASIS OF INDEX

	var courses=[]string{"C","Python","Java","Go"}
	index:=1
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)


}
