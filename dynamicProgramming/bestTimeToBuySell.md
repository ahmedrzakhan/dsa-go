Lines of Code Explained

cash = max(cash, hold + A[i] - fee)

Logic: This line calculates the potential profit if you decide to sell your stock on the current day. There are two scenarios:
You were already not holding a stock: Your profit remains the same (cash).
You were holding a stock: Your profit becomes the previous day's hold value + the current selling price (A[i]) - the transaction fee (fee).
Key: The max function ensures you keep track of the highest possible profit without a stock.
hold = max(hold, cash - A[i])

Logic: This line calculates the potential profit if you decide to buy a stock on the current day. Again, two scenarios:
You were already holding a stock: Your profit remains the same (hold).
You were not holding a stock: Your profit decreases from the previous day's cash by the current price (A[i]) to reflect the purchase.
Key: The max function ensures you keep track of the highest possible profit while holding a stock.
