package leet

import (
	"fmt"
	"log"
)

type Fact struct {
	src   string
	coeff float32
	dest  string
}

type Query struct {
	value float32
	src   string
	dest  string
}

type Graph struct {
	metric string
	edges  *[]Fact
}

func reverse(f Fact) Fact {
	return Fact{
		src:   f.dest,
		dest:  f.src,
		coeff: -f.coeff,
	}
}

func createGraph(facts []Fact) map[string]Graph {
	adj := map[string]Graph{}
	for _, f := range facts {
		src, ok_src := adj[f.src]
		if !ok_src {
			src = Graph{
				metric: f.src,
				edges:  &[]Fact{},
			}
			adj[f.src] = src
		}

		dest, ok_dest := adj[f.dest]
		if !ok_dest {
			dest = Graph{
				metric: f.dest,
				edges:  &[]Fact{},
			}
			adj[f.dest] = dest
		}

		*src.edges = append(*src.edges, f)
		*dest.edges = append(*dest.edges, reverse(f))
	}
	return adj
}

func dfs(graph map[string]Graph, head Graph, query Query, visited map[string]bool) (float32, error) {
	if query.dest == head.metric {
		return query.value, nil
	}

	if visited[head.metric] {
		return query.value, nil
	}

	visited[head.metric] = true

	for i, n := range *graph[head.metric].edges {
		log.Println(i, n)
		if visited[n.dest] {
			continue
		}
		if n.coeff > 0 {
			query.value *= n.coeff
		} else {
			query.value /= -n.coeff
		}
		return dfs(graph, graph[n.dest], query, visited)
	}
	return 0, fmt.Errorf("Path not found")
}

func AnserQuery(query Query, facts []Fact) (float32, error) {
	graph := createGraph(facts)
	start, ok := graph[query.src]
	if !ok {
		return 0, fmt.Errorf("Unknown start metric")
	}

	visited := map[string]bool{}
	return dfs(graph, start, query, visited)
}
