# 📅 GUÍA PARA NARUTRADER - PATRÓN DE REPORTES DIARIOS

Esta guía explica cómo Narutrader debe generar reportes cada día.

---

## 🔄 PATRÓN DIARIO A REPETIR

Para cada nuevo día (Día 2, Día 3, etc.), Narutrader debe:

### PASO 1: Crear Carpeta
```bash
mkdir reports/Dia_[N]/
```

### PASO 2: Ejecutar Tareas del Día
Lee `reports/Dia_[N-1]/TODO_DIA_[N].md` y ejecuta todas las tareas.

### PASO 3: Generar 5 Reportes en `reports/Dia_[N]/`

#### 3.1 Dia_[N]_Log.md
- Decisiones técnicas del día
- Fragmentos de código relevantes
- Problemas encontrados y soluciones
- Formato: Markdown detallado

#### 3.2 Resumen_[DATE].md (ej: Resumen_26_11_2025.md)
- Resumen ejecutivo del día
- Qué se hizo
- Resultados de tests
- Cambios principales
- Stats (líneas de código, archivos modificados)

#### 3.3 Flashcards_[DATE].md (ej: Flashcards_26_11_2025.md)
- **MÍNIMO 15 tarjetas Anki**
- Basadas en **CÓDIGO REAL** del día
- Preguntas: "¿Qué hace este código?" + "¿Por qué?"
- Respuestas: Explican decisiones técnicas
- Comparaciones entre lenguajes cuando aplique
- Notas contextuales

#### 3.4 TODO_DIA_[N+1].md (ej: TODO_DIA_3.md)
- Plan de tareas para el siguiente día
- Tareas detalladas en orden
- Cómo testear cada una
- Qué output esperar
- Prioridades (CRITICAL, HIGH, MEDIUM)

#### 3.5 Resumen_for_dummies.md
- Actualizar con lo que se hizo hoy
- Explicación simple (nivel niño de 5 años)
- Qué se va a hacer mañana
- Actualización DENTRO de `Dia_[N]/`

#### 3.6 (Opcional) CIERRE_DIA_[N].md
- Solo el primer día
- Resumen de cierre
- Estado final del día
- Se puede omitir en días siguientes

---

## 📋 CHECKLIST PARA CADA CIERRE DE DÍA

Antes de decir "Día completado", verificar:

- ✅ Carpeta `reports/Dia_[N]/` creada
- ✅ `Dia_[N]_Log.md` con decisiones técnicas
- ✅ `Resumen_[DATE].md` con summary ejecutivo
- ✅ `Flashcards_[DATE].md` con 15+ tarjetas
- ✅ `TODO_DIA_[N+1].md` con plan automático
- ✅ `Resumen_for_dummies.md` actualizado
- ✅ Tests ejecutados y documentados
- ✅ Código commiteado a GitHub
- ✅ Archivo actualizado en ACCESOS_RAPIDOS.md

---

## 🎯 EJEMPLO: DÍA 2 (26/11/2025)

Cuando termines Day 2:

```
reports/
├── Dia_1/
│   ├── Dia_1_Log.md
│   ├── Resumen_25_11_2025.md
│   ├── Flashcards_25_11_2025.md
│   ├── TODO_DIA_2.md
│   ├── Resumen_for_dummies.md
│   └── ESTRUCTURA_Y_ACCESOS.txt
│
└── Dia_2/                        ← NUEVA CARPETA
    ├── Dia_2_Log.md
    ├── Resumen_26_11_2025.md
    ├── Flashcards_26_11_2025.md
    ├── TODO_DIA_3.md
    └── Resumen_for_dummies.md    ← ACTUALIZADO
```

---

## 🔐 INFORMACIÓN A INCLUIR EN CADA REPORTE

### Dia_[N]_Log.md
```
# Narutrader - DÍA [N] - LOG TÉCNICO
Fecha: [DATE]
Supervisor: Narutrader
Usuario: Luchito

## 📋 RESUMEN DEL DÍA
- ✅ Tarea 1
- ✅ Tarea 2
- ⏳ Tarea 3

## 🔴 MÓDULO X (LENGUAJE)
**Archivos creados/modificados**:
- file1.ext
- file2.ext

**Decisión técnica X.X**: Explicación
- Razón: ...
- Trade-off: ...
- Código: ...

## 📊 RESUMEN DE DECISIONES
| Decisión | Módulo | Estado | Notas |

## 🎯 PRÓXIMOS PASOS
1. ...
2. ...
```

### Resumen_[DATE].md
```
# 📊 RESUMEN - DÍA [N] ([DATE])

## 🎯 OBJETIVO
[Qué se intentó lograr]

## ✅ TAREAS COMPLETADAS
1. ...
2. ...

## 📈 STATS
- Líneas de código: X
- Archivos: Y
- Módulos: Z
- Tests: Pass/Fail

## 🎓 APRENDIZAJES
- ...

## 📌 ESTADO ACTUAL
- Módulo 1: Status
- Módulo 2: Status

## 🚀 PRÓXIMA SESIÓN
1. ...
```

### Flashcards_[DATE].md
```
# 🎓 FLASHCARDS ANKI - DÍA [N]
## Keisan Trading - Código Real y Decisiones Técnicas

## FLASHCARD 1 - [TÍTULO]

**Pregunta:**
[Pregunta basada en código real]

**Código:**
[Fragmento de código relevante]

**Respuesta:**
[Explicación clara + decisión técnica]

**Código pregunta:**
[ID único para Anki]

**Notas:**
[Contexto adicional]

---

## 📊 RESUMEN DE FLASHCARDS
- Total: 15+
- Cobertura: [Qué módulos/temas]
- Enfoque: Código real + decisiones técnicas
```

### TODO_DIA_[N+1].md
```
# 🚀 PLAN PARA DÍA [N+1] ([DATE])

**Ubicación de este archivo**: `reports/Dia_[N]/TODO_DIA_[N+1].md`

## ⚡ QUICK START PARA LUCHITO
```
"Narutrader, haz lo de hoy"
```

## 📋 TAREAS DÍA [N+1]

### ✅ TAREA [N+1].1: [Nombre]
**¿Qué hacer?**
- Paso 1
- Paso 2

**Test Command:**
```bash
[Comando para probar]
```

**Expected Output:**
```
[Output esperado]
```

**Status:** [Pendiente/Ready/Done]

---

## 🎯 PRIORIDADES
**CRITICAL**: ...
**HIGH**: ...
**MEDIUM**: ...

---

## 📊 CHECKLIST PARA LUCHITO
```bash
# Para probar manualmente
[comandos de test]
```
```

### Resumen_for_dummies.md
```
# 🧒 RESUMEN PARA NIÑOS DE 5 AÑOS - DÍA [N]

## QUÉ HICIMOS HOY
[Explicación super simple]

## EN NÚMEROS
- ✅ X completado
- ✅ Y completado

## CÓMO FUNCIONA
[Diagrama simple si aplica]

## QUÉ SE VA A HACER MAÑANA
### Las Tareas
1. ...
2. ...

## DÓNDE ENCONTRAR TODO
[Rutas simples]

## EL COMANDO MÁGICO
```
"Narutrader, haz lo de hoy"
```

---

DÍA [N]: [Resumen super breve]
DÍA [N+1]: [Plan super breve]
DÍA [N+2]+: [Plan general]
```

---

## 🔄 INTEGRACIÓN CON GITHUB

Al final de cada día:

1. `git add .`
2. `git commit -m "Day [N] - [descripción breve]"`
3. `git push origin main`

Incluir en el commit message:
- Qué módulos se modificaron
- Stats principales (líneas, archivos)
- Estado de tests

---

## 📝 COMANDOS PARA NARUTRADER

### Al iniciar un nuevo día:
```bash
# Crear carpeta
mkdir reports/Dia_[N]/

# Leer el plan del día anterior
cat reports/Dia_[N-1]/TODO_DIA_[N].md

# Ejecutar tareas
[ejecutar todas las tareas]

# Generar reportes
# - Dia_[N]_Log.md
# - Resumen_[DATE].md
# - Flashcards_[DATE].md
# - TODO_DIA_[N+1].md
# - Resumen_for_dummies.md

# Actualizar ACCESOS_RAPIDOS.md
# Pushear a GitHub
```

---

## 🎓 CHECKLIST FINAL PARA CADA DÍA

```
☐ Leer TODO_DIA_[N].md del día anterior
☐ Ejecutar todas las tareas en orden
☐ Pedir confirmación a Luchito después de cada tarea
☐ Crear carpeta reports/Dia_[N]/
☐ Generar Dia_[N]_Log.md (decisiones técnicas)
☐ Generar Resumen_[DATE].md (resumen ejecutivo)
☐ Generar Flashcards_[DATE].md (15+ tarjetas)
☐ Generar TODO_DIA_[N+1].md (plan automático)
☐ Generar Resumen_for_dummies.md (actualizado)
☐ Actualizar ACCESOS_RAPIDOS.md
☐ Ejecutar tests finales
☐ Pushear a GitHub
☐ Avisar a Luchito que el día está cerrado
```

---

**Generado por**: Narutrader
**Propósito**: Plantilla para reportes diarios consistentes
**Validez**: Todos los días hasta final del proyecto
