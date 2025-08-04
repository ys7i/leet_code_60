package main

type UnionFind struct {
    depths []int
    parents []int
    areaSizes []int
}

const (
    SEA = 0
    ISLAND = 1
)

func maxAreaOfIsland2(grid [][]int) int {
    rowCount, colCount := len(grid), len(grid[0])
    depths := make([]int, rowCount * colCount)
    parents := make([]int, rowCount * colCount)
    areaSizes := make([]int, rowCount * colCount)

    uf := &UnionFind{ depths: depths, parents: parents, areaSizes: areaSizes }

    directions := [2][2]int{ {-1, 0}, {0, -1}}

    for i, row := range grid {
        for j, value := range row {
            if value == SEA {
                continue
            }

            index := colCount * i + j
            uf.parents[index] = index
            uf.areaSizes[index] = 1

           for _, direction := range directions {
                otherI := i + direction[0]
                otherJ := j + direction[1]
                if otherI < 0 || otherI >= rowCount {
                    continue
                }
                if otherJ < 0 || otherJ >= colCount {
                    continue
                }

                if grid[otherI][otherJ] == ISLAND {
                    otherIndex := otherI * colCount + otherJ
                    uf.Union(index, otherIndex)
                }
           }
        }
    }

    maxArea := 0
    for _, value := range uf.areaSizes {
        if value > maxArea {
            maxArea = value
        }
    }

    return maxArea
}

func (uf *UnionFind) Union (i, j int) {
    iRoot := uf.Find(i)
    jRoot := uf.Find(j)

    if iRoot == jRoot {
        return
    }

    if uf.depths[iRoot] > uf.depths[jRoot] {
        uf.parents[jRoot] = uf.parents[iRoot]
        uf.areaSizes[iRoot] += uf.areaSizes[jRoot]
    } else if uf.depths[jRoot] > uf.depths[iRoot] {
        uf.parents[iRoot] = uf.parents[jRoot]
        uf.areaSizes[jRoot] += uf.areaSizes[iRoot]
    } else {
        uf.parents[jRoot] = uf.parents[iRoot]
        uf.areaSizes[iRoot] += uf.areaSizes[jRoot]
        uf.depths[iRoot] += 1
    }
}

func (uf *UnionFind) Find(i int) int {
    if uf.parents[i] != i {
        root := uf.Find(uf.parents[i])
        uf.parents[i] = root
    }
    return uf.parents[i]
}