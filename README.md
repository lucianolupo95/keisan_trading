# 🚀 KEISAN TRADING - PROYECTO UNIFICADO

**Carpeta oficial**: `C:\Cosas_Lucho\Programacion\Proyectos\Keisan_trading\`

---

## 📂 ESTRUCTURA UNIFICADA

```
Keisan_trading/
│
├── .claude/                          [Configuración de Claude Code]
│   ├── agents/
│   │   └── narutrader.md            [Agente supervisor del proyecto]
│   ├── config.json                  [Configuración principal]
│   └── settings.local.json          [Permisos locales]
│
├── fuuton-go/                        [🔵 VIENTO - Orquestador]
│   ├── main.go                       [CSV reader + Signal generator]
│   └── data.csv                      [Datos de test - 3 velas]
│
├── katon-python/                     [🟠 FUEGO - Features]
│   └── ping.py                       [Transformaciones simples]
│
├── suiton-r/                         [🔵 AGUA - Estadística]
│   └── ping.R                        [Análisis estadístico]
│
├── doton-csharp/                     [🟢 TIERRA - Ejecución]
│   └── KeisanBridge.cs              [Puente para NinjaTrader]
│
├── reports/                          [📊 REPORTES Y PLANES DIARIOS]
│   ├── Dia_1_Log.md                 [Decisiones técnicas Day 1]
│   ├── Dia_[N]_Log.md               [Nuevo cada día]
│   ├── Resumen_25_11_2025.md        [Resumen ejecutivo Day 1]
│   ├── Resumen_[DATE].md            [Nuevo cada día]
│   ├── Flashcards_25_11_2025.md     [15 tarjetas de aprendizaje]
│   ├── Flashcards_[DATE].md         [Nuevo cada día]
│   ├── CIERRE_DIA_1.md              [✅ Cierre del Día 1]
│   ├── TODO_DIA_2.md                [🎯 Plan para el Día 2]
│   ├── TODO_DIA_[N].md              [Nuevo cada día]
│   └── ESTRUCTURA_Y_ACCESOS.txt     [🗺️ Mapa del proyecto]
│
└── README.md                         [Este archivo]
```

---

## 🎯 ACCESOS RÁPIDOS

### Para Luchito
```
Carpeta raíz:  C:\Cosas_Lucho\Programacion\Proyectos\Keisan_trading\
```

### Dónde encontrar qué
| Elemento | Ruta |
|----------|------|
| **LOGS técnicos** | `reports/Dia_1_Log.md` |
| **Resumen ejecutivo** | `reports/Resumen_25_11_2025.md` |
| **Flashcards Anki** | `reports/Flashcards_25_11_2025.md` |
| **Plan Day 2** | `reports/TODO_DIA_2.md` |
| **Estructura del proyecto** | `reports/ESTRUCTURA_Y_ACCESOS.txt` |
| **Cierre del Día 1** | `reports/CIERRE_DIA_1.md` |
| **Configuración Claude** | `.claude/config.json` |
| **Agente Narutrader** | `.claude/agents/narutrader.md` |

---

## 🚀 QUICK START

### Para Mañana (Comando Mágico)
Abre Claude Code en esta carpeta y di:
```
"Narutrader, haz lo de hoy"
```

Narutrader automáticamente:
1. Lee `reports/TODO_DIA_2.md`
2. Ejecuta todas las tareas del día
3. Genera logs + resumen + flashcards (en `reports/`)
4. Actualiza `reports/TODO_DIA_[N].md` para el siguiente día

### Para Testear Manual
```bash
cd C:\Cosas_Lucho\Programacion\Proyectos\Keisan_trading\fuuton-go
go run main.go
```

Expected output:
```
Fuuton activo
Ping response: Fuuton OK
Velas leídas: 3
Vela 1 - ... | Signal: BUY
Vela 2 - ... | Signal: BUY
Vela 3 - ... | Signal: BUY
```

---

## 📊 ESTADO ACTUAL

| Módulo | Status | Detalles |
|--------|--------|----------|
| **Fuuton (Go)** | ✅ FUNCIONAL | CSV reader + signal generator |
| **Katon (Python)** | ✅ FUNCIONAL | Feature transformations |
| **Suiton (R)** | ⏳ LISTO | Esperando Rscript |
| **Doton (C#)** | ⏳ LISTO | Esperando compilador |
| **Documentación** | ✅ COMPLETA | Logs + resumen + flashcards |

---

## 🔐 CONFIGURACIÓN CLAUDE CODE

El proyecto está configurado con:
- **Agente por defecto**: `narutrader`
- **Modelo**: Claude Sonnet 4.5
- **Color**: Naranja
- **Permisos**: Total (Bash, Go, Python, etc.)

Para cambiar configuración, edita `.claude/config.json`

---

## 📝 NOTAS IMPORTANTES

1. **Unified Folder**: Todo está en `Keisan_trading` (no hay duplicados)
2. **Claude Config**: Dentro del proyecto (`.claude/`)
3. **Daily Updates**: Cada día genera (en `reports/`):
   - `Dia_[N]_Log.md` - Decisiones técnicas del día
   - `Resumen_[DATE].md` - Resumen ejecutivo
   - `Flashcards_[DATE].md` - 15+ tarjetas de aprendizaje
   - `TODO_DIA_[N].md` - Plan para el siguiente día
   - `CIERRE_DIA_[N].md` - Resumen de cierre

---

## 🎓 DÍA 1 SUMMARY

- ✅ 4 módulos creados (Go, Python, R, C#)
- ✅ 2 módulos funcionales (Go, Python)
- ✅ 168 líneas de código
- ✅ 12 flashcards generadas
- ✅ Logs técnicos documentados
- ✅ Plan Day 2 creado

---

## 🚀 PRÓXIMOS PASOS

**Day 2**:
- Probar R (Suiton)
- Probar C# (Doton)
- Integración Go → Python
- Mejorar signal generator

**Day 3+**:
- Integración completa 4/4
- Backtesting
- Trading en vivo (SIM)

---

**Proyecto**: Keisan Trading
**Supervisor**: Narutrader
**Usuario**: Luchito
**Fecha de creación**: 25/11/2025
**Estado**: Unificado y listo para Day 2
