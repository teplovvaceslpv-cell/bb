package main
import(
	"net/http"
	"fmt"
)

var url string = "none"

func main(){
	http.FuncHandler("/register", register)
	http.FuncHandler("/check", check)
	http.HandleFunc("/run", run)
	http.HandleFunc("/stop", stop)

	http.ListenAndServe(":443", nil)
}

func register(w http.ResponseWriter, r *http.Request){
	fmt.Println("Новый клиент подключился.")
}

func check(w http.ReponseWriter, r *http.Request){
	fmt.Fprintf(url)
}
	
func run(){
	url = r.FormValue("url")
}

func stop(){
	url = "none"	
}
