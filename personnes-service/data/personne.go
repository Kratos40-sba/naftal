package data

import "fmt"

type Person struct {
	Name, LastName, Service, BirthDate string // 2 , 3 , 1 ,4
}

func (p Person) String() {
	fmt.Printf("Nom : %s | Prenom : %s | Service : %s | DateN : %s \n", p.Name, p.LastName, p.Service, p.BirthDate)
}
func SearchByID(id int, cache map[int]Person) (Person, bool) {
	if p, ok := cache[id]; ok {
		return p, true
	}
	return Person{}, false
}
