# ============================================================================
# R IPC (Inter-Process Communication) - SIMPLIFIED VERSION
# Para comunicación Go ↔ R via JSON
# Day 5 - 2025-11-27
# ============================================================================

# Load only e1071 library
suppressMessages({
  library(e1071)
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
# MANUAL JSON OUTPUT FUNCTION
# ============================================================================

list_to_json <- function(lst, indent = "") {
  # Simple recursive function to convert R list to JSON string
  if (is.null(lst)) {
    return("null")
  }

  if (is.logical(lst)) {
    return(tolower(as.character(lst)))
  }

  if (is.numeric(lst) && length(lst) == 1) {
    return(as.character(lst))
  }

  if (is.character(lst) && length(lst) == 1) {
    return(paste0('"', lst, '"'))
  }

  if (is.list(lst)) {
    # Check if it's a named list (object) or unnamed (array)
    if (is.null(names(lst))) {
      # Array
      items <- lapply(seq_along(lst), function(i) {
        paste0(indent, "  ", list_to_json(lst[[i]], paste0(indent, "  ")))
      })
      return(paste0("[\n", paste(items, collapse = ",\n"), "\n", indent, "]"))
    } else {
      # Object
      items <- lapply(names(lst), function(name) {
        value_json <- list_to_json(lst[[name]], paste0(indent, "  "))
        paste0(indent, '  "', name, '": ', value_json)
      })
      return(paste0("{\n", paste(items, collapse = ",\n"), "\n", indent, "}"))
    }
  }

  return(as.character(lst))
}

# ============================================================================
# PARSE JSON FROM FILE
# ============================================================================

parse_json_file <- function(filename) {
  # Read JSON file manually
  content <- readLines(filename)
  json_str <- paste(content, collapse = "")

  # Very simple JSON parser for {"prices":[...]}
  tryCatch({
    # Extract the prices array
    prices_start <- gregexpr('"prices"\\s*:\\s*\\[', json_str)[[1]][1]
    if (prices_start == -1) {
      stop("prices not found in JSON")
    }

    # Find the array
    array_start <- prices_start + nchar('"prices":[')
    array_end <- gregexpr('\\]', json_str)[[1]][1]

    prices_str <- substr(json_str, array_start, array_end - 1)

    # Split by comma and convert to numeric
    prices <- as.numeric(unlist(strsplit(prices_str, ",")))

    return(list(prices = prices))
  }, error = function(e) {
    cat(sprintf('{"error": "Failed to parse JSON: %s"}\n', e$message))
    quit(status = 1)
  })
}

# ============================================================================
# MAIN IPC FUNCTION
# ============================================================================

main_ipc <- function() {
  # Get command line arguments
  args <- commandArgs(trailingOnly = TRUE)

  if (length(args) == 0) {
    # Demo mode
    set.seed(42)
    prices <- 1500 + seq(0, 50, length.out = 35) + rnorm(35, mean = 0, sd = 5)
  } else {
    # Read from JSON file
    input_file <- args[1]

    if (!file.exists(input_file)) {
      cat(sprintf('{"error": "File not found: %s"}\n', input_file))
      quit(status = 1)
    }

    json_data <- parse_json_file(input_file)
    prices <- json_data$prices
  }

  # Perform analysis
  distribution <- analyze_distribution(prices)
  normality <- test_normality(prices)
  correlation <- calculate_correlation(prices)

  # Build result
  result <- list(
    distribution = distribution,
    normality = normality,
    correlation = correlation,
    timestamp = as.character(Sys.time())
  )

  # Convert to JSON and print
  json_output <- list_to_json(result)
  cat(json_output)
  cat("\n")
}

# Execute if running as script
if (!interactive()) {
  main_ipc()
}
