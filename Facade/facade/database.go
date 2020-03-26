package facade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Database is struct
type Database struct {
}

// Person is struct
type Person struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

var maildata []Person

func (d *Database) getProperties(dbname string) []Person {
	filename := dbname + ".json"
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = json.Unmarshal(jsonString, &maildata); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return maildata
}
