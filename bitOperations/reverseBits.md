To understand how the bits are getting added at the
�
i-th position in the reverseBits function, it's crucial to recognize how the bit shifting and OR operations work together. Let's break down the operation res = res | (bit << (31 - i)):

Understanding Bit Shifting
Shifting bit Left by 31 - i:
The expression bit << (31 - i) shifts bit to the left by 31 - i positions.
In a 32-bit number, if we want to place a bit at the
�
i-th position from the right (where
�
i starts from 0), it needs to go to the position which is 31 - i from the left.
This is because the most significant bit (MSB) is at the leftmost position (position 31 from the left or position 0 from the right), and the least significant bit (LSB) is at the rightmost position (position 0 from the left or position 31 from the right).
For example, if
�
=
1
i=1, we're processing the second least significant bit from the original number. In the reversed number, this bit should be the second most significant bit, which is at position 31 - 1 = 30 from the left.
Understanding Bitwise OR
Using Bitwise OR | with res:
The operation res | ... combines res with the shifted bit.
Bitwise OR (|) sets each bit in the result to 1 if either of the corresponding bits in the two operands is 1.
If res has a 0 at the position where bit is being added, the OR operation will set that position to the value of bit.
If res already has a 1 at that position, it remains 1, regardless of the value of bit.
This allows each processed bit to be correctly placed in the reversed number without altering the other bits that have already been set in res.
Example
Let's revisit the example of reversing the bits of num = 10 (00000000000000000000000000001010 in binary):

Second Iteration (i = 1):

bit := (num >> 1) & 1 extracts the second least significant bit (which is 1).
res = res | (bit << (31 - 1)): Here, bit is 1, and 31 - 1 = 30.
bit << 30 shifts 1 to the 30th position from the left (or second position from the right in the reversed number).
Using |, this 1 is placed in res at that position.
Fourth Iteration (i = 3):

Similarly, for the fourth bit of num (1), bit << (31 - 3) will shift it to the 28th position from the left, and the OR operation will place it correctly in res.
Through these operations, each bit from the original number num is placed in its corresponding reversed position in res.
