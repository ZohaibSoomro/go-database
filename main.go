package main

import (
	// "encoding/json"
	"fmt"
	"sync"

	"github.com/zohaibsoomro/go-database/model"
)

func main() {
	// u := model.User{
	// 	Id: json.Number("21"), Name: "Zohaib khan", City: "Mirpur Mathelo",
	// }
	// _, err := u.SaveData(&sync.Mutex{})

	// if err != nil {
	// 	print("error: ", err)
	// }

	d, _ := model.ReadAlldata("./", &sync.Mutex{})
	for _, v := range d {
		fmt.Printf("%+v\n", v)
	}
	b, er := model.DeleteAll(&sync.Mutex{})
	if er != nil {
		print(b)
		print(er)
	}
}
