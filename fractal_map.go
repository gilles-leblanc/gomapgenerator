// TODO Change random number generation, and decrease number generation on succesive calls
// TODO Do no overwrite already existing values
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const CornerStartingHeight = 50
const Roughness = 4

type FractalParams struct {
	LowX         int
	LowY         int
	HighX        int
	HighY        int
	RandomFactor float64
}

// Generate an height map using a fractal algorithm
func generate(size int) ([][]float64, error) {
	// check if size is odd
	if size%2 == 0 {
		return nil, errors.New("Size must be odd.")
	}

	var x, y = size, size

	// Generate our slice to hold the height map data
	heightMap := make([][]float64, y, y)

	for i := range heightMap {
		heightMap[i] = make([]float64, x, x)
	}

	// Assign a basic value to each of the four corners
	heightMap[0][0] = CornerStartingHeight
	heightMap[0][y-1] = CornerStartingHeight
	heightMap[x-1][0] = CornerStartingHeight
	heightMap[x-1][y-1] = CornerStartingHeight

	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomFactor := float64(Roughness * x / 10)
	fractalGeneration(heightMap, FractalParams{0, 0, x - 1, y - 1, randomFactor}, randomGen)

	return heightMap, nil
}

// Applies a fractal algorithm to a sub-section of the map
func fractalGeneration(heightMap [][]float64, params FractalParams,
	randomGen *rand.Rand) {
	// check for end of recursion
	fmt.Printf("%d %d %d %d \n", params.LowX, params.LowY, params.HighX, params.HighY)
	// if params.LowX >= params.HighX-1 || params.LowY >= params.HighY-1 {
	// 	fmt.Printf("Exit\n")
	// 	time.Sleep(1000 * time.Millisecond)
	// 	return
	// }

	// assign center value step
	var averageCenter = (heightMap[params.LowX][params.LowY] +
		heightMap[params.LowX][params.HighY] +
		heightMap[params.HighX][params.LowY] +
		heightMap[params.HighX][params.HighY]) / 4

	// we multiply the RandomFactor by 2 when we assign the center point because
	// the center points needs to be more randomized than the corner midpoints
	heightMap[(params.LowX+params.HighX)/2][(params.LowY+params.HighY)/2] = averageCenter +
		0 //(randomGen.NormFloat64() * params.RandomFactor * 2)

	// assign corner midpoints step
	heightMap[params.LowX][params.HighY/2] = (heightMap[params.LowX][params.LowY]+
		heightMap[params.LowX][params.HighY])/2 + 0 //randomGen.NormFloat64()*params.RandomFactor

	heightMap[params.HighX][params.HighY/2] = (heightMap[params.HighX][params.LowY]+
		heightMap[params.HighX][params.HighY])/2 + 0 //randomGen.NormFloat64()*params.RandomFactor

	heightMap[params.HighX/2][params.LowY] = (heightMap[params.LowX][params.LowY]+
		heightMap[params.HighX][params.LowY])/2 + 0 //randomGen.NormFloat64()*params.RandomFactor

	heightMap[params.HighX/2][params.HighY] = (heightMap[params.LowX][params.HighY]+
		heightMap[params.HighX][params.HighY])/2 + 0 //randomGen.NormFloat64()*params.RandomFactor

	// check for end of recursion
	// if params.HighX-params.LowX <= 2 {
	// 	return
	// }

	// Recalculate RandomFactor so it gets lower with each iteration
	// var newRandomFactor = params.RandomFactor

	// Recursively call fractal generation
	if params.HighX-params.LowX > 2 {
		fractalGeneration(heightMap, FractalParams{params.LowX, params.LowY,
			params.HighX / 2, params.HighY / 2,
			params.RandomFactor}, randomGen)

		fractalGeneration(heightMap, FractalParams{params.HighX / 2, params.LowY,
			params.HighX, params.HighY / 2,
			params.RandomFactor}, randomGen)

		fractalGeneration(heightMap, FractalParams{params.LowX, params.HighY / 2,
			params.HighX / 2, params.HighY,
			params.RandomFactor}, randomGen)

		fractalGeneration(heightMap, FractalParams{params.HighX / 2, params.HighY / 2,
			params.HighX, params.HighY,
			params.RandomFactor}, randomGen)
	}
}
