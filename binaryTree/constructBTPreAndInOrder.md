The preorder and inorder traversal arrays can be used to determine the exact structure of the tree, including whether a node is a left or right child. The key is in understanding the order in which nodes appear in these arrays. For instance, in the preorder array, a node appears before its children, and in the inorder array, a node appears between its left and right children. This information is used to reconstruct the tree accurately in problems like the one from LeetCode you mentioned.

firstly builds left tree completely and then builds right tree
