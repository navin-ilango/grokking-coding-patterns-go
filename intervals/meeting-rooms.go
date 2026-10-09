package intervals

import "sort"

func findSets(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	meetingsRooms := newMinHeap()
	meetingsRooms.Push(intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= meetingsRooms.Top().(int) {
			meetingsRooms.Pop()
		}
		meetingsRooms.Push(intervals[i][1])
	}
	return meetingsRooms.Len()
}
