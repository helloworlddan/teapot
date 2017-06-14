package main

import (
	"log"
	"net/http"

	"io"

	"github.com/gorilla/mux"
)

const pot = `
                       (
            _           ) )
         _,(_)._        ((
    ___,(_______).        )
  ,'__.   /       \    /\_
 /,' /  |""|       \  /  /
| | |   |__|       |,'  /
 \'.|                  /
  '. :           :    /
    '.            :.,'
      '-.________,-'
`

func main() {
	log.Println("Teapot booting..")
	router := mux.NewRouter().StrictSlash(false)
	router.PathPrefix("/").HandlerFunc(serve)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusTeapot)
	io.WriteString(w, pot)
}
