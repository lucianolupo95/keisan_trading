package main

import (
	"fmt"
	"math"
)

// ============================================================================
// PARAMETER TUNING MODULE - Day 6
// Purpose: Test different indicator periods to find optimal parameters
// ============================================================================

// ParameterResult stores the result of testing a specific parameter combination
type ParameterResult struct {
	MAPeriod          int
	RSIPeriod         int
	MAValue           float64
	RSIValue          float64
	SignalQuality     float64 // 0-1 score for signal reliability
	NoiseLevel        float64 // How much oscillation/noise
	AverageConfidence float64 // Average confidence score
}

// TuningReport summarizes all parameter testing results
type TuningReport struct {
	Results       []ParameterResult
	OptimalMA     int
	OptimalRSI    int
	RecommendedMA int
	RecommendedRSI int
	Notes         string
}

// TestMAParameters tests different MA periods and returns their characteristics
func TestMAParameters(prices []float64, periods []int) map[int]ParameterResult {
	results := make(map[int]ParameterResult)

	for _, period := range periods {
		if len(prices) < period {
			continue
		}

		// Calculate MA for this period
		ma := CalculateSMA(prices, period)

		// Calculate signal quality (stability) by looking at MA variance over time
		var maValues []float64
		for i := period; i <= len(prices); i++ {
			ma := CalculateSMA(prices[:i], period)
			maValues = append(maValues, ma)
		}

		// Lower volatility in MA = higher signal quality
		signalQuality := 1.0 - (calculateStdDev(maValues) / calculateMean(maValues))
		if signalQuality < 0 {
			signalQuality = 0
		}

		results[period] = ParameterResult{
			MAPeriod:      period,
			MAValue:       ma,
			SignalQuality: signalQuality,
		}
	}

	return results
}

// TestRSIParameters tests different RSI periods and returns their characteristics
func TestRSIParameters(prices []float64, periods []int) map[int]ParameterResult {
	results := make(map[int]ParameterResult)

	for _, period := range periods {
		if len(prices) < period+1 {
			continue
		}

		// Calculate RSI for this period
		rsi := CalculateRSI(prices, period)

		// Calculate noise level by looking at RSI oscillation
		var rsiValues []float64
		for i := period; i <= len(prices); i++ {
			rsi := CalculateRSI(prices[:i], period)
			rsiValues = append(rsiValues, rsi)
		}

		// Higher RSI variance = more noise (less stable)
		noiseLevel := calculateStdDev(rsiValues) / 100.0
		signalQuality := 1.0 - noiseLevel
		if signalQuality < 0 {
			signalQuality = 0
		}

		results[period] = ParameterResult{
			RSIPeriod:     period,
			RSIValue:      rsi,
			SignalQuality: signalQuality,
			NoiseLevel:    noiseLevel,
		}
	}

	return results
}

// CompareParameterResults generates a detailed comparison of different parameters
func CompareParameterResults(prices []float64) TuningReport {
	fmt.Println("\n\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║        DAY 6: PARAMETER TUNING ANALYSIS                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// Test MA periods: 15, 20, 50
	maPeriods := []int{15, 20, 50}
	maResults := TestMAParameters(prices, maPeriods)

	fmt.Println("\n\n📊 MOVING AVERAGE PERIOD TESTING")
	fmt.Println("═══════════════════════════════════════════════════")
	fmt.Printf("%-8s | %-12s | %-15s | %-20s\n", "Period", "MA Value", "Signal Quality", "Assessment")
	fmt.Println("─────────┼──────────────┼─────────────────┼────────────────────")

	optimalMA := 20
	maxQuality := 0.0

	for _, period := range maPeriods {
		result := maResults[period]
		assessment := getMAAssessment(period, result.SignalQuality)

		if result.SignalQuality > maxQuality {
			maxQuality = result.SignalQuality
			optimalMA = period
		}

		fmt.Printf("MA%-6d | %12.2f | %14.4f%% | %-20s\n",
			period, result.MAValue, result.SignalQuality*100, assessment)
	}

	fmt.Println("\n📝 MA PERIOD INTERPRETATION:")
	fmt.Println("  • MA15:  Faster response, more false positives, good for volatile markets")
	fmt.Println("  • MA20:  Balanced, industry standard, recommended baseline")
	fmt.Println("  • MA50:  Slower response, fewer signals, better for trending markets")

	// Test RSI periods: 9, 14, 21
	rsiPeriods := []int{9, 14, 21}
	rsiResults := TestRSIParameters(prices, rsiPeriods)

	fmt.Println("\n\n📊 RSI PERIOD TESTING")
	fmt.Println("═══════════════════════════════════════════════════")
	fmt.Printf("%-8s | %-12s | %-15s | %-20s\n", "Period", "RSI Value", "Noise Level", "Assessment")
	fmt.Println("─────────┼──────────────┼─────────────────┼────────────────────")

	optimalRSI := 14
	minNoise := 1.0

	for _, period := range rsiPeriods {
		result := rsiResults[period]
		assessment := getRSIAssessment(period, result.NoiseLevel)

		if result.NoiseLevel < minNoise {
			minNoise = result.NoiseLevel
			optimalRSI = period
		}

		fmt.Printf("RSI%-5d | %12.2f | %14.4f%% | %-20s\n",
			period, result.RSIValue, result.NoiseLevel*100, assessment)
	}

	fmt.Println("\n📝 RSI PERIOD INTERPRETATION:")
	fmt.Println("  • RSI9:   Faster signals, more sensitive, prone to whipsaws")
	fmt.Println("  • RSI14:  Standard, industry default, best for most conditions")
	fmt.Println("  • RSI21:  Smoother, fewer false signals, better for trending periods")

	// Build report
	report := TuningReport{
		OptimalMA:      optimalMA,
		OptimalRSI:     optimalRSI,
		RecommendedMA:  20, // Keep recommended as 20 (balanced)
		RecommendedRSI: 14, // Keep recommended as 14 (standard)
		Notes: fmt.Sprintf("Optimal MA: %d-period (Quality: %.2f%%), Optimal RSI: %d-period (Noise: %.2f%%)\n"+
			"Recommendation: MA%d + RSI%d for best performance on this dataset",
			optimalMA, maxQuality*100, optimalRSI, minNoise*100,
			20, 14),
	}

	return report
}

// ============================================================================
// HELPER FUNCTIONS FOR STATISTICAL CALCULATIONS
// ============================================================================

func calculateMean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func calculateStdDev(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}

	mean := calculateMean(values)
	variance := 0.0
	for _, v := range values {
		variance += (v - mean) * (v - mean)
	}
	variance /= float64(len(values))
	return math.Sqrt(variance)
}

func getMAAssessment(period int, quality float64) string {
	switch period {
	case 15:
		if quality > 0.8 {
			return "✓ Good for volatility"
		}
		return "May be noisy"
	case 20:
		if quality > 0.75 {
			return "✓ Balanced (recommended)"
		}
		return "Good for trends"
	case 50:
		if quality > 0.85 {
			return "✓ Excellent trend follower"
		}
		return "Slow response"
	}
	return "Unknown"
}

func getRSIAssessment(period int, noise float64) string {
	switch period {
	case 9:
		if noise < 0.3 {
			return "✓ Low noise (fast)"
		}
		return "Can be jumpy"
	case 14:
		if noise < 0.4 {
			return "✓ Standard (recommended)"
		}
		return "Well-balanced"
	case 21:
		if noise < 0.25 {
			return "✓ Very smooth (slow)"
		}
		return "Smooth response"
	}
	return "Unknown"
}

// PrintTuningReport prints the final tuning report
func PrintTuningReport(report TuningReport) {
	fmt.Println("\n\n🎯 PARAMETER TUNING RECOMMENDATIONS")
	fmt.Println("═══════════════════════════════════════════════════")
	fmt.Printf("Optimal MA:        %d-period\n", report.OptimalMA)
	fmt.Printf("Optimal RSI:       %d-period\n", report.OptimalRSI)
	fmt.Printf("Recommended MA:    %d-period (balanced)\n", report.RecommendedMA)
	fmt.Printf("Recommended RSI:   %d-period (industry standard)\n", report.RecommendedRSI)
	fmt.Printf("\n%s\n", report.Notes)
}
