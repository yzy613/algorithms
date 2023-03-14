package sudoku

func dfs(addr int) (success bool) {
	if addr >= dealArray.Len() {
		return true
	}
	var currNode *node
	{
		temp, ok := dealArray.Get(addr)
		if !ok {
			return
		}
		currNode = temp.(*node)
	}
	for _, val := range *currNode.Candidates {
		if isValid(currNode.Addr, val) {
			matrix[currNode.Addr].Val = val
			if dfs(addr + 1) {
				return true
			}
			matrix[currNode.Addr].Val = 0
		}
	}
	return
}
