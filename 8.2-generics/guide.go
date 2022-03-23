package main

/* ################ (1)
- change it to generics type
*/

//type Article struct {
//	Title   string
//	Content string
//}
//
//func (a Article) Slug() string {
//	r := regexp.MustCompile(" +")
//	slug := strings.ToLower(a.Title)
//	slug = strings.TrimSpace(slug)
//
//	return r.ReplaceAllString(slug, "-")
//}
//
//type Person struct {
//	FirstName string
//	LastName  string
//}
//
//func (p Person) Slug() string {
//	r := regexp.MustCompile(" +")
//	slug := strings.ToLower(p.FirstName + " " + p.LastName)
//	slug = strings.TrimSpace(slug)
//
//	return r.ReplaceAllString(slug, "-")
//}
//
//type Slugger interface {
//	Slug() string
//}
//
// // since 1.18, we can use any instead interface{}
// // Here is Generic Constraints. Go only knows `s` as simple slices
// // type parameters
//func SlugifyAll[T any](s []T) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
//		// v.Slug undefined (type T has no field or method Slug)
//		r[i] = v.Slug()
//	}
//
//	return r
//}
//
//func main() {
//	r := SlugifyAll([]Slugger{
//		Article{Title: " Lorem   Ip sum "},
//		Person{FirstName: "   John", LastName: " Doe "},
//	})
//
//	fmt.Println(r)
//}

/* ################ (2)
- implement slugger interface
*/

//// change
//func SlugifyAll[T Slugger](s []T) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
//		r[i] = v.Slug()
//	}
//
//	return r
//}

/* ################ (3)
- การอนุมานชนิดข้อมูลด้วย (Type Inference)
*/

//func main() {
// // We call what is in the square bracket as Type Arguments
//	SlugifyAll[Slugger]([]Slugger{
//		Article{Title: " Lorem   Ip sum "},
//		Person{FirstName: "Somchai", LastName: "Haha"},
//	})
//
//	SlugifyAll[Article]([]Article{
//		Article{Title: " Lorem   Ip sum "},
//	})
//
//	SlugifyAll[Person]([]Person{
//		Person{FirstName: "Somchai", LastName: "Haha"},
//	})
//}

/* ################ (3)
- การใช้งาน type parameters หลายค่า
*/

//func PrefixSlugifyAll[S Slugger, P any](s []S, p []P) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
//		r[i] = fmt.Sprintf("%v: %v", p[i], v.Slug())
//	}
//
//	return r
//}
//
//func main() {
//	r := PrefixSlugifyAll[Slugger, string]([]Slugger{
//		Article{Title: " Lorem   Ip sum "},
//		Person{FirstName: "Somchai", LastName: "Haha"},
//	}, []string{"Article", "Person"})
//
//  // ละส่วนของ Type Arguments ได้
//	r := PrefixSlugifyAll([]Slugger{
//		Article{Title: " Lorem   Ip sum "},
//		Person{FirstName: "Somchai", LastName: "Haha"},
//	}, []string{"Article", "Person"})
//
//	fmt.Println(r)
//}
