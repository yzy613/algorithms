package baShuMa

import (
	"fmt"
)

/*
 * 小笔记：
 * 对于 arr[position]，position = row * 3 + column
 * bfs 去重的算法思路先不变，后面再搞 IDA* 算法（主要是我想记录过程，好阅读还原）
 * 注意 康托展开算法
 * 考虑 verify() 函数的用途
 */

func BSM() (err error) {
	srcZeroPosition, _, err := inputAndPreprocess()
	if err != nil {
		return
	}
	//viewMatrix(srcArr)
	//viewMatrix(dstArr)
	// 这个 verify(srcZeroPosition, dstZeroPosition int) (err error) 无意义，无解也就 15s 运行完
	/*err = verify(srcZeroPosition, dstZeroPosition)
	if err != nil {
		return
	}*/
	srcCantor := CantorExpansion(srcArr)
	nodeCantorArr = make([]nodeInfo, factorialArr[matrixRow*matrixRow], factorialArr[matrixRow*matrixRow])
	nodeCantorArr[srcCantor] = nodeInfo{
		Step:         1,
		ParentCantor: 0,
		Arr:          srcArr,
		ZeroPosition: srcZeroPosition,
		Arrival:      false,
	}
	ansNode, err := bsmBFS(nodeCantorArr[srcCantor], CantorExpansion(dstArr))
	if err != nil {
		return
	}
	if ansNode.Arrival {
		ansNode.Step--
		fmt.Println()
		restoreMatrix(ansNode.ParentCantor)
		viewMatrix(ansNode.Arr)
		fmt.Printf("%+v\n", ansNode)
	} else {
		fmt.Println("\n无解")
	}
	return
}

func restoreMatrix(nextNodeCantor int) {
	if nextNodeCantor == 0 {
		return
	}
	restoreMatrix(nodeCantorArr[nextNodeCantor].ParentCantor)
	viewMatrix(nodeCantorArr[nextNodeCantor].Arr)
}

func bsmBFS(firstNode nodeInfo, targetCantor int) (ansNode nodeInfo, err error) {
	queue := make([]nodeInfo, 0)
	queue = append(queue, firstNode)
	for {
		if len(queue) == 0 {
			break
		}
		currentNode := queue[0]
		// pop
		queue = queue[1:]

		currentNodeCantor := CantorExpansion(currentNode.Arr)
		if currentNodeCantor == targetCantor {
			currentNode.Arrival = true
			ansNode = currentNode
			break
		}

		// view status
		fmt.Printf("\r%+v | 剩余队列：%v", currentNode, len(queue))

		for i := 0; i < 4; i++ {
			nextRow, nextColumn := currentNode.ZeroPosition/matrixRow, currentNode.ZeroPosition%matrixRow
			nextRow, nextColumn = nextRow+moveRule[0][i], nextColumn+moveRule[1][i]
			if nextRow < 0 || nextColumn < 0 || nextRow >= matrixRow || nextColumn >= matrixRow {
				continue
			}
			nextArr := new([]int)
			err = DeepCopy(currentNode.Arr, &nextArr)
			nextZeroPosition := nextRow*matrixRow + nextColumn
			(*nextArr)[currentNode.ZeroPosition], (*nextArr)[nextZeroPosition] = (*nextArr)[nextZeroPosition], (*nextArr)[currentNode.ZeroPosition]
			if err != nil {
				return
			}
			nextNodeStep := currentNode.Step + 1
			nextNodeArrCantor := CantorExpansion(*nextArr)
			if nodeCantorArr[nextNodeArrCantor].Step > nextNodeStep || nodeCantorArr[nextNodeArrCantor].Step == 0 {
				nextNode := nodeInfo{
					Step:         nextNodeStep,
					ParentCantor: currentNodeCantor,
					Arr:          *nextArr,
					ZeroPosition: nextZeroPosition,
					Arrival:      false,
				}
				nodeCantorArr[nextNodeArrCantor] = nextNode
				queue = append(queue, nextNode)
			}
		}
	}
	return
}
