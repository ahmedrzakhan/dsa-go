package main

import "fmt"

/**
706. Design HashMap

Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:

void add(key) Inserts the value key into the HashSet.
bool contains(key) Returns whether the value key exists in the HashSet or not.
void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

Example 1:

Input
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
Output
[null, null, null, true, false, null, true, null, false]

Explanation
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // return True
myHashSet.contains(3); // return False, (not found)
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // return True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // return False, (already removed)

Constraints:

0 <= key <= 106
At most 104 calls will be made to add, remove, and contains.
*/

const setSizeDS = 10000 // Assuming a set size of 10^4

type MyHashSetDS struct {
	set map[int]bool
}

// TC - O(N), SC - O(N)
func ConstructorDS() MyHashSetDS {
	return MyHashSetDS{set: make(map[int]bool)}
}

// node setup not really required for this problem
func (hs *MyHashSetDS) add(key int) {
	hs.set[key] = true
}

func (hs *MyHashSetDS) remove(key int) {
	// index := key % setSize
	// hs.set[index] = false
	delete(hs.set, key)
}

func (hs *MyHashSetDS) contains(key int) bool {
	val, _ := hs.set[key]
	return val
}

func mainHS() {
	hashSet := ConstructorDS()
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
