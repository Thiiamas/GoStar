package utils


func Contains(s []*Grid, v *Grid) (bool, int) {
	for i, a := range s {
		if a.Equals(v) {
			return true, i
		}
	}
	return false, -1 
}

func GridKeys(m map[*Grid]int)[]*Grid {
	keys := make([]*Grid,0 ,len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func QuickSort(slice []int, lo, hi int) ([]int, int) {
	count := 1
	if lo < hi {
		p, nCount := Partition(slice, lo, hi)
		count += nCount
		QuickSort(slice, lo, p - 1)
		QuickSort(slice, p + 1, hi)
	}
	return slice, count
}

func Partition(slice []int, lo, hi int) (int,int) {
	pivot := slice[hi]
	i := lo
	count := 0
	for j:= lo; j <= hi; j += 1 {
		if slice[j] <  pivot {
			v1 := slice[i]
			v2 := slice[j]
			slice[i] = v2
			slice[j] = v1
			i += 1
			count += 1
		}
	}
	v1 := slice[i]
	v2 := slice[hi]
	slice[i] = v2
	slice[hi] = v1
	count += 1
	return i, count
}