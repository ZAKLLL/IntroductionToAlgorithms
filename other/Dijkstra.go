package main

import (
	"github.com/liyue201/gostl/ds/priorityqueue"
)

type Pair struct {
	key   interface{}
	value interface{}
}

func init_distance(graph map[byte]map[byte]int, s byte) map[byte]int {
	distance := make(map[byte]int)
	distance[s] = 0
	for vertex := range graph {
		if s != vertex {
			distance[vertex] = 100000000
		}
	}
	return distance
}
func dijkstra(graph map[byte]map[byte]int, s byte) map[byte]int {
	pq := priorityqueue.New(priorityqueue.WithComparator(pairComparator), priorityqueue.WithGoroutineSafe())
	pq.Push(Pair{0, s})
	seen := make(map[byte]bool) // New empty set
	parent := make(map[byte]byte)
	distance := init_distance(graph, s)
	for !pq.Empty() {
		pop := pq.Pop()
		dist := pop.(Pair).key.(int)
		vertex := pop.(Pair).value.(byte)
		seen[vertex] = true
		nodes := graph[vertex]
		for w := range nodes {
			if _, ok := seen[w]; !ok {
				curDist := dist + graph[vertex][w]
				if curDist < distance[w] {
					distance[w] = curDist
					pq.Push(Pair{curDist, w})
					parent[w] = vertex
				}
			}
		}
	}
	return distance
}

func pairComparator(a, b interface{}) int {
	return a.(Pair).key.(int) - b.(Pair).key.(int)
}
func main() {
	aMap := make(map[byte]int)
	aMap['b'] = 1
	aMap['c'] = 2
	bMap := make(map[byte]int)
	bMap['a'] = 1
	bMap['c'] = 3
	cMap := make(map[byte]int)
	cMap['a'] = 2
	cMap['b'] = 3
	graph := make(map[byte]map[byte]int)
	graph['a'] = aMap
	graph['b'] = bMap
	graph['c'] = cMap
	m := dijkstra(graph, 'a')
	for k, v := range m {
		println(string(k), "-", v)
	}

}
