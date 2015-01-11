package main

import "fmt"

func main() {
	fmt.Printf("Generating map\n")
	var heightMap = generate(20, 20)

	for i := range heightMap {
		for j := range heightMap[i] {
			fmt.Printf("%f, ", heightMap[i][j])
		}

		fmt.Printf("\n")
	}
}
