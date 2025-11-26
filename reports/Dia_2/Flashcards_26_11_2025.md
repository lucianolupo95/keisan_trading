# 🎓 FLASHCARDS ANKI - DÍA 2

## Keisan Trading - Código Real y Decisiones Técnicas
**Generado**: 26/11/2025
**Total Flashcards**: 18
**Cobertura**: Integración Go-Python, Filtros de Trading, CSV Management

---

## FLASHCARD 1 - Go-Python Subprocess Execution

**Pregunta:**
¿Cuál es la forma correcta de ejecutar un script Python desde Go y capturar su output?

**Código:**
```go
func CallKaton(pythonPath string) (string, error) {
    cmd := exec.Command("python", pythonPath)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(output), nil
}
```

**Respuesta:**
Usar `exec.Command()` para crear un comando que ejecute Python, y `CombinedOutput()` para capturar tanto stdout como stderr. Esto permite que Go invoque subprocesos Python y procese sus respuestas. La función retorna el output completo como string y cualquier error de ejecución.

**Código Pregunta**: go-exec-001

**Notas:**
- `exec.Command()` es agnóstico del SO pero requiere que `python` esté en PATH
- `CombinedOutput()` espera a que el comando termine (bloqueante)
- Alternativa: `Output()` para solo stdout (ignora stderr)
- En Windows, usar "python" en vez de "python3"

---

## FLASHCARD 2 - Parseo de Respuestas Simples

**Pregunta:**
¿Cómo extraer un valor específico ("Katon OK") del output de un comando ejecutado?

**Código:**
```go
func ParseKatonResponse(output string) string {
    lines := strings.Split(output, "\n")
    for _, line := range lines {
        if strings.Contains(line, "Ping response:") {
            return strings.TrimSpace(strings.ReplaceAll(line, "Ping response:", ""))
        }
    }
    return "NO RESPONSE"
}
```

**Respuesta:**
Dividir el output en líneas usando `strings.Split()`, buscar la línea que contiene el patrón ("Ping response:"), y extraer el valor usando `ReplaceAll()` + `TrimSpace()`. Retornar un default ("NO RESPONSE") si no se encuentra el patrón.

**Código Pregunta**: go-parse-002

**Notas:**
- Este es un parseo simple y frágil; en producción usar regex o JSON
- `TrimSpace()` elimina espacios en blanco al inicio/final
- Necesario para extraer datos de output no estructurado
- Alternativa robusta: cambiar Python para retornar JSON

---

## FLASHCARD 3 - Lectura de CSV en Go

**Pregunta:**
¿Cuál es el proceso para leer un archivo CSV y convertirlo a una estructura Go?

**Código:**
```go
func ReadCSV(filename string) ([]Candle, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    // ... conversión a Candle
    return candles, nil
}
```

**Respuesta:**
1. Abrir el archivo con `os.Open()`
2. Crear un `csv.NewReader()`
3. Leer todos los registros con `ReadAll()`
4. Iterar sobre los registros y convertir a la estructura deseada
5. Usar `defer` para cerrar el archivo automáticamente

**Código Pregunta**: go-csv-003

**Notas:**
- `csv.NewReader()` es flexible para diferentes delimitadores
- `ReadAll()` carga todo en memoria (problema para archivos muy grandes)
- Alternativa: `Read()` para procesar línea a línea
- Manejo de errores es crítico (archivo no existe, formato inválido)

---

## FLASHCARD 4 - Conversión de Strings a Números en Go

**Pregunta:**
¿Cómo convertir un string a float64 e int64 en Go, manejando posibles errores?

**Código:**
```go
record := records[i]
open, _ := strconv.ParseFloat(record[1], 64)
close, _ := strconv.ParseFloat(record[2], 64)
volume, _ := strconv.ParseInt(record[5], 10, 64)
```

**Respuesta:**
Usar `strconv.ParseFloat()` para floats (segundo parámetro 64 = precisión double) y `strconv.ParseInt()` para enteros (segundo parámetro 10 = base decimal). En este caso, ignorar errores con `_`, pero en producción deberías manejarlos.

**Código Pregunta**: go-strconv-004

**Notas:**
- `ParseFloat(string, bitSize)` - bitSize 64 es IEEE 754 double
- `ParseInt(string, base, bitSize)` - base 10 es decimal
- ❌ Mala práctica: ignorar errores con `_`
- ✅ Mejor: `if err != nil { /* handle */ }`
- Considerar usar librerías como `encoding/json` para mapeo automático

---

## FLASHCARD 5 - Estructura de Datos para Candlestick

**Pregunta:**
¿Cuáles son los campos mínimos necesarios en una estructura Candle para análisis técnico básico?

**Código:**
```go
type Candle struct {
    Timestamp string
    Open      float64
    Close     float64
    High      float64
    Low       float64
    Volume    int64
}
```

**Respuesta:**
**OHLCV** (Open, High, Low, Close, Volume):
- **Open**: Precio de apertura
- **High**: Precio máximo del período
- **Low**: Precio mínimo del período
- **Close**: Precio de cierre
- **Volume**: Volumen transaccionado
- **Timestamp**: Para ordenar cronológicamente

Estos son suficientes para la mayoría de indicadores técnicos (MA, RSI, MACD, etc.).

**Código Pregunta**: go-candle-005

**Notas:**
- OHLCV es el estándar en trading
- Go utiliza tipos primitivos (float64, int64) para precisión
- Timestamp como string es flexible pero considera `time.Time` en producción
- Volume es int64 para soportar millones de transacciones

---

## FLASHCARD 6 - Filtro de Volumen en Trading

**Pregunta:**
¿Por qué es importante filtrar señales por volumen mínimo en trading?

**Código:**
```go
const minVolume int64 = 1300

if candle.Volume < minVolume {
    return "HOLD"  // Volumen insuficiente
}
```

**Respuesta:**
El volumen confirma la fortaleza de un movimiento de precio:
- **Bajo volumen**: Pocos traders, movimiento no confiable, riesgo de reversal
- **Alto volumen**: Muchos traders, movimiento confirmado, más probable que continúe
- Filtro previene **false positives** en velas débiles

Ejemplo: Si solo 10 traders hacen un movimiento vs 1000 traders, el segundo es mucho más confiable.

**Código Pregunta**: trading-volume-006

**Notas:**
- Threshold (1300) debe ajustarse según mercado/timeframe
- En BTC: 1000+ es razonable; en penny stocks: 10K+ necesario
- Volumen es especialmente importante en reversal patterns
- Importante: verificar volumen DURANTE el movimiento, no después

---

## FLASHCARD 7 - Filtro de Movimiento Porcentual

**Pregunta:**
¿Cómo calcular el movimiento porcentual entre open y close? ¿Por qué es importante?

**Código:**
```go
movePercent := ((candle.Close - candle.Open) / candle.Open) * 100

if candle.Close > candle.Open && movePercent >= minMovePercent {
    return "BUY"
}
```

**Respuesta:**
**Fórmula**: `((Close - Open) / Open) * 100`

**Ejemplo**: Open=1500, Close=1502
- Cambio = 1502 - 1500 = 2
- Porcentaje = (2 / 1500) * 100 = 0.133%

**Por qué es importante**:
- Normaliza movimientos (1% en BTC vs 1% en penny stock)
- Filtra movimientos insignificantes (0.01% = ruido)
- Threshold de 0.1% = movimiento mínimo significativo

**Código Pregunta**: trading-percent-007

**Notas:**
- Usar porcentaje en lugar de precio absoluto para comparabilidad
- Close > Open = movimiento alcista
- Close < Open = movimiento bajista
- 0.1% = umbral razonable; ajustar según estrategia
- Si se usa leverage, este umbral debería ser mayor

---

## FLASHCARD 8 - Lógica Combinada de Filtros

**Pregunta:**
¿Cuál es la lógica completa de GenerateSignal con múltiples filtros?

**Código:**
```go
func GenerateSignal(candle Candle) string {
    const minVolume int64 = 1300
    const minMovePercent float64 = 0.1

    // Filtro 1: Volumen
    if candle.Volume < minVolume {
        return "HOLD"
    }

    // Filtro 2: Movimiento porcentual
    movePercent := ((candle.Close - candle.Open) / candle.Open) * 100

    // Filtro 3: Dirección + Magnitud
    if candle.Close > candle.Open && movePercent >= minMovePercent {
        return "BUY"
    }
    return "HOLD"
}
```

**Respuesta:**
**3 Filtros en AND lógico**:
1. ✅ Volumen ≥ 1300 (confirma movimiento)
2. ✅ Close > Open (dirección alcista)
3. ✅ Movimiento ≥ 0.1% (magnitud mínima)

Solo si TODOS los filtros pasan → BUY
Cualquier filtro falla → HOLD (sin acción)

**Codigo Pregunta**: trading-logic-008

**Notas:**
- Filtros múltiples = señales más confiables pero menos frecuentes
- Trade-off: Calidad vs Frecuencia
- Test con 35 velas: 18 BUY, 17 HOLD (51% ratio)
- Considerar agregar más filtros (volatility, trend) en futuro

---

## FLASHCARD 9 - Expansión de Dataset de Trading

**Pregunta:**
¿Por qué es importante expandir de 3 a 35 velas en el dataset de testing?

**Respuesta:**
**Razones**:
1. **Representatividad**: 3 velas = accidental, 35 velas = patrón real
2. **Testing de filtros**: Ver cómo se comportan los filtros en variedad de condiciones
3. **Detección de edge cases**: Casos raros que solo aparecen con más datos
4. **Estadística**: 35 muestras da confianza en results (3 es insuficiente)

**Datos reales vs Fake**:
- ❌ 3 velas fake: Todos "lucky" cases
- ✅ 35 velas realistas: Mezcla de BUY/HOLD, volatility variada

**Código Pregunta**: testing-dataset-009

**Notas:**
- Dataset expandido permite benchmarking más robusto
- En ML, se usan miles/millones de samples
- Para trading: 30-100 velas por timeframe es start razonable
- Incluir datos de volatility baja, media, alta para robustez

---

## FLASHCARD 10 - Rango de Precios Realistas

**Pregunta:**
¿Cuál es un rango realista para BTC en el CSV, y por qué se usa ese rango?

**Código:**
```csv
2025-11-25T09:00:00Z,1500.0,1501.5,1502.0,1499.5,1000
...
2025-11-25T17:30:00Z,1516.0,1518.5,1519.0,1515.0,3000
```

**Respuesta:**
**Rango**: 1500.0 - 1518.5 (18.5 pips de movimiento en 8.5 horas)

**Por qué**:
- 1500-1550 es realista para BTC (comparar con precio actual)
- 18.5 pips = ~1.2% movimiento total = realista para mercado normal
- Evita precios fake (10,000+ o 0.01)
- Volatility gradual (no saltos abruptos) = más representativo

**En producción**: Usar datos históricos reales (exchanges, APIs)

**Código Pregunta**: market-prices-010

**Notas:**
- BTC típicamente: 1000-500,000 (últimos años)
- 1500 es razonable para simulación
- En backtesting: SIEMPRE usar datos históricos reales
- Gaps de precio sin volumen = problema en datos

---

## FLASHCARD 11 - Progresión de Volúmenes

**Pregunta:**
¿Por qué los volúmenes en el CSV van de 1000 a 3000 (progresivo)?

**Respuesta:**
**Razones para progresión**:
1. **Realismo de mercado**: Hora 9 (mercado abre) = bajo volumen; Hora 14-16 = pico volumen
2. **Testing de filtros**: Algunos datos con bajo volumen fallan filtro; con alto volumen pasan
3. **Identificar umbrales**: Ver exactamente dónde cae minVolume=1300

**Ejemplo**:
- Vela 1: Volume=1000 → HOLD (volume insuficiente)
- Vela 2: Volume=1200 → HOLD (aún insuficiente)
- Vela 3: Volume=1500 → Posible BUY (si otros filtros pasan)

**Código Pregunta**: volume-progression-011

**Notas:**
- Volatility también varía (high-low range es mayor con volumen alto)
- En datos reales: volumen sigue patrón intra-day (bajo en US market open, alto en NY close)
- Threshold 1300 elegido entre 1000 y 1500 por razón

---

## FLASHCARD 12 - CSV Headers y Parseo

**Pregunta:**
¿Por qué es importante incluir headers en CSV y cómo se manejan en el parseo?

**Código:**
```go
reader := csv.NewReader(file)
records, err := reader.ReadAll()

// Saltar header
for i := 1; i < len(records); i++ {  // Comienza en índice 1
    record := records[i]
    open, _ := strconv.ParseFloat(record[1], 64)
    // record[0] = timestamp
    // record[1] = open
    // record[2] = close
    // ...
}
```

**Respuesta:**
**Headers** (`timestamp,open,close,high,low,volume`):
- ✅ Documentan qué es cada columna
- ✅ Previenen bugs si se reordenan columnas
- ✅ Hacen legible el CSV para humanos

**Parseo**:
- Saltar primer row (índice 0 = headers)
- Comenzar iteración desde índice 1
- Indexar por posición: `record[0]`, `record[1]`, etc.

**Código Pregunta**: csv-headers-012

**Notas:**
- ❌ Sin headers: Fácil de confundir columnas
- ✅ Con headers: Self-documenting
- Alternativa: Usar librerías que mapean automáticamente (encoding/json)
- En Python: pandas automáticamente parsea headers

---

## FLASHCARD 13 - Integración Inter-Lenguajes

**Pregunta:**
¿Cuáles son las ventajas de poder integrar Go con Python?

**Código:**
```go
// Go invoca Python
output, _ := CallKaton("../katon-python/ping.py")
fmt.Println("Response:", output)
```

**Respuesta:**
**Ventajas**:
- **Go**: Performance, concurrency, ejecución simple
- **Python**: Librería rich (pandas, numpy, ML), prototipado rápido

**Combinación permite**:
1. Go como orquestador/scheduler principal
2. Python para cálculos complejos (análisis técnico, ML)
3. Máximo valor de cada lenguaje

**Ejemplo arquitectura Keisan**:
- Fuuton (Go) = Motor principal + scheduler
- Katon (Python) = Feature engineering + análisis
- Suiton (R) = Estadística + back-test
- Doton (C#) = Integración con broker

**Código Pregunta**: inter-lang-013

**Notas:**
- Overhead: exec subprocess es más lento que llamada directa
- En producción: Considerar gRPC, REST API, o sockets
- Ventaja: Desarrollo independiente en lenguajes distintos

---

## FLASHCARD 14 - Testeo de Integración Go-Python

**Pregunta:**
¿Cuál fue el output esperado vs actual del test de integración Go-Python?

**Esperado:**
```
Fuuton activo
Calling Katon (Python)...
Katon response: Katon OK
```

**Actual:**
```
Fuuton activo
Calling Katon (Python)...

Katon response: Katon OK

Katon output (full):
Katon activo
Ping response: Katon OK
simple_feature(5) = 10
```

**Respuesta:**
✅ **Integración exitosa**:
- Go ejecutó Python correctamente
- Capturó el output
- Parseó la respuesta
- Mostró output completo para debugging

**Problemas encontrados**:
- Primera intención usó "python3" → Error 9009 en Windows
- Solución: cambiar a "python" (en PATH estándar)

**Código Pregunta**: integration-test-014

**Notas:**
- Windows usa "python", Linux usa "python3"
- Error 9009 = "command not found"
- Importante: Test simple antes de usar en lógica crítica

---

## FLASHCARD 15 - Ratio BUY vs HOLD en 35 Velas

**Pregunta:**
¿Cuál fue el resultado de aplicar GenerateSignal a 35 velas? ¿Es balanceado?

**Respuesta:**
**Resultado**:
- BUY signals: 18/35 = 51.4%
- HOLD signals: 17/35 = 48.6%

**Análisis**:
- ✅ Muy balanceado (cercano a 50/50)
- ✅ Indica que filtros funcionan correctamente
- ⚠️ Ratio 50/50 puede ser bajista (buscamos más BUY en mercado alcista)

**Interpretación**:
- Datos usados: Mercado neutral (pequeño uptrend)
- Filtros no son demasiado restrictivos (sí generan signals)
- Pero tampoco son permisivos (no todas son BUY)

**En producción**:
- Backtesting histórico: medir Sharpe ratio, max drawdown
- Ratio esperado: Varía según timeframe y volatility

**Código Pregunta**: buy-hold-ratio-015

**Notas:**
- 51/49 es sospechosamente perfecto (data fake, pero intencional)
- Datos reales: Varía mucho según condiciones de mercado
- Si ratio es 90% BUY: Quizás filtros muy permisivos
- Si ratio es 5% BUY: Quizás filtros muy restrictivos

---

## FLASHCARD 16 - Csv vs Base de Datos para Trading Data

**Pregunta:**
¿Cuáles son las limitaciones de usar CSV vs una base de datos real?

**CSV (Usado en Día 2)**:
- ✅ Simple, portable
- ❌ Lento para búsquedas
- ❌ Sin índices
- ❌ Sin concurrent access
- ❌ Sin ACID garantías

**Base de Datos (Recomendado para producción)**:
- ✅ Índices → búsquedas rápidas
- ✅ Queries complejas
- ✅ Concurrent access
- ✅ Backup/recovery
- ❌ Más complejidad

**Código Pregunta**: csv-vs-db-016

**Notas:**
- Para 35 velas: CSV está bien
- Para 100K+ velas: Base de datos necesaria (TimescaleDB, InfluxDB)
- En producción: MariaDB, PostgreSQL, o time-series DB (InfluxDB)
- CSV útil para: Testing, exportación, reportes

---

## FLASHCARD 17 - Error Handling Faltante

**Pregunta:**
¿Qué error handling falta en el código actual de CSV parsing?

**Código Actual**:
```go
open, _ := strconv.ParseFloat(record[1], 64)  // ❌ Error ignorado
close, _ := strconv.ParseFloat(record[2], 64)
```

**Mejor Práctica**:
```go
open, err := strconv.ParseFloat(record[1], 64)
if err != nil {
    log.Printf("Error parsing open price in row %d: %v", i, err)
    continue  // o return error
}
```

**Problemas**:
- ❌ Si CSV corrupto (non-numeric) → programa silenciosamente toma 0
- ❌ Difícil de debuggear
- ❌ Datos inválidos = resultados incorrectos

**Código Pregunta**: error-handling-017

**Notas:**
- Go idiom: "Errors are values, handle them"
- Ignorar errores con `_` = Bad practice
- En producción: Logging + alerting crítico
- Considerar: Validación de CSV al inicio

---

## FLASHCARD 18 - Próximos Pasos Arquitectónicos

**Pregunta:**
¿Cuáles son los próximos módulos/mejoras recomendadas para Día 3+?

**Respuesta:**
**Prioritario (Día 3-4)**:
1. **Shoton** (Python avanzado): Análisis estadístico real (Bollinger Bands, MACD)
2. **Backtesting framework**: Calcular Sharpe, Max Drawdown, Win Rate
3. **Bridge multi-módulo**: Orquestación de Fuuton, Katon, Suiton, Doton

**Importante (Día 5-7)**:
1. Instalar R + habilitar Suiton
2. Instalar .NET + habilitar Doton
3. Integración con broker API (mock primero)

**Largo plazo**:
1. Machine Learning para predicción
2. Real-time trading execution
3. Risk management avanzado
4. Dashboard/UI para monitoreo

**Código Pregunta**: roadmap-018

**Notas:**
- Arquitectura es sound; expandible a múltiples lenguajes
- Go como orquestador es buena decisión
- Python + ML es próximo step natural
- Considerar: Trading live simulado antes de real

---

## 📊 RESUMEN DE FLASHCARDS

| Card | Tema | Dificultad | Tipo |
|------|------|-----------|------|
| 1 | Subprocess Go | Medio | Code |
| 2 | String Parsing | Medio | Code |
| 3 | CSV Reading | Medio | Code |
| 4 | Type Conversion | Fácil | Code |
| 5 | Data Structures | Fácil | Theory |
| 6 | Volume Filtering | Medio | Trading |
| 7 | % Calculation | Fácil | Trading |
| 8 | Logic Combination | Difícil | Code |
| 9 | Dataset Expansion | Medio | Testing |
| 10 | Price Realism | Fácil | Trading |
| 11 | Volume Progression | Medio | Data |
| 12 | CSV Headers | Fácil | Code |
| 13 | Inter-Language Integration | Difícil | Architecture |
| 14 | Integration Testing | Medio | Testing |
| 15 | Signal Statistics | Medio | Trading |
| 16 | Storage Solutions | Difícil | Architecture |
| 17 | Error Handling | Difícil | Best Practices |
| 18 | Roadmap | Difícil | Planning |

---

## 🎯 DISTRIBUCIÓN DE TEMAS

- **Go/Programming**: 45% (Subprocess, CSV, Parsing, Type Conversion)
- **Trading/Finance**: 33% (Filters, %calc, Signals, Prices)
- **Architecture/Design**: 22% (Integration, Storage, Error Handling)

---

**Generado por**: Narutrader 🤖
**Fecha**: 26/11/2025
**Total Flashcards**: 18
**Formato**: Markdown → Anki-compatible
