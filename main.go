package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/YoshikiGit/poc-go-clean-architecture/adapters/presenters"
	"github.com/YoshikiGit/poc-go-clean-architecture/adapters/repository"
	"github.com/YoshikiGit/poc-go-clean-architecture/usecases"
)

func main() {
	os.Remove("./poc_clean_architecture.sql")

	// controller
	http.HandleFunc("/", defaultRoute)
	http.HandleFunc("/create-user", createUser)
	http.HandleFunc("/route2", route2)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	server.ListenAndServe()
}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is default route.")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is /create-user.")

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Println(string(body))

	var createUserInputData usecases.CreateUserInput

	// requestbodyをjsonからstructへ変換
	if err := json.Unmarshal([]byte(body), &createUserInputData); err != nil {
		fmt.Println(err)
		w.Write([]byte("Bad request!\n"))
		return
	}

	uc := usecases.NewCreateUserInteractor(
		// Interfaceに対してオブジェクトを注入
		repository.NewUserSqlite("./infra/poc_clean_architecture.db"),
		presenters.NewCreateUserPresenter(),
	)

	result, _ := uc.CreateUserInteractor(createUserInputData)

	responce_json, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(responce_json))
	return
}

func route2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is /route2.")
}
