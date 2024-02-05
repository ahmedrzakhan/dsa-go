package main

import "fmt"

const bucketSize = 1000

type node struct {
	key, value int
	next       *node
}

type MyHashMapA struct {
	buckets [](*node)
}

func Constructor() MyHashMapA {
	return MyHashMapA{buckets: make([](*node), bucketSize)}
}

func (m *MyHashMapA) hash(key int) int {
	return key % bucketSize
}

func (m *MyHashMapA) Put(key int, value int) {
	index := m.hash(key)
	if m.buckets[index] == nil {
		m.buckets[index] = &node{key, value, nil}
		return
	}
	current := m.buckets[index]
	for current.next != nil {
		if current.next.key == key {
			current.next.value = value
			return
		}
		current = current.next
	}
	current.next = &node{key, value, nil}
}

func (m *MyHashMapA) Get(key int) int {
	index := m.hash(key)
	current := m.buckets[index]
	for current != nil {
		if current.key == key {
			return current.value
		}
		current = current.next
	}
	return -1
}

func (m *MyHashMapA) Remove(key int) {
	index := m.hash(key)
	if m.buckets[index] == nil {
		return
	}
	if m.buckets[index].key == key {
		m.buckets[index] = m.buckets[index].next
		return
	}
	current := m.buckets[index]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func mainHMA() {
	hashMap := Constructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Println(hashMap.Get(1)) // returns 1
	fmt.Println(hashMap.Get(3)) // returns -1 (not found)
	hashMap.Put(2, 1)           // update the existing value
	fmt.Println(hashMap.Get(2)) // returns 1
	hashMap.Remove(2)           // remove the mapping for 2
	fmt.Println(hashMap.Get(2)) // returns -1 (not found)
}
