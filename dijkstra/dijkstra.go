package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//生成图
	obj := make(map[string][]edge)
	g := graph{
		obj: obj,
	}
	g.addEdge("0", "1", 10)
	g.addEdge("0", "4", 15)
	g.addEdge("1", "2", 15)
	g.addEdge("1", "3", 2)
	g.addEdge("3", "2", 1)
	g.addEdge("2", "5", 5)
	g.addEdge("3", "5", 12)
	g.addEdge("4", "5", 10)

	// fmt.Print(obj)
	// return

	//起点
	dijkstra("0", "5", obj)
}

func dijkstra(s, t string, obj map[string][]edge) {
	inqueue := make(map[string]*Item)
	path := make(map[string]string)
	start := &Item{
		dist: 0,
		dot:  s,
	}

	pq := PriorityQueue{start}
	heap.Init(&pq)

	inqueue[s] = start

	for len(pq) > 0 {
		// printPirorityQueue(pq)
		minLen := heap.Pop(&pq).(*Item)

		// fmt.Printf("out queue: %s\n", minLen.dot)

		if t == minLen.dot {
			// fmt.Print(*minLen)
			// fmt.Println()
			break
		}
		for _, e := range obj[minLen.dot] {

			t := minLen.dist + e.w
			if ip, ok := inqueue[e.t]; ok {
				if t < ip.dist {
					path[e.t] = e.s
					// fmt.Print("t:", t, "ip:", ip.dist)
					// fmt.Println()
					pq.Update(ip, t)
				}
			} else {
				path[e.t] = e.s
				t := &Item{
					dot:  e.t,
					dist: t,
				}
				heap.Push(&pq, t)
				// fmt.Printf("in queue: %s\n", t.dot)
				inqueue[e.t] = t
			}
		}
	}

	fmt.Print(path)
	fmt.Println()
	// fmt.Print(minLen)
}

func printPirorityQueue(queue PriorityQueue) {
	for _, v := range queue {
		fmt.Print(*v)
		fmt.Println()
	}
}
