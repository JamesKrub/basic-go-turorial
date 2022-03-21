package main

///* ################ (1)
//- use genetic as return type
//*/
//
//func Head[T any](v []T) T {
//	return v[0]
//}
//
//func main() {
//	fmt.Println(Head([]string{"1", "2"}))
//	fmt.Println(Head([]int64{1, 2}))
//	fmt.Println(Head([]float64{1.0, 2.0}))
//}

///* ################ (2)
//- Union Types and Generic Operators (1)
//*/
//
//func Add[T any](a, b T) float64 {
//	return float64(a) + float64(b)
//}
//
//func main() {
//
//}

/* ################ (3)
- Union Types and Generic Operators (2)
*/

//func Add[T int | int32 | int64 | float32 | float64](a, b T) float64 {
//	return float64(a) + float64(b)
//}
//
//func main() {
//	fmt.Println(Add(1, 2))
//
//	//// What do we do if we want to pass with different types of parameters?
//	//fmt.Println(Add(1, 2.2))
//}

/* ################ (4)
-  pass with different types of parameters
*/

//func Add[T int | int32 | int64 | float32 | float64, V int | int32 | int64 | float32 | float64](a T, b V) float64 {
//	return float64(a) + float64(b)
//}
//
//func main() {
//	fmt.Println(Add(1, 2.2))
//}

/* ################ (5)
- make it better by using Constraint elements in interface
*/
//
//type Number interface {
//	int | int32 | int64 | float32 | float64
//}
//
//func Add[T, V Number](a T, b V) float64 {
//	return float64(a) + float64(b)
//}
