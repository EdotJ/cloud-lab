package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/edotj/cloud-lab/internal/hashes"
	"github.com/gorilla/mux"
)

func Sha256Handler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	hashed, _ := hashes.GetHash("Sha256", username)
	//username := mux.Vars(r)["username"]

	// to calculate sha256 hash of a string, use internal/hashes package and function GetHash

	// to send the response, use the following:
	// fmt.Fprint(w, "your response here")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, hashed)
}

func GithubUsernameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, os.Getenv("GITHUB_USERNAME"))
}
