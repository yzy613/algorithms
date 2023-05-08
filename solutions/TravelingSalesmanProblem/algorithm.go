package TravelingSalesmanProblem

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
)

type branch struct {
	Length  int
	Path    []int
	Visited []bool
}

func bfs(cityMap [][]int, startCity int) (path []int, length int) {
	cityNum := len(cityMap)
	// 优先队列
	queue := garray.NewSortedArray(func(a, b any) int {
		aAgent := a.(*branch)
		bAgent := b.(*branch)
		if aAgent.Length < bAgent.Length {
			return -1
		}
		if aAgent.Length == bAgent.Length {
			if len(aAgent.Path) > len(bAgent.Path) {
				return -1
			}
			if len(aAgent.Path) == len(bAgent.Path) {
				return 0
			}
		}
		return 1
	})
	// 初始分支
	startB := &branch{
		Length:  0,
		Path:    []int{startCity},
		Visited: make([]bool, cityNum, cityNum),
	}
	startB.Visited[startCity] = true
	queue.Add(startB)
	// 预备剪枝
	currLength := 0
	currPathLenMinLimit := 0
	// bfs
	for queue.Len() > 0 {
		// 取出 Length 最小的分支
		bTemp, found := queue.PopLeft()
		if !found {
			return
		}
		b := bTemp.(*branch)
		// 剪枝
		if b.Length > currLength {
			currLength = b.Length
			// 调整剪枝期望，后面的分数越小越精准
			currPathLenMinLimit = len(b.Path) * 7 / 10
		} else if len(b.Path) < currPathLenMinLimit {
			continue
		}
		// debug
		//progressBar(b, cityNum)
		// 判断已经遍历完所有城市
		switch len(b.Path) {
		case cityNum:
			b.Length += cityMap[b.Path[len(b.Path)-1]][startCity]
			b.Path = append(b.Path, startCity)
			queue.Add(b)
			continue
		case cityNum + 1:
			path = b.Path
			length = b.Length
			return
		}
		// 将分支的所有子分支加入队列
		for i := 0; i < cityNum; i++ {
			if b.Visited[i] || cityMap[b.Path[len(b.Path)-1]][i] == 0 {
				continue
			}
			// 深拷贝
			nextPath := append([]int{}, b.Path...)
			nextVisited := append([]bool{}, b.Visited...)
			nextPath = append(nextPath, i)
			nextVisited[i] = true
			// 加入队列
			queue.Add(&branch{
				Length:  b.Length + cityMap[b.Path[len(b.Path)-1]][i],
				Path:    nextPath,
				Visited: nextVisited,
			})
		}
	}
	return
}

func bfsSlow(cityMap [][]int, startCity int) (path []int, length int) {
	cityNum := len(cityMap)
	// 优先队列
	queue := garray.NewSortedArray(func(a, b any) int {
		aAgent := a.(*branch)
		bAgent := b.(*branch)
		if aAgent.Length < bAgent.Length {
			return -1
		}
		if aAgent.Length == bAgent.Length {
			if len(aAgent.Path) > len(bAgent.Path) {
				return -1
			}
			if len(aAgent.Path) == len(bAgent.Path) {
				return 0
			}
		}
		return 1
	})
	// 初始分支
	startB := &branch{
		Length:  0,
		Path:    []int{startCity},
		Visited: make([]bool, cityNum, cityNum),
	}
	startB.Visited[startCity] = true
	queue.Add(startB)
	// bfs
	for queue.Len() > 0 {
		// 取出 Length 最小的分支
		bTemp, found := queue.PopLeft()
		if !found {
			return
		}
		b := bTemp.(*branch)
		// debug
		//progressBar(b, cityNum)
		// 判断已经遍历完所有城市
		switch len(b.Path) {
		case cityNum:
			b.Length += cityMap[b.Path[len(b.Path)-1]][startCity]
			b.Path = append(b.Path, startCity)
			queue.Add(b)
			continue
		case cityNum + 1:
			path = b.Path
			length = b.Length
			return
		}
		// 将分支的所有子分支加入队列
		for i := 0; i < cityNum; i++ {
			if b.Visited[i] || cityMap[b.Path[len(b.Path)-1]][i] == 0 {
				continue
			}
			// 深拷贝
			nextPath := append([]int{}, b.Path...)
			nextVisited := append([]bool{}, b.Visited...)
			nextPath = append(nextPath, i)
			nextVisited[i] = true
			// 加入队列
			queue.Add(&branch{
				Length:  b.Length + cityMap[b.Path[len(b.Path)-1]][i],
				Path:    nextPath,
				Visited: nextVisited,
			})
		}
	}
	return
}

func progressBar(b *branch, cityNum int) {
	fmt.Print("\r", "Length:", b.Length, "    ")
	for i := range b.Path {
		fmt.Printf("%2d ", b.Path[i]+1)
	}
	for i := 0; i < cityNum-len(b.Path); i++ {
		fmt.Print("-- ")
	}
}
