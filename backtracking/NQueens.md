Positive Diagonals (PosDiag):
Identification:

Positive diagonals on a chessboard run from the bottom-left to the top-right.
They can be uniquely identified by the sum of the row and column indices (row + col).
Example:

Consider a queen placed at row 1, column 2 on a 4x4 chessboard.
The key for the positive diagonal in this case is 1 + 2 = 3.
All squares where the sum of the row and column indices equals 3 are on the same diagonal.
Usage in Code:

When a queen is placed, the corresponding positive diagonal is marked in PosDiag. For our example, solver.PosDiag[3] is set to 1.
This marking helps to quickly check if a subsequent queen placement would conflict on any positive diagonal.

Negative Diagonals (NegDiag):
Identification:

Negative diagonals run from the top-right to the bottom-left.
In theory, they could be identified by col - row, but this can lead to negative values, which are unsuitable for map keys.
Offsetting the Key:

To ensure all keys are positive, an offset is added. This offset is typically the board size minus one (solver.BoardSize - 1).
The formula becomes solver.BoardSize + col - row - 1.
Example:

Again, consider the queen at row 1, column 2 on a 4x4 chessboard.
The key for the negative diagonal is 4 + 2 - 1 - 1 = 4.
Any square where this formula equals 4 is on the same diagonal as our example queen.
Usage in Code:

When a queen is placed, the corresponding negative diagonal is marked in NegDiag. For our example, solver.NegDiag[4] is set to 1.
This enables quick checks against conflicts on negative diagonals for subsequent queen placements.

```
Since no queen was placed in row 2, the function Backtrack(2) completes, and the execution returns to Backtrack(1).
```

how is this being detected

In the N-Queens solver, the detection of "no queen placed in a row" is not done through an explicit check but rather as a result of how the algorithm navigates through the chessboard. When all columns in a row are checked and found unsafe (triggering the continue statement for each), the loop completes, and the function ends without further recursive calls, leading to an implicit return and thus backtracking to the previous row.
-> if it continues (for loop continue) backtrack didnt get called again causing last functional call to return

Space for Storing Solutions (solutions):

Space for Recursive Call Stack: As mentioned before, the maximum depth of the recursive call stack is n (the board size), contributing O(n) space complexity.

Each solution is a representation of the board, which is an n x n grid. However, each solution is stored as a slice of strings, where each string represents a row. Since there are n rows and each row is represented by a string of length n, each solution requires O(n²) space.
The number of solutions for the N-Queens problem can be quite large but is generally much less than n!. For a given board size n, let's denote the number of solutions as S(n). The exact number of solutions grows with n, but not as rapidly as n!.
Therefore, the total s

While using negative values as map keys in Go will work and isn't incorrect per se, opting for non-negative keys in a problem like the N-Queens solver is a choice driven by considerations of clarity, consistency, ease of debugging, and potential extensibility of the code. It makes the mapping from the problem domain (chessboard positions) to the solution domain (data structures in your code) more transparent.

postion 0,0 negative diagnoal is 1,1 and 2,2 etc why is 3 set to 1 then

ChatGPT
In the N-Queens problem, the calculation for the key in negDiag (for negative diagonals) when a queen is placed at position (0,0) might seem counterintuitive at first. The key calculation formula boardSize + col - row - 1 is designed to ensure a unique, non-negative identifier for each negative diagonal (from top-right to bottom-left).

However, this calculation does not directly correspond to the intuitive understanding of diagonals on a chessboard. For example, the diagonal passing through (0,0), (1,1), (2,2), etc., would naturally be considered the 'first' negative diagonal. But according to the formula, this diagonal is assigned the key 3 on a 4x4 board.

This discrepancy arises because the formula is designed to work universally across the board without yielding negative indices, rather than to match our intuitive numbering of diagonals. The primary goal is to ensure each diagonal has a unique and non-negative identifier, rather than to provide a 'natural' numbering of the diagonals. This approach simplifies the implementation, especially for conflict checking and backtracking in the algorithm.

=> ok so if hashmap allows
then col-row and row-col are same thing?

Yes, in the context of identifying diagonals on a chessboard, both col - row and row - col can be used to uniquely identify a diagonal, provided that the data structure (like a map or hashmap) allows for negative keys. While these expressions yield values with opposite signs, they both map each square on the chessboard to a unique identifier for its diagonal.
