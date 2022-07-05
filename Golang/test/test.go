package main
import(
	// "fmt"
	"Data-Structures-and-Algorithms/Golang/DSAlib"
)


func main(){
	pq := DSAlib.PriorityQueue[string]{}
	pq.Push("Hello", 1)
	pq.Push("World", 2)
}