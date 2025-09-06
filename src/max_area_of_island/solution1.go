package main

const (
    SEA = 0
    ISLAND = 1
)

func maxAreaOfIsland2(grid [][]int) int {
    visitedArea := make(map[[2]int]struct{}, len(grid) * len(grid[0]))

    maxArea := 0

    for i, row := range grid {
        for j, _ := range row {
            area := calcAreaAndRemove(grid, visitedArea, i, j, 0)
            if area > maxArea {
                maxArea = area
            }
        }
    }
    return maxArea
}

func calcAreaAndRemove(grid[][]int, visitedArea map[[2]int]struct{}, i, j, area int) int {
    if i >= len(grid) || j >= len(grid[0]) {
        return area
    }
    if i < 0 || j < 0 {
        return area
    }

    if grid[i][j] == SEA {
        return area
    }

    if _, ok := visitedArea[[2]int{i, j}]; ok {
        return area
    }

    visitedArea[[2]int{i, j}] = struct{}{}
    area += 1
    area = calcAreaAndRemove(grid, visitedArea, i + 1, j, area)
    area = calcAreaAndRemove(grid, visitedArea, i - 1, j, area)
    area = calcAreaAndRemove(grid, visitedArea, i, j + 1, area)
    area = calcAreaAndRemove(grid, visitedArea, i, j - 1, area)

    return area
}

// 時間計算量 O(n * m)
// 空間計算量 O(n * m)