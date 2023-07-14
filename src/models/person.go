package models

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func EncodeJSON(p Person) []byte {
	jsonPerson, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error with encoding Person to JSON")
	}

	return jsonPerson
}

func DecodeJSON(j []byte) Person {
	var resPers Person

	err := json.Unmarshal(j, &resPers)
	if err != nil {
		fmt.Println("Error with decoding Person from JSON")
	}

	return resPers
}

func (p Person) String() string {
	str := fmt.Sprintf("\nName: %s\n", p.Name)
	str += fmt.Sprintf("Age: %d\n", p.Age)
	str += fmt.Sprintf("Email: %s\n", p.Email)
	return str
}
