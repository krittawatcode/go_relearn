package main // special for start
import (
	"fmt"
	"log"
	"myFmt"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println("Hellow World")
	myFmt.Println("Hellow World")
	fmt.Println(stringutil.Reverse("Hellow World"))
	log.Println("I am here")
}
