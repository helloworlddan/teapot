package main

import (
	"io"
	"log"
	"net/http"
)

const pot = `HTTP 418 I am a teapot!

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusTeapot)
		io.WriteString(w, pot)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
