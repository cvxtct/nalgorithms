package hckrnk

func CountingSort(n int, arr []int) []int {
	count := make([]int, 100)
	seen := make(map[int]bool)
	res := []int{}

	for i := range arr {
		if !seen[arr[i]] {
			seen[arr[i]] = true
		}
		if seen[arr[i]] {
			count[arr[i]]++
		}
	}

	for i := range arr {
		if count[i] > 0 {
			cnt := count[i]
			for cnt > 0 {
				res = append(res, i)
				cnt--
			}
		}
	}

	return res
}
