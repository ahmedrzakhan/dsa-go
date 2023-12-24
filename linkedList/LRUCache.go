package main

import "fmt"

/**
146. LRU Cache

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair
to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.

Example 1:

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4

Constraints:

1 <= capacity <= 3000
0 <= key <= 104
0 <= value <= 105
At most 2 * 105 calls will be made to get and put.
*/

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

// TC - O(1), SC - O(capacity)
func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, found := this.cache[key]; found {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, found := this.cache[key]; found {
		node.value = value
		this.moveToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		delete(this.cache, this.tail.key)
		this.removeTail()
	}

	newNode := &Node{key: key, value: value}
	this.cache[key] = newNode
	this.addToFront(newNode)
}

func (this *LRUCache) addToFront(node *Node) {
	node.next = this.head
	node.prev = nil

	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) moveToFront(node *Node) {
	if node == this.head {
		return
	}

	// Detach the node
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// If node is the tail, update the tail
	if node == this.tail {
		this.tail = node.prev
	}

	// Move the node to the front
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	// If the list was empty, also set the tail
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) removeTail() {
	if this.tail == nil {
		return
	}

	if this.tail.prev != nil {
		this.tail.prev.next = nil
	} else {
		this.head = nil
	}
	this.tail = this.tail.prev
}

func mainLRU() {
	cache := ConstructorLRU(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	fmt.Println(cache.Get(1)) // Returns 1
	cache.Put(4, 4)           // Evicts key 2
	fmt.Println(cache.Get(2)) // Returns -1 (not found)
	fmt.Println(cache.Get(3)) // Returns 3
	fmt.Println(cache.Get(4)) // Returns 4
}
