package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type DollarQuote struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	quote, err := getDollarQuote(ctx)
	if err != nil {
		log.Fatalf("Erro ao buscar cotação: %v", err)
	}

	err = saveToFile("cotocao.txt", quote.Bid)
	if err != nil {
		log.Fatalf("Erro ao salvar cotação no arquivo: %v", err)
	}

	fmt.Println("Cotação do dolar salva com sucesso!")
}

func getDollarQuote(ctx context.Context) (*DollarQuote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao consultar cotação: status %d", resp.StatusCode)
	}

	var quote DollarQuote
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		return nil, err
	}

	return &quote, nil
}

func saveToFile(filename, bid string) error {
	data := fmt.Sprintf("Dolar: %s", bid)
	return os.WriteFile(filename, []byte(data), 0644)
}
