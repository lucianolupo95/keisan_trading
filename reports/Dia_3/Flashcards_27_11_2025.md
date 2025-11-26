# 🎓 FLASHCARDS ANKI - DÍA 3
## Keisan Trading - Indicadores Técnicos, Backtesting y Orquestación

---

## FLASHCARD 1 - SMA (Simple Moving Average)

**Pregunta:**
¿Qué es SMA y cómo se calcula en Shoton?

**Código:**
```python
def calculate_sma(prices, period=5):
    return sum(prices[-period:]) / period
```

**Respuesta:**
SMA es el promedio aritmético de los últimos N precios. En Shoton con período=5, suma los últimos 5 precios y divide entre 5. Suaviza el ruido del mercado. Ejemplo: SMA([100,101,102,101,103]) = 101.4

**Concepto clave**: SMA es un indicador lagging (retrasado), sigue el precio, no lo anticipa.

**Anki ID**: SMA_001

---

## FLASHCARD 2 - Bollinger Bands: Upper y Lower

**Pregunta:**
¿Cómo calcula Shoton las Bollinger Bands? ¿Qué indican los extremos?

**Código:**
```python
def calculate_bollinger_bands(prices, period=5):
    sma = calculate_sma(prices, period)
    std = calculate_std(prices, period)
    return {
        'upper': sma + 2 * std,
        'middle': sma,
        'lower': sma - 2 * std
    }
```

**Respuesta:**
Bollinger Bands = SMA ± 2 desviaciones estándar.
- **Upper**: SMA + 2*std → Zona de OVERBOUGHT (sobrecompra)
- **Middle**: SMA → Línea central
- **Lower**: SMA - 2*std → Zona de OVERSOLD (sobreventa)

Si precio > upper → mercado sobrecomprado → VENDER
Si precio < lower → mercado sobrevendido → COMPRAR

**Anki ID**: BOLLINGER_001

---

## FLASHCARD 3 - RSI (Relative Strength Index)

**Pregunta:**
¿Qué mide RSI? ¿Cuáles son los umbrales de overbought/oversold?

**Código:**
```python
def calculate_rsi(prices, period=14):
    gains = sum(max(prices[i] - prices[i-1], 0) for i in range(-period, 0))
    losses = sum(abs(min(prices[i] - prices[i-1], 0)) for i in range(-period, 0))
    rs = gains / losses
    rsi = 100 - (100 / (1 + rs))
    return rsi  # Valor entre 0-100
```

**Respuesta:**
RSI mide el MOMENTUM (velocidad de cambio de precio).
- RSI = 0-30: OVERSOLD (débil, posible rebote al alza)
- RSI = 30-70: Neutral
- RSI = 70-100: OVERBOUGHT (fuerte, posible corrección a la baja)

Fórmula: RSI = 100 - (100/(1+RS)) donde RS = ganancias_promedio/pérdidas_promedio

**Señal de trading**:
- RSI < 30 → BUY signal
- RSI > 70 → SELL signal

**Anki ID**: RSI_001

---

## FLASHCARD 4 - Backtester: Cálculo de P&L

**Pregunta:**
¿Cómo calcula el Backtester la ganancia/pérdida de un trade?

**Código:**
```go
func (b *Backtester) RunBacktest() BacktestResult {
    // BUY signal
    openTrade.EntryPrice = candle.Close

    // SELL signal
    openTrade.ExitPrice = candle.Close
    openTrade.ProfitLoss = (ExitPrice - EntryPrice) * (b.TradingAmount / EntryPrice)
    openTrade.ProfitLossPC = ((ExitPrice - EntryPrice) / EntryPrice) * 100

    b.CurrentCapital += openTrade.ProfitLoss
}
```

**Respuesta:**
P&L = (ExitPrice - EntryPrice) × (cantidad / EntryPrice)

Ejemplo:
- EntryPrice = $1502, TradingAmount = $1000
- ExitPrice = $1510
- P&L = ($1510 - $1502) × (1000/1502) = $8 × 0.666 = $5.33
- P&L% = ((1510-1502)/1502) × 100 = 0.53%

El capital se actualiza: CurrentCapital += P&L

**Anki ID**: BACKTEST_001

---

## FLASHCARD 5 - Win Rate vs Success Rate

**Pregunta:**
¿Cuál es la diferencia entre Win Rate y Success Rate en el reporte del backtester?

**Respuesta:**
- **Win Rate**: (Winning Trades / Total Trades) × 100%
  - Ejemplo: 3 trades ganadores de 5 totales = 60% win rate

- **Success Rate** (en Orchestrator): Similar a win rate pero calculado desde resultados del backtester

- **Diferencia**: Win rate es métrica individual, success rate es métrica agregada del sistema

**Interpretación**:
- 100% win rate = estrategia perfecta (cuidado con overfitting)
- 50% win rate = coin flip
- < 50% win rate = estrategia pierde dinero
- > 60% win rate = estrategia rentable (con risk management)

**Anki ID**: WINRATE_001

---

## FLASHCARD 6 - Max Drawdown

**Pregunta:**
¿Qué es Max Drawdown y por qué es importante en trading?

**Código:**
```go
func (b *Backtester) calculateMaxDrawdown() float64 {
    peakCapital := b.InitialCapital
    maxDrawdown := 0.0

    for _, trade := range b.Trades {
        currentCap += trade.ProfitLoss
        drawdown := ((peakCapital - currentCap) / peakCapital) * 100
        if drawdown > maxDrawdown {
            maxDrawdown = drawdown
        }
    }
    return maxDrawdown
}
```

**Respuesta:**
Max Drawdown = Peor pérdida desde el máximo capital alcanzado.

Ejemplo:
- Capital inicial: $10,000
- Peak capital: $12,000
- Worst point: $9,500
- Drawdown: ($12,000 - $9,500) / $12,000 = 20.83%

**Importancia**:
- Mide riesgo de la estrategia
- Si drawdown > 50%, estrategia muy riesgosa
- Risk management minimiza drawdown

**Anki ID**: DRAWDOWN_001

---

## FLASHCARD 7 - Sharpe Ratio

**Pregunta:**
¿Qué mide el Sharpe Ratio y cuál es su fórmula?

**Código:**
```go
func (b *Backtester) CalculateSharpeRatio(riskFreeRate float64) float64 {
    avgReturn := promedio(returns)
    stdDev := desviacion_estandar(returns)
    sharpe := (avgReturn - riskFreeRate) / stdDev
    return sharpe
}
```

**Respuesta:**
Sharpe Ratio = (Retorno promedio - Tasa libre de riesgo) / Desviación estándar

Mide RETORNO AJUSTADO POR RIESGO.

**Interpretación**:
- Sharpe < 1: Pobre rendimiento ajustado por riesgo
- Sharpe 1-2: Aceptable
- Sharpe > 2: Excelente (muy buen retorno con poco riesgo)
- Sharpe < 0: Estrategia pierde dinero

**Ejemplo**:
- Retorno promedio: +10%
- Volatilidad (std): +8%
- Sharpe = (10% - 0%) / 8% = 1.25 (aceptable)

**Anki ID**: SHARPE_001

---

## FLASHCARD 8 - Katon Bridge: IPC entre Go y Python

**Pregunta:**
¿Cómo comunica Fuuton (Go) con Katon (Python)? ¿Qué formato usa?

**Código:**
```go
func CallKaton(prices []float64) (KatonAnalysis, error) {
    pythonScript := `...código Python...`
    pricesJSON, _ := json.Marshal(prices)
    cmd := exec.Command("python", "-c", pythonScript, string(pricesJSON))
    output, _ := cmd.Output()

    var analysis KatonAnalysis
    json.Unmarshal(output, &analysis)
    return analysis, nil
}
```

**Respuesta:**
Usa **JSON sobre IPC (Inter-Process Communication)**:

1. Go serializa precios a JSON: `[1500.0, 1501.5, ...]`
2. Ejecuta Python con script inline: `python -c "script" "[json]"`
3. Python calcula indicadores y retorna JSON
4. Go parsea resultado en struct Go

**Ventajas**:
- ✅ Desacoplado (Go ≠ Python)
- ✅ Extensible (otros lenguajes)
- ✅ Fácil debuguear

**Desventajas**:
- ⚠️ Overhead de IPC
- ⚠️ Requiere Python instalado

**Anki ID**: KATON_BRIDGE_001

---

## FLASHCARD 9 - Signal Enhancement: Score System

**Pregunta:**
¿Cómo combina el Katon Bridge los indicadores técnicos para mejorar la señal?

**Código:**
```go
func EnhancedSignal(candle Candle, prices []float64) string {
    signal := GenerateSignal(candle)  // BUY/HOLD/SELL básica
    analysis, _ := CallKaton(prices)

    score := 0
    if analysis.RSI < 30 { score += 2 }        // Strong BUY
    if analysis.BBPosition == "oversold" { score += 2 }  // BUY
    if analysis.Trend == "uptrend" { score += 1 }        // Weak BUY

    if score >= 2 { return "BUY" }
    if score <= -2 { return "SELL" }
    return "HOLD"
}
```

**Respuesta:**
Sistema de puntuación (score):
- **RSI < 30**: +2 (Strong BUY)
- **RSI < 40**: +1 (Weak BUY)
- **RSI > 70**: -2 (Strong SELL)
- **RSI > 60**: -1 (Weak SELL)
- **BB Oversold**: +2 (BUY)
- **BB Overbought**: -2 (SELL)
- **Uptrend**: +1 (BUY bias)
- **Downtrend**: -1 (SELL bias)

Final:
- score ≥ 2 → BUY
- score ≤ -2 → SELL
- else → HOLD

**Ventaja**: Múltiples confirmaciones antes de actuar

**Anki ID**: ENHANCED_SIGNAL_001

---

## FLASHCARD 10 - Orchestrator: Flujo Completo

**Pregunta:**
¿Cuál es el flujo completo que ejecuta el Orchestrator?

**Respuesta:**
```
1. Lee todas las velas del CSV
2. Para cada vela:
   - Genera signal (básica o mejorada)
   - Recolecta precio para análisis

3. Calcula métricas de mercado:
   - Precio promedio
   - Rango de precios
   - Volatilidad %
   - Tendencia (uptrend/downtrend/sideways)

4. Reporta distribución de signals:
   - BUY count
   - SELL count
   - HOLD count

5. Ejecuta backtester:
   - Simula trades según signals
   - Calcula P&L
   - Retorna resultado

6. Consolida reporte final:
   - Market analysis
   - Signal distribution
   - Backtest results
   - Trade details
```

**Resultado**: Un solo reporte unificado con toda la información del sistema

**Anki ID**: ORCHESTRATOR_FLOW_001

---

## FLASHCARD 11 - Volatility Calculation

**Pregunta:**
¿Cómo calcula el Orchestrator la volatilidad del mercado?

**Código:**
```go
func (o *Orchestrator) calculateVolatility(prices []float64) float64 {
    avg := o.calculateAvgPrice(prices)
    variance := 0.0
    for _, p := range prices {
        variance += (p - avg) * (p - avg)
    }
    variance /= float64(len(prices))
    stdDev := math.Sqrt(variance)
    volatilityPercent := (stdDev / avg) * 100
    return volatilityPercent
}
```

**Respuesta:**
Volatilidad = (Desviación estándar / Precio promedio) × 100%

Pasos:
1. Calcular precio promedio (mean)
2. Calcular varianza: suma((precio - mean)²) / n
3. Calcular std dev: √varianza
4. Volatilidad % = (std dev / mean) × 100%

**Interpretación**:
- 0.1-0.5%: Bajo (mercado calmo)
- 0.5-2%: Moderado
- 2-5%: Alto (mercado volátil)
- > 5%: Muy alto (riesgo extremo)

En Día 3: Volatilidad = 0.34% (mercado muy calmo)

**Anki ID**: VOLATILITY_001

---

## FLASHCARD 12 - Trend Detection: 3 Períodos

**Pregunta:**
¿Cómo detecta el Orchestrator la tendencia del mercado?

**Código:**
```go
func (o *Orchestrator) calculateTrend(prices []float64) string {
    thirtyPercent := len(prices) / 3
    sixtyPercent := (len(prices) * 2) / 3

    period1Avg := avgPrice(prices[:thirtyPercent])
    period2Avg := avgPrice(prices[thirtyPercent:sixtyPercent])
    period3Avg := avgPrice(prices[sixtyPercent:])

    if period1Avg < period2Avg && period2Avg < period3Avg {
        return "uptrend"
    }
    if period1Avg > period2Avg && period2Avg > period3Avg {
        return "downtrend"
    }
    return "sideways"
}
```

**Respuesta:**
Divide las velas en 3 períodos iguales:
- **Período 1** (0-33%): primeras velas
- **Período 2** (33-66%): velas del medio
- **Período 3** (66-100%): últimas velas

Calcula precio promedio de cada período:
- Si P1 < P2 < P3 → **UPTREND** (sube)
- Si P1 > P2 > P3 → **DOWNTREND** (baja)
- Si no sigue patrón → **SIDEWAYS** (lateral)

**Ejemplo** (Día 3):
- P1 avg = $1505
- P2 avg = $1510
- P3 avg = $1515
- Resultado: UPTREND ✅

**Anki ID**: TREND_DETECTION_001

---

## FLASHCARD 13 - Python sin dependencias externas

**Pregunta:**
¿Por qué Shoton (Python) no usa numpy ni pandas? ¿Cuáles son las ventajas y desventajas?

**Respuesta:**
**Por qué no usar dependencias**:
- Shoton usa math puro (suma, media, varianza)
- Evita instalación de numpy/pandas
- Reduce complejidad

**Ventajas**:
- ✅ Ninguna dependencia externa
- ✅ Fácil de debuguear (código simple)
- ✅ Rápido de inicializar
- ✅ Portable (funciona en cualquier lado)

**Desventajas**:
- ⚠️ Más lento (no vectorizado)
- ⚠️ Sin caché (recalcula cada vez)
- ⚠️ Más código para escribir (sum(x)/len(x) vs np.mean(x))

**Escalabilidad**:
- Día 3: 35 velas → OK
- Día 10+: 10,000+ velas → considerar numpy

**Anki ID**: PYTHON_DEPS_001

---

## FLASHCARD 14 - Go Backtester vs Python Backtrader

**Pregunta:**
¿Por qué implementamos backtester en Go (no Python) si este es el lenguaje de trading?

**Respuesta:**
**Decisión de arquitectura**:

Go (Fuuton) es el orquestador central:
- ✅ Lee CSV (fast)
- ✅ Genera signals (fast)
- ✅ Ejecuta backtester (coordina todo)
- ✅ Reporta resultados (consolidado)

Python (Katon) es especialista:
- Indicadores técnicos
- Feature engineering
- Machine Learning (Día 5+)

**Por qué Go para backtester**:
1. Fuuton ya tenía el CSV y signals
2. Evita duplicar lógica
3. Más rápido (Go es compiled)
4. Integración natural

**Comparación**:
- Go: Fast, compiled, tight control
- Python: Flexible, data science friendly

En arquitectura multi-lenguaje: cada lenguaje hace lo que mejor sabe.

**Anki ID**: ARCH_DECISION_001

---

## FLASHCARD 15 - From Backtesting to Live Trading

**Pregunta:**
¿Cuál es el siguiente paso después de validar una estrategia con backtesting?

**Respuesta:**
Flujo típico en trading algorítmico:

1. **Backtesting** ← Día 3 (validación en datos históricos)
2. **Paper Trading** ← Día 4-5 (simulación en tiempo real SIN dinero)
3. **Live Trading (SIM)** ← Día 8-9 (dinero simulado, broker real)
4. **Live Trading (REAL)** ← Día 10+ (dinero real, riesgo real)

**Validación en cada paso**:
- Backtest: ¿Funciona el algoritmo?
- Paper: ¿Funciona en tiempo real?
- Live SIM: ¿Maneja el broker real?
- Live REAL: ¿Genera dinero?

**Hoy** (Día 3): ✅ Backtesting OK (+0.11% en 35 velas)
**Mañana** (Día 4): Mejoras con R + más datos
**Día 8-9**: Paper trading + C# bridge
**Día 10+**: Live en NinjaTrader

**Anki ID**: TRADING_PIPELINE_001

---

## 📊 RESUMEN DE FLASHCARDS

| Tema | Count | Coverage |
|------|-------|----------|
| **Indicadores técnicos** | 3 | SMA, Bollinger, RSI |
| **Backtesting** | 3 | P&L, Win Rate, Drawdown |
| **Métricas** | 2 | Sharpe Ratio, Volatility |
| **Arquitectura** | 4 | Katon Bridge, Orchestrator, Trend, Signal |
| **Decisiones técnicas** | 3 | Python deps, Go choice, Pipeline |
| **Total** | 15 | Cobertura completa Día 3 |

---

**Generado por**: Narutrader
**Fecha**: 27/11/2025
**Para**: Anki - Estudio y retención del conocimiento técnico
**Validez**: Código actual del proyecto (referencia Día 3)
