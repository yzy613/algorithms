package ShipLoading

import (
	"fmt"
	"math/rand"
)

func Run() (err error) {
	goodsNum, err := inputFromConsole()
	if err != nil {
		return
	}
	goodsWeight, cargoShipLoadCapacity := generateRandData(goodsNum)
	printData(goodsWeight, cargoShipLoadCapacity)
	success, cargoShipLoadGoods := loadGoods(goodsWeight, cargoShipLoadCapacity)
	if !success {
		fmt.Println("无装载方案")
		return
	}
	for i := 0; i < len(cargoShipLoadCapacity); i++ {
		fmt.Printf("货船 %d:\n", i+1)
		fmt.Print("装载的货物为：")
		for j := 0; j < len(cargoShipLoadGoods[i]); j++ {
			if cargoShipLoadGoods[i][j] != 0 {
				fmt.Printf("%d ", j+1)
			}
		}
		fmt.Println()
		fmt.Printf("剩余载重量：%d\n", cargoShipLoadCapacity[i])
		if i != len(cargoShipLoadCapacity)-1 {
			fmt.Println("---")
		}
	}
	return
}

func inputFromConsole() (goodsNum int, err error) {
	_, err = fmt.Scan(&goodsNum)
	if goodsNum <= 10 {
		err = fmt.Errorf("货物数量必须大于 10")
	}
	return
}

func generateRandData(goodsNum int) (goodsWeight, cargoShipLoadCapacity []int) {
	totalWeight := 0
	for i := 0; i < goodsNum; i++ {
		goodsWeight = append(goodsWeight, rand.Intn(20)+1)
		totalWeight += goodsWeight[i]
	}
	// cargo ship 1
	floorLoadCapacity := totalWeight * 3 / 10
	floatingRange := totalWeight*6/10 - floorLoadCapacity
	cargoShipLoadCapacity = append(cargoShipLoadCapacity, rand.Intn(floatingRange)+floorLoadCapacity)
	// cargo ship 2
	floorLoadCapacity = totalWeight - cargoShipLoadCapacity[0]
	floatingRange = totalWeight*12/10 - cargoShipLoadCapacity[0] - floorLoadCapacity
	cargoShipLoadCapacity = append(cargoShipLoadCapacity, rand.Intn(floatingRange)+floorLoadCapacity)
	return
}

func loadGoods(goodsWeight, cargoShipLoadCapacity []int) (success bool, cargoShipLoadGoods [][]int) {
	goodsChoice := make([]bool, len(goodsWeight), len(goodsWeight))
	cargoShipLoadGoods = make([][]int, len(cargoShipLoadCapacity), len(cargoShipLoadCapacity))
	for i := 0; i < len(cargoShipLoadCapacity); i++ {
		cargoShipLoadGoods[i] = make([]int, len(goodsWeight), len(goodsWeight))
	}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(goodsWeight) {
			success = true
			return
		}
		for i := 0; i < len(cargoShipLoadCapacity); i++ {
			if cargoShipLoadCapacity[i] >= goodsWeight[index] && !goodsChoice[index] {
				cargoShipLoadCapacity[i] -= goodsWeight[index]
				goodsChoice[index] = true
				cargoShipLoadGoods[i][index] = goodsWeight[index]
				dfs(index + 1)
				if success {
					return
				}
				cargoShipLoadCapacity[i] += goodsWeight[index]
				goodsChoice[index] = false
				cargoShipLoadGoods[i][index] = 0
			}
		}
	}
	dfs(0)
	return
}

func printData(goodsWeight, cargoShipLoadCapacity []int) {
	fmt.Println("---")
	totalWeight := 0
	fmt.Println("货物重量：")
	for i := 0; i < len(goodsWeight); i++ {
		fmt.Printf("%d ", goodsWeight[i])
		totalWeight += goodsWeight[i]
	}
	fmt.Println()
	fmt.Println("货船载重量：")
	for i := 0; i < len(cargoShipLoadCapacity); i++ {
		fmt.Printf("%d ", cargoShipLoadCapacity[i])
	}
	//s := struct {
	//	TotalWeight          int
	//	FloorLoadCapacity    int
	//	FloatingRange        int
	//	FloorLoadCapacitySec int
	//	FloatingRangeSec     int
	//}{
	//	TotalWeight: totalWeight,
	//}
	//s.FloorLoadCapacity = totalWeight * 3 / 10
	//s.FloatingRange = totalWeight*6/10 - s.FloorLoadCapacity
	//s.FloorLoadCapacitySec = totalWeight - cargoShipLoadCapacity[0]
	//s.FloatingRangeSec = totalWeight*12/10 - cargoShipLoadCapacity[0] - s.FloorLoadCapacitySec
	//fmt.Printf("\n%+v", s)
	fmt.Println("\n---")
}
