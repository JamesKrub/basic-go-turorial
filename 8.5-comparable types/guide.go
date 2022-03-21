package main

/* ################ (1)
- comparable type
- go 1.18 come with constraint named comparable
*/

//func CountOfOccurrences[T comparable](s []T, h T) int {
//	var count int
//
//	for _, v := range s {
//		if v == h {
//			count++
//		}
//	}
//
//	return count
//}
//
//func main() {
//	fmt.Println(CountOfOccurrences([]int{1, 2, 3, 2, 2, 2, 1}, 2))         // 4
//	fmt.Println(CountOfOccurrences([]string{"high", "low", "low"}, "low")) // 2
//}
