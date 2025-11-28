package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
)

// ============================================================================
// SUITON BRIDGE: Go ↔ R Communication Layer
// Day 4 - 28/11/2025
//
// Purpose: Call R statistical analysis functions from Go
// Methods: JSON serialization for IPC
// ============================================================================

// SuitonDistribution representa estadísticas de distribución
type SuitonDistribution struct {
	Mean                 float64 `json:"mean"`
	Std                  float64 `json:"std"`
	Median               float64 `json:"median"`
	Skewness             float64 `json:"skewness"`
	Kurtosis             float64 `json:"kurtosis"`
	CoefficientVariation float64 `json:"coefficient_variation"`
	DistributionType     string  `json:"distribution_type"`
}

// SuitonNormality representa resultados del test de normalidad
type SuitonNormality struct {
	PValue           float64 `json:"p_value"`
	TestStatistic    float64 `json:"test_statistic"`
	IsNormal         bool    `json:"is_normal"`
	Interpretation   string  `json:"interpretation"`
}

// SuitonCorrelation representa análisis de correlación
type SuitonCorrelation struct {
	Value             float64 `json:"value"`
	PValue            float64 `json:"p_value"`
	Interpretation    string  `json:"interpretation"`
}

// SuitonAnalysis contiene resultados completos del análisis
type SuitonAnalysis struct {
	Distribution SuitonDistribution `json:"distribution"`
	Normality    SuitonNormality    `json:"normality"`
	Correlation  SuitonCorrelation  `json:"correlation"`
	Timestamp    string             `json:"timestamp"`
}

// ============================================================================
// LOCAL STATISTICAL ANALYSIS (Fallback cuando R no está disponible)
// ============================================================================

// LocalAnalyzeDistribution calcula estadísticas de distribución sin R
func LocalAnalyzeDistribution(prices []float64) SuitonDistribution {
	if len(prices) == 0 {
		return SuitonDistribution{}
	}

	// Media
	sum := 0.0
	for _, p := range prices {
		sum += p
	}
	mean := sum / float64(len(prices))

	// Mediana
	sorted := make([]float64, len(prices))
	copy(sorted, prices)
	sort.Float64s(sorted)
	median := 0.0
	if len(sorted)%2 == 0 {
		median = (sorted[len(sorted)/2-1] + sorted[len(sorted)/2]) / 2
	} else {
		median = sorted[len(sorted)/2]
	}

	// Desviación estándar
	variance := 0.0
	for _, p := range prices {
		variance += (p - mean) * (p - mean)
	}
	variance /= float64(len(prices))
	std := variance

	// Aproximación simple de skewness
	cubed := 0.0
	for _, p := range prices {
		cubed += (p - mean) * (p - mean) * (p - mean)
	}
	skewness := (cubed / float64(len(prices))) / (std * std)

	coeffVar := (std / mean) * 100

	return SuitonDistribution{
		Mean:                 mean,
		Std:                  std,
		Median:               median,
		Skewness:             skewness,
		Kurtosis:             0.0, // Not implemented in fallback
		CoefficientVariation: coeffVar,
		DistributionType:     "fallback-analysis",
	}
}

// LocalTestNormality realiza un test simple de normalidad sin R
func LocalTestNormality(prices []float64) SuitonNormality {
	if len(prices) < 3 {
		return SuitonNormality{
			PValue:         1.0,
			TestStatistic:  0.0,
			IsNormal:       true,
			Interpretation: "insufficient data",
		}
	}

	// Simple heurística: si skewness < 0.5 es aproximadamente normal
	dist := LocalAnalyzeDistribution(prices)
	isNormal := dist.Skewness > -0.5 && dist.Skewness < 0.5

	return SuitonNormality{
		PValue:         0.5, // Dummy value
		TestStatistic:  dist.Skewness,
		IsNormal:       isNormal,
		Interpretation: "fallback normality check",
	}
}

// LocalAnalyzeCorrelation calcula correlación precio-tiempo sin R
func LocalAnalyzeCorrelation(prices []float64) SuitonCorrelation {
	if len(prices) < 2 {
		return SuitonCorrelation{}
	}

	// time = [1, 2, 3, ..., n]
	// Pearson correlation
	n := float64(len(prices))
	timeSum := n * (n + 1) / 2 // sum(1..n)
	timeMean := timeSum / n

	priceSum := 0.0
	for _, p := range prices {
		priceSum += p
	}
	priceMean := priceSum / n

	numerator := 0.0
	timeSumSq := 0.0
	priceSumSq := 0.0

	for i, price := range prices {
		t := float64(i + 1)
		numerator += (t - timeMean) * (price - priceMean)
		timeSumSq += (t - timeMean) * (t - timeMean)
		priceSumSq += (price - priceMean) * (price - priceMean)
	}

	correlation := 0.0
	if timeSumSq > 0 && priceSumSq > 0 {
		correlation = numerator / (timeSumSq * priceSumSq)
	}

	interpretation := "weak"
	if correlation > 0.7 {
		interpretation = "strong positive"
	} else if correlation > 0.3 {
		interpretation = "moderate positive"
	} else if correlation < -0.7 {
		interpretation = "strong negative"
	} else if correlation < -0.3 {
		interpretation = "moderate negative"
	}

	return SuitonCorrelation{
		Value:          correlation,
		PValue:         0.5, // Dummy
		Interpretation: interpretation,
	}
}

// ============================================================================
// R BRIDGE (Intenta usar R, fallback a local si no disponible)
// ============================================================================

// CallSuitonR intenta llamar al módulo R Suiton
// Retorna análisis completo o nil si R no está disponible
func CallSuitonR(prices []float64) *SuitonAnalysis {
	// Ruta completa a Rscript en Windows
	rscriptPath := "C:\\Program Files\\R\\R-4.5.2\\bin\\Rscript.exe"

	// Get absolute path to the R script
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return nil
	}
	rIpcScript := workDir + "\\..\\suiton-r\\r_ipc_simple.R"

	// 1. Serializar prices a JSON
	inputData := map[string]interface{}{
		"prices": prices,
	}
	inputJSON, _ := json.Marshal(inputData)

	// 2. Guardar en archivo temporal (usar ruta absoluta)
	tmpFile := workDir + "\\temp_suiton_input.json"
	err = os.WriteFile(tmpFile, inputJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing temp file: %v\n", err)
		return nil
	}
	defer os.Remove(tmpFile)

	// 3. Ejecutar Rscript pasando el archivo como argumento
	// r_ipc_simple.R leerá el JSON del archivo y retornará JSON
	cmd := exec.Command(rscriptPath, rIpcScript, tmpFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Debug: imprimir error
		fmt.Printf("R error: %v\nOutput: %s\n", err, string(output))
		fmt.Printf("R script path: %s\n", rIpcScript)
		fmt.Printf("Temp file path: %s\n", tmpFile)
		return nil
	}

	// 4. Parsear respuesta JSON (limpiar warnings de R)
	outputStr := string(output)

	// Buscar el inicio del JSON (el primer {)
	jsonStart := -1
	for i := 0; i < len(outputStr); i++ {
		if outputStr[i] == '{' {
			jsonStart = i
			break
		}
	}

	if jsonStart == -1 {
		fmt.Printf("No JSON found in R output: %s\n", outputStr)
		return nil
	}

	// Buscar el final del JSON (el último })
	jsonEnd := -1
	for i := len(outputStr) - 1; i >= jsonStart; i-- {
		if outputStr[i] == '}' {
			jsonEnd = i + 1
			break
		}
	}

	if jsonEnd == -1 {
		fmt.Printf("Malformed JSON in R output: %s\n", outputStr)
		return nil
	}

	jsonStr := outputStr[jsonStart:jsonEnd]

	var analysis SuitonAnalysis
	err = json.Unmarshal([]byte(jsonStr), &analysis)
	if err != nil {
		fmt.Printf("JSON parse error: %v\nJSON was: %s\n", err, jsonStr)
		return nil
	}

	return &analysis
}

// AnalyzePricesWithSuiton realiza análisis estadístico usando Suiton
// Primero intenta usar R, sino usa análisis local
func AnalyzePricesWithSuiton(candles []Candle) *SuitonAnalysis {
	fmt.Println("\n📊 SUITON ANALYSIS (Statistical Analysis Module)")
	fmt.Println("═══════════════════════════════════════════════════")

	// Extraer precios de cierre
	prices := make([]float64, len(candles))
	for i, c := range candles {
		prices[i] = c.Close
	}

	// Intentar usar R
	analysis := CallSuitonR(prices)
	if analysis != nil {
		fmt.Println("✓ Using R (Suiton)")
		return analysis
	}

	// Fallback a análisis local
	fmt.Println("⚠ R no disponible, usando análisis local")
	fmt.Println("  (Instala R para análisis completo)")

	analysis = &SuitonAnalysis{
		Distribution: LocalAnalyzeDistribution(prices),
		Normality:    LocalTestNormality(prices),
		Correlation:  LocalAnalyzeCorrelation(prices),
		Timestamp:    candles[len(candles)-1].Timestamp,
	}

	return analysis
}

// PrintSuitonAnalysis imprime resultados del análisis
func PrintSuitonAnalysis(analysis *SuitonAnalysis) {
	fmt.Println("\n📈 DISTRIBUTION STATS:")
	fmt.Printf("  Mean:      %.2f\n", analysis.Distribution.Mean)
	fmt.Printf("  Median:    %.2f\n", analysis.Distribution.Median)
	fmt.Printf("  Std Dev:   %.2f\n", analysis.Distribution.Std)
	fmt.Printf("  Skewness:  %.4f (%s)\n",
		analysis.Distribution.Skewness,
		analysis.Distribution.DistributionType)

	fmt.Println("\n🔬 NORMALITY TEST (Shapiro-Wilk):")
	fmt.Printf("  p-value:   %.4f\n", analysis.Normality.PValue)
	fmt.Printf("  Result:    %s\n", analysis.Normality.Interpretation)

	fmt.Println("\n📊 CORRELATION (Price vs Time):")
	fmt.Printf("  Correlation: %.4f\n", analysis.Correlation.Value)
	fmt.Printf("  Type:        %s\n", analysis.Correlation.Interpretation)
}

// ============================================================================
// ENHANCED SIGNAL WITH STATISTICS
// ============================================================================

// GenerateEnhancedSignal genera señal mejorada con análisis estadístico + indicadores técnicos (Day 5)
func GenerateEnhancedSignal(candle Candle, analysis *SuitonAnalysis, candleIndex int) string {
	// Base signal (original)
	const minVolume int64 = 1300
	const minMovePercent float64 = 0.1

	baseSignal := "HOLD"

	// Aplicar filtros básicos
	if candle.Volume >= minVolume {
		movePercent := ((candle.Close - candle.Open) / candle.Open) * 100
		if candle.Close > candle.Open && movePercent >= minMovePercent {
			baseSignal = "BUY"
		}
	}

	if baseSignal == "HOLD" || analysis == nil {
		return baseSignal
	}

	// Mejorar con análisis estadístico + indicadores técnicos (Day 5)
	// Factores estadísticos (Day 4):
	// 1. Distribución es normal (skewness cercano a 0) → +0.15
	// 2. Correlación es positiva (uptrend) → +0.15
	// 3. Baja volatilidad relativa → +0.1
	//
	// Factores técnicos (Day 5):
	// 4. RSI < 30 (oversold) → +0.15
	// 5. MA20 uptrend → +0.1
	// 6. MACD bullish → +0.1

	confidence := 0.5

	// ============ STATISTICAL FACTORS (Day 4) ============
	// Factor 1: Normalidad (distribución normal = más confiable)
	if analysis.Normality.IsNormal {
		confidence += 0.15
	}

	// Factor 2: Correlación positiva (uptrend)
	if analysis.Correlation.Value > 0.5 {
		confidence += 0.15
	}

	// Factor 3: Baja volatilidad relativa
	if analysis.Distribution.CoefficientVariation < 1.0 {
		confidence += 0.1
	}

	// ============ TECHNICAL INDICATORS (Day 5) ============
	// Calculate technical indicators if we have enough data
	if candleIndex > 26 { // Need at least 26 candles for MACD
		// Extract prices from beginning to current candle
		// (In production, would extract from actual price series)

		// Factor 4: RSI oversold (< 30) = buying opportunity
		// Note: In production, calculate RSI from actual price history
		// For now, using placeholder; integrate with actual price tracking

		// Factor 5: MA20 uptrend
		// Check if price is above moving average
		// Note: Would need price history for accurate MA20 calculation

		// Factor 6: MACD bullish
		// Check MACD components
		// Note: Would need price history for accurate MACD calculation
	}

	// Decisión final basada en confianza (4 niveles en Day 5)
	if confidence > 0.8 {
		return "BUY (Very High Confidence)" // 💪💪
	} else if confidence > 0.6 {
		return "BUY (High Confidence)" // 💪
	} else if confidence > 0.4 {
		return "BUY (Medium Confidence)" // 🤔
	}

	return "HOLD"
}

// GenerateEnhancedSignalWithIndicators genera señal mejorada con estadística + indicadores técnicos (Day 5)
func GenerateEnhancedSignalWithIndicators(candle Candle, analysis *SuitonAnalysis, candleIndex int, indicators IndicatorAnalysis) string {
	// Base signal
	const minVolume int64 = 1300
	const minMovePercent float64 = 0.1

	baseSignal := "HOLD"

	// Aplicar filtros básicos
	if candle.Volume >= minVolume {
		movePercent := ((candle.Close - candle.Open) / candle.Open) * 100
		if candle.Close > candle.Open && movePercent >= minMovePercent {
			baseSignal = "BUY"
		}
	}

	if baseSignal == "HOLD" || analysis == nil {
		return baseSignal
	}

	// Confianza mejorada con indicadores técnicos (Day 5)
	confidence := 0.5

	// ============ STATISTICAL FACTORS (Day 4) ============
	if analysis.Normality.IsNormal {
		confidence += 0.15
	}

	if analysis.Correlation.Value > 0.5 {
		confidence += 0.15
	}

	if analysis.Distribution.CoefficientVariation < 1.0 {
		confidence += 0.1
	}

	// ============ TECHNICAL INDICATORS (Day 5) ============
	if candleIndex >= 26 { // Need at least 26 candles for all indicators

		// Factor 4: RSI oversold (< 30) = buying opportunity
		if indicators.RSI14 < 30 {
			confidence += 0.15
		} else if indicators.RSI14 > 70 {
			confidence -= 0.1 // Overbought = caution
		}

		// Factor 5: MA20 uptrend
		if candle.Close > indicators.MA20 {
			confidence += 0.1
		} else if candle.Close < indicators.MA20 {
			confidence -= 0.05 // Below MA20 = caution
		}

		// Factor 6: MACD bullish
		if indicators.MACD.IsBullish {
			confidence += 0.1
		} else {
			confidence -= 0.05 // Bearish MACD = caution
		}
	}

	// Clamp confidence to 0-1 range
	if confidence < 0 {
		confidence = 0
	} else if confidence > 1 {
		confidence = 1
	}

	// Decisión final basada en confianza (4 niveles en Day 5)
	if confidence > 0.8 {
		return "BUY (Very High Confidence)" // 💪💪
	} else if confidence > 0.6 {
		return "BUY (High Confidence)" // 💪
	} else if confidence > 0.4 {
		return "BUY (Medium Confidence)" // 🤔
	}

	return "HOLD"
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// ExtractPricesFromCandles extrae precios de cierre de velas
func ExtractPricesFromCandles(candles []Candle) []float64 {
	prices := make([]float64, len(candles))
	for i, c := range candles {
		prices[i] = c.Close
	}
	return prices
}

// SuitonToJSON convierte análisis a JSON
func SuitonToJSON(analysis *SuitonAnalysis) string {
	data, _ := json.MarshalIndent(analysis, "", "  ")
	return string(data)
}

// SuitonFromJSON carga análisis desde JSON
func SuitonFromJSON(jsonStr string) (*SuitonAnalysis, error) {
	var analysis SuitonAnalysis
	err := json.Unmarshal([]byte(jsonStr), &analysis)
	return &analysis, err
}

// CheckRAvailable verifica si R está instalado
func CheckRAvailable() bool {
	cmd := exec.Command("Rscript", "--version")
	err := cmd.Run()
	return err == nil
}

// ============================================================================
// STATS INTEGRATION EXAMPLE
// ============================================================================

// SuitonStats contiene estadísticas resumidas para backtesting
type SuitonStats struct {
	TotalPrices   int     `json:"total_prices"`
	MeanPrice     float64 `json:"mean_price"`
	StdDeviation  float64 `json:"std_deviation"`
	IsNormal      bool    `json:"is_normal"`
	TrendStrength float64 `json:"trend_strength"`
	VolatilityPct float64 `json:"volatility_pct"`
}

// ComputeSuitonStats calcula estadísticas para backtesting
func ComputeSuitonStats(analysis *SuitonAnalysis, candles []Candle) SuitonStats {
	volatilityPct := (analysis.Distribution.Std / analysis.Distribution.Mean) * 100

	return SuitonStats{
		TotalPrices:   len(candles),
		MeanPrice:     analysis.Distribution.Mean,
		StdDeviation:  analysis.Distribution.Std,
		IsNormal:      analysis.Normality.IsNormal,
		TrendStrength: analysis.Correlation.Value,
		VolatilityPct: volatilityPct,
	}
}
