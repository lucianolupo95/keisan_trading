package main

import (
	"fmt"
	"math"
)

// OrchestratorConfig contiene configuración del orquestador
type OrchestratorConfig struct {
	InitialCapital     float64
	UseKaton           bool // Usar análisis técnico
	UseEnhancedSignals bool // Usar signals mejoradas
	Verbose            bool // Output detallado
}

// OrchestratorResult contiene resultados de la orquestación
type OrchestratorResult struct {
	BacktestResult     BacktestResult
	AveragePrice       float64
	PriceRange         float64
	VolatilityPercent  float64
	TrendDirection     string
	BUYCount           int
	SELLCount          int
	HOLDCount          int
	SuccessRate        float64
}

// Orchestrator coordina múltiples módulos
type Orchestrator struct {
	Candles []Candle
	Config  OrchestratorConfig
}

// NewOrchestrator crea una nueva instancia del orquestador
func NewOrchestrator(candles []Candle, config OrchestratorConfig) *Orchestrator {
	return &Orchestrator{
		Candles: candles,
		Config:  config,
	}
}

// Run ejecuta la orquestación completa
func (o *Orchestrator) Run() OrchestratorResult {
	result := OrchestratorResult{}

	// Recolectar señales
	var allPrices []float64
	var signals []string

	for _, candle := range o.Candles {
		allPrices = append(allPrices, candle.Close)
		signal := GenerateSignal(candle)

		// Usar señales mejoradas si está habilitado
		if o.Config.UseKaton && o.Config.UseEnhancedSignals {
			signal = EnhancedSignal(candle, allPrices)
		}

		signals = append(signals, signal)

		if signal == "BUY" {
			result.BUYCount++
		} else if signal == "SELL" {
			result.SELLCount++
		} else {
			result.HOLDCount++
		}
	}

	// Calcular métricas de mercado
	result.AveragePrice = o.calculateAvgPrice(allPrices)
	result.PriceRange = o.calculatePriceRange(allPrices)
	result.VolatilityPercent = o.calculateVolatility(allPrices)
	result.TrendDirection = o.calculateTrend(allPrices)

	// Ejecutar backtest
	backtester := NewBacktester(o.Candles, o.Config.InitialCapital)
	result.BacktestResult = backtester.RunBacktest()

	// Calcular tasa de éxito
	if len(result.BacktestResult.Trades) > 0 {
		result.SuccessRate = (float64(result.BacktestResult.WinningTrades) / float64(len(result.BacktestResult.Trades))) * 100
	}

	return result
}

// calculateAvgPrice calcula el precio promedio
func (o *Orchestrator) calculateAvgPrice(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}
	sum := 0.0
	for _, p := range prices {
		sum += p
	}
	return sum / float64(len(prices))
}

// calculatePriceRange calcula el rango de precios
func (o *Orchestrator) calculatePriceRange(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxPrice := prices[0]

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		if p > maxPrice {
			maxPrice = p
		}
	}

	return maxPrice - minPrice
}

// calculateVolatility calcula volatilidad como desviación estándar porcentual
func (o *Orchestrator) calculateVolatility(prices []float64) float64 {
	if len(prices) < 2 {
		return 0
	}

	avg := o.calculateAvgPrice(prices)
	variance := 0.0

	for _, p := range prices {
		variance += (p - avg) * (p - avg)
	}
	variance /= float64(len(prices))

	stdDev := math.Sqrt(variance)
	volatilityPercent := (stdDev / avg) * 100

	return math.Round(volatilityPercent*100) / 100
}

// calculateTrend calcula la tendencia general
func (o *Orchestrator) calculateTrend(prices []float64) string {
	if len(prices) < 5 {
		return "insufficient_data"
	}

	// Dividir en 3 períodos
	thirtyPercent := len(prices) / 3
	sixtyPercent := (len(prices) * 2) / 3

	period1Avg := o.calculateAvgPrice(prices[:thirtyPercent])
	period2Avg := o.calculateAvgPrice(prices[thirtyPercent:sixtyPercent])
	period3Avg := o.calculateAvgPrice(prices[sixtyPercent:])

	if period1Avg < period2Avg && period2Avg < period3Avg {
		return "uptrend"
	} else if period1Avg > period2Avg && period2Avg > period3Avg {
		return "downtrend"
	}

	return "sideways"
}

// PrintReport imprime un reporte completo
func (r OrchestratorResult) PrintReport() {
	separator := "============================================================"
	fmt.Println("\n" + separator)
	fmt.Println("ORCHESTRATOR FULL REPORT - KEISAN TRADING")
	fmt.Println(separator + "\n")

	// Market Analysis
	fmt.Println("📊 MARKET ANALYSIS")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("Average Price:    $%.2f\n", r.AveragePrice)
	fmt.Printf("Price Range:      $%.2f\n", r.PriceRange)
	fmt.Printf("Volatility:       %.2f%%\n", r.VolatilityPercent)
	fmt.Printf("Trend:            %s\n\n", r.TrendDirection)

	// Signal Distribution
	fmt.Println("🎯 SIGNAL DISTRIBUTION")
	fmt.Println("------------------------------------------------------------")
	totalSignals := r.BUYCount + r.SELLCount + r.HOLDCount
	fmt.Printf("Total Signals:    %d\n", totalSignals)
	fmt.Printf("BUY Signals:      %d (%.1f%%)\n", r.BUYCount, float64(r.BUYCount)/float64(totalSignals)*100)
	fmt.Printf("SELL Signals:     %d (%.1f%%)\n", r.SELLCount, float64(r.SELLCount)/float64(totalSignals)*100)
	fmt.Printf("HOLD Signals:     %d (%.1f%%)\n\n", r.HOLDCount, float64(r.HOLDCount)/float64(totalSignals)*100)

	// Backtest Results
	fmt.Println("💰 BACKTEST RESULTS")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("Initial Capital:  $%.2f\n", r.BacktestResult.InitialCapital)
	fmt.Printf("Final Capital:    $%.2f\n", r.BacktestResult.FinalCapital)
	fmt.Printf("Total Return:     $%.2f (%.2f%%)\n", r.BacktestResult.TotalReturn, r.BacktestResult.TotalReturnPC)
	fmt.Printf("Total Trades:     %d\n", r.BacktestResult.TotalTrades)
	fmt.Printf("Winning Trades:   %d\n", r.BacktestResult.WinningTrades)
	fmt.Printf("Success Rate:     %.2f%%\n", r.SuccessRate)
	fmt.Printf("Max Drawdown:     %.2f%%\n\n", r.BacktestResult.MaxDrawdown)

	// Trade Details
	if len(r.BacktestResult.Trades) > 0 {
		fmt.Println("📈 TRADES EXECUTED")
		fmt.Println("------------------------------------------------------------")
		for i, trade := range r.BacktestResult.Trades {
			fmt.Printf("Trade %d: Vela %d → %d | $%.2f → $%.2f | P&L: $%.2f (%.2f%%)\n",
				i+1,
				trade.EntryIndex+1,
				trade.ExitIndex+1,
				trade.EntryPrice,
				trade.ExitPrice,
				trade.ProfitLoss,
				trade.ProfitLossPC)
		}
	}

	fmt.Println("\n" + separator + "\n")
}
