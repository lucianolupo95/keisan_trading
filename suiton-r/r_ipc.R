# ============================================================================
# R IPC (Inter-Process Communication) Module
# Para comunicación Go ↔ R via JSON
# Day 4 - 28/11/2025
# ============================================================================

# Load libraries
suppressMessages({
  library(e1071)
  library(jsonlite)
})

# ============================================================================
# FUNCIONES DE ANÁLISIS (reutilizadas de suiton.R)
# ============================================================================

calculate_correlation <- function(prices) {
  if (length(prices) < 2) {
    stop("Need at least 2 price points")
  }

  time <- 1:length(prices)
  cor_result <- cor.test(time, prices, method = "pearson")

  list(
    correlation = as.numeric(cor_result$estimate),
    p_value = as.numeric(cor_result$p.value),
    interpretation = ifelse(cor_result$estimate > 0.7, "strong positive",
                           ifelse(cor_result$estimate > 0.3, "moderate positive",
                                  ifelse(cor_result$estimate < -0.7, "strong negative",
                                         ifelse(cor_result$estimate < -0.3, "moderate negative", "weak"))))
  )
}

analyze_distribution <- function(prices) {
  if (length(prices) < 3) {
    stop("Need at least 3 price points")
  }

  mean_price <- mean(prices)
  median_price <- median(prices)
  std_price <- sd(prices)
  skew_price <- skewness(prices)
  kurt_price <- kurtosis(prices)

  skew_interpretation <- ifelse(skew_price > 0.5, "right-skewed (positive)",
                               ifelse(skew_price < -0.5, "left-skewed (negative)", "approximately symmetric"))

  list(
    mean = as.numeric(mean_price),
    median = as.numeric(median_price),
    std = as.numeric(std_price),
    skewness = as.numeric(skew_price),
    kurtosis = as.numeric(kurt_price),
    coefficient_variation = as.numeric((std_price / mean_price) * 100),
    distribution_type = skew_interpretation
  )
}

test_normality <- function(prices) {
  if (length(prices) < 3) {
    stop("Need at least 3 price points")
  }

  sw_test <- shapiro.test(prices)
  is_normal <- sw_test$p.value > 0.05

  list(
    p_value = as.numeric(sw_test$p.value),
    test_statistic = as.numeric(sw_test$statistic),
    is_normal = is_normal,
    interpretation = ifelse(is_normal, "distribution is normal (p > 0.05)", "distribution is NOT normal (p < 0.05)")
  )
}

# ============================================================================
# IPC MAIN FUNCTION
# ============================================================================

main_ipc <- function() {
  # Leer argumentos de línea de comandos
  args <- commandArgs(trailingOnly = TRUE)

  if (length(args) == 0) {
    # Si no hay argumentos, usar datos de demo
    set.seed(42)
    prices <- 1500 + seq(0, 50, length.out = 35) + rnorm(35, mean = 0, sd = 5)
  } else {
    # Leer precios desde archivo JSON
    input_file <- args[1]

    if (!file.exists(input_file)) {
      stop(sprintf("Input file not found: %s", input_file))
    }

    # Parsear JSON
    tryCatch({
      data <- fromJSON(input_file)
      prices <- as.numeric(data$prices)

      if (length(prices) < 3) {
        stop("Need at least 3 prices")
      }
    }, error = function(e) {
      stop(sprintf("Error parsing JSON: %s", e$message))
    })
  }

  # ========== ANÁLISIS ==========
  distribution <- analyze_distribution(prices)
  normality <- test_normality(prices)
  correlation <- calculate_correlation(prices)

  # ========== RESULTADO ==========
  result <- list(
    distribution = distribution,
    normality = normality,
    correlation = correlation,
    timestamp = Sys.time()
  )

  # ========== OUTPUT JSON ==========
  output_json <- toJSON(result, auto_unbox = TRUE, pretty = TRUE)
  cat(output_json)
}

# ============================================================================
# EJECUTAR
# ============================================================================

# Ejecutar IPC si está corriendo como script
if (!interactive()) {
  main_ipc()
}
