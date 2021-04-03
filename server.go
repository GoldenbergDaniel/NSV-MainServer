package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var userList []User
var pathToData string = "users.json"

// User represents a user
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	fmt.Println("The server is active...")
	fmt.Println(" * Running on http://localhost:8080/")
	fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: 8080")

	r := mux.NewRouter()

	r.HandleFunc("/api/get-user", getUserResponse).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func getUserResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	data, err := ioutil.ReadFile("./" + pathToData)
	checkErr(err)

	err = json.Unmarshal(data, &userList)
	checkErr(err)

	var user *User = getUser(params["username"])

	fmt.Println("Returned user " + user.Username)

	json.NewEncoder(w).Encode(&user)
}

// Returns user of given username
func getUser(username string) *User {
	for i := 0; i < len(userList); i++ {
		if userList[i].Username == username {
			return &userList[i]
		}
	}

	return nil
}

// Check for and logs errors
func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
