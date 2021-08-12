package xslice

import (
	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=$GOFILE -out=int.go gen "KeyType=int8,uint8,int16,uint16,int,uint,int32,uint32,int64,uint64"

//go:generate genny -in=$GOFILE -out=float.go gen "KeyType=float64"

//go:generate genny -in=$GOFILE -out=string.go gen "KeyType=string"

type KeyType generic.Type

// ContainKeyType check if target is in ss
func ContainKeyType(s []KeyType, target KeyType) bool {
	for _, val := range s {
		if val == target {
			return true
		}
	}
	return false
}

// RemoveKeyType remove empty target elements from ss
func RemoveKeyType(s []KeyType, target KeyType) []KeyType {
	var ret = make([]KeyType, len(s) -1)
	var targetFind bool
	for index, val := range s {
		if val == target {
			targetFind = true
			continue
		}

		if targetFind {
			ret[index-1] = val
		}else {
			ret[index] = val
		}
	}
	return ret
}

// ReverseKeyType reverse the input slice
func ReverseKeyType(s []KeyType) []KeyType {
	if len(s) < 2 {
		return s
	}

	start := 0
	end := len(s) - 1
	for start < end {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp
		start++
		end--
	}
	return s
}

// DiffKeyType computes the difference of two slices
// return a slice containing all the entries from s1 but not in s2
func DiffKeyType(s1 []KeyType, s2 []KeyType) []KeyType {
	ret := make([]KeyType, 0)
	for _, val := range s1 {
		if !ContainKeyType(s2, val) {
			ret = append(ret, val)
		}
	}
	return ret
}

// IntersectKeyType computes the intersection of two slices
// return a slice containing the entries exist in s1 and s2
func IntersectKeyType(s1 []KeyType, s2 []KeyType) []KeyType {
	ret := make([]KeyType, 0)
	for _, val := range s1 {
		if ContainKeyType(s2, val) {
			ret = append(ret, val)
		}
	}
	return ret
}

// UniqueKeyType Removes duplicate values from slice
func UniqueKeyType(s1 []KeyType) []KeyType {
	unique := make(map[KeyType]interface{})
	for _, val := range s1 {
		unique[val] = nil
	}

	ret := make([]KeyType, 0, len(unique))
	for key := range unique {
		ret = append(ret, key)
	}
	return ret
}

// MergeKeyType merge one or more arrays
func MergeKeyType(s1 []KeyType, s2 ...[]KeyType) []KeyType {
	if len(s2) == 0 {
		return s1
	}
	size := len(s1)
	for _, s := range s2 {
		size += len(s)
	}
	ret := make([]KeyType, 0,size)
	ret = append(ret,s1...)
	for _, s := range s2 {
		ret = append(ret,s...)
	}
	return ret
}
