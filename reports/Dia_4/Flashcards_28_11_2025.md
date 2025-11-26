# 🎓 FLASHCARDS DÍA 4 - 28/11/2025

**Total de flashcards**: 20
**Temas**: R, Estadística, Integration Go-R, Trading, Suiton
**Dificultad**: Intermedia

---

## ESTADÍSTICA Y DISTRIBUCIONES

### 📌 FC-01: ¿Qué es la media (mean) en una distribución?

**Pregunta**: ¿Cuál es la definición de media?

**Respuesta**:
La **media** es el promedio aritmético de todos los valores. Se calcula sumando todos los valores y dividiéndolos por la cantidad total.

Fórmula: `mean = Σ(x) / n`

**Ejemplo**:
- Precios: [1500, 1510, 1520]
- Mean = (1500+1510+1520)/3 = 1510

**En trading**: Precio promedio del período analizado

---

### 📌 FC-02: ¿Qué es la desviación estándar (std)?

**Pregunta**: ¿Qué mide la desviación estándar?

**Respuesta**:
La **desviación estándar** mide cuánto se desvían los datos respecto a la media. Valores altos indican mayor dispersión; valores bajos indican datos más concentrados.

Fórmula: `std = √(Σ(x-mean)² / n)`

**Interpretación**:
- std pequeño: datos concentrados (mercado estable)
- std grande: datos dispersos (mercado volátil)

**En trading**: Medida de volatilidad

---

### 📌 FC-03: ¿Qué es el skewness (asimetría)?

**Pregunta**: ¿Cuál es el significado del skewness?

**Respuesta**:
El **skewness** mide la asimetría de una distribución:
- **Skewness = 0**: Distribución simétrica (normal)
- **Skewness > 0**: Asimetría derecha (cola derecha larga)
- **Skewness < 0**: Asimetría izquierda (cola izquierda larga)

**Interpretación**:
|Skewness| < 0.5: Aprox. simétrica
|Skewness| 0.5-1.0: Moderadamente sesgada
|Skewness| > 1.0: Altamente sesgada

**En trading**: Indica posible sesgo en movimientos de precio

---

### 📌 FC-04: ¿Qué es el kurtosis?

**Pregunta**: ¿Qué mide el kurtosis?

**Respuesta**:
El **kurtosis** mide el "peso" de las colas de una distribución:
- **Kurtosis = 3** (normal): Colas normales
- **Kurtosis > 3**: Colas pesadas (más extremos)
- **Kurtosis < 3**: Colas ligeras (menos extremos)

**Implicación**:
- Alto kurtosis: Mayor riesgo de eventos extremos
- Bajo kurtosis: Movimientos más predecibles

**En trading**: Riesgo de saltos grandes de precio

---

### 📌 FC-05: ¿Qué es el coefficient of variation (CV)?

**Pregunta**: ¿Cuál es la utilidad del CV?

**Respuesta**:
El **Coefficient of Variation** normaliza la volatilidad respecto al precio promedio:

Fórmula: `CV = (std / mean) × 100%`

**Ventaja**: Permite comparar volatilidad entre instrumentos con diferentes precios.

**Interpretación**:
- CV < 1%: Muy estable
- CV 1-5%: Estable
- CV > 10%: Muy volátil

**En trading**: Mejor que std raw para comparaciones

---

## TESTING Y NORMALIDAD

### 📌 FC-06: ¿Qué es el test de Shapiro-Wilk?

**Pregunta**: ¿Para qué sirve el Shapiro-Wilk test?

**Respuesta**:
El **Shapiro-Wilk test** verifica si una distribución es normal (gaussiana).

**Hipótesis**:
- H0: La distribución ES normal
- H1: La distribución NO es normal

**Interpretación del p-value**:
- p-value > 0.05: Aceptamos H0 (ES normal)
- p-value < 0.05: Rechazamos H0 (NO es normal)

**En trading**: Importante para asumir distribuciones normales en modelos

---

### 📌 FC-07: ¿Qué es el p-value?

**Pregunta**: ¿Cómo se interpreta el p-value?

**Respuesta**:
El **p-value** es la probabilidad de observar datos tan extremos (o más) bajo la hipótesis nula.

**Interpretación**:
- p > 0.05 (5%): Resultado NO significativo
- p < 0.05: Resultado significativo
- p < 0.01: Muy significativo
- p < 0.001: Altamente significativo

**En estadística trading**:
- p > 0.05 para normalidad: Podemos asumir distribución normal
- p < 0.05 para normalidad: Debemos considerar distribuciones alternas

---

### 📌 FC-08: ¿Qué es la mediana en una distribución?

**Pregunta**: ¿Cómo se calcula la mediana?

**Respuesta**:
La **mediana** es el valor que divide la distribución en 50% arriba y 50% abajo.

**Cálculo**:
- Si n es impar: Elemento en posición (n+1)/2
- Si n es par: Promedio de elementos n/2 y (n/2)+1

**Diferencia con media**:
- Media: Sensible a outliers
- Mediana: Robusta a outliers

**En trading**: Mediana suele ser mejor que media cuando hay spikes

---

## CORRELACIÓN

### 📌 FC-09: ¿Qué es la correlación de Pearson?

**Pregunta**: ¿Cuál es la definición de correlación Pearson?

**Respuesta**:
La **correlación de Pearson** mide la relación lineal entre dos variables.

Valores: -1.0 a +1.0
- **r = +1**: Correlación perfecta positiva
- **r = 0**: Sin correlación
- **r = -1**: Correlación perfecta negativa

**Interpretación de magnitud**:
- |r| > 0.7: Correlación fuerte
- |r| 0.3-0.7: Correlación moderada
- |r| < 0.3: Correlación débil

**En trading**: Precio vs Tiempo muestra trend

---

### 📌 FC-10: ¿Cómo se interpreta correlación positiva?

**Pregunta**: ¿Qué significa correlación positiva en precio-tiempo?

**Respuesta**:
**Correlación positiva** entre precio y tiempo significa:
- Conforme avanza el tiempo, el precio tiende a subir
- **Trend: UPTREND**
- Mayor confianza en BUY signals

**Ejemplo**:
- r = +0.8: Strong uptrend, excelente para entradas BUY
- r = +0.3: Weak uptrend, cautela necesaria
- r = 0.0: Sideways, sin dirección clara

**En la estrategia Day 4**: Factor +0.2 en confidence score

---

## R PROGRAMMING (SUITON)

### 📌 FC-11: ¿Cuál es la estructura de suiton.R?

**Pregunta**: ¿Qué funciones principales tiene Suiton?

**Respuesta**:
```r
# 1. calculate_correlation(prices)
#    → Retorna: correlation, p_value, interpretation

# 2. analyze_distribution(prices)
#    → Retorna: mean, std, median, skewness, kurtosis, CV

# 3. test_normality(prices)
#    → Retorna: p_value, is_normal, interpretation

# Main execution:
main() → Crea datos de ejemplo
        → Ejecuta 3 análisis
        → Retorna JSON-ready list
```

**Librerías usadas**:
- `e1071`: Para skewness/kurtosis

---

### 📌 FC-12: ¿Cuál es la diferencia entre Shapiro test y check local?

**Pregunta**: ¿Qué hace Shapiro test vs heurística local en Go?

**Respuesta**:
**En R (Suiton)**:
- Shapiro-Wilk test real
- P-value preciso
- Mejor precisión

**En Go (fallback local)**:
- Heurística simple: |skewness| < 0.5 → normal
- P-value dummy (0.5)
- Rápido, sin dependencias

**Cuándo usar cada uno**:
- R disponible: Usar R (mejor precisión)
- R no disponible: Usar fallback (sigue funcionando)

**Conclusión**: Go fallback = graceful degradation ✅

---

## GO INTEGRATION (BRIDGES)

### 📌 FC-13: ¿Cuál es la arquitectura del Suiton Bridge?

**Pregunta**: ¿Cómo está estructurado suiton_bridge.go?

**Respuesta**:
```go
// 1. TIPOS (Data structures)
type SuitonAnalysis struct
type SuitonDistribution struct
type SuitonNormality struct
type SuitonCorrelation struct

// 2. ANÁLISIS LOCAL (Sin R)
func LocalAnalyzeDistribution()
func LocalTestNormality()
func LocalAnalyzeCorrelation()

// 3. R BRIDGE (Con R)
func CallSuitonR() // TODO: Implementar IPC
func AnalyzePricesWithSuiton() // Intenta R, fallback

// 4. SIGNAL GENERATION
func GenerateEnhancedSignal() // Con confianza estadística

// 5. UTILIDADES
func PrintSuitonAnalysis()
func ComputeSuitonStats()
func SuitonToJSON() / SuitonFromJSON()
```

---

### 📌 FC-14: ¿Cuál es el flujo de datos en AnalyzePricesWithSuiton()?

**Pregunta**: ¿Cuál es el orden de operaciones?

**Respuesta**:
```
1. Input: []Candle (velas de precios)
   ↓
2. Extraer: []float64 (solo precios de cierre)
   ↓
3. Intentar: CallSuitonR(prices)
   ├─ Si R disponible: Usar Rscript
   └─ Si R NO disponible: Retorna nil
   ↓
4. Si nil: Usar análisis local
   ├─ LocalAnalyzeDistribution(prices)
   ├─ LocalTestNormality(prices)
   └─ LocalAnalyzeCorrelation(prices)
   ↓
5. Output: *SuitonAnalysis (JSON-serializable)
   ├─ Distribution stats
   ├─ Normality test
   └─ Correlation analysis
```

**Ventaja**: Funciona con o sin R ✅

---

### 📌 FC-15: ¿Cómo funciona GenerateEnhancedSignal()?

**Pregunta**: ¿Cuál es la lógica de confidence scoring?

**Respuesta**:
```
Confidence = 0.5 (base)

IF distribution.IsNormal:
    confidence += 0.2  // Distribución predecible

IF correlation.Value > 0.5:
    confidence += 0.2  // Uptrend fuerte

IF distribution.CV < 1.0:
    confidence += 0.1  // Baja volatilidad

──────────────────────────
IF confidence > 0.7: "BUY (High Confidence)"
ELSE IF confidence > 0.5: "BUY (Medium Confidence)"
ELSE: "HOLD"
```

**Interpretación**:
- Combina 4 factores
- Score máximo: 1.0
- Umbral: > 0.7 para confianza alta

---

## TRADING STRATEGY

### 📌 FC-16: ¿Cuál es la evolución de signals del Day 3 al Day 4?

**Pregunta**: ¿Qué cambió en la generación de señales?

**Respuesta**:
**Day 3**:
```
IF (Volume >= 1300 AND ClosePercent >= 0.1%):
    signal = "BUY"
ELSE:
    signal = "HOLD"
```

**Day 4**:
```
Base signal (Day 3 logic)
IF base_signal == "BUY":
    confidence = calculate_confidence(distribution, correlation, volatility)
    IF confidence > 0.7:
        signal = "BUY (High Confidence)"
    ELIF confidence > 0.5:
        signal = "BUY (Medium Confidence)"
    ELSE:
        signal = "HOLD"
```

**Mejora**: Filtrado estadístico de falsas señales

---

### 📌 FC-17: ¿Por qué el trade del Day 4 perdió dinero?

**Pregunta**: ¿Cuál fue el resultado del backtest?

**Respuesta**:
**Entrada**: Vela 3, precio $1502.00
**Salida**: Vela 100, precio $1480.50
**P&L**: -$14.31 (-1.43%)

**Razones**:
1. **Mercado sideways**: Correlación ≈ 0 (sin trend)
2. **Volatilidad alta**: Luego de entrada, precio bajó
3. **Stop loss no implementado**: Sin protección
4. **Dataset adverso**: Patrones que favorecen downside

**Conclusión**: Sistema funciona, datos necesitan mejor tendencia

---

### 📌 FC-18: ¿Cuáles son las métricas clave del Day 4?

**Pregunta**: ¿Cuál es el resumen de resultados?

**Respuesta**:
| Métrica | Valor |
|---------|-------|
| Initial Capital | $10,000 |
| Final Capital | $9,985.69 |
| Return | -0.14% |
| Total Trades | 1 |
| Winning Trades | 0 |
| Success Rate | 0% |
| Max Drawdown | 0.14% |
| BUY Signals | 35/100 (35%) |
| HOLD Signals | 65/100 (65%) |

**Interpretación**: Sistema estable, necesita mejor generación de signals

---

## ARQUITECTURA Y DISEÑO

### 📌 FC-19: ¿Cuál es el diagrama de integración Go-R?

**Pregunta**: ¿Cómo se conectan los módulos?

**Respuesta**:
```
Fuuton (Go)
│
├─ main.go
│  ├─ ReadCSV (data.csv)
│  ├─ AnalyzePricesWithSuiton()
│  │  └─ suiton_bridge.go
│  │     ├─ CallSuitonR() → Suiton (R)
│  │     │  └─ suiton.R (statistical analysis)
│  │     │
│  │     └─ LocalAnalyze*() [Fallback]
│  │
│  ├─ GenerateEnhancedSignal()
│  │  └─ usa SuitonAnalysis
│  │
│  └─ RunOrchestrator()
│     └─ backtest.go → backtest.csv
│
└─ Integración Katon (Python)
   └─ katon_bridge.go (para indicators)
```

**Status**: 3/4 módulos integrados ✅

---

### 📌 FC-20: ¿Cuáles son los próximos pasos después de Day 4?

**Pregunta**: ¿Qué hace falta para completar el sistema?

**Respuesta**:
**Day 5-6 (CRÍTICO)**:
- Instalar R en Windows
- Implementar R Bridge IPC (subprocess + JSON)
- Machine Learning prep

**Day 5-6 (HIGH)**:
- Feature engineering (MA, RSI, MACD)
- Better risk management (stop loss, take profit)
- Improve signal logic

**Day 7-8**:
- Integrar Doton (C#) - 4/4 módulos
- NinjaTrader bridge

**Day 9+**:
- Live trading setup
- Real market data
- Risk monitoring

**Vision**: Sistema completo en Day 9-10 ✅

---

## RESUMEN RÁPIDO

**Temas Day 4**:
1. Distribuciones (mean, std, skewness, kurtosis)
2. Normalidad (Shapiro-Wilk, p-value)
3. Correlación (Pearson, trend detection)
4. R programming (suiton.R)
5. Go integration (suiton_bridge.go)
6. Enhanced signals (confidence scoring)
7. Backtesting (100 velas)
8. Architecture (3/4 módulos)

**Conceptos clave**: Statistical rigor + Software integration = Better trading system

---

**Generated**: 28/11/2025
**Flashcards**: 20
**Difficulty**: Intermediate
**Study time**: ~30-45 min
**Next**: Flashcards Day 5 (ML prep)

> "Las flashcards de Day 4 cubren los fundamentos estadísticos y la arquitectura de integración. Estudia estos conceptos para entender ML del Day 5."

---
