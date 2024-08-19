package main

const helloPrefix = "Hello, "
func hello(name string) string {
	if name == ""{
		name = "World"
	}
	return helloPrefix + name
}

 
func main() {
	hello("abc")
	

}
