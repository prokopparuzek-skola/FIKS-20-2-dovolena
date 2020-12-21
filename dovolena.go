// psáno v go
package main

import "fmt"

type peoples struct {
	muzi  [][]int
	muziI []map[int]int
	zeny  [][]int
	zenyI []map[int]int
}

func pair(group [][]int) (pairs []int) {
	var block []bool

	block = make([]bool, len(group))
	pairs = make([]int, len(group))
	for i := 0; i < len(group); i++ {
		for j := 0; j < len(group); j++ {
			if block[group[i][j]] == false {
				block[group[i][j]] = true
				pairs[i] = j
				break
			}
		}
	}
	return
}

func main() {
	var N int
	var relations peoples

	fmt.Scan(&N)
	relations.muzi = make([][]int, N)
	relations.zeny = make([][]int, N)
	relations.muziI = make([]map[int]int, N)
	relations.zenyI = make([]map[int]int, N)
	for i := 0; i < N; i++ {
		relations.muzi[i] = make([]int, N)
		relations.muziI[i] = make(map[int]int)
		for j := 0; j < N; j++ {
			fmt.Scan(&relations.muzi[i][j])
			relations.muziI[i][relations.muzi[i][j]] = j
		}
	}
	for i := 0; i < N; i++ {
		relations.zeny[i] = make([]int, N)
		relations.zenyI[i] = make(map[int]int)
		for j := 0; j < N; j++ {
			fmt.Scan(&relations.zeny[i][j])
			relations.zenyI[i][relations.zeny[i][j]] = j
		}
	}
}
