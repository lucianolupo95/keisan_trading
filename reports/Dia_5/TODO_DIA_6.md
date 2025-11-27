# 🚀 PLAN PARA DÍA 6 (30/11/2025)

**Ubicación de este archivo**: `reports/Dia_5/TODO_DIA_6.md`

---

## ⚡ QUICK START PARA LUCHITO

Cuando entres mañana (Día 6), simplemente dile a Narutrader:
```
"Haz lo de hoy (Day 6)"
```

Y yo automáticamente:
1. Leo este archivo (TODO_DIA_6.md)
2. Ejecuto tareas en orden
3. Genero documentación + logs + flashcards
4. Creo TODO_DIA_7.md para el próximo día

---

## 📋 TAREAS DÍA 6 (30/11/2025)

### ✅ TAREA 6.1: Fijar R Bridge IPC

**¿Qué hacer?**
- Investigar segmentation fault en r_ipc.R
- Opción 1: Eliminar dependencia jsonlite (usar formato manual)
- Opción 2: Usar rjson en su lugar
- Opción 3: Usar system2() en R para JSON output
- Testar desde Go subprocess

**Comandos a probar**:
```bash
# Debug R issue
cd fuuton-go
go run main.go ... 2>&1 | grep -A5 "R error"

# Test R bridge directo
Rscript suiton-r/r_ipc_fixed.R temp.json

# Si OK, recompilar y ejecutar
go build -o fuuton_v6.exe ...
./fuuton_v6.exe
```

**Expected Output**:
```
✓ Using R (Suiton)  ← Cambio de "fallback" a real R
```

**Status**: Bloqueante para mejor performance
**Difficulty**: Media
**Time estimate**: 20-30 minutos

---

### ✅ TAREA 6.2: Parameter Tuning

**¿Qué hacer?**
- Test diferentes períodos de MA (15, 20, 50)
- Test diferentes períodos de RSI (9, 14, 21)
- Optimizar MACD parameters si necesario
- Comparar backtest results con cada combinación

**Pseudocódigo**:
```go
// Test MA periods
for _, period := range []int{15, 20, 50} {
    ma := CalculateSMA(prices, period)
    // Compare signal quality
}

// Test RSI periods
for _, period := range []int{9, 14, 21} {
    rsi := CalculateRSI(prices, period)
    // Compare detection accuracy
}
```

**Expected Output**:
```
MA15:  Better for short-term, more false positives
MA20:  Balanced (current)
MA50:  Better for long-term trends

RSI9:  Faster but more noise
RSI14: Current (industry standard)
RSI21: Smoother but slower
```

**Status**: Optimización
**Difficulty**: Media
**Time estimate**: 30 minutos

---

### ✅ TAREA 6.3: Multiple Dataset Testing

**¿Qué hacer?**
- Generar nuevos datasets con diferentes patrones
- Patrón 1: Strong uptrend (testerá si compra en subidas)
- Patrón 2: Strong downtrend (testerá si evita pérdidas)
- Patrón 3: Volatile sideways (testerá robustez)
- Correr backtest con cada patrón

**Pseudocódigo**:
```go
// Pattern 1: Uptrend
// Generar 100 precios: 1500 → 1600 (subida constante)
prices := generateUptrend(1500, 1600, 100)

// Pattern 2: Downtrend
prices := generateDowntrend(1600, 1500, 100)

// Pattern 3: Sideways volatile
prices := generateSideways(1550, 1550±50, 100)

// Correr backtest para cada
```

**Expected Output**:
```
UPTREND TEST:
- Debería generar muchas BUY señales
- Todas con HIGH confidence
- Backtest ROI > +5%

DOWNTREND TEST:
- Debería generar pocas BUY señales
- HOLD debería ser mayoría
- Backtest ROI > -2%

SIDEWAYS TEST:
- Debería balanced
- Algunos falsos positivos (esperado)
- Backtest ROI cerca 0%
```

**Status**: Validación
**Difficulty**: Fácil
**Time estimate**: 25 minutos

---

### ✅ TAREA 6.4: Machine Learning Prep

**¿Qué hacer?**
- Investigar qué features usar para ML
- Feature selection (qué indicadores + cuáles)
- Feature engineering (crear nuevas features)
- Crear plan para Day 7 (ML implementation)

**Features Candidatos**:
```
Existing:
- MA20 value
- RSI value
- MACD values
- Volatility
- Correlation

New Features to Create:
- Price momentum (% change)
- MA slope (dirección de MA)
- RSI extreme count (cuántas veces < 30)
- MACD zero crosses (cuántos cruces)
- Volume weighted indicators
```

**Pseudocódigo**:
```go
type Features struct {
    MA20Value      float64
    RSIValue       float64
    MACDValue      float64
    PriceMomentum  float64  // NEW
    MASlope        float64  // NEW
    RSIExtremes    int      // NEW
    MACDCrosses    int      // NEW
    Volatility     float64
    // ... más features
}
```

**Status**: Investigación + planning
**Difficulty**: Media-Alto
**Time estimate**: 45 minutos

---

### ✅ TAREA 6.5: Documentation + Logs

**¿Qué hacer?**
- Crear `reports/Dia_6/Dia_6_Log.md`
- Crear `reports/Dia_6/Comparativa_Day5_vs_Day6.md`
- Crear `reports/Dia_6/Resumen_30_11_2025.md`
- Crear flashcards Day 6

**Status**: Automático
**Difficulty**: Trivial
**Time estimate**: 45 minutos

---

### ✅ TAREA 6.6: Generar TODO_DIA_7.md

**¿Qué hacer?**
- Plan para Day 7 (Machine Learning)
- Basado en resultados de Day 6
- Próximas mejoras

**Preview Day 7**:
- Machine Learning setup (feature extraction)
- Model training (XGBoost o similar)
- Backtesting con modelo
- Model validation

**Status**: Automático
**Difficulty**: Trivial
**Time estimate**: 20 minutos

---

### ✅ TAREA 6.7: Commit to GitHub

**¿Qué hacer?**
- `git add .`
- `git commit -m "Day 6 - R Bridge Fix + Parameter Tuning + ML Prep"`
- `git push origin main`

**Status**: Final step
**Time estimate**: 5 minutos

---

## 🎯 PRIORIDADES

**CRITICAL (Necesario)**:
1. R Bridge fix (6.1) - Sin esto, seguimos en fallback
2. Parameter tuning (6.2) - Optimizar indicadores

**HIGH (Útil)**:
3. Multiple testing (6.3) - Validar robustez
4. ML prep (6.4) - Plan para siguiente phase

**MEDIUM (Documentación)**:
5. Logs (6.5) - Registrar trabajo
6. TODO_DIA_7 (6.6) - Plan siguiente

**LOW (Housekeeping)**:
7. Commit (6.7) - Version control

---

## 🔄 DEPENDENCIAS

```
6.1 (R Bridge) - Independiente
  ↓
  ├─→ 6.2 (Parameters) - Mejor con real R
  │     ↓
  │     ├─→ 6.3 (Testing) - Validar tuning
  │
6.4 (ML Prep) - Independiente
  ↓
  └─→ 6.5-6.7 (Documentation) - Final steps
```

---

## 📊 MÉTRICAS ESPERADAS DÍA 6

| Métrica | Esperado |
|---------|----------|
| R Bridge Fixed | ✅ Yes |
| Parameters Tuned | ✅ MA20 + RSI14 confirmed |
| Datasets Tested | 3 (uptrend, downtrend, sideways) |
| ML Features Identified | 8-10 |
| Code Changes | +200 líneas (testing) |
| Tests Pass | ✅ 100% |
| Documentation | Complete |
| Ready for ML | YES |

---

## 🔮 VISIÓN DÍA 6+

**Después de Day 6**:
- R Bridge funcionando en tiempo real
- Indicadores optimizados
- Sistema validado con múltiples patrones
- ML architecture diseñada

**Después de Day 7-8**:
- Machine Learning model trained
- Feature engineering completo
- Better signal accuracy esperada

**Después de Day 9-10**:
- C# bridge integrado
- Risk management robusto
- Ready para live trading

---

## 📚 ARCHIVOS DEL PROYECTO (POST-DÍA 6)

```
reports/
├── Dia_1-5/ (anterior)
└── Dia_6/  ✨ NUEVA CARPETA
    ├── Dia_6_Log.md
    ├── Comparativa_Day5_vs_Day6.md
    ├── Resumen_30_11_2025.md
    ├── Flashcards_30_11_2025.md
    └── TODO_DIA_7.md

fuuton-go/
├── main.go (sin cambios)
├── indicators.go (sin cambios)
├── suiton_bridge.go (sin cambios)
├── r_ipc_testing.go ✨ NUEVO (testing utilities)
├── features.go ✨ NUEVO (ML feature extraction)
└── fuuton_v6.exe ✨ COMPILADO
```

---

## 🎓 CONCEPTOS A APRENDER DÍA 6

1. **R Bridge Debugging**
   - Windows subprocess issues
   - JSON communication
   - Error handling

2. **Parameter Optimization**
   - Period sensitivity
   - Backtesting frameworks
   - Metric comparison

3. **Feature Engineering**
   - Feature selection
   - Normalization
   - Correlation analysis

4. **ML Prep**
   - Training/test split
   - Feature scaling
   - Model selection

---

## ⚠️ NOTAS IMPORTANTES

- **R Bridge**: Crítico pero con fallback seguro
- **Parameters**: Diferentes mercados pueden necesitar diferentes values
- **Testing**: Importante validar con patrones diversos
- **ML**: Esperar buena mejora en accuracy
- **Próximo**: Day 7 empezará ML model training

---

## 🎯 COMANDO MÁGICO PARA MAÑANA

Simplemente escribe:
```
"Narutrader, haz lo de hoy (Day 6)"
```

Y yo automáticamente:
1. Leo este archivo (TODO_DIA_6.md)
2. Ejecuto tareas 6.1-6.4 en orden (críticas)
3. Valido con tests
4. Genero logs + documentación (6.5-6.6)
5. Commit a GitHub (6.7)
6. Creo TODO_DIA_7.md

---

**Generado por**: Narutrader
**Fecha**: 27/11/2025
**Para**: Día 6 (30/11/2025)
**Status**: Ready to execute

> "Day 6 fija los cimientos para Machine Learning. Cuando termine, el sistema será lo suficientemente inteligente para aprender patrones."

---

## 🔗 REFERENCIAS RÁPIDAS

- **R Issue**: Probable jsonlite segfault en Windows WSL
- **Parameter Guide**: https://en.wikipedia.org/wiki/Technical_analysis
- **ML Features**: Standard deviation, momentum, rate of change
- **Go Testing**: Built-in testing framework

---

⚠️ **IMPORTANTE**: Si alguna tarea falla durante Day 6, pauso y pido ayuda antes de continuar.

---

**¡Narutrader listo para Día 6! 🤖**
