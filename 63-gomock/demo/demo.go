package demo

//go:generate mockgen.exe --source=demo.go -destination=demo_mock.go -package=demo
type Reverser interface {
	Reverse(str string) string
}

func Reverse(str string) string {
	rstr := ""
	for _, v := range str {
		rstr = string(v) + rstr
	}
	return rstr
}

// mockgen -source=demo.go -destination=demo_mock.go -package=demo
