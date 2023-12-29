Time Complexity (TC):
O(3^N \* 4^M):

The worst-case time complexity of this algorithm is influenced by the number of digits in the input (digits) and the number of letters each digit maps to.
For digits that map to 3 letters (like 2, 3, 4, 5, 6, 8), let's say there are N such digits. For digits that map to 4 letters (like 7 and 9), let's say there are M such digits.
The algorithm needs to explore 3 choices for each of the N digits and 4 choices for each of the M digits. Thus, the time complexity is O(3^N \* 4^M).
Example:

For an input like "23" (where both digits map to 3 letters), the time complexity is O(3^2) = O(9).
Space Complexity (SC):
O(N + M):

Space complexity is primarily determined by two factors: the depth of the recursion (which contributes to the call stack size) and the space needed for the curSet variable in each function call.
The maximum depth of the recursive call stack will be N + M, where N and M are defined as above.
Since curSet is modified in place and only one is maintained throughout the recursion, it does not add significantly to the space complexity beyond the depth of the recursion.
Output Storage:

The space required to store the output (all combinations) is not typically counted in the auxiliary space complexity. However, it's worth noting that this can be substantial if there are many combinations. The length of each combination is N + M, and there are 3^N \* 4^M combinations.
