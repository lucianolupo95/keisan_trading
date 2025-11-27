# 🎓 FLASHCARDS - DAY 5 TECHNICAL INDICATORS

**Total Flashcards**: 25
**Topic**: MA20, RSI, MACD, IPC, Signal Integration
**Review**: Diario para retención

---

## 📚 MOVING AVERAGE (MA)

### Card 1: ¿Qué es una Moving Average?
**Q**: Define Moving Average (MA) en trading
**A**: Promedio de precios en un período, suaviza fluctuaciones. Tendencia del precio.
- MA20 = promedio últimas 20 velas
- MA50 = promedio últimas 50 velas
- Mayores períodos = más suave

### Card 2: ¿Cuál es la diferencia SMA vs EMA?
**Q**: Diferencias entre Simple MA y Exponential MA
**A**:
- **SMA**: Peso igual a todos los precios. Ejemplo: (P1+P2+...+P20)/20
- **EMA**: Pesa más los precios recientes. Fórmula: EMA = P*multiplier + prevEMA*(1-multiplier)
- EMA responde más rápido a cambios

### Card 3: ¿Cómo usar MA20 como filtro de compra?
**Q**: Aplicación de MA20 en señales de compra
**A**:
- Si precio > MA20 → uptrend (bullish)
- Si precio < MA20 → downtrend (bearish)
- Comprar cuando precio cruza sobre MA20
- Vender cuando precio cae bajo MA20

### Card 4: ¿Cuál es el período típico de MA?
**Q**: Cuáles son los períodos más comunes?
**A**:
- MA20 = 20 días = corto plazo
- MA50 = 50 días = mediano plazo
- MA200 = 200 días = largo plazo
- Day Trading: MA5-MA10

### Card 5: ¿Cómo calcular MA20?
**Q**: Pseudocódigo para calcular MA20
**A**:
```go
func CalculateMA20(prices []float64) float64 {
    if len(prices) < 20 {
        return average(prices) // todas disponibles
    }
    return average(prices[len-20:len]) // últimas 20
}
```

---

## 📈 RSI (RELATIVE STRENGTH INDEX)

### Card 6: ¿Qué es RSI?
**Q**: Define RSI en trading
**A**: Oscillator de momentum que mide fuerza de cambios de precio
- Rango: 0-100
- Indica velocidad y magnitud de cambios
- Detecta condiciones extremas

### Card 7: ¿Cuáles son los niveles críticos de RSI?
**Q**: Qué significan los valores de RSI
**A**:
- RSI < 30: Oversold (posible compra)
- RSI 30-70: Neutral
- RSI > 70: Overbought (posible venta)
- RSI = 50: Neutral

### Card 8: ¿Cómo se calcula RSI?
**Q**: Fórmula de RSI
**A**:
- RS = avgGain / avgLoss (últimas 14 velas)
- RSI = 100 - (100 / (1 + RS))
- Período estándar = 14

### Card 9: ¿Cuándo comprar según RSI?
**Q**: Señales de compra con RSI
**A**:
- Mejor oportunidad cuando RSI < 30 (oversold)
- Divergencia bajista en RSI = venta próxima
- Divergencia alcista en RSI = compra próxima
- No es 100% preciso, usar con otros indicadores

### Card 10: ¿Qué es divergencia en RSI?
**Q**: Explica divergencia en RSI
**A**:
- Precio hace nuevo low, pero RSI no (alcista)
- Precio hace nuevo high, pero RSI no (bajista)
- Señal de cambio de tendencia próximo
- Muy preciso cuando se confirma

---

## 📊 MACD (MOVING AVERAGE CONVERGENCE DIVERGENCE)

### Card 11: ¿Qué es MACD?
**Q**: Define MACD
**A**: Indicador de tendencia que combina 2 EMAs
- Detecta cambios en momentum y tendencia
- Mejor para trends, no para sideways
- Componentes: MACD line, Signal line, Histogram

### Card 12: ¿Cuáles son los componentes de MACD?
**Q**: Qué calcula MACD
**A**:
- **MACD Line**: EMA12 - EMA26 (momentum)
- **Signal Line**: EMA9 del MACD (promedio suavizado)
- **Histogram**: MACD - Signal (diferencia)
- Si Histogram > 0 y MACD > Signal = bullish

### Card 13: ¿Cómo genera señales MACD?
**Q**: Señales de compra/venta con MACD
**A**:
- **Compra**: Cuando MACD cruza sobre Signal line (bullish)
- **Compra**: Cuando MACD > Signal line + RSI < 50
- **Venta**: Cuando MACD cruza bajo Signal line (bearish)
- **Confirmación**: Histogram > 0 (fuerza de tendencia)

### Card 14: ¿Diferencia MACD Bullish vs Bearish?
**Q**: Cómo detectar dirección en MACD
**A**:
- **Bullish**: MACD > Signal line (arriba)
- **Bullish**: Histogram positivo y creciendo
- **Bearish**: MACD < Signal line (abajo)
- **Bearish**: Histogram negativo y bajando

### Card 15: ¿Cuándo es mejor MACD?
**Q**: Condiciones óptimas para usar MACD
**A**:
- Excelente en mercados trending
- Pobre en mercados laterales (sideways)
- Mejor con timeframes largos (diarios)
- Peor con scalping (muy frecuente)

---

## 🔗 IPC (INTER-PROCESS COMMUNICATION)

### Card 16: ¿Qué es IPC en el contexto de Keisan?
**Q**: IPC en nuestro sistema de trading
**A**: Comunicación entre Go y R:
- Go envía precios como JSON
- R analiza estadísticas
- R retorna resultados como JSON
- Go parsea y usa resultados

### Card 17: ¿Cómo serializar datos para IPC?
**Q**: Qué es serialización y por qué importa
**A**:
- Convertir datos a formato universal (JSON)
- Go: `json.Marshal(data)`
- R: `jsonlite::toJSON(data)`
- Permite comunicación entre lenguajes

### Card 18: ¿Qué paso si R no está disponible?
**Q**: Estrategia de fallback en Keisan
**A**:
- Intenta usar R IPC primero
- Si falla, automáticamente usa análisis local (Go)
- Graceful degradation = sistema sigue funcionando
- No hay pérdida de funcionalidad core

### Card 19: ¿Por qué JSON y no binario?
**Q**: Por qué JSON es mejor que formato binario
**A**:
- JSON es legible por humanos
- JSON es estándar universal
- JSON no requiere parsing binario complejo
- Fácil de debug (ver exactamente qué se envía)

### Card 20: ¿Cómo ejecutar subprocess en Go?
**Q**: Sintaxis para ejecutar Rscript desde Go
**A**:
```go
cmd := exec.Command("Rscript", "script.R", "args...")
output, err := cmd.CombinedOutput()
if err != nil {
    // manejar error
}
```

---

## 🎯 CONFIDENCE SCORING & SIGNALS

### Card 21: ¿Cómo calcula confianza el sistema Day 5?
**Q**: Factores en confidence scoring
**A**:
- Base: 0.5
- Statistical: +0.15 cada (normal, correlation, volatility)
- Technical: +0.15 (RSI), +0.1 (MA20), +0.1 (MACD)
- Penalidades: -0.1 (overbought), -0.05 (otros)

### Card 22: ¿Cuáles son los 4 niveles de signal?
**Q**: Niveles de confianza en Day 5
**A**:
- Confidence > 0.8: "BUY (Very High)" 💪💪
- Confidence > 0.6: "BUY (High)" 💪
- Confidence > 0.4: "BUY (Medium)" 🤔
- Confidence ≤ 0.4: "HOLD" 😐

### Card 23: ¿Por qué combinar múltiples indicadores?
**Q**: Ventajas de usar MA + RSI + MACD juntos
**A**:
- Reduce false signals
- Confirma tendencias
- MA = tendencia, RSI = momentum, MACD = cambio
- Mayor confianza en decisiones

### Card 24: ¿Cómo evitar falsos positivos?
**Q**: Técnicas para filtrar bad signals
**A**:
- Usar múltiples confirmaciones
- Requerer RSI dentro de rango
- Verificar MA alignment
- Usar MACD para confirmar dirección
- Mejor: 2-3 confirmaciones antes de actuar

### Card 25: ¿Cuándo usar qué indicador?
**Q**: Aplicación de cada indicador
**A**:
- **MA20**: Detectar uptrend/downtrend claro
- **RSI < 30**: Buscar dips para compra
- **MACD > Signal**: Confirmar nuevo uptrend
- **Todos juntos**: Máxima confianza

---

## 🎓 RESUMEN RÁPIDO

### MA20 Recap
- Suaviza precio
- > MA20 = bullish
- < MA20 = bearish

### RSI Recap
- < 30 = oversold (compra)
- > 70 = overbought (venta)
- 0-100 range

### MACD Recap
- MACD > Signal = bullish
- MACD < Signal = bearish
- Histogram = fuerza

### Confianza Recap
- Múltiples factores = mejor decisión
- 4 niveles > 2 niveles
- Combine para máxima precisión

---

## 📝 NOTAS DE ESTUDIO

### Importante Recordar
1. MA es trend-follower, no predictor
2. RSI es oscillator, trabaja bien en sideways
3. MACD es trend indicator, falla en laterales
4. Nunca uses UN solo indicador
5. Confía en las combinaciones

### Ejercicios Prácticos
1. Calcula MA20 manualmente para 20 precios
2. Interpreta RSI=25 vs RSI=75
3. Dibuja MACD histogram creciendo/cayendo
4. Combina 2 indicadores para generar señal
5. Identifica falsos positivos en chart real

### Conceptos Clave
- **Oversold**: Precio muy bajo, bounce próximo
- **Overbought**: Precio muy alto, drop próximo
- **Bullish**: Dirección alcista (subida)
- **Bearish**: Dirección bajista (caída)
- **Divergence**: Precio vs indicador no alinean

---

## 🔄 REVIEW SCHEDULE

**Recomendación**:
- **Hoy**: Leer todas
- **Mañana**: Review 1-10
- **Día 3**: Review 11-20
- **Día 4**: Review 21-25
- **Día 5**: Test completo
- **Semanal**: Repaso general

---

## ✨ CONNECTIONS TO TRADING

### Real World Application
1. **Entrada**: RSI < 30 + MA20 bullish = COMPRA
2. **Confirmación**: MACD > Signal = sí, compra
3. **Salida**: RSI > 70 o MACD < Signal = VENDA
4. **Gestión Riesgo**: Stop-loss debajo de MA20

### Ejemplo Trade
```
Precio: $100
MA20: $99 (arriba de precio = bajista)
RSI: 28 (oversold = compra potencial)
MACD: MACD < Signal pero subiendo = cambio próximo

Decisión: ESPERAR
Razón: MA20 bearish cancela RSI oversold signal
Acción: Esperar MACD cruce + MA20 cruce
```

---

**Status**: ✅ COMPLETE
**Created**: 27-11-2025
**For**: Daily Review & Mastery
**Next**: Apply en backtest real

---

¡Memoriza estas 25 flashcards y serás 10x mejor en trading técnico! 💪
