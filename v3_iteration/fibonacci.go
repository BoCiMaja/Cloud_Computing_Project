
import(
	"fmt"
	"log"
	"time"
	"net/http"
)

func fibonacci(n int) int {
	f := make([]int, n+2)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
        
func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, fibonacci(9))  // Output: 34
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}