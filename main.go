package main

import "fmt"

func main() {
	fmt.Printf("Generating map\n")
	var heightMap, err = generate(129) // 9, 17, 33, 129, 257

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range heightMap {
		for j := range heightMap[i] {
			fmt.Printf("%f, ", heightMap[i][j])
		}

		fmt.Printf("\n")
	}
}
