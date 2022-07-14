package moody

type alive interface {
	walk()
}
type People struct {
	name string
	sex  string
}

func (p People) walk() {

}
