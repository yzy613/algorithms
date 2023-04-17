package OptimalDelivery

import "github.com/gogf/gf/v2/container/gqueue"

var moveRule = [2][4]int{
	//右,下,左,上
	{0, 1, 0, -1}, // row
	{1, 0, -1, 0}} // column

func bfs(matrix []point, size, customerNum int, shopIndex []int) (cost int) {
	queue := gqueue.New()
	for _, index := range shopIndex {
		queue.Push(index)
	}
	for queue.Len() > 0 {
		index := queue.Pop().(int)
		for i := 0; i < 4; i++ {
			nextRow := index/size + moveRule[0][i]
			nextCol := index%size + moveRule[1][i]
			if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size {
				continue
			}
			nextIndex := nextRow*size + nextCol
			if matrix[nextIndex].Role == obstaclePoint || matrix[nextIndex].Cost != 0 {
				continue
			}
			matrix[nextIndex].Cost = matrix[index].Cost + 1
			if matrix[nextIndex].Role == customerPoint {
				cost += matrix[nextIndex].Cost * matrix[nextIndex].NeedDeliver
				customerNum--
			}
			if customerNum == 0 {
				return
			}
			queue.Push(nextIndex)
		}
	}
	return
}
