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

func makePairs(relations *peoples, mWants []int, zWants []int) (pairs [][2]int) {
	pairs = make([][2]int, 0)
	for i, z := range mWants {
		if zWants[z] == i {
			pairs = append(pairs, [2]int{i, z})
		} else if zWants[z] < i {
			pairs = append(pairs, [2]int{i, z})
		} else if relations.muziI[zWants[z]][z] < relations.muziI[zWants[z]][mWants[zWants[z]]] {
			mWants[i] = mWants[zWants[z]]
			mWants[zWants[z]] = z
			pairs = append(pairs, [2]int{i, mWants[i]})
		} else {
			pairs = append(pairs, [2]int{i, z})
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
			relations.muzi[i][j]--
			relations.muziI[i][relations.muzi[i][j]] = j
		}
	}
	for i := 0; i < N; i++ {
		relations.zeny[i] = make([]int, N)
		relations.zenyI[i] = make(map[int]int)
		for j := 0; j < N; j++ {
			fmt.Scan(&relations.zeny[i][j])
			relations.zeny[i][j]--
			relations.zenyI[i][relations.zeny[i][j]] = j
		}
	}
	mWants := pair(relations.muzi)
	zWants := pair(relations.zeny)
	pairs := makePairs(&relations, mWants, zWants)
	for _, p := range pairs {
		fmt.Printf("%d + %d\n", p[0], p[1])
	}
}
