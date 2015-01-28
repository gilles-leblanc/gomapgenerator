// TODO Do no overwrite already existing values
// TODO Check not for odd numbers but for 2 pow n + 1
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const CornerStartingHeight = 50
const Roughness = 10

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

	// Initialize randomness
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomFactor := float64(Roughness * x / 10)

	// Assign a basic value to each of the four corners
	heightMap[0][0] = CornerStartingHeight + generateRandomNumber(randomFactor, randomGen)
	heightMap[0][y-1] = CornerStartingHeight + generateRandomNumber(randomFactor, randomGen)
	heightMap[x-1][0] = CornerStartingHeight + generateRandomNumber(randomFactor, randomGen)
	heightMap[x-1][y-1] = CornerStartingHeight + generateRandomNumber(randomFactor, randomGen)

	// Recursively generate height map
	fractalGeneration(heightMap, FractalParams{0, 0, x - 1, y - 1, randomFactor}, randomGen)

	return heightMap, nil
}

// Applies a fractal algorithm to a sub-section of the map
func fractalGeneration(heightMap [][]float64, params FractalParams,
	randomGen *rand.Rand) {
	fmt.Printf("%d %d %d %d \n", params.LowX, params.LowY, params.HighX, params.HighY)

	// assign center value step
	var averageCenter = (heightMap[params.LowX][params.LowY] +
		heightMap[params.LowX][params.HighY] +
		heightMap[params.HighX][params.LowY] +
		heightMap[params.HighX][params.HighY]) / 4

	// we multiply the RandomFactor by 2 when we assign the center point because
	// the center points needs to be more randomized than the corner midpoints
	heightMap[(params.LowX+params.HighX)/2][(params.LowY+params.HighY)/2] = averageCenter +
		generateRandomNumber(params.RandomFactor*2, randomGen)

	xMidPoint := (params.LowX + params.HighX) / 2
	yMidPoint := (params.LowY + params.HighY) / 2

	// assign corner midpoints step
	heightMap[params.LowX][yMidPoint] = (heightMap[params.LowX][params.LowY]+
		heightMap[params.LowX][params.HighY])/2 + generateRandomNumber(params.RandomFactor, randomGen)

	heightMap[params.HighX][yMidPoint] = (heightMap[params.HighX][params.LowY]+
		heightMap[params.HighX][params.HighY])/2 + generateRandomNumber(params.RandomFactor, randomGen)

	heightMap[xMidPoint][params.LowY] = (heightMap[params.LowX][params.LowY]+
		heightMap[params.HighX][params.LowY])/2 + generateRandomNumber(params.RandomFactor, randomGen)

	heightMap[xMidPoint][params.HighY] = (heightMap[params.LowX][params.HighY]+
		heightMap[params.HighX][params.HighY])/2 + generateRandomNumber(params.RandomFactor, randomGen)

	// Recalculate RandomFactor so it gets lower with each iteration
	if params.RandomFactor > 1 {
		params.RandomFactor = params.RandomFactor * 0.9
	}

	// Recursively call fractal generation
	if params.HighX-params.LowX > 2 {
		fractalGeneration(heightMap, FractalParams{params.LowX, params.LowY,
			xMidPoint, yMidPoint,
			params.RandomFactor}, randomGen)

		fractalGeneration(heightMap, FractalParams{xMidPoint, params.LowY,
			params.HighX, yMidPoint,
			params.RandomFactor}, randomGen)

		fractalGeneration(heightMap, FractalParams{params.LowX, yMidPoint,
			xMidPoint, params.HighY,
			params.RandomFactor}, randomGen)

		fractalGeneration(heightMap, FractalParams{xMidPoint, yMidPoint,
			params.HighX, params.HighY,
			params.RandomFactor}, randomGen)
	}
}

func generateRandomNumber(randomFactor float64, randomGen *rand.Rand) float64 {
	return randomGen.NormFloat64() * randomFactor
}
