package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("Sujata", "Panchal")
	fmt.Println(a, b)

}
