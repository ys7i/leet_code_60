package main

import (
	"cmp"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
    sortedStrToStrSlices := make(map[string][]string)

    for _, str := range strs {
        runes := []rune(str)
        slices.SortFunc(runes, cmp.Compare)
        sortedStr := string(runes)

        if strSlices, ok := sortedStrToStrSlices[sortedStr]; ok {
            sortedStrToStrSlices[sortedStr] = append(strSlices, str)
        } else {
            sortedStrToStrSlices[sortedStr] = []string { str }
        }
    }

    groupsStrs := make([][]string, 0)

    for _, strSlices := range sortedStrToStrSlices {
        groupsStrs = append(groupsStrs, strSlices)
    }

    return groupsStrs
}

// 時間計算量: strsの長さをnとすると O(n), strs[i]をsortしているが、strs[i].length <= 100なので固定として考えて良さそう
// 仮にstrs[i] = kとおくならO(n * k log k)