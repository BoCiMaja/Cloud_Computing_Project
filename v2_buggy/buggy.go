
import(
	"fmt"
	"log"
	"time"
	"net/http"
)

func fibo_buggy(n uint) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		panic("Buggy program")
    	// Calling Sleep method
    	time.Sleep(1 * time.Second) 
		return fib(n-1) + fib(n-2)
	}
}
        
func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, fibo_buggy(9))  // Output: 34
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}