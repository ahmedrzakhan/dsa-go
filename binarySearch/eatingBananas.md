Why L <= H?
Including Both Bounds: The condition L <= H ensures that the search includes both the lower and upper bounds of the search space. This is important because the solution could be at either bound.

Single Element Consideration: When L and H converge to a single point (i.e., L == H), it implies that there's only one candidate speed left to check. If we use L < H, the loop would exit without checking this last candidate, potentially missing the correct answer.

Correct Termination: The loop continues until L passes H. This is the correct termination condition for binary search, as it means the entire search space has been exhaustively and correctly explored.
