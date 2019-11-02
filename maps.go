package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2 = empty map

	var m3 map[string]int // m3 = nil

	fmt.Println("m, m2 ,m3 :")
	fmt.Println(m, m2, m3)
	//add
	m["qqq"] = "www"
	m["aaa"] = "sss"
	//foreach
	fmt.Println("Traversing map m")
	for key, value := range m {
		fmt.Println(key, value)
	}
	//get
	name := m["name"]
	fmt.Println(name)

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key 'name' does not exist")
	}
	//delete
	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
	//update
	m["name"] = "test"
	fmt.Println(m["name"])
}
