package main

/* ################ (1)
 */
//
//type Integer interface {
//	int | int8 | int16 | int32 | int64 |
//		uint | uint8 | uint16 | uint32 | uint64
//}
//
//// ฟังก์ชัน Next ที่ทำการคืนค่าเลขถัดไป พร้อมกำหนด Type Parameter T เป็น Interger
//func Next[T Integer](n T) T {
//	return n + 1
//}
//
//func main() {
//	fmt.Println(Next(1))
//}

/* ################ (2)
- Step has Underlying Types as int
*/

//type Integer interface {
//	int | int8 | int16 | int32 | int64 |
//		uint | uint8 | uint16 | uint32 | uint64
//}
//
//type Step int
//
//func Next[T Integer](n T) T {
//	return n + 1
//}
//
//func main() {
//	// Step does not implement Integer
//	Next(Step(1))
//}

/* ################ (3)
- Step has Underlying Types as int
*/

//type Integer interface {
//	~int | ~int8 | ~int16 | ~int32 | ~int64 |
//		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
//}
//
//type Step int
//
//func Next[T Integer](n T) T {
//	return n + 1
//}
//
//func main() {
//	fmt.Println(Next(Step(1)))
//}
