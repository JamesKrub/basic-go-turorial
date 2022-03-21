package main

/* ################ (1)
- add Article struct
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
//func SlugifyAll(s []Article) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
//		r[i] = v.Slug()
//	}
//
//	return r
//}
//
//func main() {
//	r := SlugifyAll([]Article{
//		{Title: " Lorem   Ip sum "},
//		{Title: "Babel Coder"},
//	})
//
//	fmt.Println(r)
//}

/* ################ (2)
- add Person struct
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
//func SlugifyAllArticle(s []Article) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
//		r[i] = v.Slug()
//	}
//
//	return r
//}
//
//func SlugifyAllPerson(s []Person) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
//		r[i] = v.Slug()
//	}
//
//	return r
//}

/* ################ (3)
- made it as interface
*/

//
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
//
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
//func SlugifyAll(s interface{}) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
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

/* ################ (4)
- add type slugger interface
*/

//
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
//
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
//func SlugifyAll(s []Slugger) []string {
//	r := make([]string, len(s))
//
//	for i, v := range s {
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
