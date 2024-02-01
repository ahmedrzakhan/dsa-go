Count Connected Components: After all edges have been processed, the function iterates through all nodes from 0 to n-1. For each node, it calls the Find function to get the root parent. If a node is its own parent (uf.Find(i) == i), it means this node is the root of a connected component. The function increments componentCount each time it finds such a root node.

Brute force:
Initialization: Represent the graph as an adjacency list for easy access to connected nodes. Initially, consider each node to be in its own connected component.

Full Graph Traversal for Each Node:

For each node in the graph, perform a Depth-First Search (DFS) starting from that node.
Mark all reachable nodes as part of the same connected component. This can be done by assigning a unique component ID to each node reached during the traversal.
If a node has already been marked with a component ID in a previous traversal, it means it's already been visited and is part of a known connected component, so you can skip the traversal from this node.
Counting Connected Components:

After performing DFS/BFS from each node, count the number of unique component IDs assigned to the nodes. This count represents the number of connected components in the graph.
