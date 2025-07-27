package main

func numIslands1(grid [][]byte) int {
    visitedPoint := make(map[[2]int]struct{}, len(grid) * len(grid[0]))
    islandCount := 0

    for i, _ := range grid {
        for j, _ := range grid[i] {
            _, isVisited := visitedPoint[[2]int{i, j}]
            if isVisited {
                continue
            }

            if grid[i][j] == '1' {
                islandCount += 1;
                removeIsland(grid, i, j, visitedPoint)
            }
        }
    }
    
    return islandCount
}

func removeIsland(grid [][]byte, row int, column int, visitedPoint map[[2]int]struct{}) {
    if _, ok := visitedPoint[[2]int{row, column}]; ok {
        return
    }
    if row < 0 || len(grid) <= row {
        return
    }
    if column < 0 || len(grid[row]) <= column {
        return
    }
    if grid[row][column] == '0' {
        return
    }
    visitedPoint[[2]int{row, column}] = struct{}{}
    grid[row][column] = 0;
    removeIsland(grid, row - 1, column, visitedPoint)
    removeIsland(grid, row + 1, column, visitedPoint)
    removeIsland(grid, row, column - 1, visitedPoint)
    removeIsland(grid, row, column + 1, visitedPoint)
}

// 時間計算量 O(n * m)
// 空間計算量 O(n * m)