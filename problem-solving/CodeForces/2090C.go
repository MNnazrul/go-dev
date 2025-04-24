package main

import (
	."fmt"
	"os"
	"bufio"
	"sort"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	cell []Gr
	stSeat []Gr
)
const INF int64 = 1e18
const N int = 3e5 + 5

type Gr struct {
	x int 
	y int 
}

func dist(a, b int) int {
	sm := a + b 
	if a % 3 == 2 && b % 3 == 2 {
		sm += 2
	}
	return sm
}

func compare(a , b Gr) bool {
	sm1, sm2 := dist(a.x, a.y), dist(b.x, b.y)
	if sm1 < sm2 {
		return true
	}  
	if sm1 == sm2 {
		if a.x < b.x {
			return true
		}
		if a.x == b.x && a.y <= b.y {
			return true
		}
	}
	return false
}

func pre() {
	mp := make(map[int][]Gr)
	ar := make([]int, 0)
	for i := 1; i < N; i++ {
		if i % 3 == 0 {
			continue
		}
		ar = append(ar, i)
	}
	i1, cnt := 0, N 
	for cnt > 0 {
		i11, i2 := i1, 0 
		for i11 >= 0 {
			sm := dist(ar[i11], ar[i2])
			mp[sm] = append(mp[sm], Gr{ar[i2], ar[i11]})
			i11--
			i2++
			cnt--
		}
		i1++
	}
	keys := make([]int, 0)
	for key := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, i := range keys {
		vec, ok := mp[i]
		if !ok || len(vec) == 0 {
			continue
		}
		sort.Slice(vec, func(i, j int) bool {
			return compare(vec[i], vec[j])
		})
		for _, item := range vec {
			cell = append(cell, item)
			if item.x % 3 == 1 && item.y % 3 == 1 {
				stSeat = append(stSeat, item)
			}
		}
	}
}

func setMap(mp map[int]map[int]int,i, j int) {
	if mp[i] == nil {
		mp[i] = make(map[int]int)
	}
	mp[i][j] = 1
}

func chkMap(mp map[int]map[int]int, i, j int) bool {
	if _, ok := mp[i]; ok {
		_, ok1 := mp[i][j]
		return ok1
	}
	return false
}

func solve() {
	var n int  
	Fscan(in, &n)
	i1, i2 := 0, 0
	mp := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		var flg bool
		Fscan(in, &flg)
		if !flg {
			for chkMap(mp, stSeat[i2].x, stSeat[i2].y) {
				i2++
			}
			setMap(mp, stSeat[i2].x, stSeat[i2].y)
			Fprintln(out, stSeat[i2].x, stSeat[i2].y)
			continue
		}
		for chkMap(mp, cell[i1].x, cell[i1].y) {
			i1++
		}
		setMap(mp, cell[i1].x, cell[i1].y)
		Fprintln(out, cell[i1].x, cell[i1].y)
	}

}

func main() {

	defer out.Flush()

	pre()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}