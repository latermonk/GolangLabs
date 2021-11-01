#  encoding/json



```
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee1 struct {
	Name   string `json:"n"`
	Age    int    `json:"a"`
	salary int    `json:"s"`
}

type employee2 struct {
	Name   string
	Age    int
	salary int
}

func main() {
	e1 := employee1{
		Name:   "John",
		Age:    21,
		salary: 1000,
	}
	j, err := json.Marshal(e1)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Printf("employee1 JSON: %s\n", string(j))
	e2 := employee2{
		Name:   "John",
		Age:    21,
		salary: 1000,
	}
	j, err = json.Marshal(e2)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Printf("\nemployee2 JSON: %s\n", string(j))

}


```

