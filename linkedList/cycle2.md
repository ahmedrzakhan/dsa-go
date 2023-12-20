Here's a step-by-step breakdown of the math:

Distance Traveled by Slow: slow has moved D + M steps when they meet. Here, D is the distance from the head of the list to the start of the cycle, and M is the distance from the start of the cycle to the meeting point.

Distance Traveled by Fast: fast has traveled 2(D + M) steps, which is twice the distance of slow.

Equating Distances: The equation 2(D + M) = D + M + nC is established. nC represents the additional cycles fast has completed.

Simplifying the Equation: Simplifying 2(D + M) = D + M + nC gives D + M = nC.

Finding the Cycle Start: This implies D = nC - M. When slow is reset to the head and moves D steps, and fast moves D steps from the meeting point, they meet at the cycle start.

The key is realizing that the extra distance fast covers (compared to slow) is exactly one or more full cycles (nC). This understanding leads to the conclusion that moving both pointers D steps from their respective positions will bring them to the cycle's start.

When slow starts from the head, it moves D steps to reach the cycle's start.
Simultaneously, fast moves D steps from the meeting point. Since D is equivalent to the cycle's length minus the distance inside the cycle to the meeting point (nC - M), fast effectively completes the cycle(s) and reaches the start.
Thus, their meeting point after these D steps is guaranteed to be the cycle's start.
