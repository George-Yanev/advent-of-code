package cmd

import "fmt"

func Root(inputs []string) {
	hash := &HashMapTree{
		hash: make(map[string]*Node),
	}
	myMap := ConvertInputToMap(inputs)

	hash.Insert("COM", "COM", myMap["COM"], myMap)
	// fmt.Println(hash.hash)
	youPath := hash.findNodeRootPath("YOU")
	sanPath := hash.findNodeRootPath("SAN")
	totalPath := append(youPath, sanPath...)

	seen := make(map[string]struct{})
	for _, v := range totalPath {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
		} else {
			delete(seen, v)
		}
	}
	result := make([]string, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}

	fmt.Println()
	fmt.Println("Result: ", result)
	fmt.Println("Oribtal Transfers: ", len(result))
}
