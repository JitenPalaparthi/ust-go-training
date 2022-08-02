package demo

//go:generate mockgen.exe --source=demo.go -destination=demo_mock.go -package=demo
type Reverser interface {
	Reverse(str string) string
	Count(str string) int
}

func Reverse(str string) string {
	rstr := ""
	for _, v := range str {
		rstr = string(v) + rstr
	}
	return rstr
}

func Count(str string) int {
	return len(str)
}

// mockgen -source=demo.go -destination=demo_mock.go -package=demo
