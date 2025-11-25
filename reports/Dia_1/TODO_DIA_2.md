# 🚀 PLAN PARA DÍA 2 (26/11/2025)

**Ubicación de este archivo**: `reports/TODO_DIA_2.md`

## ⚡ QUICK START PARA LUCHITO
Cuando entres mañana, simplemente dile a Narutrader:
```
"Haz lo de hoy"
```

Y yo automáticamente ejecutaré todas estas tareas en orden.

---

## 📋 TAREAS DÍA 2

### ✅ TAREA 2.1: Instalar/Probar R (Suiton)
**¿Qué hacer?**
- Compilar/ejecutar `suiton-r/ping.R` con Rscript
- Verificar output: "Suiton OK" + "simple_stat(c(10,20,30)) = 20"

**Test Command:**
```bash
cd C:\Cosas_Lucho\Programacion\Proyectos\KeisanTrading\suiton-r
Rscript ping.R
```

**Expected Output:**
```
Suiton activo
Ping response: Suiton OK
simple_stat(c(10,20,30)) = 20
```

**Status:** ⏳ Pendiente instalación de Rscript

---

### ✅ TAREA 2.2: Instalar/Compilar C# (Doton)
**¿Qué hacer?**
- Compilar `doton-csharp/KeisanBridge.cs`
- Ejecutar: Verificar output "Doton OK"

**Test Commands:**
```bash
cd C:\Cosas_Lucho\Programacion\Proyectos\KeisanTrading\doton-csharp
csc KeisanBridge.cs
KeisanBridge.exe
```

**Expected Output:**
```
Doton activo
Ping response: Doton OK
```

**Status:** ⏳ Pendiente compilador C# (csc o dotnet)

---

### ✅ TAREA 2.3: Integración Fuuton → Katon
**¿Qué hacer?**
- Crear `fuuton-go/katon_caller.go`
- Fuuton llama Python `ping.py` usando `os/exec`
- Leer response y verificar que Katon responda correctamente

**Expected Flow:**
```
Fuuton (Go)
  ↓ [exec Python]
Katon (Python)
  ↓ [response back]
Fuuton (Go) [prints response]
```

**Test Output:**
```
Fuuton activo
Calling Katon...
Katon response: Katon OK
```

---

### ✅ TAREA 2.4: Crear data.csv más realista
**¿Qué hacer?**
- Expandir `data.csv` a 20-50 velas
- Agregar casos mixtos: BUY signals + HOLD signals
- Usar precios más realistas (p.ej., 1500-1550 para BTC)

**Example:**
```csv
timestamp,open,close,high,low,volume
2025-11-25T09:00:00Z,1500.0,1501.5,1502.0,1499.5,1000
2025-11-25T09:15:00Z,1501.5,1500.0,1503.0,1500.0,1200  [HOLD]
2025-11-25T09:30:00Z,1500.0,1502.0,1503.0,1499.0,1500  [BUY]
...
```

**Status:** Requiere decisión de rango de precios y volúmenes

---

### ✅ TAREA 2.5: Mejorar GenerateSignal en Fuuton
**¿Qué hacer?**
- Agregar filtro de VOLUMEN mínimo
- Agregar filtro de MOVIMIENTO porcentual mínimo
- Output mejorado: mostrar reasoning de signal

**Example:**
```go
func GenerateSignal(candle Candle) string {
    minVolume := int64(500)
    minMove := 0.5  // 0.5%

    if candle.Volume < minVolume {
        return "HOLD"  // Volume insuficiente
    }

    movePercent := ((candle.Close - candle.Open) / candle.Open) * 100
    if candle.Close > candle.Open && movePercent > minMove {
        return "BUY"
    }
    return "HOLD"
}
```

**Status:** Ready to implement

---

### ✅ TAREA 2.6: Registrar decisiones Day 2
**¿Qué hacer?**
- Crear `logs/Dia_2_Log.md` con todas las decisiones
- Documentar cada cambio y su reasoning
- Registrar problemas encontrados

**Status:** Automático al finalizar Day 2

---

### ✅ TAREA 2.7: Generar Resumen Day 2
**¿Qué hacer?**
- Crear `docs/Resumen_26_11_2025.md`
- Resumen ejecutivo con cambios principales
- Screenshots/outputs de tests exitosos

**Status:** Automático al finalizar Day 2

---

### ✅ TAREA 2.8: Generar Flashcards Day 2
**¿Qué hacer?**
- Crear 10+ nuevas flashcards sobre:
  - Integración Go-Python
  - Mejoras a GenerateSignal
  - Filtros de volumen y movimiento
  - Próximas arquitecturas (Suiton integration)

**Status:** Automático al finalizar Day 2

---

## 🎯 PRIORIDADES

**CRITICAL (Hace el sistema funcional 4/4):**
1. Instalar R → Test Suiton
2. Instalar/compilar C# → Test Doton

**HIGH (Mejora funcionalidad):**
3. Integración Fuuton → Katon
4. Mejorar GenerateSignal con filtros

**MEDIUM (Data quality):**
5. CSV más realista (50 velas)

---

## 📊 CHECKLIST PARA LUCHITO

Si quieres probar solo una tarea manualmente:
```bash
# Para probar Fuuton con data más realista
cd C:\Cosas_Lucho\Programacion\Proyectos\KeisanTrading\fuuton-go
go run main.go

# Para ver los nuevos filtros en action (cuando esté implementado)
cd C:\Cosas_Lucho\Programacion\Proyectos\KeisanTrading\fuuton-go
go run main.go --verbose
```

---

## 🔄 COMANDO MÁGICO PARA MAÑANA

Simplemente escribe:
```
"Narutrader, haz lo de hoy"
```

Y yo automáticamente:
1. Leo este archivo (TODO_MAÑANA.md)
2. Ejecuto todas las tareas en orden
3. Pido confirmación después de cada una
4. Genero logs + resumen + flashcards
5. Actualizo TODO_MAÑANA.md para Day 3

---

## 📍 ARCHIVOS DEL PROYECTO

```
KeisanTrading/
├── fuuton-go/
│   ├── main.go           (Orquestador + CSV reader)
│   ├── data.csv          (3 velas, será expandido a 50+)
│   └── katon_caller.go   [PENDING - Day 2]
├── katon-python/
│   └── ping.py
├── suiton-r/
│   └── ping.R            [PENDING TEST - necesita Rscript]
├── doton-csharp/
│   └── KeisanBridge.cs   [PENDING COMPILE - necesita csc]
├── docs/
│   ├── Resumen_25_11_2025.md
│   ├── Flashcards_25_11_2025.md
│   ├── Resumen_26_11_2025.md   [PENDING - Day 2]
│   └── Flashcards_26_11_2025.md [PENDING - Day 2]
├── logs/
│   ├── Dia_1_Log.md
│   └── Dia_2_Log.md      [PENDING - Day 2]
└── TODO_MAÑANA.md        [Este archivo - actualizado cada día]
```

---

## 🎓 NOTAS TÉCNICAS

- **Go version**: 1.x+ (testeado con tu setup)
- **Python version**: 3.x+ (testeado con tu setup)
- **R version**: TBD (necesita instalación)
- **C# version**: TBD (necesita dotnet o csc)

---

**Generado por**: Narutrader
**Fecha**: 25/11/2025
**Para**: Día 2 (26/11/2025)

⚠️ **IMPORTANTE**: Si alguna tarea falla durante Day 2, yo pausaré y te pediré ayuda antes de continuar.
