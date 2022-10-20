package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	row := len(mat)
	col := len(mat[0])
	if row*col !=  r*c {
		return mat
	}
	var res = make([][]int, r)
	for t := 0; t < r; t++ {
		res[t] = make([]int, c)
	}
	i, j := 0, 0
	for m := 0; m < r; m++ {
		for n := 0; n < c; n++ {
			if j == col {
				j = 0
				i++
			}
			res[m][n] = mat[i][j]
			j++
		}
	}
	return res
}

func main() {
	var mat =  [][]int{{1,2}}
	fmt.Println(matrixReshape(mat,1,1))
}
