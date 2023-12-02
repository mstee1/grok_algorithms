package algorithms

import "fmt"

type person struct {
	name     string
	friends  []string
	searched bool
	seller   bool
}

func Breadth_first_search() {

	peoples := make(map[string]person)

	peoples["you"] = person{
		name: "you",
		friends: []string{
			"alice",
			"bob",
			"claire",
		},
		searched: false,
		seller:   false,
	}

	peoples["alice"] = person{
		name: "alice",
		friends: []string{
			"peggy",
			"you",
		},
		searched: false,
		seller:   false,
	}

	peoples["bob"] = person{
		name: "bob",
		friends: []string{
			"anuj",
			"you",
			"peggy",
		},
		searched: false,
		seller:   false,
	}

	peoples["claire"] = person{
		name: "claire",
		friends: []string{
			"thom",
			"jonny",
		},
		searched: false,
		seller:   false,
	}

	peoples["thom"] = person{
		name:     "thom",
		friends:  []string{},
		searched: false,
		seller:   true,
	}

	var friends []string
	friends = append(friends, "you")
	friends = append(friends, "unknown")
	for i := 0; i < len(friends); i++ {

		if _, ok := peoples[friends[i]]; ok {

			if peoples[friends[i]].searched {
				continue
			}

			if !search(peoples, peoples[friends[i]]) {
				fmt.Printf("%s is not a seller, adding friends - %s\n",
					friends[i], peoples[friends[i]].friends)
				friends = append(friends, peoples[friends[i]].friends...)
				continue
			}

			fmt.Println("Found seller -", friends[i])
			break
		}

		fmt.Printf("%s not found in friends\n", friends[i])

	}

}

func search(peoples map[string]person, people person) bool {

	people.searched = true
	peoples[people.name] = people

	return people.seller
}
