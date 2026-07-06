package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 🚨 VULNERABILIDAD ADREDE: Token hardcodeado para que salte el escáner de secretos (SAST)
	aws_secret_key := "AKIAIOSFODNN7EXAMPLE-FAKE-TOKEN-VAL-D4"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "secure", "day": 4, "framework": "ArgoCD", "change": "one"}`)
	})

	fmt.Println("Servidor DevSecOps inicializado en el puerto 8080...", aws_secret_key)
	http.ListenAndServe(":8080", nil)
}
