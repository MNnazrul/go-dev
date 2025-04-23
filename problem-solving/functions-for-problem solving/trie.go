package main

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type TrieNode struct {
	children [2]*TrieNode
	endIdx int 
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{endIdx: -1}
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func isSet(num, bt int) int {
	if (num & (1 << bt)) > 0 {
		return 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b
}

func (t *Trie) Insert(num, idx int) {
	cur := t.root
	for i := 30; i >= 0; i-- {
		bt := isSet(num, i)
		if cur.children[bt] == nil {
			cur.children[bt] = NewTrieNode()
		}
		cur.endIdx = idx 
		cur = cur.children[bt]
	} 
	cur.endIdx = idx 
}

func (t *Trie) FindAnswer(num, k int) int {
	cur := t.root
	ans := 0
	flg := true
	for i := 30; i >= 0; i-- {
		bt1, bt2 := isSet(num, i), isSet(k, i)
		if bt2 == 0 {
			if cur.children[bt1^1] != nil {
				ans = max(ans, cur.children[bt1^1].endIdx)
			}
		}
		if cur.children[bt1^bt2] == nil {
			flg = false
			break;
		} 
		cur = cur.children[bt1^bt2]
	}
	if flg {
		ans = max(ans, cur.endIdx)
	}
	return ans
}

func solve() {
	trie := NewTrie()
	var n, k, val int 
	Fscan(in, &n, &k)
	idx := -1
	for i := 1; i <= n; i++ {
		Fscan(in, &val)
		trie.Insert(val, i)
		id := trie.FindAnswer(val, k)
		if id > 0 && (idx == -1 || i - id + 1 < idx) {
			idx = i - id + 1
		}
	}
	Fprintln(out, idx)
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}