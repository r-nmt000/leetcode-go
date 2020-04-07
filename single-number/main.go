package main

func main() {

}

func singleNumber(nums []int)int {
	noDuplicateMap := make(map[int]int)
	for _, v := range nums {
		_, ok := noDuplicateMap[v]
		if ok {
			delete(noDuplicateMap, v)
		} else {
			noDuplicateMap[v] = 1
		}
	}
	for k, _ := range noDuplicateMap {
		return k
	}
	return 0
}
