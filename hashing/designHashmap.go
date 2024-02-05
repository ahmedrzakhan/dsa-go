package main

import "fmt"

/**
706. Design HashMap

Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap.
If the key already exists in the map, update the corresponding value.
int get(int key) returns the value to which the specified key is mapped,
or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map
contains the mapping for the key.

Example 1:

Input
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
Output
[null, null, null, 1, -1, null, 1, null, -1]

Explanation
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // The map is now [[1,1]]
myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]

Constraints:

0 <= key, value <= 106
At most 104 calls will be made to put, get, and remove.
*/

// TC - O(1), SC - O(N)
type PairHM struct {
	Key   int
	Value int
}

type MyHashMap struct {
	items []*PairHM
}

func ConstructorHM() MyHashMap {
	return MyHashMap{items: make([]*PairHM, 1000001)} // Assuming we have keys in range [0, 1000000]
}

func (hm *MyHashMap) Put(key int, value int) {
	hm.items[key] = &PairHM{Key: key, Value: value}
}

func (hm *MyHashMap) Get(key int) int {
	if hm.items[key] != nil {
		return hm.items[key].Value
	}
	return -1 // Return -1 if the key does not exist
}

func (hm *MyHashMap) Remove(key int) {
	hm.items[key] = nil
}

func mainHM() {
	hashMap := ConstructorHM()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Println(hashMap.Get(1)) // returns 1
	fmt.Println(hashMap.Get(3)) // returns -1 (not found)
	hashMap.Put(2, 1)           // update the existing value
	fmt.Println(hashMap.Get(2)) // returns 1
	hashMap.Remove(2)           // remove the mapping for 2
	fmt.Println(hashMap.Get(2)) // returns -1 (not found)
}
