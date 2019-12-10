package main

type edge struct {
	s string
	t string
	w int
}

type graph struct {
	obj map[string][]edge
}

func (g *graph) addEdge(s string, t string, w int) {
	e := edge{
		s: s,
		t: t,
		w: w,
	}
	if edgs, ok := g.obj[s]; ok {
		g.obj[s] = append(edgs, e)
	} else {
		g.obj[s] = []edge{e}
	}
}
