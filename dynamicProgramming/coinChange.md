Let's walk through an example step by step using the coins {1, 3, 4, 5} and an amount of 7. The goal is to find the minimum number of coins that add up to 7 using these denominations.

Initially, the changes array is initialized with values greater than 7 (let's say 8 as an infinity value) and changes[0] = 0 since zero coins are needed for amount zero. The array looks like this: [0, 8, 8, 8, 8, 8, 8, 8].

Now, let's iterate through the amounts and the coins:

Amount = 1:

Coins = 1: changes[1] becomes min(8, 1 + changes[1-1]) = min(8, 1 + 0) = 1.
Coins = 3, 4, 5: Not used as they are greater than 1.
changes is now [0, 1, 8, 8, 8, 8, 8, 8].
Amount = 2:

Coins = 1: changes[2] becomes min(8, 1 + changes[2-1]) = min(8, 1 + 1) = 2.
Coins = 3, 4, 5: Not used as they are greater than 2.
changes is now [0, 1, 2, 8, 8, 8, 8, 8].
Amount = 3:

Coins = 1: changes[3] becomes min(8, 1 + changes[3-1]) = min(8, 1 + 2) = 3.
Coins = 3: changes[3] becomes min(3, 1 + changes[3-3]) = min(3, 1 + 0) = 1.
Coins = 4, 5: Not used as they are greater than 3.
changes is now [0, 1, 2, 1, 8, 8, 8, 8].
Amount = 4:

Coins = 1: changes[4] becomes min(8, 1 + changes[4-1]) = min(8, 1 + 1) = 2.
Coins = 3: changes[4] stays 2 as min(2, 1 + changes[4-3]) = min(2, 1 + 2) = 2.
Coins = 4: changes[4] becomes min(2, 1 + changes[4-4]) = min(2, 1 + 0) = 1.
Coins = 5: Not used as it is greater than 4.
changes is now [0, 1, 2, 1, 1, 8, 8, 8].
Amount = 5:

Follow a similar process. Each coin is evaluated.
Finally, changes[5] will be 1, as the coin 5 alone makes the amount.
Amount = 6:

Similarly, check each coin.
changes[6] will be 2, using coins 1 + 5 or 3 + 3.
Amount = 7:

Coins = 1: changes[7] becomes min(8, 1 + changes[7-1]) = min(8, 1 + 2) = 3.
Coins = 3: changes[7] becomes min(3, 1 + changes[7-3]) = min(3, 1 + 1) = 2.
Coins = 4, 5: changes[7] remains 2, as using these won't give a better result.
changes is now [0, 1, 2, 1, 1, 1, 2, 2].
At the end, changes[7] is 2, which means the minimum number of coins required to make the amount 7 with the given denominations is 2. The coins used are 3 and 4. This solution ensures that all possible combinations are considered to find the minimum number of coins for each amount.

The expression 1 + changes[i-coin] in the coin change algorithm is crucial for understanding how the minimum number of coins is calculated for a given amount. Let's break down its components:

changes[i-coin]: This part of the expression refers to the number of coins needed to make up the amount that remains after using one coin of the current denomination. For example, if i is the current amount we are trying to make and coin is the value of the coin we are considering, then i - coin is the remaining amount after using one such coin. changes[i-coin] gives us the minimum number of coins needed for this remaining amount, as previously calculated by the algorithm.

1 +: The 1 in this expression represents the current coin that we are considering. By adding 1, we are essentially saying, "if we use one coin of this denomination (coin), how many more coins do we need to make up the rest of the amount (i - coin)?". This is why we add 1 – to account for the current coin being used.

Putting it together, 1 + changes[i-coin] calculates the total number of coins needed if we choose to use one coin of the current denomination. The algorithm then compares this number with the current minimum number of coins known for making up the amount i (changes[i]) and updates changes[i] to the smaller of these two values.

This process is repeated for each coin denomination and for each amount up to the target amount. As a result, the algorithm is able to find the minimum number of coins needed for each amount incrementally, ensuring an optimal solution through dynamic programming.

The +1 is added to account for the current coin we are considering. It is saying, "if we use one coin of this denomination, then how many additional coins are needed to make up the remaining amount?"
In our example (i = 6, coin = 4), 1 + changes[2] calculates the total number of coins needed if we include one coin of value 4. The 1 accounts for this coin, and changes[2] tells us how many more coins we need for the remaining amount of 2.

what if we sorted the array and added check to coins at a number, would it improve TC?

Dynamic Programming Nature: The essence of the solution lies in its dynamic programming approach, where the algorithm iteratively builds up the minimum number of coins needed for each amount from 1 to the target amount. This process requires considering all possible combinations of coins, regardless of their order.
ALSO not expected to have 1000's of coins

Sorting Overhead: Sorting the coins array introduces an additional computational overhead. The best sorting algorithms have a time complexity of O(N log N), where N is the number of coins. This preprocessing step would add to the overall time complexity without reducing the complexity of the main algorithm, which remains O(amount \* N).

Impact on Worst-Case Scenario: The worst-case scenario for this algorithm occurs when the coin values are small relative to the amount, requiring the algorithm to consider many combinations. Sorting the coins doesn't change this, as the algorithm still needs to explore all possible combinations to find the minimum number of coins for each sub-amount.
