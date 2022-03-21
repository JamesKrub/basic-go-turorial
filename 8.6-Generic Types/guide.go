package main

/* ################ (1)
- generics type
*/

//type Integer interface {
//	~int | ~int8 | ~int16 | ~int32 | ~int64 |
//		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
//}
//
//type Point2D[T Integer] struct {
//	x T
//	y T
//}
//
//func main() {
//	// cannot use generic type Point2D[T Integer] without instantiation
//	a := Point2D{1, 2}
//	// cannot use generic type Point2D[T Integer] without instantiation
//	b := Point2D{3, 4}
//
//	fmt.Println(a, b)
//}

/* ################ (2)
- fix (1)
*/

//type Integer interface {
//	~int | ~int8 | ~int16 | ~int32 | ~int64 |
//		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
//}
//
//type Point2D[T Integer] struct {
//	x T
//	y T
//}
//
//func main() {
//	a := Point2D[int]{1, 2}
//	b := Point2D[int{3, 4}
//
//	fmt.Println(a, b)
//}
//
//func (p Point2D[T]) Slope(o Point2D[T]) float64 {
//	return float64((o.y - p.y) / (o.x - p.x))
//}

// Thank you: https://www.babelcoder.com/blog/articles/go-generics?fbclid=IwAR0Ke4e6GeXMpeBq_iN22sST6TRIIBOWfSsHNaHp0_qE4dQNJwGqkVnz3Dc
