Remember the core goals: the maxHeap should store the smaller half of the numbers we've seen so far, and its top element should be the largest of these smaller numbers. The minHeap should store the larger half. By inserting into the maxHeap first and then carefully redistributing if necessary, we preserve these crucial properties!

Key Note: The FindMedian function relies on this structure when calculating the median; ensuring its consistency is what AddNum is meant to do.

## AddNum function:

Scenario: Suppose our MedianFinder currently looks like this:

maxHeap: [3, 6]
minHeap: [7]
Let's add the number 4 using mf.AddNum(4).

Step by Step:

Push onto maxHeap:

maxHeap becomes [3, 6, 4].
Pop and Re-insert:

val := heap.Pop(mf.maxHeap).(int): The value 6 is popped from the maxHeap.
heap.Push(mf.minHeap, val): The value 6 is pushed onto the minHeap.
Balancing Check:

if mf.maxHeap.Len() < mf.minHeap.Len(): Our maxHeap now has 2 elements, and our minHeap also has 2 elements. The condition is False.
Balancing (Skipped): Since the heap lengths are balanced, we skip the balancing step.

Final State:

After AddNum(4), our heaps look like this:

maxHeap: [3, 4]
minHeap: [6, 7]
Observation: Notice that by adding 4 and carefully rearranging, we still maintain our core properties:

maxHeap: Holds the smaller half of numbers seen.
minHeap: Holds the larger half of numbers seen.
Balance: maxHeap has exactly the same number of elements as minHeap.
