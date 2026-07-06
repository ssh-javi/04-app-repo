package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ✅ Código limpio de secretos para pasar el filtro de Trivy

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "secure", "day": 4, "framework": "ArgoCD"}`)
	})

	// 👇 Quitamos "aws_secret_key" de acá para que Go compile sin errores
	fmt.Println("Servidor DevSecOps inicializado en el puerto 8080 y verificado seguro...")
	http.ListenAndServe(":8080", nil)
}
