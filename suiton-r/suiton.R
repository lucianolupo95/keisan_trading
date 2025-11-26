# ============================================================================
# SUITON: R Advanced Statistical Analysis Module for KeisanTrading
# Day 4 - 28/11/2025
#
# Purpose: Provide statistical analysis capabilities including:
# - Correlation matrix analysis
# - Distribution statistics
# - Normality testing (Shapiro-Wilk)
# ============================================================================

# Load required libraries
suppressMessages({
  library(e1071)  # For skewness/kurtosis
})

# ============================================================================
# 1. CORRELATION ANALYSIS
# ============================================================================
#' Calculate Correlation Matrix
#'
#' Analyzes correlation between price and time
#' @param prices Numeric vector of price data
#' @return Matrix with correlation and p-value
calculate_correlation <- function(prices) {
  if (length(prices) < 2) {
    stop("Need at least 2 price points")
  }

  # Create time index
  time <- 1:length(prices)

  # Calculate Pearson correlation
  cor_result <- cor.test(time, prices, method = "pearson")

  # Return results as list
  list(
    correlation = cor_result$estimate,
    p_value = cor_result$p.value,
    interpretation = ifelse(cor_result$estimate > 0.7, "strong positive",
                           ifelse(cor_result$estimate > 0.3, "moderate positive",
                                  ifelse(cor_result$estimate < -0.7, "strong negative",
                                         ifelse(cor_result$estimate < -0.3, "moderate negative", "weak"))))
  )
}

# ============================================================================
# 2. DISTRIBUTION ANALYSIS
# ============================================================================
#' Analyze Distribution Statistics
#'
#' Computes mean, median, std deviation, skewness, and kurtosis
#' @param prices Numeric vector of price data
#' @return List with distribution statistics
analyze_distribution <- function(prices) {
  if (length(prices) < 3) {
    stop("Need at least 3 price points")
  }

  mean_price <- mean(prices)
  median_price <- median(prices)
  std_price <- sd(prices)
  skew_price <- skewness(prices)
  kurt_price <- kurtosis(prices)

  # Interpret skewness
  skew_interpretation <- ifelse(skew_price > 0.5, "right-skewed (positive)",
                               ifelse(skew_price < -0.5, "left-skewed (negative)", "approximately symmetric"))

  list(
    mean = mean_price,
    median = median_price,
    std = std_price,
    skewness = skew_price,
    kurtosis = kurt_price,
    coefficient_variation = (std_price / mean_price) * 100,
    distribution_type = skew_interpretation
  )
}

# ============================================================================
# 3. NORMALITY TESTING
# ============================================================================
#' Test Normality using Shapiro-Wilk Test
#'
#' Performs Shapiro-Wilk normality test
#' @param prices Numeric vector of price data
#' @return List with test results
test_normality <- function(prices) {
  if (length(prices) < 3) {
    stop("Need at least 3 price points")
  }

  # Shapiro-Wilk test
  sw_test <- shapiro.test(prices)

  # Interpret result (alpha = 0.05)
  is_normal <- sw_test$p.value > 0.05

  list(
    p_value = sw_test$p.value,
    test_statistic = sw_test$statistic,
    is_normal = is_normal,
    interpretation = ifelse(is_normal, "distribution is normal (p > 0.05)", "distribution is NOT normal (p < 0.05)")
  )
}

# ============================================================================
# MAIN EXECUTION - DEMO MODE
# ============================================================================
main <- function() {
  cat("╔════════════════════════════════════════════════════════════╗\n")
  cat("║           SUITON: Statistical Analysis Module              ║\n")
  cat("║              KeisanTrading - Day 4                          ║\n")
  cat("╚════════════════════════════════════════════════════════════╝\n\n")

  # Generate sample price data (simulating candlestick prices)
  # Uptrend with normal distribution
  set.seed(42)
  base_price <- 1500
  noise <- rnorm(35, mean = 0, sd = 5)
  trend <- seq(0, 50, length.out = 35)
  prices <- base_price + trend + noise

  cat("✓ Suiton activo\n")
  cat("✓ Datos: 35 precios generados\n\n")

  # ========== DISTRIBUTION ANALYSIS ==========
  cat("📊 DISTRIBUTION ANALYSIS\n")
  cat("─────────────────────────\n")

  dist_stats <- analyze_distribution(prices)
  cat(sprintf("  Media:      %.2f\n", dist_stats$mean))
  cat(sprintf("  Mediana:    %.2f\n", dist_stats$median))
  cat(sprintf("  Std Dev:    %.2f\n", dist_stats$std))
  cat(sprintf("  Skewness:   %.4f (%s)\n", dist_stats$skewness, dist_stats$distribution_type))
  cat(sprintf("  Kurtosis:   %.4f\n", dist_stats$kurtosis))
  cat(sprintf("  CV:         %.2f%%\n\n", dist_stats$coefficient_variation))

  # ========== NORMALITY TEST ==========
  cat("🔬 NORMALITY TEST (Shapiro-Wilk)\n")
  cat("──────────────────────────────────\n")

  norm_test <- test_normality(prices)
  cat(sprintf("  p-value:    %.4f\n", norm_test$p_value))
  cat(sprintf("  Statistic:  %.4f\n", norm_test$test_statistic))
  cat(sprintf("  Result:     %s\n\n", norm_test$interpretation))

  # ========== CORRELATION ANALYSIS ==========
  cat("📈 CORRELATION ANALYSIS (Price vs Time)\n")
  cat("────────────────────────────────────────\n")

  corr_result <- calculate_correlation(prices)
  cat(sprintf("  Correlation: %.4f\n", corr_result$correlation))
  cat(sprintf("  p-value:     %.4f\n", corr_result$p_value))
  cat(sprintf("  Type:        %s\n\n", corr_result$interpretation))

  # ========== SUMMARY FOR INTEGRATION ==========
  cat("📋 SUMMARY FOR GO INTEGRATION\n")
  cat("──────────────────────────────\n")
  cat(sprintf("Distribution: media=%.2f, std=%.2f, skewness=%.4f\n",
              dist_stats$mean, dist_stats$std, dist_stats$skewness))
  cat(sprintf("Normality: p-value=%.4f (%s)\n",
              norm_test$p_value, ifelse(norm_test$is_normal, "normal", "not normal")))
  cat(sprintf("Correlation: precios vs tiempo = %.4f (%s)\n",
              corr_result$correlation, corr_result$interpretation))
  cat("\n✅ Suiton analysis complete\n")

  # Return as JSON-ready list for Go bridge
  invisible(list(
    distribution = list(
      mean = dist_stats$mean,
      std = dist_stats$std,
      median = dist_stats$median,
      skewness = dist_stats$skewness,
      kurtosis = dist_stats$kurtosis
    ),
    normality = list(
      p_value = norm_test$p_value,
      is_normal = norm_test$is_normal
    ),
    correlation = list(
      value = as.numeric(corr_result$correlation),
      p_value = corr_result$p_value
    )
  ))
}

# Execute main function if running as script
if (!interactive()) {
  main()
}
