package moody

import "fmt"

type alive interface {
	walk()
}
type People struct {
	name string
	sex  string
}

func (p People) walk() {

}

func main() {
	i := string("2222")
	var a int
	a = 8
	fmt.Println(a, i)

}