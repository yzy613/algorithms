package BaShuMa

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/util/gconv"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	foundCtx  = "found"
	cancelCtx = "cancel"
)

func start(beginNode *node) (err error) {
	// 初始化散列表
	hashTable = gmap.NewStrAnyMap(true)
	hashTable.Set(matrix2Key(*beginNode.Matrix), beginNode)
	// 初始化队列
	msgQueue = gqueue.New()
	queue = gqueue.New()
	// 放入第一个节点
	queue.Push(beginNode)
	// 设置上下文
	ctx, cancel := context.WithCancel(context.Background())
	cancelF := func() {
		cancel()
		time.Sleep(10 * time.Millisecond)
		msgQueue.Close()
		queue.Close()
	}
	var foundF foundFunc = func(dstNode *node) {
		cancelF()
		time.Sleep(100 * time.Millisecond)
		printResult(dstNode)
	}
	ctx = context.WithValue(ctx, foundCtx, foundF)
	ctx = context.WithValue(ctx, cancelCtx, cancelF)
	// 设置并发等待组
	wg := sync.WaitGroup{}
	numCpu := runtime.NumCPU()
	wg.Add(numCpu)
	// 并发执行
	for i := 0; i < numCpu; i++ {
		go bfs(ctx, &wg)
	}
	wg.Add(1)
	go printMessage(ctx, &wg)
	wg.Wait()
	return
}

func bfs(ctx context.Context, wg *sync.WaitGroup) {
	for {
		// 处理上下文
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
		}
		// 队列为空
		if queue.Len() <= 0 {
			continue
		}
		var curr *node
		{
			temp := queue.Pop()
			// 检查队列是否关闭
			if temp == nil {
				continue
			}
			curr = temp.(*node)
		}
		// 当前节点的 key
		currKey := matrix2Key(*curr.Matrix)
		// 找到最终节点
		if currKey == dstMatrixKey {
			ctx.Value(foundCtx).(foundFunc)(curr)
			wg.Done()
			return
		}
		// 推送消息
		msgQueue.Push(packMessage(curr, queue.Len()))
		// 遍历所有可能的移动
		for i := 0; i < 4; i++ {
			nextRow, nextColumn := curr.ZeroPosition/matrixSize+moveRule[0][i], curr.ZeroPosition%matrixSize+moveRule[1][i]
			if nextRow < 0 || nextRow >= matrixSize || nextColumn < 0 || nextColumn >= matrixSize {
				continue
			}
			nextZeroPosition := nextRow*matrixSize + nextColumn
			nextMatrix := make([]int, matrixSize*matrixSize)
			copy(nextMatrix, *curr.Matrix)
			nextMatrix[curr.ZeroPosition], nextMatrix[nextZeroPosition] = nextMatrix[nextZeroPosition], nextMatrix[curr.ZeroPosition]
			nextNode := &node{
				Step:         curr.Step + 1,
				ZeroPosition: nextZeroPosition,
				Matrix:       &nextMatrix,
				ParentKey:    currKey,
			}
			nextNodeKey := matrix2Key(nextMatrix)
			if hashTable.Contains(nextNodeKey) {
				front := hashTable.Get(nextNodeKey).(*node)
				if front.Step > nextNode.Step {
					hashTable.Set(nextNodeKey, nextNode)
				}
			} else {
				hashTable.Set(nextNodeKey, nextNode)
				queue.Push(nextNode)
			}
		}
	}
}

func matrix2Key(matrix []int) (key string) {
	keyBuilder := strings.Builder{}
	for i := 0; i < len(matrix); i++ {
		keyBuilder.WriteString(gconv.String(matrix[i]))
	}
	key = keyBuilder.String()
	return
}

func printMessage(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			if msgQueue.Len() <= 0 {
				time.Sleep(100 * time.Millisecond)
				if msgQueue.Len() <= 0 {
					ctx.Value(cancelCtx).(func())()
					fmt.Println("\n可能无解")
					continue
				}
			}
			fmt.Printf("\r%v", msgQueue.Pop().(string))
		}
	}
}

func packMessage(curr *node, queueLen int64) (msg string) {
	msg = fmt.Sprintf("Step:%v ZeroPosition:%v Matrix:%v ParentKey:%v | 剩余队列：%v",
		curr.Step, curr.ZeroPosition, *curr.Matrix, curr.ParentKey, queueLen)
	return
}
