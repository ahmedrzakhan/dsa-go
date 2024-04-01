```go
adj[n2] = append(adj[n2], Edge{node1: n2, node2: n1, weight: weight})
```

This line is almost identical to the previous one, but it updates the adjacency list for node n2.
Since the graph is undirected, the edge needs to be represented in the adjacency lists of both nodes.

Consistency: Focusing on N2 provides a simple and consistent rule. Consider these scenarios:

Scenario 1: When processing edge (1, 2), you check if node '2' has been visited.
Scenario 2: Later, when processing edge (2, 1), you again check if node '1' has been visited.

Why not N1 or both?

N1 is arbitrary: Since edges are undirected, either node could be considered the starting point for the edge.
