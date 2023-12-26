The Need for Full Traversal:

The kth smallest element is not necessarily in the left subtree alone. It depends on the total number of nodes in the left subtree and the value of k.
If the left subtree of a node contains fewer than k nodes, the kth smallest element cannot be in that subtree. The search must continue to the root node and potentially into the right subtree.
For example, if the left subtree has only 2 nodes and k is 4, you need to count the root node and continue to the right subtree to find the 4th smallest element.

Time Complexity (TC):
O(H + k): This is the time complexity where H is the height of the tree.
Initially, the algorithm traverses down to the leftmost node, which takes O(H) time (H is the height of the tree).
Then, it starts popping elements from the stack and visiting nodes. In the worst case, it might have to pop and visit k nodes.
So, in total, the worst-case time complexity is O(H + k).
In a balanced BST, H is O(log n), making the time complexity O(log n + k). In a skewed BST (which is more like a list), H is O(n), making the time complexity O(n) in the worst case.

Space Complexity (SC):
O(H): The space complexity is mainly due to the stack used in the algorithm.
The stack will hold at most H TreeNode pointers, where H is the height of the tree.
In a balanced BST, this would be O(log n), and in a skewed BST, this would be O(n).
