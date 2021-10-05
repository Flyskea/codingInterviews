package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	nums      []int
	numsIndex map[int]int
	rand      *rand.Rand
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	right := RandomizedSet{
		numsIndex: make(map[int]int),
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return right
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.numsIndex[val]; ok {
		return false
	}
	rs.numsIndex[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.numsIndex[val]; !ok {
		return false
	}
	index := rs.numsIndex[val]
	t := rs.nums[len(rs.nums)-1]
	rs.nums[index] = t
	rs.numsIndex[t] = index

	delete(rs.numsIndex, val)
	rs.nums = rs.nums[:len(rs.nums)-1]
	return true
}

/** Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	index := rs.rand.Intn((len(rs.nums)))
	return rs.nums[index]
}

func main() {
	rs := Constructor()
	rs.Insert(2)
	rs.Insert(4)
	fmt.Println(rs.GetRandom())
	rs.Remove(2)
	fmt.Println(rs.GetRandom())
}
