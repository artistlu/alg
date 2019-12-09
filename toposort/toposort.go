package main

import "fmt"

type graph struct {
	Ats map[string][]string
}

func (g *graph) AddEdg(s, t string) {
	if v, ok := g.Ats[s]; ok {
		g.Ats[s] = append(v, t)
	} else {
		t := []string{t}
		g.Ats[s] = t
	}
}

func main() {

	g := new(graph)
	g.Ats = make(map[string][]string)

	g.AddEdg("a", "b")
	g.AddEdg("b", "c")
	g.AddEdg("c", "d")
	g.AddEdg("d", "x")
	g.AddEdg("f", "z")
	g.AddEdg("f", "b")
	g.AddEdg("s", "t")

	for i, v := range g.Ats {
		fmt.Print(i, ": ")
		for _, vv := range v {
			fmt.Printf("%s ", vv)

		}
		fmt.Println()

	}

	// topoSortByKahn(g.Ats)
	topoSortByDFS(g.Ats)
	// fmt.Print(g)

	// m := make(map[string]string)

	// m["age"] = "8"
	// fmt.Print(m)

}

func topoSortByKahn(ats map[string][]string) {

	//统计每个点的入度
	inDegree := make(map[string]int)
	for i, at := range ats {
		if _, ok := inDegree[i]; !ok {
			inDegree[i] = 0
		}
		for _, v := range at {
			inDegree[v]++
		}
	}

	//入度为0的加入队列
	zeroQueue := make([]string, 0)
	for s, c := range inDegree {
		if c == 0 {
			zeroQueue = append(zeroQueue, s)
		}
	}

	//打印入度为0的点，并减少依赖改点的点的入度值
	for len(zeroQueue) != 0 {
		s := zeroQueue[0]
		zeroQueue = zeroQueue[1:]
		fmt.Printf("->%s", s)
		for _, v := range ats[s] {
			inDegree[v]--
			if inDegree[v] == 0 {
				zeroQueue = append(zeroQueue, v)
			}
		}
	}

	fmt.Println()
}

func topoSortByDFS(ats map[string][]string) {

	reverAts := make(map[string][]string)
	for i, v := range ats {
		for _, s := range v {
			if vv, ok := reverAts[s]; ok {
				reverAts[s] = append(vv, i)
			} else {
				reverAts[s] = []string{i}

			}
		}
	}

	visited := make(map[string]bool)
	for i, _ := range reverAts {
		if _, ok := visited[i]; ok {
			continue
		}
		visited[i] = true
		dfs(i, reverAts, &visited)
	}
}

func dfs(s string, reverAts map[string][]string, visited *map[string]bool) {
	for _, k := range reverAts[s] {
		if _, ok := (*visited)[k]; ok {
			continue
		}
		(*visited)[k] = true
		dfs(k, reverAts, visited)
	}

	fmt.Printf("->%s", s)
}
