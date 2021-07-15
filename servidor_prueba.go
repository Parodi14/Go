package  main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola"))

	}
	// rutas del servidor
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/contacto", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hola mucho gusto sapo hdp"))
			http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hola mucho gusto sapo hpt2"))

		})
	})
		// initial el servidor
		fmt.Println("Servidor inicializado en localhost: 3000")
		http.ListenAndServe(":3000", nil)
}

