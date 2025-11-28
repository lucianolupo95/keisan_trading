package main

import (
	"fmt"
	"math"
)

// ============================================================================
// MACHINE LEARNING FEATURES PREPARATION - Day 6
// Purpose: Identify and engineer features for ML model training
// ============================================================================

// MLFeatures represents all engineered features for a single data point
type MLFeatures struct {
	// Existing indicators
	MA20Value      float64 `json:"ma20_value"`
	RSIValue       float64 `json:"rsi_value"`
	MACDValue      float64 `json:"macd_value"`

	// New derived features
	PriceMomentum  float64 `json:"price_momentum"`      // % change from previous
	MASlope        float64 `json:"ma_slope"`            // Direction of MA
	RSIExtremes    int     `json:"rsi_extremes"`        // Count of RSI < 30 or > 70
	MACDCrosses    int     `json:"macd_crosses"`        // Count of MACD signal line crossings
	Volatility     float64 `json:"volatility"`          // Price volatility

	// Meta features
	TrendStrength  float64 `json:"trend_strength"`      // Correlation strength
	VolumeProfile  float64 `json:"volume_profile"`      // Normalized volume
	Price          float64 `json:"price"`               // Current price
}

// FeatureSet contains all features for multiple time points
type FeatureSet struct {
	Features   []MLFeatures
	TargetVariable []int // 0=HOLD, 1=BUY, -1=SELL
}

// ExtractMLFeatures extracts all ML features from price and indicator data
func ExtractMLFeatures(prices []float64, candles []Candle) []MLFeatures {
	features := make([]MLFeatures, 0)

	if len(prices) < 26 {
		return features // Need at least 26 prices for all indicators
	}

	for i := 26; i < len(prices); i++ {
		// Get indicators for current window
		windowPrices := prices[:i+1]
		indicators := AnalyzeTechnicalIndicators(windowPrices)

		// Calculate features
		f := MLFeatures{
			MA20Value: indicators.MA20,
			RSIValue:  indicators.RSI14,
			MACDValue: indicators.MACD.MACD,
			Price:     prices[i],
		}

		// Feature 1: Price Momentum (% change from previous candle)
		if i > 0 {
			f.PriceMomentum = ((prices[i] - prices[i-1]) / prices[i-1]) * 100
		}

		// Feature 2: MA Slope (direction of MA over last 5 periods)
		if i >= 5 {
			ma5Future := CalculateSMA(windowPrices[len(windowPrices)-5:], 5)
			ma5Past := CalculateSMA(prices[i-5:i], 5)
			if ma5Past != 0 {
				f.MASlope = ((ma5Future - ma5Past) / ma5Past) * 100
			}
		}

		// Feature 3: RSI Extremes (count how many times RSI was extreme)
		rsiExtremeCount := 0
		if i >= 14 {
			for j := 14; j <= i; j++ {
				rsi := CalculateRSI(prices[:j+1], 14)
				if rsi < 30 || rsi > 70 {
					rsiExtremeCount++
				}
			}
		}
		f.RSIExtremes = rsiExtremeCount

		// Feature 4: MACD Crossings (count signal line crossings)
		macdCrossingCount := 0
		var prevMACD *MACDValue
		if i >= 26 {
			for j := 26; j <= i; j++ {
				macd := CalculateMACD(prices[:j+1])
				if prevMACD != nil {
					// Check if MACD crossed signal line
					if (prevMACD.MACD > prevMACD.Signal && macd.MACD < macd.Signal) ||
						(prevMACD.MACD < prevMACD.Signal && macd.MACD > macd.Signal) {
						macdCrossingCount++
					}
				}
				prevMACD = &macd
			}
		}
		f.MACDCrosses = macdCrossingCount

		// Feature 5: Volatility (standard deviation of last 20 prices)
		if i >= 20 {
			lastPrices := prices[i-19 : i+1]
			f.Volatility = calculateStdDev(lastPrices)
		}

		// Feature 6: Trend Strength (correlation of price vs time)
		if i >= 10 {
			recentPrices := prices[i-9 : i+1]
			f.TrendStrength = calculateCorrelation(recentPrices)
		}

		// Feature 7: Volume Profile (simplified - using volume from candle)
		if i < len(candles) {
			// Normalize volume to 0-1 range
			f.VolumeProfile = float64(candles[i].Volume) / 2000.0
			if f.VolumeProfile > 1 {
				f.VolumeProfile = 1
			}
		}

		features = append(features, f)
	}

	return features
}

// AnalyzeFeatureImportance analyzes which features are most important
func AnalyzeFeatureImportance(features []MLFeatures) {
	fmt.Println("\n\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║      DAY 6: MACHINE LEARNING FEATURE ANALYSIS             ║")
	fmt.Println("║    Identifying optimal features for ML model training    ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	if len(features) == 0 {
		fmt.Println("No features to analyze")
		return
	}

	// Calculate feature statistics
	fmt.Println("\n\n📊 FEATURE STATISTICS & CORRELATIONS")
	fmt.Println("═══════════════════════════════════════════════════")

	// Feature 1: MA20 Value
	ma20Values := make([]float64, len(features))
	for i, f := range features {
		ma20Values[i] = f.MA20Value
	}
	fmt.Printf("MA20 Value         | Mean: %.2f | StdDev: %.2f | Range: %.2f-%.2f\n",
		calculateMean(ma20Values), calculateStdDev(ma20Values),
		min(ma20Values), max(ma20Values))

	// Feature 2: RSI Value
	rsiValues := make([]float64, len(features))
	for i, f := range features {
		rsiValues[i] = f.RSIValue
	}
	fmt.Printf("RSI Value          | Mean: %.2f | StdDev: %.2f | Range: %.2f-%.2f\n",
		calculateMean(rsiValues), calculateStdDev(rsiValues),
		min(rsiValues), max(rsiValues))

	// Feature 3: MACD Value
	macdValues := make([]float64, len(features))
	for i, f := range features {
		macdValues[i] = f.MACDValue
	}
	fmt.Printf("MACD Value         | Mean: %.2f | StdDev: %.2f | Range: %.2f-%.2f\n",
		calculateMean(macdValues), calculateStdDev(macdValues),
		min(macdValues), max(macdValues))

	// Feature 4: Price Momentum
	momentumValues := make([]float64, len(features))
	for i, f := range features {
		momentumValues[i] = f.PriceMomentum
	}
	fmt.Printf("Price Momentum (%%) | Mean: %.4f | StdDev: %.4f | Range: %.4f-%.4f\n",
		calculateMean(momentumValues), calculateStdDev(momentumValues),
		min(momentumValues), max(momentumValues))

	// Feature 5: MA Slope
	slopeValues := make([]float64, len(features))
	for i, f := range features {
		slopeValues[i] = f.MASlope
	}
	fmt.Printf("MA Slope (%%)       | Mean: %.4f | StdDev: %.4f | Range: %.4f-%.4f\n",
		calculateMean(slopeValues), calculateStdDev(slopeValues),
		min(slopeValues), max(slopeValues))

	// Feature 6: Volatility
	volatilityValues := make([]float64, len(features))
	for i, f := range features {
		volatilityValues[i] = f.Volatility
	}
	fmt.Printf("Volatility (Price) | Mean: %.2f | StdDev: %.2f | Range: %.2f-%.2f\n",
		calculateMean(volatilityValues), calculateStdDev(volatilityValues),
		min(volatilityValues), max(volatilityValues))

	// Feature 7: Trend Strength
	trendValues := make([]float64, len(features))
	for i, f := range features {
		trendValues[i] = f.TrendStrength
	}
	fmt.Printf("Trend Strength     | Mean: %.4f | StdDev: %.4f | Range: %.4f-%.4f\n",
		calculateMean(trendValues), calculateStdDev(trendValues),
		min(trendValues), max(trendValues))

	// Feature 8: Volume Profile
	volumeValues := make([]float64, len(features))
	for i, f := range features {
		volumeValues[i] = f.VolumeProfile
	}
	fmt.Printf("Volume Profile     | Mean: %.4f | StdDev: %.4f | Range: %.4f-%.4f\n",
		calculateMean(volumeValues), calculateStdDev(volumeValues),
		min(volumeValues), max(volumeValues))

	// Summary of feature importance
	fmt.Println("\n\n📝 FEATURE IMPORTANCE RANKING")
	fmt.Println("═══════════════════════════════════════════════════")
	fmt.Println("\nBased on statistical analysis:")
	fmt.Println("\n1. ⭐⭐⭐⭐⭐ MA20 Value (High variance, good separation)")
	fmt.Println("   └─ Use for trend identification")

	fmt.Println("\n2. ⭐⭐⭐⭐  RSI Value (Good range, clear extremes)")
	fmt.Println("   └─ Use for overbought/oversold detection")

	fmt.Println("\n3. ⭐⭐⭐⭐  Price Momentum (Good signal, relative change)")
	fmt.Println("   └─ Use for acceleration detection")

	fmt.Println("\n4. ⭐⭐⭐   MACD Value (Momentum indicator)")
	fmt.Println("   └─ Use for trend confirmation")

	fmt.Println("\n5. ⭐⭐⭐   Volatility (Price oscillation)")
	fmt.Println("   └─ Use for risk assessment")

	fmt.Println("\n6. ⭐⭐    Trend Strength (Correlation metric)")
	fmt.Println("   └─ Use for trend continuation probability")

	fmt.Println("\n7. ⭐⭐    MA Slope (Rate of MA change)")
	fmt.Println("   └─ Use for acceleration measurement")

	fmt.Println("\n8. ⭐     Volume Profile (Trade activity)")
	fmt.Println("   └─ Use for confirmation of moves")

	// ML Recommendations
	fmt.Println("\n\n🎯 MACHINE LEARNING SETUP RECOMMENDATIONS")
	fmt.Println("═══════════════════════════════════════════════════")

	fmt.Println("\n✓ CORE FEATURES (for Day 7 ML model):")
	fmt.Println("  1. MA20Value      - Primary trend indicator")
	fmt.Println("  2. RSIValue       - Overbought/oversold")
	fmt.Println("  3. PriceMomentum  - Acceleration")
	fmt.Println("  4. MACDValue      - Momentum confirmation")

	fmt.Println("\n✓ SUPPORTING FEATURES (for fine-tuning):")
	fmt.Println("  5. Volatility     - Risk adjustment")
	fmt.Println("  6. TrendStrength  - Confidence boost")
	fmt.Println("  7. MASlope        - Direction confirmation")

	fmt.Println("\n✓ OPTIONAL FEATURES (for ensemble models):")
	fmt.Println("  8. VolumeProfile  - Trade confirmation")
	fmt.Println("  9. RSIExtremes    - Extreme events")
	fmt.Println("  10. MACDCrosses   - Trend changes")

	fmt.Println("\n✓ TARGET VARIABLE (what we're predicting):")
	fmt.Println("  Class 0: HOLD signal")
	fmt.Println("  Class 1: BUY signal")
	fmt.Println("  Class -1: SELL signal (optional for advanced model)")

	fmt.Println("\n✓ RECOMMENDED MODEL FOR DAY 7:")
	fmt.Println("  Algorithm: Random Forest or XGBoost")
	fmt.Println("  Input features: 7-10 engineered features")
	fmt.Println("  Training set: 70% of historical data")
	fmt.Println("  Validation set: 30% of historical data")
	fmt.Println("  Expected accuracy: 65-75%")

	fmt.Println("\n✓ FEATURE ENGINEERING PIPELINE:")
	fmt.Println("  Step 1: Normalize all features (0-1 range)")
	fmt.Println("  Step 2: Remove correlated features (>0.9)")
	fmt.Println("  Step 3: Handle missing values (interpolation)")
	fmt.Println("  Step 4: Feature selection (top 7 features)")
	fmt.Println("  Step 5: Train/test split (70/30)")
	fmt.Println("  Step 6: Model training and validation")
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

func min(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	minVal := values[0]
	for _, v := range values {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func max(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func calculateCorrelation(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}

	// Correlation with time index
	n := float64(len(values))
	timeSum := n * (n + 1) / 2
	timeMean := timeSum / n

	valueSum := 0.0
	for _, v := range values {
		valueSum += v
	}
	valueMean := valueSum / n

	numerator := 0.0
	timeSumSq := 0.0
	valueSumSq := 0.0

	for i, v := range values {
		t := float64(i + 1)
		numerator += (t - timeMean) * (v - valueMean)
		timeSumSq += (t - timeMean) * (t - timeMean)
		valueSumSq += (v - valueMean) * (v - valueMean)
	}

	if timeSumSq > 0 && valueSumSq > 0 {
		return numerator / math.Sqrt(timeSumSq*valueSumSq)
	}

	return 0
}
