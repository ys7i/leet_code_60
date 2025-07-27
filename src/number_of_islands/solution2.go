package main

type UnionFind struct {
    ranks []int
    parents []int
    count int
}

func numIslands2(grid [][]byte) int {
    rows := len(grid)
    cols := len(grid[0])

    uf := NewUnionFind(rows * cols)

    directions := [2][2]int{ [2]int{-1, 0}, [2]int{ 0, -1 } }

    for i, row := range grid {
        for j, value := range row {
            index := i * cols + j

            if value != '1' {
                continue
            }

            uf.parents[index] = index
            uf.count++

            for _, direction := range directions {
                neighborRow := i + direction[0]
                neighborCols := j + direction[1]
                neighborIndex := neighborRow * cols + neighborCols
                if neighborRow >= 0 && neighborCols >= 0 && grid[neighborRow][neighborCols] == '1' {
                    uf.Union(index, neighborIndex)
                }
            }
        }
    }

    return uf.count
}

func NewUnionFind(n int) *UnionFind {
    ranks := make([]int, n)
    parents := make([]int, n)

    return &UnionFind { ranks, parents, 0 }
}

func (uf *UnionFind) Union(i, j int) {
    iRoot := uf.Find(i)
    jRoot := uf.Find(j)

    if iRoot == jRoot {
        return
    }

    if uf.ranks[iRoot] > uf.ranks[jRoot] {
        uf.parents[jRoot] = iRoot
        uf.count--
    } else if uf.ranks[iRoot] < uf.ranks[jRoot] {
        uf.parents[iRoot] = jRoot
        uf.count--
    } else {
        uf.parents[jRoot] = iRoot
        uf.ranks[iRoot]++
        uf.count--
    }
}

func (uf *UnionFind) Find(i int) int {
    if i != uf.parents[i] {
        uf.parents[i] = uf.Find(uf.parents[i])
    }
    return uf.parents[i]
}

// 時間計算量 O(n * m)
// 空間計算量 O(n * m)