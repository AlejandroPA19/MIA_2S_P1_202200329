package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "log"
    "os"
    "os/signal"
    "syscall"
	"github.com/AlejandroPA19/MIA_2S_P1_202200329/Analyzer"
	"github.com/AlejandroPA19/MIA_2S_P1_202200329/DiskManagement" 
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
    // Configurar la captura de señales para la terminación del programa
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    // Inicializar el servidor
    http.Handle("/", http.FileServer(http.Dir("../front/build")))
    http.HandleFunc("/api/analyze", analyzeHandler)

    fmt.Println("Servidor escuchando en el puerto 8080")

    // Ejecutar el servidor en una goroutine
    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error al iniciar el servidor: %v", err)
        }
    }()

    // Esperar a recibir una señal de terminación
    <-quit

    // Limpiar las particiones al terminar
    DiskManagement.UnmountDisks()

    fmt.Println("Servidor detenido.")
}