package SingleSourceShortestPath

const (
	unreachable = -1
	none        = -1
)

type singleSourceShortestPath struct {
	Length int
	Path   []int
}

func dijkstra(graph [][]int, start int) (result []singleSourceShortestPath) {
	result = make([]singleSourceShortestPath, len(graph), len(graph))
	result[start] = singleSourceShortestPath{
		Length: 0,
		Path:   []int{start},
	}
	// 已找出最短路径的点的集合
	picked := make(map[int]struct{}, len(graph))
	picked[start] = struct{}{}
	// 候选点的集合
	candidate := make(map[int]struct{}, len(graph)-1)
	for i := range graph {
		if i == start {
			continue
		}
		candidate[i] = struct{}{}
	}
	for len(candidate) > 0 {
		// 更新距离
		for i := range picked {
			for j := range candidate {
				if graph[i][j] != unreachable {
					if result[j].Length == 0 || result[i].Length+graph[i][j] < result[j].Length {
						result[j] = singleSourceShortestPath{
							Length: result[i].Length + graph[i][j],
							Path:   append(append([]int{}, result[i].Path...), j),
						}
					}
				}
			}
		}
		// 选出一个最短路径的点
		min := unreachable
		minIndex := none
		for i := range candidate {
			if result[i].Length > 0 && (min == unreachable || result[i].Length < min) {
				min = result[i].Length
				minIndex = i
			}
		}
		if minIndex == none {
			break
		}
		picked[minIndex] = struct{}{}
		delete(candidate, minIndex)
	}
	return
}
