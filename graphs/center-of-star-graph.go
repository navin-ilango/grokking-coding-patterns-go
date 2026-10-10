package graphs

func findCenter(edges [][]int) int {
	edgeCounter := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		node := edges[i]
		if edgeCounter[node[0]] == 0 {
			edgeCounter[node[0]] = 1
		} else {
			edgeCounter[node[0]] = edgeCounter[node[0]] + 1
		}

		if edgeCounter[node[1]] == 0 {
			edgeCounter[node[1]] = 1
		} else {
			edgeCounter[node[1]] = edgeCounter[node[1]] + 1
		}
	}

	for k, v := range edgeCounter {
		if v > 1 {
			return k
		}
	}

	return -1
}
