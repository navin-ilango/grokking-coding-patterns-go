package topologicalsort

func verifyAlienDictionary(words []string, order string) bool {
	if len(words) == 1 {
		return true
	}

	orderMap := make(map[byte]int)

	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}

			if words[i][j] != words[i+1][j] {
				if orderMap[words[i][j]] > orderMap[words[i+1][j]] {
					return false
				} else {
					break
				}
			}
		}
	}

	return true
}
