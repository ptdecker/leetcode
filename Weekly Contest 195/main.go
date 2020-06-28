// Weekly Contest 195
// Solved one (isPathCrossing()) out of 4

package main

import (
	"fmt"
	"math"
)

func isPathCrossing(path string) bool {
	var posX, posY int
	visited := make(map[string]bool)
	visited[fmt.Sprintf("(%d,%d)", posX, posY)] = true
	for _, dir := range path {
		switch dir {
		case 'N':
			posY++
		case 'E':
			posX++
		case 'S':
			posY--
		case 'W':
			posX--
		}
		_, ok := visited[fmt.Sprintf("(%d,%d)", posX, posY)]
		if ok {
			return true
		}
		visited[fmt.Sprintf("(%d,%d)", posX, posY)] = true
	}
	return false
}

type pair struct {
	x int
	y int
}

func pick(pics []pair, start int, result []pair) {
	fmt.Println(result)
	if start == len(pics) {
		fmt.Println(result)
		return
	}
	result = append(result, pics[start])
	for i := start + 1; start < len(pics); i++ {
		pick(pics, i, result)
		result = nil
	}
}

func canArrange(arr []int, k int) bool {
	result := []pair{}
	allPairs := []pair{}
	for x := 0; x < len(arr); x++ {
		for y := (x + 1); y < len(arr); y++ {
			if (arr[x]+arr[y])%k == 0 {
				allPairs = append(allPairs, pair{x: arr[x], y: arr[y]})
			}
		}
	}
	fmt.Println(allPairs)
	pick(allPairs, 0, result)
	return false
}

func findMaxValueOfEquation(points [][]int, k int) int {
	var yi, yj, xi, xj int64
	max := int64(-10 ^ 8)
	for i, point := range points {
		xj = int64(point[0])
		yj = int64(point[1])
		if i > 0 {
			diff := int64(math.Abs(float64(xi) - float64(xj)))
			if diff > int64(k) {
				continue
			}
			val := yi + yj + diff
			if val > max {
				max = val
			}
		}
		xi = xj
		yi = yj
	}
	return int(max)
}

func main() {
	//	fmt.Println(isPathCrossing("NES"))
	//	fmt.Println(isPathCrossing("NESWW"))
	//	fmt.Println(findMaxValueOfEquation([][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, 1))
	//	fmt.Println(findMaxValueOfEquation([][]int{{0, 0}, {3, 0}, {9, 2}}, 3))
	//	fmt.Println(findMaxValueOfEquation([][]int{{-19, 9}, {-15, -19}, {-5, -8}}, 10))
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 6}, 7))
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 6}, 10))
	fmt.Println(canArrange([]int{-10, 10}, 2))
	fmt.Println(canArrange([]int{-1, 1, -2, 2, -3, 3, -4, 4}, 3))
}
