package main

const MAX_NUM = 10_000

type IntHeap []int

type KthLargest1 struct {
    intHeap IntHeap
    k int
}

func IntHeapConstructor(nums []int) IntHeap {
    intHeap := make(IntHeap, len(nums))
    for i, value := range nums {
        intHeap[i] = value
        intHeap.sortFromLeave(i)
    }

    return intHeap
}

func (this IntHeap) Min() int {
    return this[0]
}

func (this *IntHeap) Push(val int) {
    *this = append(*this, val)
    this.sortFromLeave(len(*this) - 1)
}

func (h *IntHeap) Pop() int {
    maxValue := (*h)[len(*h) - 1]
    h.swap(0, len(*h) - 1)
    *h = (*h)[0:(len(*h) - 1)]
    h.sortFromRoot()

    return maxValue
}

func (this *IntHeap) sortFromRoot() {
    
    index := 0

    for {
        leftChildIndex := index * 2 + 1
        rightChildIndex := index * 2 + 2
        leftChild := MAX_NUM
        rightChild := MAX_NUM
        if leftChildIndex < len(*this) {
            leftChild = (*this)[leftChildIndex]
        }

        if rightChildIndex < len(*this) {
            rightChild = (*this)[rightChildIndex]
        }

        if leftChild <= rightChild && leftChild < (*this)[index] {
            this.swap(leftChildIndex, index)
            index = leftChildIndex
            continue
        } 

        if rightChild < leftChild && rightChild < (*this)[index] {
            this.swap(rightChildIndex, index)
            index = rightChildIndex
            continue
        }
        return
    }
}

func (this *IntHeap) sortFromLeave(leafIndex int) {
    for {
        if leafIndex == 0 {
            return
        }
        parentIndex := (leafIndex - 1) / 2
        if (*this)[parentIndex] < (*this)[leafIndex] {
            return
        }
        this.swap(parentIndex, leafIndex)
        leafIndex = parentIndex
    }
}

func (this *IntHeap) swap (i int, j int) {
    (*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func Constructor3(k int, nums []int) KthLargest1 {
    intHeap := IntHeapConstructor(nums)
    return KthLargest1{ intHeap, k }
}


func (this *KthLargest1) Add3(val int) int {
    this.intHeap.Push(val)
    for len(this.intHeap) > this.k {
        this.intHeap.Pop()
    }

    return this.intHeap.Min()
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */