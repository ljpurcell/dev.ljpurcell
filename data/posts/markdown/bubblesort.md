---
title: Bubblesort
slug: bubblesort
tags: algorithm, complexity analysis, go
---

Some algorithms are fast. Some are simple. 

Bubblesort belongs to the latter category. However, just because it isn't the fastest does not mean it should be ignored. Bubblesort's relative simplicity make it a useful algorithm to study when learning about sorting algorithms.

Bubblesort demonstrates the core criteria of a sorting algorithm, namely:
- It looks at all elements
- It makes comparisons
- It swaps elements that are in the wrong order

The different sorting algorithms all follow the same general steps. They just have their own unique process for how they go about it, which determines their runtime and memory efficiency.

## How bubblesort works

Bubblesort ultimately results in a sorted sequence by starting at the start, and moving the largest element to the end. It then begins another run, starting at the start again, and completes the run by placing the largest element it can find (which is the second largest element) in the second last spot. The sequence is then sorted when all elements from the largest, down to the (second) smallest, have been placed in the correct location.

In short, bubblesort works by bubbling up the largest element to the _"ceiling"_ location.

At a more tactical level, bubblesort achieves this by assessing elements in a pairwise manner. It takes the first element and compares that to the second. If the first is larger than the second, it swaps them. It then compares the second (which may have been originally the first) and third element. If the second is bigger, it swaps them. If not, the third element stays in place, becoming part of the next comparison with the fourth element.

Extrapolating out this out, the largest element within the entire sequence will move all the way to the top (or end). This is because in every comparison it will be swapped to the right (or will have been in the last spot already and wont be swapped when compared). The process is then repeated from the start, but the comparisons are only necessary up to the second last element, as we know the last element is the largest and has already been placed.

How this looks written in Go:

```go
package main

import (
	"fmt"
	"cmp"
)

func bubbleSort[T cmp.Ordered](sequence []T) {
	if sequence == nil || len(sequence) < 2 {
		return
	}

	// Bubblesort algorithm
	for run := 0; run < len(sequence); run++ {
		for i := 0; i < len(sequence)-1-run; i++ {
			if sequence[i] > sequence[i+1] {
				sequence[i], sequence[i+1] = sequence[i+1], sequence[i]
			}
		}
	}
}


func main() {
	words := []string{"The", "cat", "had", "a", "bad", "attitude"}
    fmt.Printf("Unsorted words: %v\n", words)
	bubbleSort(words)
    fmt.Printf("Sorted words: %v\n\n", words)

	nums := []int{12, 3, 1, 0, 6, 15, 23, 7}
    fmt.Printf("Unsorted numbers: %v\n", nums)
	bubbleSort(nums)
    fmt.Printf("Sorted numbers: %v\n\n", nums)
}
```

The output of this program is:
```
Unsorted words: [The cat had a bad attitude]
Sorted words: [The a attitude bad cat had]

Unsorted numbers: [12 3 1 0 6 15 23 7]
Sorted numbers: [0 1 3 6 7 12 15 23]
```

This demonstrates the successful sorting of our sequences.

## Conclusion

As I said, bubblesort is a useful algorithm to understand. It provides a useful base to compare and contrast other algorithms against. 

For instance, although each run ultimately finds and places the largest element, it doesn't maintain a `currentMax` or `currentMin` value. It simply _brings the largest element with it_ as it goes along, and the inner-iterator counter becomes the de facto pointer.

This differs from an algorithm like [selectionsort](/post/selectionsort), that iterates over the entire list, _"remembering"_ the largest value it has seen so far, and then places it once it has looked at all items. 

Bubblesort also does not segment or divide the sequence in any drastic way (other than the sorted section), making it differ from an algorithms like quicksort. 

These differences are one of the reasons that bubblesort is a rather unoptimised algorithm. It's utility is found in its simplicity, not speed. From an asymptotic analysis standpoint, it has a run-time of $O(n^2)$.
