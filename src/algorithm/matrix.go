package algorithm

import (
	"fmt"
	"github.com/pkg/errors"
)

func printMatrix(matrix []int, m, n int) (err error) {
	if m*n != len(matrix) || m*n == 0 {
		err = errors.New("error")
		return
	}

	var cursorPool = [][]int{
		[]int{1, 0},
		[]int{0, 1},
		[]int{-1, 0},
		[]int{0, -1},
	}

	var cursorIdx = 0
	var startIdx = []int{0, 0}

	printed := make([]int, m*n)

	for {
		fmt.Println(matrix[startIdx[0]*m+startIdx[1]])
		printed[startIdx[0]*m + startIdx[1]] = 1

		lastIdx := cursorIdx
		for {
			nextIdx := []int{startIdx[0] + cursorPool[lastIdx][0], startIdx[1] + cursorPool[lastIdx][1]}
			if nextIdx[0] >= m || nextIdx[0] < 0 || nextIdx[1] >= n || nextIdx[1] < 0 {
				cursorIdx = (cursorIdx + 1) % 4
			}else if printed[nextIdx[0]*m + nextIdx[1]] != 0{
				cursorIdx = (cursorIdx + 1) % 4
			}else{
				break
			}

			if cursorIdx == lastIdx{
				return
			}
		}

		cursorIdx = lastIdx
		startIdx[0] = startIdx[0] + cursorPool[cursorIdx][0]
		startIdx[1] = startIdx[1] + cursorPool[cursorIdx][1]
	}
}
