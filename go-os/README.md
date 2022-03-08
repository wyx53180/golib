
```
func main() {
	// w
	f := Open("asdf", "w")
	defer f.Close()
	f.Write("asdf1234\n")
	f.Write("asdf1234\n")

	// r
	f = Open("asdf", "r")
	fmt.Print(f.Read())

	// readlines
	f = Open("asdf", "r")
	for i := range f.ReadLines() {
		fmt.Println(i)
	}

	// walk
	for i := range Walk("./") {
		fmt.Println(i)
	}
}
```
