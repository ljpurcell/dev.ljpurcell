---
title: Selectionsort
slug: selectionsort
tags: algorithm, complexity analysis, go
---

Like [bubblesort](/post/bubblesort), selectionsort is another of the simple, but not fast (relatively speaking) algorithms. It operates in a manner similar to bubblesort, but makes *fewer* swaps. Thus, it can be can considered an optimisation of sorts, as it requires fewer operations over all. 

It achieves this by tracking the **smallest** value it finds and then places that element into its final position with a single swap. Unlike bubblesort which walks an element to its final location, every step of the way.

Consider the first element of a sequence that needs to eventually end up in the tenth location. Bubblesort will iterate and swap the element progressively through locations two, three, four... and so on. 

Selectionsort, however, skips all unnecessary swaps. It still needs to iterate over those elements in between, but when it finds the right location, it moves the element there in a single swap.

So, how does it do this?

## How selectionsort works

Like bubblesort, it uses two loops. The outer loop we can consider as a pointer to the *"slot"* for this round of inner iterations. We then set the minimum to point to the same value as the `slot`. 

The inner loop then iterates over all the elements (starting at the next element). If the inner loop discovers a smaller value than the current one, it updates the `smallest` variable to point to the new minimum.

Once this iteration through the inner loop is complete, we check if the `smallest` value has changed. If it has, we know that we found an even smaller element and it must be placed in the current slot.

The process then repeats for the next slot until the entire sequence has been sorted.

In Go code, the selectionsort algorithm is as follows:

```go
package main

import (
	"fmt"
	"cmp"
)

func selectionSort[T cmp.Ordered](sequence []T) {
	if sequence == nil || len(sequence) < 2 {
		return
	}

	// Selectionsort algorithm
	for slot := 0; slot < len(sequence)-1; slot++ {
		smallest := slot
		for i := slot + 1; i < len(sequence); i++ {
			if sequence[i] < sequence[smallest] {
				smallest = i
			}
		}

		if smallest != slot {
			sequence[smallest], sequence[slot] = sequence[slot], sequence[smallest]
		}
	}
}

func main() {
	words := []string{"The", "cat", "had", "a", "bad", "attitude"}
	fmt.Printf("Unsorted words: %v\n", words)
	selectionSort(words)
	fmt.Printf("Sorted words: %v\n\n", words)

	nums := []int{12, 3, 1, 0, 6, 15, 23, 7}
	fmt.Printf("Unsorted numbers: %v\n", nums)
	selectionSort(nums)
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

Successfully sorted.

## Conclusion

Within the selectionsort algorithm we do see a minor increase in code complexity in comparisaon to bubblesort. It's not much, but there's an additional variable and if-gate. These are the trade-offs we make in order to save operations.

Though, with that said, the fewer operations are trivial in terms of asymptotic analysis and the run-time of selectionsort is $\mathcal{O}(n^2)$, like bubblesort; even if it does have comparatively fewer operations.

As is the case in any engineering discipline, you must determine whether the difference is meaningful for your specific context.
