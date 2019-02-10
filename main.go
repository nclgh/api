package main
import (
	"fmt"
	"github.com/koding/kite"
	"time"
)

func main() {
	k := kite.New("second", "1.0.0")
	client := k.NewClient("http://localhost:8080/kite")
	client.Dial()
	response, _ := client.Tell("kite.ping")
	fmt.Println(response.MustString())
	time.Sleep(100*time.Second)
}