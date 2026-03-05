package main

func main() {

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{k, j}
		}
		m[v] = k
	}
	return nil
}
