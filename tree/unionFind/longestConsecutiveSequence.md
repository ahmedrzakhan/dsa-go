Brute-Force Steps:
Iterate Over Each Number: Go through each number in the array.

Build Consecutive Sequences: For each number, attempt to build the longest consecutive sequence by:

Incrementing the current number and checking if the incremented number exists in the array.
Continue incrementing and checking for the presence of consecutive numbers until a number is not found in the array.
Track Maximum Length: Keep track of the length of each consecutive sequence formed and update the maximum length found so far.

Return Maximum Length: After all numbers have been processed, return the length of the longest consecutive sequence found.

// sorted approach
Sort the Array: First, sort the given array. This will arrange all the numbers in ascending order, making it easier to find consecutive sequences.

Iterate Through the Sorted Array: Start iterating from the first element of the sorted array.

Find Consecutive Sequences: For each number, check if the next number in the sorted array is consecutive. A number is considered consecutive if it is either the same as the current number (to handle duplicates) or exactly one more than the current number.

Track the Longest Sequence: Keep track of the current sequence length. If you encounter a number that is not consecutive (i.e., more than one greater than the current number), check if the current sequence length is longer than the maximum sequence length found so far. If it is, update the maximum sequence length.

Return the Maximum Sequence Length: After iterating through the entire array, return the maximum sequence length found.
