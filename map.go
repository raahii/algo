package main

import "sort"

/* Map */

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortMapByValue(hashmap map[string]int) PairList {
	pl := make(PairList, len(hashmap))

	i := 0
	for k, v := range hashmap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}
