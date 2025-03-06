package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type GitHubWebhook struct {
    Action string `json:"action"`
    Workflow string `json:"workflow"`
    Repository struct {
        FullName string `json:"full_name"`
    } `json:"repository"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    var payload GitHubWebhook
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    log.Printf("Received CI/CD event: %s on %s", payload.Action, payload.Repository.FullName)
    fmt.Fprintf(w, "Received event")
}

func main() {
    http.HandleFunc("/webhook", webhookHandler)
    log.Println("🚀 CI/CD Backend running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

