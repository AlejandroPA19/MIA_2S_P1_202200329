package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/AlejandroPA19/MIA_2S_P1_202200329/Analyzer"
)

var mensajes []string

type AnalyzeRequest struct {
	Text string `json:"text"`
}

type AnalyzeResponse struct {
	Messages []string `json:"messages"`
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	var req AnalyzeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error parsing JSON request", http.StatusBadRequest)
		return
	}

	// Limpiamos la lista antes de analizar
	mensajes = []string{}

	// Aquí analizamos el texto y llenamos la lista de mensajes
	Analyzer.Analyze(req.Text, &mensajes)

	// Respondemos con la lista de mensajes
	resp := AnalyzeResponse{Messages: mensajes}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../front/build")))
	http.HandleFunc("/api/analyze", analyzeHandler)

	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
