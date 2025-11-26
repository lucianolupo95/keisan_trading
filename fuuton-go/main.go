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

// GenerateSignal genera una señal BUY con filtros de volumen y movimiento
func GenerateSignal(candle Candle) string {
	const minVolume int64 = 1300
	const minMovePercent float64 = 0.1  // 0.1%

	// Filtro de volumen mínimo
	if candle.Volume < minVolume {
		return "HOLD"  // Volumen insuficiente
	}

	// Calcular movimiento porcentual
	movePercent := ((candle.Close - candle.Open) / candle.Open) * 100

	// Filtro de movimiento mínimo
	if candle.Close > candle.Open && movePercent >= minMovePercent {
		return "BUY"
	}
	return "HOLD"
}

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║           FUUTON: Trading System - Day 4                   ║")
	fmt.Println("║       Go + Python (Katon) + R (Suiton) Integration         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	response := Ping()
	fmt.Printf("✓ %s\n\n", response)

	// Leer CSV
	candles, err := ReadCSV("data.csv")
	if err != nil {
		log.Fatalf("Error al leer CSV: %v", err)
	}

	fmt.Printf("✓ Velas leídas: %d\n\n", len(candles))

	// ========== SUITON ANALYSIS ==========
	analysis := AnalyzePricesWithSuiton(candles)
	if analysis != nil {
		PrintSuitonAnalysis(analysis)
	}

	// ========== GENERAR SEÑALES ==========
	fmt.Println("\n\n=== SEÑALES GENERADAS (CON ANÁLISIS ESTADÍSTICO) ===\n")
	buyCount := 0
	for i, candle := range candles {
		var signal string
		if analysis != nil {
			signal = GenerateEnhancedSignal(candle, analysis, i)
		} else {
			signal = GenerateSignal(candle)
		}

		if signal != "HOLD" {
			buyCount++
		}

		// Mostrar solo cada 10 velas para no saturar output
		if (i+1)%10 == 1 || (i+1)%10 == 0 {
			fmt.Printf("Vela %d - %s | Open: %.2f | Close: %.2f | Signal: %s\n",
				i+1, candle.Timestamp, candle.Open, candle.Close, signal)
		}
	}
	fmt.Printf("\nTotal BUY signals: %d / %d velas\n", buyCount, len(candles))

	// ========== EJECUTAR ORQUESTADOR ==========
	fmt.Println("\n\n=== EJECUTANDO ORQUESTADOR ===\n")
	config := OrchestratorConfig{
		InitialCapital:     10000.0,
		UseKaton:           false, // Deshabilitado por ahora
		UseEnhancedSignals: analysis != nil, // Usar señales mejoradas si análisis disponible
		Verbose:            true,
	}
	orchestrator := NewOrchestrator(candles, config)
	orchestratorResult := orchestrator.Run()
	orchestratorResult.PrintReport()

	// ========== ESTADÍSTICAS FINALES ==========
	if analysis != nil {
		stats := ComputeSuitonStats(analysis, candles)
		fmt.Println("\n\n=== ESTADÍSTICAS FINALES ===\n")
		fmt.Printf("Total de precios analizados: %d\n", stats.TotalPrices)
		fmt.Printf("Precio promedio: %.2f\n", stats.MeanPrice)
		fmt.Printf("Desviación estándar: %.2f\n", stats.StdDeviation)
		fmt.Printf("Volatilidad: %.2f%%\n", stats.VolatilityPct)
		fmt.Printf("Distribución normal: %v\n", stats.IsNormal)
		fmt.Printf("Fuerza del trend: %.4f\n", stats.TrendStrength)
	}
}
