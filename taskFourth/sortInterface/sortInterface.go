package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

const (
	layoutISO = "2006-01-2"
)

type Person struct {
	firstName string
	lastName  string
	birthday  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}
func (p People) Less(i, j int) bool {
	if p[i].birthday.Equal(p[j].birthday) {
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthday.After(p[j].birthday)
}
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func parseDate(d string) time.Time {
	date, err := time.Parse(layoutISO, d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return date
}

func print(message string, people People) {
	fmt.Println(message)
	for i, person := range people {
		fmt.Println(i+1, person.firstName, person.lastName, person.birthday.String())
	}
}

func main() {

	people := People{
		{"Ivan", "Ivanov", parseDate("2005-08-10")},
		{"Ivan", "Ivanov", parseDate("2003-08-5")},
		{"Artiom", "Ivanov", parseDate("2005-08-10")},
		{"Pavel", "Ivanov", parseDate("2006-09-10")},
		{"Anton", "Ivanov", parseDate("2006-09-10")},
	}
	print("List: ", people)
	sort.Sort(people)
	print("Sorted list: ", people)
}
