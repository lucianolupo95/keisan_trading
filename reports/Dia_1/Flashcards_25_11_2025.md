# 🎓 FLASHCARDS ANKI - DÍA 1
## Keisan Trading - Código Real y Decisiones Técnicas

Preguntas sobre el código real, decisiones técnicas y herramientas utilizadas.

---

## FLASHCARD 1 - ESTRUCTURA CANDLE

**Pregunta:**
En `fuuton-go/main.go`, ¿cuál es la estructura `Candle` y por qué elegimos esos tipos de datos?

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
- **Timestamp (string)**: Guarda el momento de la vela (ej: "2025-11-25T09:00:00Z")
- **Open, Close, High, Low (float64)**: Precios con decimales (ej: 101.50)
- **Volume (int64)**: Cantidad de transacciones (números enteros grandes)

**Decisión técnica**: Usamos `float64` para precios (suficientemente preciso) e `int64` para volumen (evita overflow). `Timestamp` es string para mantenerlo flexible.

**Código pregunta:**
```
candle_struct_types
```

**Notas:**
En trading, la precisión en precios es crítica pero no necesitamos más que float64. El volumen es entero porque no hay "media transacción".

---

## FLASHCARD 2 - DEFER STATEMENT

**Pregunta:**
¿Qué hace `defer file.Close()` en la línea 32 de `main.go`?

**Código:**
```go
file, err := os.Open(filename)
if err != nil {
    return nil, err
}
defer file.Close()
```

**Respuesta:**
`defer` garantiza que `file.Close()` se ejecute al FINAL de la función, sin importar si hay errores en el medio. Es como un "hazlo al irte".

**Decisión técnica**: Sin `defer`, si hay un error en el parsing del CSV, el archivo quedaría abierto (memory leak). Con `defer`, siempre se cierra, incluso si retornamos un error.

**Código pregunta:**
```
defer_file_close
```

**Notas:**
`defer` es una característica única de Go. En Python usarías `with open()`. En C# usarías `using`. La idea es igual: limpiar recursos automáticamente.

---

## FLASHCARD 3 - IGNORAR ERRORES CON _

**Pregunta:**
En la línea 44 de `main.go`, ¿por qué usamos `_` en `open, _ := strconv.ParseFloat(...)`?

**Código:**
```go
open, _ := strconv.ParseFloat(record[1], 64)
```

**Respuesta:**
`_` es el "blank identifier". Ignora el error que retorna `ParseFloat`. Si el CSV es válido (lo es), ignoramos el error.

**Decisión técnica**: Es un trade-off. En código PRODUCTION deberíamos validar:
```go
open, err := strconv.ParseFloat(record[1], 64)
if err != nil {
    log.Fatalf("Error parsing precio: %v", err)
}
```
Pero para Day 1 (CSV de test), confiamos en los datos. En futuro lo haremos robusto.

**Código pregunta:**
```
blank_identifier_underscore
```

**Notas:**
Esta es una decisión consciente de "mantener simple para Day 1". No es ideal para production, pero funciona para prototipado rápido.

---

## FLASHCARD 4 - SLICES Y APPEND

**Pregunta:**
¿Qué hace `candles = append(candles, candle)` en la línea 58?

**Código:**
```go
var candles []Candle
// ...
for i := 1; i < len(records); i++ {
    // ... crear candle ...
    candles = append(candles, candle)
}
return candles, nil
```

**Respuesta:**
- `var candles []Candle` crea un slice vacío
- `append()` agrega el nuevo `candle` al final del slice
- Al final, retornamos todos los candles acumulados

Es como una lista dinámica que crece conforme leemos el CSV.

**Decisión técnica**: Usamos slices (no arrays) porque no sabemos cuántas velas hay. Los slices crecen dinámicamente.

**Código pregunta:**
```
append_slice_grow
```

**Notas:**
En Go, `append` no modifica el slice original, retorna un nuevo slice. Por eso reasignamos: `candles = append(...)`.

---

## FLASHCARD 5 - GENERACIÓN DE SEÑAL SIMPLE

**Pregunta:**
¿Cuál es la lógica de `GenerateSignal()` en la línea 64?

**Código:**
```go
func GenerateSignal(candle Candle) string {
    if candle.Close > candle.Open {
        return "BUY"
    }
    return "HOLD"
}
```

**Respuesta:**
Comparamos el precio de cierre vs apertura:
- Si `Close > Open`: vela alcista → señal **BUY**
- Si no: → señal **HOLD** (no comprar)

**Decisión técnica**: Es la lógica más simple posible para un primer día. Una vela alcista es bullish. En futuro agregaremos:
- Media móvil (MA)
- Índice de fuerza relativa (RSI)
- Filtro de volumen
- Filtro de spread

**Código pregunta:**
```
signal_generation_close_open
```

**Notas:**
Esta es la base. En Day 2+ haremos más sofisticado. Pero un MVP funcional requiere empezar así.

---

## FLASHCARD 6 - RANGE LOOP

**Pregunta:**
¿Qué hace el `for` en la línea 87?

**Código:**
```go
for i, candle := range candles {
    signal := GenerateSignal(candle)
    fmt.Printf("Vela %d - ... Signal: %s\n", i+1, ... signal)
}
```

**Respuesta:**
- `range candles` itera sobre cada elemento del slice
- `i` es el índice (0, 1, 2, ...)
- `candle` es el valor actual
- Procesamos cada vela y mostramos su señal

**Decisión técnica**: Usamos `range` (más idiomático en Go) en lugar de índices manuales. Es más legible y menos propenso a off-by-one errors.

**Código pregunta:**
```
range_loop_idiom
```

**Notas:**
En Python sería `for i, candle in enumerate(candles)`. En Go es `for i, candle := range candles`. La idea es igual.

---

## FLASHCARD 7 - PYTHON: IF __NAME__ == "__MAIN__"

**Pregunta:**
¿Por qué en `katon-python/ping.py` (línea 14) hay `if __name__ == "__main__"`?

**Código:**
```python
def ping():
    return "Katon OK"

def simple_feature(x):
    return x * 2

if __name__ == "__main__":
    print("Katon activo")
    response = ping()
    print(f"Ping response: {response}")
```

**Respuesta:**
`if __name__ == "__main__"` significa: "ejecuta esto SOLO si corro el archivo directamente".

Si otro script hace `from ping import simple_feature`, el bloque NO se ejecuta (no verías "Katon activo").

**Decisión técnica**: Permite que las funciones se importen sin side effects. Las definiciones están disponibles, pero el print() no se ejecuta.

**Código pregunta:**
```
python_main_guard
```

**Notas:**
Es una convención Python. Go no la necesita porque tiene un único punto de entrada (`main()`).

---

## FLASHCARD 8 - F-STRING EN PYTHON

**Pregunta:**
¿Qué es `f"simple_feature({test_value}) = {result}"` en la línea 22?

**Código:**
```python
test_value = 5
result = simple_feature(test_value)
print(f"simple_feature({test_value}) = {result}")
```

**Respuesta:**
Es una f-string (formatted string literal). Interpola variables directamente en la cadena usando `{variable}`.

Sin f-string:
```python
print("simple_feature(" + str(test_value) + ") = " + str(result))
```

Con f-string: más limpio y legible.

**Decisión técnica**: F-strings (Python 3.6+) son el estándar moderno. Más legibles que .format() o %.

**Código pregunta:**
```
fstring_interpolation
```

**Notas:**
Es azúcar sintáctico. Hace el código más legible sin cambiar funcionamiento.

---

## FLASHCARD 9 - R: FUNCTION SYNTAX

**Pregunta:**
En `suiton-r/ping.R`, ¿cuál es la sintaxis de una función en R?

**Código:**
```r
simple_stat <- function(x) {
  return(mean(x))
}

test_values <- c(10, 20, 30)
result <- simple_stat(test_values)
```

**Respuesta:**
- `<-` es el operador de asignación en R
- `function(x) { ... }` define la función
- `return(...)` retorna el valor
- `c(10, 20, 30)` crea un vector (array)

`simple_stat` calcula la media de los valores.

**Decisión técnica**: R es vectorizado. `mean()` es nativa. No necesitamos loops.

**Código pregunta:**
```
r_function_syntax
```

**Notas:**
En Go/Python usarías loops. En R, `mean(c(...))` es idiomatic y rápido.

---

## FLASHCARD 10 - CSV STRUCTURE

**Pregunta:**
¿Cuál es la estructura del CSV en `fuuton-go/data.csv` y por qué ese orden?

**Código CSV:**
```csv
timestamp,open,close,high,low,volume
2025-11-25T09:00:00Z,100.0,101.5,102.0,99.5,1000
2025-11-25T09:15:00Z,101.5,102.0,103.0,101.0,1200
2025-11-25T09:30:00Z,102.0,103.5,104.0,101.5,1500
```

**Código Go (línea 42-48):**
```go
for i := 1; i < len(records); i++ {
    open, _ := strconv.ParseFloat(record[1], 64)      // índice 1
    close, _ := strconv.ParseFloat(record[2], 64)     // índice 2
    high, _ := strconv.ParseFloat(record[3], 64)      // índice 3
    low, _ := strconv.ParseFloat(record[4], 64)       // índice 4
    volume, _ := strconv.ParseInt(record[5], 10, 64)  // índice 5
}
```

**Respuesta:**
El CSV tiene columnas en orden: timestamp (0), open (1), close (2), high (3), low (4), volume (5).

El código accede por índice: `record[1]` es open, `record[2]` es close, etc.

**Decisión técnica**: Ese orden es OHLCV estándar en finanzas. Si cambias el CSV, cambia el índice también.

**Código pregunta:**
```
csv_column_order
```

**Notas:**
En futuro, podríamos parsear dinámicamente los headers. Pero por ahora (Day 1), asumimos el orden correcto.

---

## FLASHCARD 11 - ERROR HANDLING EN GO

**Pregunta:**
¿Por qué en Go retornamos `([]Candle, error)` en la línea 27? ¿Cómo se maneja en main()?

**Código:**
```go
func ReadCSV(filename string) ([]Candle, error) {  // retorna dos valores
    file, err := os.Open(filename)
    if err != nil {
        return nil, err  // retorna nil + error
    }
    // ...
    return candles, nil  // retorna candles + nil error
}

// En main():
candles, err := ReadCSV("data.csv")
if err != nil {
    log.Fatalf("Error al leer CSV: %v", err)
}
```

**Respuesta:**
Go no tiene exceptions. Retorna múltiples valores: el resultado + un error.
- Si todo OK: `(candles, nil)`
- Si falla: `(nil, error)`

En main(), verificamos con `if err != nil`.

**Decisión técnica**: Es explícito. Cada función que puede fallar debe retornar error. No hay sorpresas.

**Código pregunta:**
```
go_error_handling
```

**Notas:**
En Python usarías try/except. En Go es `if err != nil`. Ambos funcionan, Go es solo más explícito.

---

## FLASHCARD 12 - TIPOS DE DATOS: FLOAT64 VS INT64

**Pregunta:**
¿Por qué en Candle usamos `float64` para precios e `int64` para volumen?

**Código:**
```go
type Candle struct {
    Open      float64   // precio
    Close     float64   // precio
    High      float64   // precio
    Low       float64   // precio
    Volume    int64     // cantidad
}
```

**Respuesta:**
- **float64**: Precios pueden tener decimales (ej: 101.50, 1500.25). float64 es 64-bit con precisión decimal.
- **int64**: Volumen es cantidad de transacciones. No hay "media transacción". int64 es entero 64-bit (rango: -9 billones a +9 billones).

**Decisión técnica**: Trade-off de memoria vs precisión. int64 para volumen es suficiente (nunca veremos 9 billones de transacciones en 15 min). float64 para precios es estándar en finanzas.

**Código pregunta:**
```
float64_vs_int64
```

**Notas:**
Podrías usar float64 para volumen también, pero sería desperdicio. Podrías usar float32 para precios para ahorrar memoria, pero pierdes precisión.

---

## FLASHCARD 13 - LOG.FATALF

**Pregunta:**
¿Qué diferencia hay entre `log.Fatalf()` (línea 80) y `fmt.Printf()`?

**Código:**
```go
candles, err := ReadCSV("data.csv")
if err != nil {
    log.Fatalf("Error al leer CSV: %v", err)  // FATAL
}

fmt.Printf("Velas leídas: %d\n", len(candles))  // INFO
```

**Respuesta:**
- `log.Fatalf()`: Imprime el mensaje + **detiene el programa** (exit code 1)
- `fmt.Printf()`: Solo imprime, continúa ejecutando

Use `Fatalf` para errores irrecuperables (no puedo continuar sin datos).

**Decisión técnica**: Si no hay CSV, el programa no tiene razón de continuar. Mejor fallar rápido y explícitamente.

**Código pregunta:**
```
log_fatalf_vs_printf
```

**Notas:**
En Go, fallar temprano es mejor que silenciosamente tener datos vacíos.

---

## FLASHCARD 14 - SPRINTF EN R

**Pregunta:**
¿Qué hace `sprintf()` en la línea 15 de `ping.R`?

**Código:**
```r
cat(sprintf("Ping response: %s\n", response))
```

**Respuesta:**
`sprintf()` formatea una cadena (como .format() en Python o fmt.Sprintf en Go).
- `%s` es placeholder para string
- `\n` es salto de línea

Retorna la cadena formateada que `cat()` imprime.

**Decisión técnica**: R también tiene strings, pero `sprintf` + `cat` es la forma idiomática para output formateado.

**Código pregunta:**
```
r_sprintf_format
```

**Notas:**
Equivalente a:
- Python: `print(f"Ping response: {response}")`
- Go: `fmt.Printf("Ping response: %s\n", response)`

---

## FLASHCARD 15 - MEAN() EN R

**Pregunta:**
En `suiton-r/ping.R`, ¿qué calcula `mean(c(10, 20, 30))`?

**Código:**
```r
simple_stat <- function(x) {
  return(mean(x))
}

test_values <- c(10, 20, 30)
result <- simple_stat(test_values)
cat(sprintf("simple_stat(c(10,20,30)) = %f\n", result))
```

**Respuesta:**
`mean(x)` calcula la media aritmética:
- (10 + 20 + 30) / 3 = 60 / 3 = **20**

`c(10, 20, 30)` crea un vector con 3 elementos.

**Decisión técnica**: `mean()` es built-in en R. No necesitamos loops. R es vectorizado.

**Código pregunta:**
```
r_mean_function
```

**Notas:**
En futuro, usaremos `mean()` en series de precios (media móvil). Este es el primer paso.

---

## 📊 RESUMEN DE FLASHCARDS

- **Total**: 15 flashcards
- **Cobertura**:
  - Go (8 tarjetas): struct, defer, underscore, slices, signal, range, error handling, tipos
  - Python (2 tarjetas): if __name__, f-string
  - R (2 tarjetas): function syntax, sprintf, mean
  - CSV (1 tarjeta): estructura
- **Enfoque**: Código real + decisiones técnicas + herramientas

---

## 🎓 CÓMO USAR ESTAS FLASHCARDS

1. **Importa a Anki**: Copia el contenido y pega en Anki (Create → Add Note)
2. **Estudia**: 10 minutos al día aprenderás el codebase
3. **Pregunta**: Si algo no queda claro, pregúntame por la pregunta específica

---

**Generadas por**: Narutrader
**Fecha**: 25/11/2025
**Propósito**: Consolidación del codebase real y decisiones técnicas
