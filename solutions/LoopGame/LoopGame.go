package LoopGame

import "fmt"

type kid struct {
	Id   int
	Last *kid
	Next *kid
}

func Run() (err error) {
	begin, n, k, err := input()
	if err != nil {
		return
	}
	winner := loopGame(begin, n, k)
	fmt.Println(winner)
	return
}

func input() (begin *kid, n, k int, err error) {
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&k)
	if err != nil {
		return
	}
	begin = &kid{Id: 1}
	last := begin
	var head *kid
	for i := 2; i <= n; i++ {
		head = &kid{Id: i}
		head.Last = last
		last.Next = head
		last = last.Next
	}
	head.Next = begin
	begin.Last = head
	return
}

func loopGame(head *kid, n, k int) (winner int) {
	count := 0
	for {
		if n <= 1 {
			winner = head.Id
			break
		}
		count++
		if count%k == 0 || count%10 == k {
			head.Last.Next = head.Next
			head.Next.Last = head.Last
			n--
		}
		head = head.Next
	}
	return
}
