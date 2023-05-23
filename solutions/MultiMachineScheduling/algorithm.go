package MultiMachineScheduling

import "sort"

type branch struct {
	MaxCost     int
	Assigned    int
	MachineCost []int
	MachineTask [][]int
}

func greedy(tasks []*task, machine int) (result *branch) {
	// 降序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Cost > tasks[j].Cost
	})
	originB := &branch{
		MaxCost:     tasks[0].Cost,
		Assigned:    1,
		MachineCost: make([]int, machine, machine),
		MachineTask: make([][]int, machine, machine),
	}
	originB.MachineCost[0] = tasks[0].Cost
	originB.MachineTask[0] = append(originB.MachineTask[0], tasks[0].Id)
	var dfs func(b *branch)
	dfs = func(b *branch) {
		if b.Assigned == len(tasks) {
			result = b
			return
		}
		choice := -1
		for i := 0; i < machine; i++ {
			if choice == -1 || b.MachineCost[i] < b.MachineCost[choice] {
				choice = i
			}
		}
		nextB := &branch{
			MaxCost:     b.MaxCost,
			Assigned:    b.Assigned,
			MachineCost: append([]int{}, b.MachineCost...),
			MachineTask: append([][]int{}, b.MachineTask...),
		}
		nextB.MachineCost[choice] += tasks[nextB.Assigned].Cost
		nextB.MachineTask[choice] = append(nextB.MachineTask[choice], tasks[nextB.Assigned].Id)
		maxCost := 0
		for _, cost := range nextB.MachineCost {
			if cost > maxCost {
				maxCost = cost
			}
		}
		nextB.MaxCost = maxCost
		nextB.Assigned++
		dfs(nextB)
	}
	dfs(originB)
	return
}
