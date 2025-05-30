package utils

import (
	"container/heap"
)

type State struct {
	totalItems int
	totalPacks int
	counts     map[int]int
}

type Item struct {
	state *State
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	a, b := pq[i].state, pq[j].state
	if a.totalItems == b.totalItems {
		return a.totalPacks < b.totalPacks
	}
	return a.totalItems < b.totalItems
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func PackageDistribution(packs []int, target int) map[int]int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	start := &State{
		totalItems: 0,
		totalPacks: 0,
		counts:     make(map[int]int),
	}
	heap.Push(pq, &Item{state: start})

	visited := make(map[int]int)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		state := item.state

		if state.totalItems >= target {
			return state.counts
		}

		if prevPacks, ok := visited[state.totalItems]; ok && prevPacks <= state.totalPacks {
			continue
		}
		visited[state.totalItems] = state.totalPacks

		for _, pack := range packs {
			newCounts := make(map[int]int)
			for k, v := range state.counts {
				newCounts[k] = v
			}
			newCounts[pack]++

			newState := &State{
				totalItems: state.totalItems + pack,
				totalPacks: state.totalPacks + 1,
				counts:     newCounts,
			}
			heap.Push(pq, &Item{state: newState})
		}
	}
	return nil
}