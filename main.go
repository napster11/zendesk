package main
import (
	"fmt"
	"github.com/napster11/zendesk/zendeskService"
)

func main() {
	fmt.Println("statred zendeskService on port :8080")
	zendeskService.BootRouter(":8080")        //BootRouter will start the service on Port 8080
}
