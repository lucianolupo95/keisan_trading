package main

import (
	"fmt"
	"math"
)

// ============================================================================
// TECHNICAL INDICATORS MODULE - Day 5
// Purpose: Calculate technical indicators (MA, RSI, MACD) for trading signals
// ============================================================================

// ============================================================================
// MOVING AVERAGE (MA20) - Simple Moving Average
// ============================================================================

// CalculateMA20 calculates the 20-period Simple Moving Average
// MA20 = average of last 20 closing prices
func CalculateMA20(prices []float64) float64 {
	if len(prices) < 20 {
		// If fewer than 20 prices, average all available
		return CalculateSMA(prices, len(prices))
	}
	return CalculateSMA(prices[len(prices)-20:], 20)
}

// CalculateSMA calculates Simple Moving Average for n periods
func CalculateSMA(prices []float64, period int) float64 {
	if len(prices) < period || period <= 0 {
		return 0
	}

	sum := 0.0
	for i := len(prices) - period; i < len(prices); i++ {
		sum += prices[i]
	}
	return sum / float64(period)
}

// CalculateEMA calculates Exponential Moving Average
// EMA gives more weight to recent prices
func CalculateEMA(prices []float64, period int) float64 {
	if len(prices) < period || period <= 0 {
		return 0
	}

	// Multiplier for EMA
	multiplier := 2.0 / float64(period+1)

	// Start with SMA
	ema := CalculateSMA(prices[:period], period)

	// Apply EMA formula
	for i := period; i < len(prices); i++ {
		ema = prices[i]*multiplier + ema*(1-multiplier)
	}

	return ema
}

// ============================================================================
// RSI (Relative Strength Index) - Momentum Indicator
// ============================================================================

// CalculateRSI calculates the Relative Strength Index
// RSI measures momentum and the magnitude of price changes
// RSI < 30: Oversold (potential BUY signal)
// RSI > 70: Overbought (potential SELL signal)
// RSI between 30-70: Neutral
func CalculateRSI(prices []float64, period int) float64 {
	if len(prices) < period+1 {
		return 50.0 // Neutral if insufficient data
	}

	// Calculate price changes
	gains := 0.0
	losses := 0.0

	for i := len(prices) - period; i < len(prices); i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			gains += change
		} else {
			losses += -change
		}
	}

	// Avoid division by zero
	if losses == 0 {
		if gains > 0 {
			return 100.0 // All gains
		}
		return 50.0 // No change
	}

	// Calculate average gains and losses
	avgGain := gains / float64(period)
	avgLoss := losses / float64(period)

	// Calculate RS (Relative Strength)
	rs := avgGain / avgLoss

	// Calculate RSI
	rsi := 100 - (100 / (1 + rs))

	// Clamp to 0-100 range
	if rsi < 0 {
		rsi = 0
	} else if rsi > 100 {
		rsi = 100
	}

	return rsi
}

// ============================================================================
// MACD (Moving Average Convergence Divergence)
// ============================================================================

// MACDValue contains MACD components
type MACDValue struct {
	MACD      float64 // EMA12 - EMA26
	Signal    float64 // EMA9 of MACD
	Histogram float64 // MACD - Signal
	IsBullish bool    // True if MACD > Signal (uptrend)
}

// CalculateMACD calculates MACD (Moving Average Convergence Divergence)
// MACD is a trend-following momentum indicator
// MACD > Signal: Bullish (uptrend)
// MACD < Signal: Bearish (downtrend)
func CalculateMACD(prices []float64) MACDValue {
	if len(prices) < 26 {
		return MACDValue{
			MACD:      0,
			Signal:    0,
			Histogram: 0,
			IsBullish: false,
		}
	}

	// Calculate 12-period and 26-period EMAs
	ema12 := CalculateEMA(prices, 12)
	ema26 := CalculateEMA(prices, 26)

	// MACD line = EMA12 - EMA26
	macdLine := ema12 - ema26

	// For signal line, we need to calculate EMA9 of the MACD values
	// For simplicity, use EMA9 approximation with recent MACD values
	signalLine := macdLine * 0.8 // Simplified signal (in production, calculate full EMA9 of MACD series)

	// Histogram = MACD - Signal
	histogram := macdLine - signalLine

	// Determine bullish/bearish
	isBullish := macdLine > signalLine && histogram > 0

	return MACDValue{
		MACD:      macdLine,
		Signal:    signalLine,
		Histogram: histogram,
		IsBullish: isBullish,
	}
}

// ============================================================================
// INDICATOR ANALYSIS STRUCT
// ============================================================================

// IndicatorAnalysis contains all technical indicators for a set of prices
type IndicatorAnalysis struct {
	MA20     float64   // 20-period Simple Moving Average
	RSI14    float64   // 14-period RSI
	MACD     MACDValue // MACD components
	Timestamp string   // Timestamp of analysis
}

// AnalyzeTechnicalIndicators calculates all technical indicators
func AnalyzeTechnicalIndicators(prices []float64) IndicatorAnalysis {
	return IndicatorAnalysis{
		MA20:  CalculateMA20(prices),
		RSI14: CalculateRSI(prices, 14),
		MACD:  CalculateMACD(prices),
	}
}

// ============================================================================
// INDICATOR INTERPRETATION
// ============================================================================

// GetMA20Interpretation returns trading signal based on MA20
func GetMA20Interpretation(currentPrice float64, ma20 float64) string {
	if currentPrice > ma20 {
		return "Above MA20 (bullish)"
	} else if currentPrice < ma20 {
		return "Below MA20 (bearish)"
	}
	return "At MA20 (neutral)"
}

// GetRSIInterpretation returns trading signal based on RSI
func GetRSIInterpretation(rsi float64) string {
	if rsi < 30 {
		return "Oversold (buy signal)"
	} else if rsi > 70 {
		return "Overbought (sell signal)"
	}
	return "Neutral"
}

// GetMACDInterpretation returns trading signal based on MACD
func GetMACDInterpretation(macd MACDValue) string {
	if macd.IsBullish {
		if macd.Histogram > 0 {
			return "Bullish (uptrend)"
		}
		return "Bullish (MACD > Signal)"
	}
	if macd.Histogram < 0 {
		return "Bearish (downtrend)"
	}
	return "Bearish (MACD < Signal)"
}

// ============================================================================
// INDICATOR CONFIDENCE SCORING
// ============================================================================

// CalculateIndicatorConfidence returns a confidence score based on indicators
// Used to enhance signal generation
func CalculateIndicatorConfidence(candle Candle, indicators IndicatorAnalysis) float64 {
	confidence := 0.5 // Base confidence

	// RSI factor: Oversold/Overbought
	if indicators.RSI14 < 30 {
		confidence += 0.15 // Oversold = buying opportunity
	} else if indicators.RSI14 > 70 {
		confidence -= 0.15 // Overbought = caution
	}

	// MA20 factor: Price trend
	if candle.Close > indicators.MA20 {
		confidence += 0.1 // Above MA20 = uptrend
	} else if candle.Close < indicators.MA20 {
		confidence -= 0.1 // Below MA20 = downtrend
	}

	// MACD factor: Momentum
	if indicators.MACD.IsBullish {
		confidence += 0.1 // Bullish MACD = momentum up
	} else {
		confidence -= 0.1 // Bearish MACD = momentum down
	}

	// Clamp confidence to 0-1 range
	if confidence < 0 {
		confidence = 0
	} else if confidence > 1 {
		confidence = 1
	}

	return confidence
}

// ============================================================================
// PRINTING FUNCTIONS
// ============================================================================

// PrintIndicatorAnalysis prints all technical indicators in readable format
func PrintIndicatorAnalysis(indicators IndicatorAnalysis) {
	fmt.Println("\n📈 TECHNICAL INDICATORS")
	fmt.Println("═══════════════════════════════════════════════════")

	fmt.Printf("MA20 (20-period Moving Average):    %.2f\n", indicators.MA20)
	fmt.Printf("RSI14 (14-period RSI):              %.2f (%s)\n",
		indicators.RSI14, GetRSIInterpretation(indicators.RSI14))

	fmt.Printf("\nMACD Components:\n")
	fmt.Printf("  MACD Line:      %.4f\n", indicators.MACD.MACD)
	fmt.Printf("  Signal Line:    %.4f\n", indicators.MACD.Signal)
	fmt.Printf("  Histogram:      %.4f\n", indicators.MACD.Histogram)
	fmt.Printf("  Status:         %s\n", GetMACDInterpretation(indicators.MACD))
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// ExtractClosePrices extracts closing prices from candles
func ExtractClosePrices(candles []Candle) []float64 {
	prices := make([]float64, len(candles))
	for i, c := range candles {
		prices[i] = c.Close
	}
	return prices
}

// RoundFloat rounds a float64 to 2 decimal places
func RoundFloat(f float64, decimals int) float64 {
	pow10n := math.Pow(10, float64(decimals))
	return math.Round(f*pow10n) / pow10n
}
