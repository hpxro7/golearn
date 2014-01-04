package utilities

import (
		"sort"
		rand "math/rand"
		)

type sortedIntMap struct {
	m map[int]float64
	s []int
}
 
func (sm *sortedIntMap) Len() int {
	return len(sm.m)
}
 
func (sm *sortedIntMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] < sm.m[sm.s[j]]
}
 
func (sm *sortedIntMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}
 
func SortIntMap(m map[int]float64) []int {
	sm := new(sortedIntMap)
	sm.m = m
	sm.s = make([]int, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

type sortedStringMap struct {
	m map[string]int
	s []string
}
 
func (sm *sortedStringMap) Len() int {
	return len(sm.m)
}
 
func (sm *sortedStringMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] < sm.m[sm.s[j]]
}
 
func (sm *sortedStringMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}
 
func SortStringMap(m map[string]int) []string {
	sm := new(sortedStringMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func RandomArray(n int) []float64 {
	ReturnedArray := make([]float64, n)
	for i := 0; i < n; i++ {
		ReturnedArray[i] = rand.Float64() * float64(rand.Intn(7))
	}
	return ReturnedArray
}