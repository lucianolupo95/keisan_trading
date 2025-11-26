# 🚀 PLAN PARA DÍA 5 (29/11/2025)

**Ubicación de este archivo**: `reports/Dia_4/TODO_DIA_5.md`

---

## ⚡ QUICK START PARA LUCHITO

Cuando entres mañana (Día 5), simplemente dile a Narutrader:
```
"Haz lo de hoy (Day 5)"
```

Y yo automáticamente:
1. Leo este archivo (TODO_DIA_5.md)
2. Ejecuto tareas en orden
3. Genero documentación + logs + flashcards
4. Creo TODO_DIA_6.md para el próximo día

---

## 📋 TAREAS DÍA 5 (29/11/2025)

### ✅ TAREA 5.1: Instalar R en Windows

**¿Qué hacer?**
- Descargar e instalar R desde https://cran.r-project.org/
- Instalar Rscript en PATH
- Instalar paquete `e1071` (para suiton.R)
- Verificar instalación

**Comandos a ejecutar**:
```bash
# Test 1: Verificar Rscript
Rscript --version
# Expected: R scripting front-end version 4.x.x

# Test 2: Instalar e1071
Rscript -e "install.packages('e1071')"

# Test 3: Ejecutar suiton.R
cd suiton-r
Rscript suiton.R
# Expected: Output completo con estadísticas
```

**Expected Output** (suiton.R corriendo):
```
╔════════════════════════════════════════════════════════════╗
║           SUITON: Statistical Analysis Module              ║
║              KeisanTrading - Day 4                          ║
╚════════════════════════════════════════════════════════════╝

✓ Suiton activo
✓ Datos: 35 precios generados

📊 DISTRIBUTION ANALYSIS
  Media:      1508.93
  Mediana:    1510.00
  Std Dev:    6.54
  Skewness:   -0.0168
  ...
```

**Status**: Bloqueante para 5.2

**Difficulty**: Fácil
**Time estimate**: 15-20 minutos

---

### ✅ TAREA 5.2: Implementar R Bridge IPC

**¿Qué hacer?**
- Crear `fuuton-go/r_ipc.go` (Inter-Process Communication)
- Implementar `CallSuitonR()` (actualmente placeholder)
- Serializar precios a JSON
- Ejecutar Rscript como subprocess
- Parsear JSON de respuesta

**Pseudocódigo**:
```go
func CallSuitonR(prices []float64) *SuitonAnalysis {
    // 1. Crear archivo JSON temporal con precios
    inputJSON := convertPricesToJSON(prices)
    writeToFile("temp_prices.json", inputJSON)

    // 2. Ejecutar: Rscript r_ipc.R temp_prices.json
    cmd := exec.Command("Rscript", "r_ipc.R", "temp_prices.json")
    output, err := cmd.Output()

    // 3. Parsear respuesta JSON
    analysis := parseJSON(output)

    // 4. Limpiar
    deleteFile("temp_prices.json")

    return analysis
}
```

**Archivo nuevo**: `suiton-r/r_ipc.R`
```r
# Lee JSON de precios
# Ejecuta análisis
# Retorna JSON con resultados
```

**Test**:
```bash
cd fuuton-go
go build -o fuuton_v5.exe ...
./fuuton_v5.exe
# Output debe mostrar "Using R (Suiton)" en lugar de "fallback"
```

**Status**: Depende de 5.1

**Difficulty**: Media
**Time estimate**: 30-40 minutos

---

### ✅ TAREA 5.3: Feature Engineering - Technical Indicators

**¿Qué hacer?**
- Crear `fuuton-go/indicators.go`
- Implementar 3 indicadores técnicos:
  1. **Moving Average (MA)** - Precio promedio últimas N velas
  2. **RSI (Relative Strength Index)** - Momentum [0-100]
  3. **MACD (Moving Avg Convergence Divergence)** - Trend + momentum

**Funciones a implementar**:

```go
// Moving Average (20 velas)
func CalculateMA20(prices []float64) float64 {
    // Retorna promedio últimas 20 velas
}

// RSI (14 períodos)
func CalculateRSI(prices []float64) float64 {
    // Retorna RSI [0-100]
    // < 30: Oversold (bueno para BUY)
    // > 70: Overbought (bueno para SELL)
}

// MACD
type MACDValue struct {
    MACD      float64  // EMA12 - EMA26
    Signal    float64  // EMA9 del MACD
    Histogram float64  // MACD - Signal
}

func CalculateMACD(prices []float64) MACDValue {
    // Retorna MACD components
}
```

**Test**:
```bash
cd fuuton-go
# Test RSI
# Test MA20
# Test MACD
```

**Expected behavior**:
- MA20 = promedio suave de precios
- RSI < 30 en downtrends (oportunidad BUY)
- MACD positivo en uptrends

**Status**: Independiente

**Difficulty**: Media-Alto
**Time estimate**: 45-60 minutos

---

### ✅ TAREA 5.4: Integrar Indicadores en Signals

**¿Qué hacer?**
- Actualizar `GenerateEnhancedSignal()` en suiton_bridge.go
- Incorporar RSI, MA20, MACD

**Nueva lógica**:
```
Base confidence = 0.5

// Factores estadísticos (Day 4)
if distribution.IsNormal: confidence += 0.15
if correlation > 0.5: confidence += 0.15
if cv < 1.0: confidence += 0.1

// Factores técnicos (Day 5) NEW
if rsi < 30: confidence += 0.15  // Oversold = compra
if ma20_uptrend: confidence += 0.1  // Trend up
if macd_positive: confidence += 0.1  // Bullish

──────────────────────────────────
Max confidence: 1.0

> 0.8: "BUY (Very High Confidence)" 💪💪
> 0.6: "BUY (High Confidence)" 💪
> 0.4: "BUY (Medium Confidence)" 🤔
< 0.4: "HOLD" 😐
```

**Test**:
```bash
cd fuuton-go
go run main.go ... indicators.go
# Output debe mostrar confidence levels actualizados
```

**Status**: Depende de 5.3

**Difficulty**: Media
**Time estimate**: 20-30 minutos

---

### ✅ TAREA 5.5: Test Nuevo Sistema con 100 Velas

**¿Qué hacer?**
- Compilar Go con nuevas features
- Ejecutar backtest con:
  - Suiton (R) activado
  - Indicadores técnicos (MA, RSI, MACD)
  - Señales mejoradas
  - 100 velas de datos

**Test Command**:
```bash
cd fuuton-go
go build -o fuuton_v5.exe main.go backtest.go orchestrator.go \
    katon_bridge.go suiton_bridge.go indicators.go r_ipc.go
./fuuton_v5.exe
```

**Expected Output**:
```
╔═══════════════════════════════════════════════════╗
║      FUUTON v5: Full System Integration          ║
║    Go + Python (Katon) + R (Suiton)             ║
╚═══════════════════════════════════════════════════╝

✓ Fuuton OK
✓ Velas leídas: 100

📊 SUITON ANALYSIS (Statistical Analysis Module)
✓ Using R (Suiton)  ← Cambio de "fallback" a "Using R"

📈 TECHNICAL INDICATORS
  MA20:        1508.50
  RSI:         42.3 (neutral)
  MACD:        +0.15 (bullish)

=== SEÑALES GENERADAS (MEJORADAS) ===
Vela 1 - Signal: HOLD
Vela 11 - Signal: BUY (High Confidence)  ← Más confidence
Vela 21 - Signal: BUY (Very High Confidence)  ← Mejor
...

=== ORCHESTRATOR BACKTEST ===
Initial: $10,000
Final: $X,XXX
Return: +Y.YY% (mejor que -0.14%)
Trades: Z

📈 ESTADÍSTICAS FINALES
Distribution: Normal
Trend: (detecta correctamente)
RSI: 42.3
MACD Status: Bullish/Bearish
```

**Expectations**:
- Return debería mejorar (> 0%)
- Más BUY signals correctas
- Menos falsos positivos

**Status**: Depende de 5.1-5.4

**Difficulty**: Fácil (solo compilar)
**Time estimate**: 5-10 minutos

---

### ✅ TAREA 5.6: Comparar Day 4 vs Day 5

**¿Qué hacer?**
- Crear tabla comparativa
- Documenta cambios en:
  - Número de signals
  - Quality de signals
  - Return %
  - Indicadores disponibles

**Tabla esperada**:

| Métrica | Day 4 | Day 5 | Cambio |
|---------|-------|-------|--------|
| **Modules** | 3/4 (fallback) | 3/4 (full R) | - |
| **BUY Signals** | 35/100 | X/100 | +/- |
| **Confidence Levels** | 2 (High/Med) | 4 (VeryHigh/High/Med/None) | +2 |
| **Indicators** | Volume + Movement | + MA20 + RSI + MACD | +3 |
| **Return** | -0.14% | +X.XX% | Esperado + |
| **Winning Trades** | 0/1 | X/Y | Esperado + |
| **Max Drawdown** | 0.14% | X.XX% | Esperado - |

**Status**: Automático tras 5.5

**Difficulty**: Trivial
**Time estimate**: 10-15 minutos

---

### ✅ TAREA 5.7: Generar Day 5 Log

**¿Qué hacer?**
- Crear `reports/Dia_5/Dia_5_Log.md`
- Documenta:
  - R installation results
  - IPC implementation details
  - Technical indicators code
  - Signal improvements
  - Backtest results
  - Comparativa Day 4 vs 5

**Secciones**:
1. Executive Summary
2. Tasks Completed (5.1-5.6)
3. Technical Details
4. Code Examples
5. Results & Analysis
6. Next Steps

**Status**: Automático

**Difficulty**: Trivial
**Time estimate**: 30-45 minutos

---

### ✅ TAREA 5.8: Generar Resumen Day 5

**¿Qué hacer?**
- Crear `reports/Dia_5/Resumen_29_11_2025.md`
- Resumen ejecutivo
- Key metrics
- Architecture update

**Status**: Automático

**Difficulty**: Trivial
**Time estimate**: 20-30 minutos

---

### ✅ TAREA 5.9: Generar Flashcards Day 5

**¿Qué hacer?**
- Crear 20+ flashcards basadas en:
  - Moving Averages (MA)
  - RSI (Relative Strength Index)
  - MACD (convergence divergence)
  - IPC (Inter-Process Communication)
  - R Bridge implementation
  - Indicator integration
  - Trading strategy improvements

**Status**: Automático

**Difficulty**: Trivial
**Time estimate**: 30-45 minutos

---

### ✅ TAREA 5.10: Generar Resumen_for_dummies Day 5

**¿Qué hacer?**
- Actualizar explicación para "niños de 5 años"
- Qué son Moving Averages
- Qué es RSI (oversold/overbought)
- Qué es MACD
- Cómo se usan en trading

**Status**: Automático

**Difficulty**: Trivial
**Time estimate**: 20-30 minutos

---

### ✅ TAREA 5.11: Generar TODO_DIA_6.md

**¿Qué hacer?**
- Plan para Day 6 (30/11/2025)
- Basado en resultados de Day 5
- Próximas features

**Preview Day 6**:
- Machine Learning prep (feature selection)
- Better risk management (stop loss)
- C# module integration (Doton)
- Improved backtest framework

**Status**: Automático

**Difficulty**: Trivial
**Time estimate**: 15-20 minutos

---

### ✅ TAREA 5.12: Commit to GitHub

**¿Qué hacer?**
- `git add .`
- `git commit -m "Day 5 - R Bridge + Technical Indicators + IPC"`
- `git push origin main`

**Commit message format**:
```
Day 5 - R Bridge IPC + Technical Indicators (MA, RSI, MACD)

- Instalar R en Windows
- Implementar R Bridge con JSON IPC
- Crear indicadores técnicos (MA20, RSI, MACD)
- Integrar indicadores en signal generation
- Test completo con 100 velas
- Documentación y reportes

📊 Results:
- R module funcional (no fallback)
- 3 indicadores nuevos integrados
- Señales con 4 niveles de confianza
- Backtest mejorado esperado

Generated with Claude Code
```

**Status**: Final step

---

## 🎯 PRIORIDADES

**CRITICAL (Necesario para funcionar)**:
1. R installation (5.1) - Sin esto, fallback
2. R Bridge IPC (5.2) - Usar poder real de R

**HIGH (Mejora significativa)**:
3. Technical indicators (5.3) - Mejor signals
4. Integration (5.4) - Usar todos los datos
5. Testing (5.5) - Validar improvements

**MEDIUM (Validación)**:
6. Comparativa (5.6) - Medir progreso
7. Documentation (5.7-5.11) - Conocimiento

**LOW (Housekeeping)**:
8. Commit (5.12) - Version control

---

## 🔄 DEPENDENCIAS

```
5.1 (Install R)
  ↓
  ├─→ 5.2 (R Bridge IPC) ✓
  │     ↓
  │     └─→ 5.5 (Test) ✓
  │           ↓
  │           └─→ 5.6 (Compare) ✓
  │
5.3 (Indicators) [Independente]
  ↓
  ├─→ 5.4 (Integration)
  │     ↓
  │     └─→ 5.5 (Test) ✓
  │
  └─→ 5.5, 5.6 (Ambos)

Todos → 5.7-5.11 (Documentation)
      → 5.12 (Commit)
```

---

## 📊 MÉTRICAS ESPERADAS DÍA 5

| Métrica | Esperado |
|---------|----------|
| Nuevas líneas de código | 500-700 |
| Archivos nuevos | 3-4 (r_ipc.go, indicators.go, r_ipc.R) |
| R working | ✅ |
| Technical indicators | 3 (MA, RSI, MACD) |
| Signal confidence levels | 4 |
| Backtesting result | +X.XX% (esperado positivo) |
| Winning trades | > 0 |
| Code compiles | ✅ |
| Tests pass | ✅ |

---

## 🔮 VISIÓN DÍA 5

**Después de Day 5**:
- Sistema con real R integration (no fallback)
- Technical indicators implementados
- Better signal generation
- Expected positive returns en backtest

**Arquitectura mejorada**:
```
Fuuton (Go)
├─ main.go (actualizado)
├─ indicators.go (NUEVO)
├─ r_ipc.go (NUEVO)
├─ suiton_bridge.go (actualizado)
│
Suiton (R)
├─ r_ipc.R (NUEVO)
└─ suiton.R (ya existe)

Flujo mejorado:
Data → Indicators → Suiton → Signals → Backtest
       ✓ MA20     ✓ Stats   ✓ 4 levels
       ✓ RSI      ✓ Normal
       ✓ MACD     ✓ Corr
```

---

## 📚 ARCHIVOS DEL PROYECTO (POST-DÍA 5)

```
reports/
├── ...Dia_1-4 (anterior)
└── Dia_5/  ✨ NUEVA CARPETA
    ├── Dia_5_Log.md
    ├── Resumen_29_11_2025.md
    ├── Flashcards_29_11_2025.md
    ├── Resumen_for_dummies.md
    └── TODO_DIA_6.md

fuuton-go/
├── main.go (actualizado)
├── backtest.go (sin cambios)
├── orchestrator.go (sin cambios)
├── katon_bridge.go (sin cambios)
├── suiton_bridge.go (actualizado)
├── indicators.go ✨ NUEVO
├── r_ipc.go ✨ NUEVO
├── data.csv (sin cambios)
└── fuuton_v5.exe ✨ COMPILADO

suiton-r/
├── suiton.R (sin cambios)
└── r_ipc.R ✨ NUEVO
```

---

## 🎓 CONCEPTOS A APRENDER DÍA 5

1. **Moving Averages (MA)**
   - Simple MA (SMA)
   - Exponential MA (EMA)
   - MA crossovers para signals

2. **RSI (Relative Strength Index)**
   - Cálculo (14 períodos default)
   - Oversold (< 30) y Overbought (> 70)
   - Divergences

3. **MACD (Moving Average Convergence Divergence)**
   - MACD line (EMA12 - EMA26)
   - Signal line (EMA9 del MACD)
   - Histogram (MACD - Signal)
   - Crossovers

4. **IPC (Inter-Process Communication)**
   - JSON serialization
   - Subprocess execution
   - Error handling

5. **Integration Patterns**
   - Combining statistical + technical
   - Confidence scoring
   - Signal prioritization

---

## ⚠️ NOTAS IMPORTANTES

- **R es CRÍTICO**: Sin instalación, seguimos con fallback (igual funciona)
- **IPC puede ser complejo**: Si falla, usamos análisis local (graceful degradation)
- **Indicators son matemáticamente intensivos**: Revisar fórmulas cuidadosamente
- **Testing es CRUCIAL**: Cada cambio debe validarse con backtest
- **Próximo: Machine Learning** - Day 5-6 requiere feature engineering solido

---

## 🎯 COMANDO MÁGICO PARA MAÑANA

Simplemente escribe:
```
"Narutrader, haz lo de hoy (Day 5)"
```

Y yo automáticamente:
1. Leo este archivo (TODO_DIA_5.md)
2. Ejecuto tareas 5.1-5.5 en orden (críticas)
3. Valido con tests
4. Genero logs + documentación (5.6-5.11)
5. Commit a GitHub (5.12)
6. Creo TODO_DIA_6.md

---

**Generado por**: Narutrader
**Fecha**: 28/11/2025
**Para**: Día 5 (29/11/2025)
**Status**: Ready to execute

> "Day 5 es sobre agregar inteligencia técnica a la estadística. Cuando termine, el robot será una máquina bien calibrada."

---

## 🔗 REFERENCIAS RÁPIDAS

- **R installation**: https://cran.r-project.org/
- **Moving Average**: Promedio últimas N velas
- **RSI formula**: 100 - (100 / (1 + RS)), RS = avgGain/avgLoss
- **MACD formula**: EMA12 - EMA26, Signal = EMA9(MACD)
- **Go subprocess**: `exec.Command("Rscript", args...)`

---

⚠️ **IMPORTANTE**: Si alguna tarea falla durante Day 5, pauso y pido ayuda antes de continuar.

---

**¡Narutrader listo para Día 5! 🤖**

