package main

import "fmt"

//PersonInfo is an detail of man
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	// instert some data into personDB map
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101"}

	// search "123" data in map PersonInfo
	person, ok := personDB["1234"]

	// ok is bool type return, if true indecate find the corspond value
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

}
