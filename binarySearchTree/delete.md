* minimumin in right subtree and maximum in left subtree maintains the binary search tree properties

Mininmum in right subtree
* given bst

    20
   /  \
  15   25
 /  \    \
10   18   30

*  delete node 20, 25 is the minimum in right subtree 

     25
    /  \
   15   25
  /  \    \
 10   18   30

* remove duplicate 25

     25
    /  \
   15   30
  /  \
 10   18

Maximum in left subtree
* given bst

    20
   /  \
  15   25
 /  \    \
10   18   30

* delete node 20, 18 is maximum in left subtree

     18
    /  \
   15   25
  /  \    \
 10   18   30

* remove duplicate 18

     18
    /  \
   15   25
  /       \
 10       30

/**
1. if the root is nil return nil
2. find if the node exists in O(LogN) or O(H)
3. if the node to be deleted does not exist do nothng return nil (first if block)
4. if the node to be deleted does not have any descendants return nil
5. if the node to be deleted does not have any descendants in right return left
6. if the node to be deleted does not have any descendants in left return right
*/

* findMin function starts at root. Right and moves leftwards (current = current.Left) until it finds the leftmost node in this subtree. This leftmost node is the node with the smallest value in the right subtree of the original node to be deleted.
