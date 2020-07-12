package main

import "sort"

type inter [][]int

func (in *inter) Len() int {
	return len(*in)
}

func (in *inter) Less(i, j int) bool {
	return (*in)[i][0] < (*in)[j][0]
}

func (in *inter) Swap(i, j int) {
	(*in)[i], (*in)[j] = (*in)[j], (*in)[i]
}

func mergeInter(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	inter := inter(intervals)
	sort.Sort(&inter)
	res := [][]int{}
	for i := 0; i < len(inter); i++ {
		if len(res) <= 0 || (len(res) >= 1 && res[len(res)-1][1] < inter[i][0]) {
			res = append(res, inter[i])
		} else {
			res[len(res)-1][1] = maxInter(res[len(res)-1][1], inter[i][1])
		}
	}
	return res
}

func maxInter(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
