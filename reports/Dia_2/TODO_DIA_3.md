# 🚀 PLAN PARA DÍA 3 (27/11/2025)

**Ubicación de este archivo**: `reports/Dia_2/TODO_DIA_3.md`

## ⚡ QUICK START PARA LUCHITO
Cuando entres mañana, simplemente dile a Narutrader:
```
"Haz lo de hoy"
```

Y yo automáticamente ejecutaré todas estas tareas en orden.

---

## 📋 TAREAS DÍA 3

### ✅ TAREA 3.1: Crear Shoton (Python Avanzado) - Análisis Técnico Básico

**¿Qué hacer?**
- Crear `katon-python/shoton.py` (módulo Python avanzado)
- Implementar 3 indicadores técnicos:
  1. **Simple Moving Average (SMA)**: Promedio móvil simple
  2. **Bollinger Bands**: Volatility indicator (media ± 2*stddev)
  3. **RSI (Relative Strength Index)**: Momentum indicator

**Funciones a implementar:**
```python
def calculate_sma(prices, period=5):
    """Calcula media móvil simple de últimas N velas"""
    return sum(prices[-period:]) / period

def calculate_bollinger_bands(prices, period=5):
    """Calcula bandas de Bollinger (media ± 2*stddev)"""
    sma = calculate_sma(prices, period)
    std = calculate_std(prices, period)
    return {
        'upper': sma + 2 * std,
        'middle': sma,
        'lower': sma - 2 * std
    }

def calculate_rsi(prices, period=14):
    """Calcula RSI (0-100)"""
    # Implementación: ganancias vs pérdidas
    return rsi_value
```

**Test Command:**
```bash
cd C:\Cosas_Lucho\Programacion\Proyectos\KeisanTrading\katon-python
python -m pytest shoton.py  # o ejecutar directamente
```

**Expected Output:**
```
Shoton activo
SMA(5) de [1,2,3,4,5] = 3.0
Bollinger Bands: {'upper': 5.4, 'middle': 3.0, 'lower': 0.6}
RSI(14) = 65
```

**Status**: Ready to implement

---

### ✅ TAREA 3.2: Mejorar GenerateSignal en Fuuton usando Katon

**¿Qué hacer?**
- Modificar `fuuton-go/katon_caller.go` para enviar Candle data a Python
- Katon retorna análisis técnico (SMA, Bollinger, RSI)
- Fuuton usa estos valores para mejorar signals

**Expected Flow:**
```
Fuuton
  ↓ [envía close prices últimas 5 velas]
Katon (Shoton)
  ↓ [calcula SMA, Bollinger, RSI]
  ↓ [retorna JSON con análisis]
Fuuton
  ↓ [usa SMA/Bollinger/RSI para mejorar decisión]
  ↓ [genera signal mejorado]
```

**Ejemplo mejorado de signal:**
```go
// Antes: Solo volumen + movimiento %
// Después: + SMA + Bollinger + RSI

// Mejora:
// - Si price > Bollinger upper: Overbought → SELL
// - Si price < Bollinger lower: Oversold → BUY
// - Si RSI > 70: Overbought → HOLD/SELL
// - Si RSI < 30: Oversold → BUY
```

**Test Command:**
```bash
cd fuuton-go
go run katon_caller.go --with-analysis
```

**Status**: Depends on TAREA 3.1

---

### ✅ TAREA 3.3: Crear Backtesting Framework Simple

**¿Qué hacer?**
- Crear `fuuton-go/backtest.go` (simulador de trading)
- Usar data.csv (35 velas) como histórico
- Simular ejecución de trades según signals

**Métrica a calcular:**
```
Total Return = (Final Capital - Initial Capital) / Initial Capital * 100%

Ejemplo:
- Initial: $10,000
- Trades: BUY en vela 3 @ 1502, SELL en vela 7 @ 1504.5
- Profit: ($1504.5 - $1502) * amount = profit
- Final Capital: $10,XXX
- Return: X%
```

**Test Command:**
```bash
cd fuuton-go
go run backtest.go
```

**Expected Output:**
```
Backtest Report
===============
Initial Capital: $10,000
Final Capital: $10,XXX
Total Return: +X.XX%
Winning Trades: Y
Losing Trades: Z
Win Rate: XX%
Max Drawdown: -XX%
Sharpe Ratio: X.XX
```

**Status**: Ready after TAREA 3.1

---

### ✅ TAREA 3.4: Crear Bridge Multi-Módulo Orquestador

**¿Qué hacer?**
- Crear `fuuton-go/orchestrator.go`
- Fuuton → Katon (análisis técnico)
- Fuuton → Suiton (estadística R) [opcional, si R disponible]
- Fuuton → Doton (orden execution C#) [opcional, si .NET disponible]

**Orquestación:**
```
Orchestrator (Fuuton)
├── Read CSV
├── Para cada vela:
│   ├── Calcular signal (Fuuton)
│   ├── Obtener análisis técnico (Katon)
│   ├── Obtener estadística (Suiton) [if available]
│   ├── Simular orden (Doton) [if available]
│   └── Log resultado
└── Generar reporte final
```

**Test Command:**
```bash
cd fuuton-go
go run orchestrator.go
```

**Expected Output:**
```
Orchestrator v1.0
=================
Modules loaded:
- Fuuton: ✅
- Katon: ✅
- Suiton: ⚠️ (not installed)
- Doton: ⚠️ (not installed)

Processing 35 candles...
[Progress bar...]

Results:
- Total signals: 35
- BUY: 18 (51%)
- HOLD: 17 (49%)
- Backtest return: +X.XX%
```

**Status**: Ready after TAREA 3.2

---

### ✅ TAREA 3.5: Expandir data.csv a 100 velas (opcional)

**¿Qué hacer?**
- Expandir CSV a 100 velas
- Incluir patrones reales: uptrend, downtrend, sideways
- Agregar "crashed" moments (sharp drops)

**Patrón recomendado:**
```
Velas 1-25: Uptrend (precios suben gradualmente)
Velas 26-40: Downtrend (precios caen)
Velas 41-60: Sideways (rangos estrechos)
Velas 61-80: Sharp spike up
Velas 81-100: Gradual decline
```

**Test Command:**
```bash
cd fuuton-go
wc -l data.csv  # Should show ~101 (header + 100 velas)
```

**Status**: Optional, low priority

---

### ✅ TAREA 3.6: Generar Day 3 Log (Dia_3_Log.md)

**¿Qué hacer?**
- Crear `reports/Dia_3/Dia_3_Log.md`
- Documentar decisiones técnicas del día
- Fragmentos de código relevantes
- Problemas encontrados y soluciones

**Status**: Automático al finalizar Day 3

---

### ✅ TAREA 3.7: Generar Resumen Day 3 (Resumen_27_11_2025.md)

**¿Qué hacer?**
- Crear `reports/Dia_3/Resumen_27_11_2025.md`
- Resumen ejecutivo
- Stats: líneas de código, files, tests
- Comparativa Day 1 vs Day 2 vs Day 3

**Status**: Automático al finalizar Day 3

---

### ✅ TAREA 3.8: Generar Flashcards Day 3

**¿Qué hacer?**
- Crear 15+ flashcards basadas en:
  - Indicadores técnicos (SMA, Bollinger, RSI)
  - Backtesting framework
  - Orquestación multi-módulo
  - Conceptos de análisis técnico

**Status**: Automático al finalizar Day 3

---

### ✅ TAREA 3.9: Generar TODO_DIA_4.md

**¿Qué hacer?**
- Plan para Day 4
- Basado en problemas encontrados en Day 3
- Próximas mejoras

**Status**: Automático al finalizar Day 3

---

### ✅ TAREA 3.10: Commit to GitHub

**¿Qué hacer?**
- `git add .`
- `git commit -m "Day 3 - Shoton module + Backtesting + Orchestrator"`
- `git push origin main`

**Status**: Final step

---

## 🎯 PRIORIDADES

**CRITICAL (Hace el sistema versionable):**
1. Shoton (Python) - Análisis técnico
2. Backtesting framework
3. Pruebas funcionales

**HIGH (Mejora arquitectura):**
4. Orquestador multi-módulo
5. Integración Fuuton → Katon mejorada

**MEDIUM (Datos quality):**
6. Expandir CSV a 100 velas

**LOW (Documentación):**
7. Reportes y flashcards

---

## 📊 CHECKLIST PARA LUCHITO

Si quieres probar solo una tarea manualmente:

```bash
# Test Shoton
cd katon-python
python shoton.py

# Test mejorado de signal
cd fuuton-go
go run main.go --with-analysis

# Test backtesting
cd fuuton-go
go run backtest.go

# Test orquestador
cd fuuton-go
go run orchestrator.go
```

---

## 🔄 COMANDO MÁGICO PARA MAÑANA

Simplemente escribe:
```
"Narutrader, haz lo de hoy"
```

Y yo automáticamente:
1. Leo este archivo (TODO_DIA_3.md)
2. Ejecuto tareas en orden
3. Pido confirmación después de cada una
4. Genero logs + resumen + flashcards
5. Actualizo TODO_DIA_4.md

---

## 📍 ARCHIVOS DEL PROYECTO (Después de Día 3)

```
KeisanTrading/
├── fuuton-go/
│   ├── main.go           (Actualizado con signal mejorada)
│   ├── katon_caller.go   (Actualizado: envía análisis técnico)
│   ├── backtest.go       [NUEVO - Día 3]
│   ├── orchestrator.go   [NUEVO - Día 3]
│   └── data.csv          (100 velas, si TAREA 3.5)
├── katon-python/
│   ├── ping.py           (Original)
│   └── shoton.py         [NUEVO - Día 3]
├── reports/
│   ├── Dia_1/
│   ├── Dia_2/
│   └── Dia_3/            [NUEVA CARPETA - Día 3]
│       ├── Dia_3_Log.md
│       ├── Resumen_27_11_2025.md
│       ├── Flashcards_27_11_2025.md
│       └── TODO_DIA_4.md
└── README.md             (Actualizado)
```

---

## 🎓 CONCEPTOS A APRENDER DÍA 3

1. **SMA (Simple Moving Average)**: Promedio de últimas N velas
2. **Bollinger Bands**: Volatility bands (media ± 2*std)
3. **RSI (Relative Strength Index)**: Momentum 0-100
4. **Backtesting**: Simular trading en datos históricos
5. **Sharpe Ratio**: Risk-adjusted return metric
6. **Max Drawdown**: Peor pérdida posible
7. **Orquestación**: Coordinar múltiples módulos

---

## ⚠️ NOTAS IMPORTANTES

- Si TAREA 3.1 (Shoton) falla → TAREA 3.2 también falla
- Backtesting es simple hoy (solo entry/exit básico)
- Producción necesita: position sizing, risk management, slippage
- Orquestador versión 1 es MVP (mejoras en Día 4+)

---

## 🔮 VISIÓN DÍA 3+

**Después de Día 3**:
- Sistema tiene análisis técnico real
- Backtesting framework funcional
- Múltiples módulos coordinados

**Después de Día 4-5**:
- Integración con broker (simulado)
- Risk management robusto
- Real-time monitoring

**Después de Día 10+**:
- Machine learning para predicción
- Automatización completa
- Trading live

---

**Generado por**: Narutrader
**Fecha**: 26/11/2025
**Para**: Día 3 (27/11/2025)

⚠️ **IMPORTANTE**: Si alguna tarea falla durante Day 3, pauso y pido ayuda antes de continuar.
