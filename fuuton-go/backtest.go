package main

import (
	"fmt"
	"math"
)

// Trade representa un trade ejecutado
type Trade struct {
	EntryIndex   int
	ExitIndex    int
	EntryPrice   float64
	ExitPrice    float64
	EntrySignal  string
	ExitSignal   string
	ProfitLoss   float64
	ProfitLossPC float64
	Status       string // "open" o "closed"
}

// BacktestResult contiene resultados del backtest
type BacktestResult struct {
	InitialCapital float64
	FinalCapital   float64
	TotalReturn    float64
	TotalReturnPC  float64
	TotalTrades    int
	WinningTrades  int
	LosingTrades   int
	WinRate        float64
	MaxDrawdown    float64
	Trades         []Trade
}

// Backtester ejecuta un backtest sobre datos históricos
type Backtester struct {
	Candles        []Candle
	InitialCapital float64
	TradingAmount  float64 // Cantidad por trade
	CurrentCapital float64
	Trades         []Trade
	MaxCapital     float64
	MinCapital     float64
}

// NewBacktester crea una nueva instancia del backtester
func NewBacktester(candles []Candle, initialCapital float64) *Backtester {
	return &Backtester{
		Candles:        candles,
		InitialCapital: initialCapital,
		TradingAmount:  initialCapital * 0.1, // 10% por trade
		CurrentCapital: initialCapital,
		MaxCapital:     initialCapital,
		MinCapital:     initialCapital,
		Trades:         []Trade{},
	}
}

// RunBacktest ejecuta el backtest completo
func (b *Backtester) RunBacktest() BacktestResult {
	var openTrade *Trade

	for i, candle := range b.Candles {
		signal := GenerateSignal(candle)

		// Entrar en posición (BUY)
		if signal == "BUY" && openTrade == nil {
			openTrade = &Trade{
				EntryIndex:  i,
				EntryPrice:  candle.Close,
				EntrySignal: signal,
				Status:      "open",
			}
		}

		// Salir de posición (SELL)
		if signal == "SELL" && openTrade != nil {
			openTrade.ExitIndex = i
			openTrade.ExitPrice = candle.Close
			openTrade.ExitSignal = signal
			openTrade.Status = "closed"

			// Calcular P&L
			openTrade.ProfitLoss = (openTrade.ExitPrice - openTrade.EntryPrice) * (b.TradingAmount / openTrade.EntryPrice)
			openTrade.ProfitLossPC = ((openTrade.ExitPrice - openTrade.EntryPrice) / openTrade.EntryPrice) * 100

			// Actualizar capital
			b.CurrentCapital += openTrade.ProfitLoss
			b.Trades = append(b.Trades, *openTrade)

			// Actualizar máximo y mínimo
			if b.CurrentCapital > b.MaxCapital {
				b.MaxCapital = b.CurrentCapital
			}
			if b.CurrentCapital < b.MinCapital {
				b.MinCapital = b.CurrentCapital
			}

			openTrade = nil
		}
	}

	// Si hay trade abierto al final, ciérralo con el último precio
	if openTrade != nil {
		lastCandle := b.Candles[len(b.Candles)-1]
		openTrade.ExitIndex = len(b.Candles) - 1
		openTrade.ExitPrice = lastCandle.Close
		openTrade.ExitSignal = "END"
		openTrade.Status = "closed"

		openTrade.ProfitLoss = (openTrade.ExitPrice - openTrade.EntryPrice) * (b.TradingAmount / openTrade.EntryPrice)
		openTrade.ProfitLossPC = ((openTrade.ExitPrice - openTrade.EntryPrice) / openTrade.EntryPrice) * 100

		b.CurrentCapital += openTrade.ProfitLoss
		b.Trades = append(b.Trades, *openTrade)

		if b.CurrentCapital > b.MaxCapital {
			b.MaxCapital = b.CurrentCapital
		}
		if b.CurrentCapital < b.MinCapital {
			b.MinCapital = b.CurrentCapital
		}
	}

	return b.GenerateReport()
}

// GenerateReport genera un reporte detallado del backtest
func (b *Backtester) GenerateReport() BacktestResult {
	totalReturn := b.CurrentCapital - b.InitialCapital
	totalReturnPC := (totalReturn / b.InitialCapital) * 100

	winningTrades := 0
	losingTrades := 0

	for _, trade := range b.Trades {
		if trade.ProfitLoss > 0 {
			winningTrades++
		} else if trade.ProfitLoss < 0 {
			losingTrades++
		}
	}

	winRate := 0.0
	if len(b.Trades) > 0 {
		winRate = (float64(winningTrades) / float64(len(b.Trades))) * 100
	}

	maxDrawdown := b.calculateMaxDrawdown()

	return BacktestResult{
		InitialCapital: b.InitialCapital,
		FinalCapital:   b.CurrentCapital,
		TotalReturn:    totalReturn,
		TotalReturnPC:  totalReturnPC,
		TotalTrades:    len(b.Trades),
		WinningTrades:  winningTrades,
		LosingTrades:   losingTrades,
		WinRate:        winRate,
		MaxDrawdown:    maxDrawdown,
		Trades:         b.Trades,
	}
}

// calculateMaxDrawdown calcula la peor pérdida desde máximo
func (b *Backtester) calculateMaxDrawdown() float64 {
	if len(b.Trades) == 0 {
		return 0
	}

	maxDrawdown := 0.0
	peakCapital := b.InitialCapital

	currentCap := b.InitialCapital
	for _, trade := range b.Trades {
		currentCap += trade.ProfitLoss

		if currentCap > peakCapital {
			peakCapital = currentCap
		}

		drawdown := ((peakCapital - currentCap) / peakCapital) * 100
		if drawdown > maxDrawdown {
			maxDrawdown = drawdown
		}
	}

	return math.Round(maxDrawdown*100) / 100
}

// CalculateSharpeRatio calcula el ratio de Sharpe (retorno ajustado por riesgo)
func (b *Backtester) CalculateSharpeRatio(riskFreeRate float64) float64 {
	if len(b.Trades) < 2 {
		return 0
	}

	// Calcular retornos de cada trade
	returns := make([]float64, len(b.Trades))
	for i, trade := range b.Trades {
		returns[i] = trade.ProfitLossPC
	}

	// Calcular promedio de retornos
	avgReturn := 0.0
	for _, r := range returns {
		avgReturn += r
	}
	avgReturn /= float64(len(returns))

	// Calcular desviación estándar
	variance := 0.0
	for _, r := range returns {
		variance += (r - avgReturn) * (r - avgReturn)
	}
	variance /= float64(len(returns))
	stdDev := math.Sqrt(variance)

	if stdDev == 0 {
		return 0
	}

	sharpe := (avgReturn - riskFreeRate) / stdDev
	return math.Round(sharpe*100) / 100
}

// PrintReport imprime un reporte bonito en consola
func (r *BacktestResult) PrintReport() {
	separator := "============================================================"
	fmt.Println("\n" + separator)
	fmt.Println("BACKTEST REPORT")
	fmt.Println(separator + "\n")

	fmt.Printf("Initial Capital:  $%.2f\n", r.InitialCapital)
	fmt.Printf("Final Capital:    $%.2f\n", r.FinalCapital)
	fmt.Printf("Total Return:     $%.2f (%.2f%%)\n", r.TotalReturn, r.TotalReturnPC)
	fmt.Println()

	fmt.Printf("Total Trades:     %d\n", r.TotalTrades)
	fmt.Printf("Winning Trades:   %d\n", r.WinningTrades)
	fmt.Printf("Losing Trades:    %d\n", r.LosingTrades)
	fmt.Printf("Win Rate:         %.2f%%\n", r.WinRate)
	fmt.Printf("Max Drawdown:     %.2f%%\n", r.MaxDrawdown)

	fmt.Println()
	fmt.Println("TRADES EXECUTED:")
	fmt.Println("------------------------------------------------------------")

	for i, trade := range r.Trades {
		fmt.Printf("Trade %d:\n", i+1)
		fmt.Printf("  Entry:  Vela %d @ $%.2f\n", trade.EntryIndex+1, trade.EntryPrice)
		fmt.Printf("  Exit:   Vela %d @ $%.2f\n", trade.ExitIndex+1, trade.ExitPrice)
		fmt.Printf("  P&L:    $%.2f (%.2f%%)\n", trade.ProfitLoss, trade.ProfitLossPC)
		fmt.Println()
	}

	fmt.Println(separator + "\n")
}
