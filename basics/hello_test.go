package main 
import "testing"
func TestHello(t *testing.T){
	got := hello("world")
	want := "hello world"


	if got != want {
		t.Errorf("got %q,want %q",got,want)
	}
}