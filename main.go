package main

func main() {
	err := ParseConfig()
	if err != nil {
		panic(err)
	}
}
