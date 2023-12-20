- 0 can never be part of cycle
- some part would always be outside cycle

In a linked list with a cycle, it is not a strict guarantee that some nodes will be outside the cycle, but it is typical in many scenarios. A linked list can be fully cyclical (where the cycle starts at the head of the list), or it can have a non-cyclical part leading into a cycle. The presence of non-cyclical nodes depends on where the cycle begins. If the cycle starts at the head, then all nodes are in the cycle. If the cycle starts somewhere after the head, then the list has a non-cyclical part (nodes from the head up to the start of the cycle). The structure of the list is defined by how the nodes are linked together.
