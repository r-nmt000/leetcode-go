package main

func main() {

}

func reverseString(s []byte) {
	helper(s, 0, len(s) -1)
}

func helper(s []byte, left int, right int) {
	if left >= right { return }
	tmp := s[left]
	s[left] = s[right]
	s[right] = tmp
	left = left + 1
	right = right -1
	helper (s, left, right)
}