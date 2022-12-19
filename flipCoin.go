package main
import "fmt"
import "math/rand"
import "time"

func main(){
	const(
		HEAD = 1
	)
	rand.Seed(time.Now().UnixNano())
	coin := rand.Intn(2)

	if coin == HEAD {
		fmt.Println("Head")
	} else {
		fmt.Println("Tail")
	}
}