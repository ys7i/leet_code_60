package main

import "container/heap"

type PairHeapElement struct {
    index1 int
    index2 int
    sum int
}

type PairHeap []*PairHeapElement

func (h *PairHeap) Len() int {
    return len(*h)
}

func (h *PairHeap) Less(i, j int) bool {
    return (*h)[i].sum < (*h)[j].sum
}

func (h *PairHeap) Swap(i, j int) {
    (*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *PairHeap) Push(e any) {
    *h = append(*h, e.(*PairHeapElement))
}

func (h *PairHeap) Pop() any {
    length := len(*h)
    popped := (*h)[length - 1]
    *h = (*h)[:length - 1]
    return popped
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    pairHeap := &PairHeap{}
    heap.Init(pairHeap)

    smallPairs := make([][]int, 0, k)
    smallPairs = append(smallPairs, []int{nums1[0], nums2[0]})

    heaped := make(map[[2]int]struct{}, k)

    if len(nums1) > 1 {
        heap.Push(pairHeap, &PairHeapElement{ 1, 0, nums1[1] + nums2[0] })
        heaped[[2]int{1, 0}] = struct{}{}
    }
    if len(nums2) > 1 {
        heap.Push(pairHeap, &PairHeapElement{ 0, 1, nums1[0] + nums2[1] })
        heaped[[2]int{0, 1}] = struct{}{}
    }
    

    for pairHeap.Len() > 0 &&  len(smallPairs) < k {
        popped := heap.Pop(pairHeap).(*PairHeapElement)
        smallPairs = append(smallPairs, []int{ nums1[popped.index1], nums2[popped.index2] })

        if len(nums1) > popped.index1 + 1 {
            if _, ok := heaped[[2]int{ popped.index1 + 1, popped.index2 }]; !ok {
                heap.Push(pairHeap, &PairHeapElement{ popped.index1 + 1, popped.index2, nums1[popped.index1 + 1] + nums2[popped.index2]  })
                heaped[[2]int{popped.index1 + 1, popped.index2}] = struct{}{}
            }
        }

        if len(nums2) > popped.index2 + 1 {
            if _, ok := heaped[[2]int{ popped.index1, popped.index2 + 1 }]; !ok {
                heap.Push(pairHeap, &PairHeapElement{ popped.index1, popped.index2 + 1, nums1[popped.index1] + nums2[popped.index2 + 1]  })
                heaped[[2]int{popped.index1, popped.index2 + 1}] = struct{}{}
            }
        }
    }

    return smallPairs
}