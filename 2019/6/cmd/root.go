package cmd

import "fmt"

func Root(inputs []string) error {
	hash := &HashMapTree{
		hash: make(map[string]*Node),
	}
	myMap := ConvertInputToMap(inputs)

	hash.Insert("COM", "COM", myMap["COM"], myMap)
	fmt.Println(hash.hash["COM"])
	return nil
}
