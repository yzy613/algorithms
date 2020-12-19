package baShuMa

var (
	// 一个 Row * Row 的矩阵
	matrixRow      int
	// 初始矩阵
	srcArr         []int
	// 目标矩阵
	dstArr         []int
	// 存储阶乘值
	factorialArr   []int
	// 记录 node 状态
	nodeCantorArr  []nodeInfo
	// 移动规则
	moveRule       = [2][4]int{
		//右,下,左,上
		{0, 1, 0, -1}, // row
		{1, 0, -1, 0}} // column
)

type nodeInfo struct {
	Step         int
	ParentCantor int
	Arr          []int
	ZeroPosition int
	Arrival      bool
}
