package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const proxyPort = 5000

type RequestPayload struct {
	RemoteURL    string `json:"remoteURL"`
	RemoteMethod string `json:"remoteMethod"`
	RemoteBody   string `json:"remoteBody"`
}

func main() {
	// Definisci il gestore per le richieste XML in ingresso
	http.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Richiesta non valida. Deve essere una richiesta POST JSON.", http.StatusMethodNotAllowed)
			return
		}

		// Leggi il corpo della richiesta JSON in ingresso
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura del corpo della richiesta.", http.StatusInternalServerError)
			return
		}

		var requestPayload RequestPayload
		err = json.Unmarshal(body, &requestPayload)
		if err != nil {
			http.Error(w, "Errore nella decodifica del corpo JSON.", http.StatusBadRequest)
			return
		}

		// Effettua la richiesta HTTP al server remoto con le specifiche fornite
		client := &http.Client{}
		req, err := http.NewRequest(requestPayload.RemoteMethod, requestPayload.RemoteURL, bytes.NewBuffer([]byte(requestPayload.RemoteBody)))
		req.Header.Set("Content-Type", "application/xml")
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Errore nella creazione della richiesta al server remoto.", http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Errore nella richiesta al server remoto.", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Leggi la risposta dal server remoto
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura della risposta dal server remoto.", http.StatusInternalServerError)
			return
		}

		// Restituisci la risposta al client
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	})

	// Avvia il server proxy
	fmt.Printf("Server proxy XML in esecuzione su :%d\n", proxyPort)
	http.ListenAndServe(fmt.Sprintf(":%d", proxyPort), nil)
}
