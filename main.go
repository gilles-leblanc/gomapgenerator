package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	fmt.Printf("Generating map\n")
	imageSize := 1025 // 9, 17, 33, 129, 257, 513, 1025
	var heightMap, err = generate(imageSize)

	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create("./height_map.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	imgRect := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewNRGBA64(imgRect)

	for i := range heightMap {
		for j := range heightMap[i] {
			value := heightMap[i][j]
			colorValue := color.NRGBA64{value, value, value, 65535}
			img.SetNRGBA64(i, j, colorValue)
		}
	}

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Created height_map.png")
}
