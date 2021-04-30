package helpers

import (
	"encoding/csv"
	"github.com/Kratos40-sba/personnes-service/data"
	"log"
	"os"
	"strconv"
)

func ReadFromCsv(filename string) *csv.Reader {
	f, err := os.Open(filename)
	ErrorHandler(err)
	r := csv.NewReader(f)
	return r
}
func ErrorHandler(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
func ParseToMap(r *csv.Reader) (map[int]data.Person, error) {
	m := make(map[int]data.Person)
	d, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, d := range d {
		id, err := strconv.Atoi(d[0])
		if err != nil {
			log.Println("Error parsing the ID")
			continue
		}
		m[id] = data.Person{Name: d[2], LastName: d[3], Service: d[1], BirthDate: d[4]}
	}
	return m, nil
}
