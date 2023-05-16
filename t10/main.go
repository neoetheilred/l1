package main

/*
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.
*/

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -25.33, -40, 0, 1.5,
	} // Source data
	groups := make(map[int64][]float64) // Store grops in map
	for _, t := range temperatures {
		groupKey := int64(t) / 10 * 10 // Calculate group key (values stepped by 10)
		_, ok := groups[groupKey]      // Check if corresponding group is present in map
		if ok {
			groups[groupKey] = append(groups[groupKey], t) // If group is present, just append
		} else {
			groups[groupKey] = []float64{t} // Else create new slice
		}
	}

	// Print the result
	fmt.Println(groups)

	// If needed to sort by keys:
	gropsSorted := []KVPair[int64, []float64]{}

	for k, v := range groups {
		gropsSorted = append(gropsSorted, KVPair[int64, []float64]{k, v})
	}
	sort.Slice(gropsSorted, func(i, j int) bool {
		return gropsSorted[i].Key < gropsSorted[j].Key
	})

	fmt.Println(gropsSorted)
}

type KVPair[T any, U any] struct {
	Key   T
	Value U
}

func (kv *KVPair[T, U]) String() string {
	return fmt.Sprintf("{%v: %v}", kv.Key, kv.Value)
}
