package sortLib

import "math/rand"

func QuickSort(unsortedTable []int32) ([]int32, error) {
	if len(unsortedTable) < 2 {
		return unsortedTable, nil
	}

	left, right := 0, len(unsortedTable)-1

	pivot := rand.Int() % len(unsortedTable)

	unsortedTable[pivot], unsortedTable[right] = unsortedTable[right], unsortedTable[pivot]

	for i, _ := range unsortedTable {
		if unsortedTable[i] < unsortedTable[right] {
			unsortedTable[left], unsortedTable[i] = unsortedTable[i], unsortedTable[left]
			left++
		}
	}

	unsortedTable[left], unsortedTable[right] = unsortedTable[right], unsortedTable[left]

	_, _ = QuickSort(unsortedTable[:left])
	_, _ = QuickSort(unsortedTable[left+1:])

	return unsortedTable, nil
}
