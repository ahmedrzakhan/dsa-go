package main

import "fmt"

const setSize = 10000 // Assuming a set size of 10^4

type nodeHS struct {
	key  int
	next *node
}

type MyHashSet struct {
	set []*node // Using a slice of pointers to node
}

func ConstructorHS() MyHashSet {
	return MyHashSet{set: make([]*node, setSize)}
}

// node setup not really required for this problem
func (hs *MyHashSet) add(key int) {
	index := key % setSize
	if hs.set[index] == nil {
		hs.set[index] = &node{key: key}
		return
	}
	current := hs.set[index]
	for current.next != nil {
		if current.next.key == key {
			return
		}
		current = current.next
	}
	current.next = &node{key: key}
}

func (hs *MyHashSet) remove(key int) {
	index := key % setSize
	if hs.set[index] == nil {
		return
	}
	if hs.set[index].key == key {
		hs.set[index] = hs.set[index].next
		return
	}
	current := hs.set[index]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (hs *MyHashSet) contains(key int) bool {
	index := key % setSize
	current := hs.set[index]
	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	hashSet := ConstructorHS()
	hashSet.add(1)
	hashSet.add(1)
	hashSet.add(1)
	hashSet.add(2)
	fmt.Println(hashSet.contains(1)) // true
	fmt.Println(hashSet.contains(3)) // false
	hashSet.add(2)
	fmt.Println(hashSet.contains(2)) // true
	hashSet.remove(2)
	fmt.Println(hashSet.contains(2)) // false
}
