// psáno v go
package main

import "fmt"

func main() {
	var N int
	var muzi [][]int
	var zeny [][]int

	fmt.Scan(&N)
	muzi = make([][]int, N)
	zeny = make([][]int, N)
	for i := 0; i < N; i++ {
		muzi[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&muzi[i][j])
		}
	}
	for i := 0; i < N; i++ {
		zeny[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&zeny[i][j])
		}
	}
}
