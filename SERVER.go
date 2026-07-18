package main
import(
	"net/http"
	"fmt"
)

var url string = "none"

func main(){
	http.HandleFunc("/register", register)
	http.HandleFunc("/check", check)
	http.HandleFunc("/run", run)
	http.HandleFunc("/stop", stop)

	http.ListenAndServe(":443", nil)
}

func register(w http.ResponseWriter, r *http.Request){
	fmt.Println("Новый клиент подключился.")
}

func check(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, url)
}
	
func run(w http.ResponseWriter, r *http.Request){
	url = r.FormValue("url")
	r.Body.Close()
}

func stop(w http.ResponseWriter, r *http.Request){
	url = "none"	
}
