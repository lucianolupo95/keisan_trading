package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// KatonAnalysis contiene el análisis técnico de Katon
type KatonAnalysis struct {
	CurrentPrice float64            `json:"current_price"`
	SMA5         float64            `json:"sma_5"`
	SMA20        float64            `json:"sma_20"`
	BB           map[string]float64 `json:"bollinger_bands"`
	BBPosition   string             `json:"bb_position"`
	RSI          float64            `json:"rsi"`
	Trend        string             `json:"trend"`
}

// CallKaton ejecuta análisis técnico desde Python
func CallKaton(prices []float64) (KatonAnalysis, error) {
	// Crear script Python inline para evitar dependencias
	pythonScript := `
import json
import sys

def calculate_sma(prices, period=5):
    if len(prices) < period:
        return sum(prices) / len(prices) if prices else 0
    return sum(prices[-period:]) / period

def calculate_std(prices, period=5):
    if len(prices) < period:
        subset = prices
    else:
        subset = prices[-period:]
    if len(subset) < 2:
        return 0
    mean = sum(subset) / len(subset)
    variance = sum((x - mean) ** 2 for x in subset) / len(subset)
    return variance ** 0.5

def calculate_bollinger_bands(prices, period=5):
    sma = calculate_sma(prices, period)
    std = calculate_std(prices, period)
    return {
        'upper': sma + 2 * std,
        'middle': sma,
        'lower': sma - 2 * std
    }

def calculate_rsi(prices, period=14):
    if len(prices) < period + 1:
        if len(prices) < 2:
            return 50
        period = len(prices) - 1
    gains = 0
    losses = 0
    for i in range(-period, 0):
        change = prices[i] - prices[i - 1]
        if change > 0:
            gains += change
        else:
            losses += abs(change)
    avg_gain = gains / period
    avg_loss = losses / period
    if avg_loss == 0:
        return 100 if avg_gain > 0 else 50
    rs = avg_gain / avg_loss
    rsi = 100 - (100 / (1 + rs))
    return round(rsi, 2)

prices = json.loads(sys.argv[1])
current_price = prices[-1]
sma_5 = calculate_sma(prices, 5)
sma_20 = calculate_sma(prices, 20) if len(prices) >= 20 else sma_5
bb = calculate_bollinger_bands(prices, 5)
rsi = calculate_rsi(prices, 14)

if current_price > bb['upper']:
    bb_position = 'overbought'
elif current_price < bb['lower']:
    bb_position = 'oversold'
else:
    bb_position = 'normal'

if sma_5 > sma_20:
    trend = 'uptrend'
elif sma_5 < sma_20:
    trend = 'downtrend'
else:
    trend = 'neutral'

analysis = {
    'current_price': round(current_price, 2),
    'sma_5': round(sma_5, 2),
    'sma_20': round(sma_20, 2),
    'bollinger_bands': {
        'upper': round(bb['upper'], 2),
        'middle': round(bb['middle'], 2),
        'lower': round(bb['lower'], 2)
    },
    'bb_position': bb_position,
    'rsi': rsi,
    'trend': trend
}

print(json.dumps(analysis))
`

	// Convertir precios a JSON
	pricesJSON, _ := json.Marshal(prices)

	// Ejecutar Python
	cmd := exec.Command("python", "-c", pythonScript, string(pricesJSON))
	output, err := cmd.Output()
	if err != nil {
		// Si Python falla, retornar análisis vacío
		return KatonAnalysis{}, fmt.Errorf("failed to call Katon: %v", err)
	}

	// Parsear resultado
	var analysis KatonAnalysis
	err = json.Unmarshal(output, &analysis)
	if err != nil {
		return KatonAnalysis{}, fmt.Errorf("failed to parse Katon response: %v", err)
	}

	return analysis, nil
}

// EnhancedSignal genera una señal mejorada usando análisis técnico
func EnhancedSignal(candle Candle, prices []float64) string {
	// Primero, aplicar filtros básicos de Fuuton
	basicSignal := GenerateSignal(candle)

	// Obtener análisis técnico de Katon
	analysis, err := CallKaton(prices)
	if err != nil {
		// Si Katon falla, retornar signal básica
		return basicSignal
	}

	// Mejorar signal usando análisis técnico
	score := 0

	// RSI signals
	if analysis.RSI < 30 {
		score += 2 // Strong BUY
	} else if analysis.RSI < 40 {
		score += 1 // Weak BUY
	} else if analysis.RSI > 70 {
		score -= 2 // Strong SELL
	} else if analysis.RSI > 60 {
		score -= 1 // Weak SELL
	}

	// Bollinger Bands signals
	if analysis.BBPosition == "oversold" {
		score += 2 // BUY
	} else if analysis.BBPosition == "overbought" {
		score -= 2 // SELL
	}

	// Trend signals
	if analysis.Trend == "uptrend" {
		score += 1
	} else if analysis.Trend == "downtrend" {
		score -= 1
	}

	// Determine final signal
	if score >= 2 {
		return "BUY"
	} else if score <= -2 {
		return "SELL"
	}

	return basicSignal
}

// PrintAnalysis imprime el análisis de forma legible
func (a KatonAnalysis) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Current Price: $%.2f\n", a.CurrentPrice))
	sb.WriteString(fmt.Sprintf("SMA(5): %.2f | SMA(20): %.2f\n", a.SMA5, a.SMA20))
	sb.WriteString(fmt.Sprintf("Bollinger Bands: [%.2f, %.2f, %.2f]\n", a.BB["lower"], a.BB["middle"], a.BB["upper"]))
	sb.WriteString(fmt.Sprintf("RSI(14): %.2f | Trend: %s | BB Position: %s\n", a.RSI, a.Trend, a.BBPosition))
	return sb.String()
}
