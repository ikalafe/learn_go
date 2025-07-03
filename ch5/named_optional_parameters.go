package main

func main() {
	MyFunc(MyFuncOpts{
		LastName: "Dehghan",
		Age:      20,
	})

	MyFunc(MyFuncOpts{
		FirstName: "Daniyal",
		LastName:  "Dehghan",
	})

}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	// Do Something
}
