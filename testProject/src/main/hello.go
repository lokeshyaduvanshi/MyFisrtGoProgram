//func someFunc(w http.ResponseWriter, req *http.Request){
//	w.Write([]byte("Hello universe"))
//}

package main

import (
	"net/http"
	"html/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			data := TodoPageData{
				PageTitle: "My TODO list",
				Todos: []Todo{
					{Title: "Task 1", Done: false},
					{Title: "Task 2", Done: true},
					{Title: "Task 3", Done: true},
				},
			}
			tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
//package main
//
//import "fmt"
//
//func add(x , y int )int{
//      return x + y
//}
//
//func main()  {
//	fmt.Println(add(10,20))
//}

//package main
//
//import "fmt"
//
//func swap (x,y string) (string,string){
//	return  y,x
//}
//
//func main()  {
//	a,b :=swap("hello","world")
//	fmt.Println(a,b)
//}

//var exemple

//package main
//
//import "fmt"
//
//var x,y,z bool
//
//func main (){
//	var i int
//	fmt.Println(x,y,z,i)
//}

//Variables with initializers

//package main
//
//import "fmt"
//
//var x,y int = 1,2
//
//func main () {
//	var a,b,c=false,true,"test"
//	fmt.Println(x,y,a,b,c)
//}
//
//

//package main
//
//import "fmt"
//
//func main(){
//	sum:=0
//	for i:=0;i<10;i++{
//		sum+=i
//	}
//
//	fmt.Println(sum)
//
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main(){
//	fmt.Println("When's Saturday")
//	today:=time.Now().Weekday()
//	switch time.Saturday {
//	case today+0:
//		fmt.Println("Today")
//	case today+1:
//		fmt.Println("Tomorrow")
//	default:
//		fmt.Println("Too far away.")
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main()  {
//	t:=time.Now()
//	switch  {
//	case t.Hour() < 12:
//		fmt.Println("Good Moring!")
//	case t.Hour() < 17:
//		fmt.Println("Good Afternoon")
//	default:
//		fmt.Println("Good Evening")
//	}
//
//}

//defer

//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	defer fmt.Println("hello")
//	fmt.Println("world")
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	fmt.Println("continous")
//	for i:=0;i<=10;i++{
//	   defer fmt.Println("i",i)
//	}
//
//	fmt.Println("work done")
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	a:=10
//	p:=&a
//	fmt.Println("a",*p)
//	*p=20
//	fmt.Println("a",a)
//
//}

//package main
//
//import "fmt"
//
//type users struct{
//	id int
//	name string
//	email string
//	address string
//}
//
//func main(){
//	fmt.Println(users{1,"lokesh","lokesh@appointy.com","Vinayak Parisar"})
//}

//package main
//
//import "fmt"
//
//type users struct {
//	id int
//	name string
//}
//
//func main(){
//	x:=users{1,"lokesh"}
//	x.id=2
//	fmt.Println(x)
//}

//package main
//
//import "fmt"
//
//type users struct{
//	 id int
//	 name string
//}
//
//func main(){
//	v:=users{1,"test"}
//	p:=&v
//
//	fmt.Println("p",*p)
//
//}

//package main
//
//import "fmt"
//
//func main(){
//	var x[2]int
//	x[0]=1
//	x[1]=2
//
//	fmt.Println(x)
//
//	y:=[5]int{1,2,3,4,5}
//
//	fmt.Println(y)
//
//}

//package main
//
//import "fmt"
//
//func main(){
//	x:=[6]int{1,2,3,4,5,6}
//
//	var s[]int=x[3:5]
//
//	fmt.Println(s)
//
//}

//package main
//
//import "fmt"
//
//func main(){
//	x:=[5]string{
//		"rahul",
//		"amit",
//		"test",
//		"test1",
//    }
//
//    y:=x[2:4]
//    z:=x[1:3]
//
//    fmt.Println(y,z)
//
//    y[0]="lokesh"
//
//	fmt.Println(y,z)
//	fmt.Println(x)
//
//}

//package main
//import "fmt"






