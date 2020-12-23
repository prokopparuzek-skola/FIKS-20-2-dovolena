// psáno v go
package main

import "fmt"

type peoples struct {
	muzi  [][]int
	muziI []map[int]int
	zeny  [][]int
	zenyI []map[int]int
}

type block struct {
	block bool
	who   int
	where int
}

func pair(group [][]int) (pairs []int) { // jdenostraně najde páry
	var whoBlock []block

	whoBlock = make([]block, len(group))
	pairs = make([]int, len(group))
	for i := 0; i < len(group); i++ { // projde všechny subjekty na spárování
		for j := 0; j < len(group); j++ { // projde odpovídající opačná pohlaví
			if whoBlock[group[i][j]].block == true {
				if j < whoBlock[group[i][j]].where { // když už si někdo někoho zabral, ale já ho mám dřív
					tmp := whoBlock[group[i][j]]
					whoBlock[group[i][j]].who = i
					whoBlock[group[i][j]].where = j
					pairs[i] = group[i][j]
					i = tmp.who
					j = tmp.where
					continue // prohodím a tomu komu jsem sebral protějšek najdu jiný
				}
			} else { // najdu někoho volného
				whoBlock[group[i][j]].block = true
				whoBlock[group[i][j]].who = i
				whoBlock[group[i][j]].where = j
				pairs[i] = group[i][j]
				break
			}
		}
	}
	return
}

func makePairs(relations *peoples, mWants []int, zWants []int) (pairs [][2]int) { // zkompletuje páry od obou stran
	pairs = make([][2]int, 0)
	for i, z := range mWants {
		if zWants[z] == i { // oba se chtějí, není co řešit
			pairs = append(pairs, [2]int{i, z})
		} else if zWants[z] < i { // žena chce někoho už spárovaného, ten je spokojený, smůla
			pairs = append(pairs, [2]int{i, z})
		} else if relations.muziI[zWants[z]][z] < relations.muziI[zWants[z]][mWants[zWants[z]]] { // žena chce někoho ještě volného více, prohodím
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
	for i := 0; i < N; i++ { // načti data
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
	mWants := pair(relations.muzi) // jednostranné páry
	zWants := pair(relations.zeny)
	pairs := makePairs(&relations, mWants, zWants) // obostranné
	for _, p := range pairs {
		fmt.Printf("%d + %d\n", p[0]+1, p[1]+1)
	}
}
