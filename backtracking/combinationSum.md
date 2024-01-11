To understand the difference between these two combinationSum functions and why one is giving duplicates while the other is not, let's compare their logic and the helper functions they use:

First Program:

The helperComb function uses the parameter idx to control from where to start picking elements in arr.
In each iteration of the for loop, it picks an element starting from idx and recursively calls helperComb with i (not idx) as the new starting index. This means in each recursive call, it considers the current element and all elements after it.
This approach avoids duplicates because each combination includes the current element and then considers only elements that come after it in the array.
Second Program:

The helper function also uses an idx parameter, but when recursively calling itself, it mistakenly uses idx instead of i. This causes it to always start from the same index in each recursive call.
This leads to considering the same set of elements over and over again, resulting in duplicate combinations.
The key difference lies in how the index is updated and used in the recursive calls:

In the first program, by calling helperComb(i, arr, target, curSum+arr[i], curSet, subsets), it ensures that each recursive call considers a subset of the array starting from the current element, thus generating unique combinations.

In the second program, by calling helper(idx, curSum+arr[i], arr, target, curSet, subsets), it's essentially not moving forward in the array in each recursive call. This results in the same elements being considered multiple times, leading to duplicates.

To fix the second program and avoid duplicates, you should replace helper(idx, curSum+arr[i], arr, target, curSet, subsets) with helper(i, curSum+arr[i], arr, target, curSet, subsets) in the for loop. This change makes sure that each recursive call moves forward in the array.
