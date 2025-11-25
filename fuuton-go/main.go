package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Candle representa una vela (candlestick)
type Candle struct {
	Timestamp string
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Volume    int64
}

// Ping retorna un mensaje de validación
func Ping() string {
	return "Fuuton OK"
}

// ReadCSV lee el archivo data.csv y retorna slices de Candles
func ReadCSV(filename string) ([]Candle, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var candles []Candle
	// Saltar header
	for i := 1; i < len(records); i++ {
		record := records[i]
		open, _ := strconv.ParseFloat(record[1], 64)
		close, _ := strconv.ParseFloat(record[2], 64)
		high, _ := strconv.ParseFloat(record[3], 64)
		low, _ := strconv.ParseFloat(record[4], 64)
		volume, _ := strconv.ParseInt(record[5], 10, 64)

		candle := Candle{
			Timestamp: record[0],
			Open:      open,
			Close:     close,
			High:      high,
			Low:       low,
			Volume:    volume,
		}
		candles = append(candles, candle)
	}
	return candles, nil
}

// GenerateSignal genera una señal BUY si close > open
func GenerateSignal(candle Candle) string {
	if candle.Close > candle.Open {
		return "BUY"
	}
	return "HOLD"
}

func main() {
	fmt.Println("Fuuton activo")
	response := Ping()
	fmt.Printf("Ping response: %s\n", response)
	fmt.Println()

	// Leer CSV
	candles, err := ReadCSV("data.csv")
	if err != nil {
		log.Fatalf("Error al leer CSV: %v", err)
	}

	fmt.Printf("Velas leídas: %d\n", len(candles))
	fmt.Println()

	// Generar señales
	for i, candle := range candles {
		signal := GenerateSignal(candle)
		fmt.Printf("Vela %d - Timestamp: %s | Open: %.2f | Close: %.2f | Signal: %s\n",
			i+1, candle.Timestamp, candle.Open, candle.Close, signal)
	}
}
