package main

import (
	"fmt"
	"math/rand"
	"time"
)

const CornerStartingHeight = 50
const Roughness = 8

type DiamondSquareParams struct {
	LowX         int
	LowY         int
	HighX        int
	HighY        int
	RandomFactor float64
}

// Generate an height map using the diamond square algorithm
func generate(x, y int) [][]float64 {
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
	diamondSquare(heightMap, DiamondSquareParams{0, 0, x, y, randomFactor}, randomGen)

	return heightMap
}

// Applies the diamond square algorithm to a sub-section of the map
func diamondSquare(heightMap [][]float64, params DiamondSquareParams,
	randomGen *rand.Rand) {
	// check for end of recursion
	fmt.Printf("%d %d %d %d ", params.LowX, params.LowY, params.HighX, params.HighY)
	if params.LowX >= params.HighX || params.LowY >= params.HighY {
		fmt.Printf("Exit")
		return
	}

	// diamond step
	var averageDiamond = (heightMap[params.LowX][params.LowY] +
		heightMap[params.LowX][params.HighY-1] +
		heightMap[params.HighX-1][params.LowY] +
		heightMap[params.HighX-1][params.HighY-1]) / 4

	heightMap[params.HighX/2][params.HighY/2] = averageDiamond +
		(randomGen.NormFloat64() * params.RandomFactor)

	// square step
	heightMap[params.LowX][params.HighY/2] = (heightMap[params.LowX][params.LowY]+
		heightMap[params.LowX][params.HighY-1])/2 + randomGen.NormFloat64()*params.RandomFactor

	heightMap[params.HighX-1][params.HighY/2] = (heightMap[params.HighX-1][params.LowY]+
		heightMap[params.HighX-1][params.HighY-1])/2 + randomGen.NormFloat64()*params.RandomFactor

	heightMap[params.HighX/2][params.LowY] = (heightMap[params.LowX][params.LowY]+
		heightMap[params.HighX-1][params.LowY])/2 + randomGen.NormFloat64()*params.RandomFactor

	heightMap[params.HighX/2][params.HighY-1] = (heightMap[params.LowX][params.HighY-1]+
		heightMap[params.HighX-1][params.HighY-1])/2 + randomGen.NormFloat64()*params.RandomFactor

	// Recursively call diamondSquare
	diamondSquare(heightMap, DiamondSquareParams{params.LowX, params.LowY,
		params.HighX / 2, params.HighY / 2,
		params.RandomFactor}, randomGen)

	// diamondSquare(heightMap, DiamondSquareParams{params.HighX / 2, params.LowY,
	// 	params.HighX, params.HighY / 2,
	// 	params.RandomFactor}, randomGen)

	// diamondSquare(heightMap, DiamondSquareParams{params.HighX / 2, params.HighY / 2,
	// 	params.HighX, params.HighY,
	// 	params.RandomFactor}, randomGen)

	// diamondSquare(heightMap, DiamondSquareParams{params.LowX, params.HighY / 2,
	// 	params.HighX / 2, params.HighY,
	// 	params.RandomFactor}, randomGen)
}
