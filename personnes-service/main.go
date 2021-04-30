package main

import (
	"flag"
	"fmt"
	"github.com/Kratos40-sba/personnes-service/data"
	"github.com/Kratos40-sba/personnes-service/helpers"
)

var (
	csvFileName = flag.String("f", "personnes.csv", "fichier ou les donnees des pesonnes")
	idFlag      = flag.Int("id", 250084, "Matricule de personne")
	cache       map[int]data.Person
)

func main() {
	var err error
	flag.Parse()
	r := helpers.ReadFromCsv(*csvFileName)
	cache, err = helpers.ParseToMap(r)
	helpers.ErrorHandler(err)
	p, found := data.SearchByID(*idFlag, cache)
	if found {
		p.String()
	} else {
		fmt.Println("ID Mismatch")
	}

}
