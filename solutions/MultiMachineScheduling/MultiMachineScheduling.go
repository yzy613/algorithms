package MultiMachineScheduling

import (
	"fmt"
	"math/rand"
)

type task struct {
	Id   int
	Cost int
}

func Run() (err error) {
	n, m, err := inputFromConsole()
	if err != nil {
		return
	}
	tasks := generateData(n)
	printTasks(tasks)
	result := greedy(tasks, m)
	printBranch(result)
	return
}

func inputFromConsole() (n, m int, err error) {
	_, err = fmt.Scan(&n, &m)
	return
}

func generateData(n int) (tasks []*task) {
	tasks = make([]*task, n, n)
	for i := range tasks {
		tasks[i] = &task{
			Id:   i,
			Cost: rand.Intn(19) + 2,
		}
	}
	return
}

func printTasks(tasks []*task) {
	fmt.Print("[")
	for i, t := range tasks {
		fmt.Print(t.Cost)
		if i != len(tasks)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("]")
	fmt.Println()
}

func printBranch(b *branch) {
	fmt.Println("MaxCost:", b.MaxCost)
	fmt.Println("MachineCost:", b.MachineCost)
	fmt.Println("MachineTask:")
	for _, tasks := range b.MachineTask {
		fmt.Print("[")
		for i, t := range tasks {
			fmt.Print(t + 1)
			if i != len(tasks)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("]")
		fmt.Println()
	}
}
