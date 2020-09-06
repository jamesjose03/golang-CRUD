package mainimport (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	)
	func main() {
	route := mux.NewRouter()x := route.PathPrefix("/api").Subrouter() //Base Path
	//Routes
	x.HandleFunc("/createProfile", createProfile).Methods("POST")
	x.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")
	x.HandleFunc("/getUserProfile", getUserProfile).Methods("POST")
	x.HandleFunc("/updateProfile", updateProfile).Methods("PUT")
	x.HandleFunc("/deleteProfile/{id}", deleteProfile).Methods("DELETE")log.Fatal(http.ListenAndServe(":8000", x)) // Run Server
	}