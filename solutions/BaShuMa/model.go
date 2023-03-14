package BaShuMa

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gqueue"
)

var (
	// 矩阵大小
	matrixSize int
	// 源矩阵
	srcMatrix []int
	// 目标矩阵
	dstMatrix    []int
	dstMatrixKey string
	// 散列表
	hashTable *gmap.StrAnyMap
	// 队列
	queue *gqueue.Queue
	// 消息队列
	msgQueue *gqueue.Queue
	// 移动规则
	moveRule = [2][4]int{
		//右,下,左,上
		{0, 1, 0, -1}, // row
		{1, 0, -1, 0}} // column
)

type node struct {
	Step         int
	ZeroPosition int
	Matrix       *[]int
	ParentKey    string
}

type foundFunc func(dstNode *node)
