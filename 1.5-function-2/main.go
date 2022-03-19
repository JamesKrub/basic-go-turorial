package main

var myRealAge = tellMeYourAge()

func tellMeYourAge() int {
	return 18
}

func init() {
	myRealAge = 30
}

func main() {
	//fmt.Println(myRealAge)
}
