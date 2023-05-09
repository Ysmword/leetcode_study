package main

import (
	"fmt"
)

func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var si, sj, bi, bj int
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' { // 玩家的位置
				si, sj = i, j
			} else if c == 'B' { // 箱子的位置
				bi, bj = i, j
			}
		}
	}

	f := func(i, j int) int {
		return i*n + j // 为啥取n
	}

	check := func(i, j int) bool { // 判断是否可以移动
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] != '#'
	}

	q := [][]int{{f(si, sj), f(bi, bj), 0}} // 数组第一个值为玩家，第二个值为箱子，第三个值为推动箱子次数
	vis := make([][]bool, m*n)              // 玩家和箱子的位置是否同时被访问过
	for i := range vis {
		vis[i] = make([]bool, m*n)
	}
	vis[f(si, sj)][f(bi, bj)] = true
	dirs := [5]int{-1, 0, 1, 0, -1} // 神奇，这里就能够表示上、下、左、右移动
	for len(q) > 0 {
		p := q[0]
		q = q[1:]                                       // 出队
		si, sj, bi, bj = p[0]/n, p[0]%n, p[1]/n, p[1]%n // 这个是根据公式计算【x*i+j】出来的
		d := p[2]                                       // 推动的次数
		if grid[bj][bj] == 'T' {
			return d
		}

		for k := 0; k < 4; k++ {
			sx, sy := si+dirs[k], sj+dirs[k+1]
			if !check(sx, sy) {
				continue
			}

			if sx == bi && sy == bj { // 箱子的位置
				bx, by := bi+dirs[k], bj+dirs[k+1]               // 推动箱子的位置
				if !check(bx, by) || vis[f(sx, sy)][f(bx, bj)] { // 不能移动, 已经移动过
					continue
				}
				vis[f(sx, sy)][f(bx, by)] = true
				q = append(q, []int{f(sx, sj), f(bx, by), d + 1})
			} else if !vis[f(sx, sy)][f(bi, bj)] { // 没有推动箱子，但是人动
				vis[f(sx, sy)][f(bi, bj)] = true
				q = append(q, []int{f(sx, sy), f(bi, bj), d + 1})
			}
		}
	}
	return -1
}

func main() {
	s := [][]byte{
		[]byte("######"),
		[]byte("#T####"),
		[]byte("#..B.#"),
		[]byte("####.#"),
		[]byte("#...S#"),
		[]byte("######"),
	}

	fmt.Println(minPushBox(s))
}
