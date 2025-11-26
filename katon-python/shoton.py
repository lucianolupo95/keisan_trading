#!/usr/bin/env python3
"""
Shoton (Fuego Avanzado) - Análisis Técnico e Indicadores
Módulo de análisis técnico con SMA, Bollinger Bands, y RSI
"""

import json
from typing import List, Dict, Union

def calculate_sma(prices: List[float], period: int = 5) -> float:
    """
    Calcula Simple Moving Average (promedio móvil simple)

    Args:
        prices: Lista de precios (cierre)
        period: Período del promedio (default 5)

    Returns:
        float: SMA del período
    """
    if len(prices) < period:
        return sum(prices) / len(prices) if prices else 0
    return sum(prices[-period:]) / period


def calculate_std(prices: List[float], period: int = 5) -> float:
    """
    Calcula desviación estándar (volatilidad)

    Args:
        prices: Lista de precios
        period: Período (default 5)

    Returns:
        float: Desviación estándar
    """
    if len(prices) < period:
        subset = prices
    else:
        subset = prices[-period:]

    if len(subset) < 2:
        return 0

    mean = sum(subset) / len(subset)
    variance = sum((x - mean) ** 2 for x in subset) / len(subset)
    return variance ** 0.5


def calculate_bollinger_bands(prices: List[float], period: int = 5) -> Dict[str, float]:
    """
    Calcula Bollinger Bands (bandas de volatilidad)

    Args:
        prices: Lista de precios
        period: Período (default 5)

    Returns:
        dict: {'upper': valor, 'middle': SMA, 'lower': valor}
    """
    sma = calculate_sma(prices, period)
    std = calculate_std(prices, period)

    return {
        'upper': sma + 2 * std,
        'middle': sma,
        'lower': sma - 2 * std
    }


def calculate_rsi(prices: List[float], period: int = 14) -> float:
    """
    Calcula Relative Strength Index (momentum 0-100)

    Args:
        prices: Lista de precios
        period: Período (default 14)

    Returns:
        float: RSI value (0-100)
    """
    if len(prices) < period + 1:
        # Si no hay suficientes datos, usar lo que tenemos
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


def analyze_candle(prices: List[float], volumes: List[float] = None) -> Dict:
    """
    Análisis completo de una vela/período

    Args:
        prices: Lista de precios de cierre
        volumes: Lista de volúmenes (opcional)

    Returns:
        dict: Análisis técnico completo
    """
    if not prices:
        return {'error': 'No prices provided'}

    current_price = prices[-1]
    sma_5 = calculate_sma(prices, 5)
    sma_20 = calculate_sma(prices, 20) if len(prices) >= 20 else sma_5
    bb = calculate_bollinger_bands(prices, 5)
    rsi = calculate_rsi(prices, 14)

    # Determinar posición relativa a bandas de Bollinger
    if current_price > bb['upper']:
        bb_position = 'overbought'
    elif current_price < bb['lower']:
        bb_position = 'oversold'
    else:
        bb_position = 'normal'

    # Determinar tendencia
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

    if volumes:
        analysis['current_volume'] = volumes[-1] if volumes else 0

    return analysis


def generate_signal(analysis: Dict) -> str:
    """
    Genera señal de trading basada en análisis técnico

    Args:
        analysis: Dict retornado por analyze_candle()

    Returns:
        str: 'BUY', 'SELL', o 'HOLD'
    """
    signal = 'HOLD'
    score = 0

    # RSI signals
    if analysis['rsi'] < 30:
        score += 2  # Strong BUY
    elif analysis['rsi'] < 40:
        score += 1  # Weak BUY
    elif analysis['rsi'] > 70:
        score -= 2  # Strong SELL
    elif analysis['rsi'] > 60:
        score -= 1  # Weak SELL

    # Bollinger Bands signals
    if analysis['bb_position'] == 'oversold':
        score += 2  # BUY
    elif analysis['bb_position'] == 'overbought':
        score -= 2  # SELL

    # Trend signals
    if analysis['trend'] == 'uptrend':
        score += 1
    elif analysis['trend'] == 'downtrend':
        score -= 1

    # Determine final signal
    if score >= 2:
        signal = 'BUY'
    elif score <= -2:
        signal = 'SELL'

    return signal


def ping():
    """Validación de módulo"""
    return "Shoton OK"


if __name__ == "__main__":
    print("Shoton activo")
    print(f"Ping response: {ping()}\n")

    # Test data
    test_prices = [100, 101, 102, 101, 103, 104, 102, 105, 106, 104]

    print("=== INDICADORES TÉCNICOS ===\n")

    # SMA
    sma = calculate_sma(test_prices, 5)
    print(f"SMA(5) de {test_prices} = {sma:.2f}")

    # Bollinger Bands
    bb = calculate_bollinger_bands(test_prices, 5)
    print(f"Bollinger Bands: upper={bb['upper']:.2f}, middle={bb['middle']:.2f}, lower={bb['lower']:.2f}")

    # RSI
    rsi = calculate_rsi(test_prices, 14)
    print(f"RSI(14) = {rsi}\n")

    # Análisis completo
    print("=== ANÁLISIS COMPLETO ===\n")
    analysis = analyze_candle(test_prices)
    print(json.dumps(analysis, indent=2))

    # Signal
    signal = generate_signal(analysis)
    print(f"\nSeñal generada: {signal}")
