package main

import "container/heap"

type TopK struct {
	k     int
	times int
}

type TopKs []TopK

func (t *TopKs) Len() int {
	return len(*t)
}

func (t *TopKs) Less(i, j int) bool {
	return (*t)[i].times < (*t)[j].times
}

func (t *TopKs) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *TopKs) Push(x interface{}) {
	*t = append(*t, x.(TopK))
}

func (t *TopKs) Pop() interface{} {
	res := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return res
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] += 1
	}
	hp := TopKs{}
	for key, v := range mp {
		if k > len(hp) {
			heap.Push(&hp, TopK{
				k:     key,
				times: v,
			})
		} else {
			if v > hp[0].times {
				heap.Pop(&hp)
				heap.Push(&hp, TopK{
					k:     key,
					times: v,
				})
			}
		}
	}
	res := []int{}
	for len(hp) > 0 {
		res = append(res, heap.Pop(&hp).(TopK).k)
	}
	return res
}
