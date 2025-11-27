#!/usr/bin/env Rscript
# Minimal R IPC for Day 5 - Testing

# Load library
library(e1071)

# Parse command line
args <- commandArgs(trailingOnly = TRUE)

# Generate demo data
set.seed(42)
prices <- 1500 + seq(0, 50, length.out = 35) + rnorm(35, mean = 0, sd = 5)

# Simple analysis
mean_val <- mean(prices)
std_val <- sd(prices)
skew_val <- skewness(prices)
kurt_val <- kurtosis(prices)

# Test normality
sw_test <- shapiro.test(prices)
p_val <- sw_test$p.value

# Correlation
time <- 1:length(prices)
cor_test <- cor.test(time, prices)
cor_val <- cor_test$estimate

# Manual JSON output
cat('{\n')
cat('  "distribution": {\n')
cat('    "mean": ', mean_val, ',\n')
cat('    "std": ', std_val, ',\n')
cat('    "skewness": ', skew_val, ',\n')
cat('    "kurtosis": ', kurt_val, '\n')
cat('  },\n')
cat('  "normality": {\n')
cat('    "p_value": ', p_val, ',\n')
cat('    "is_normal": ', tolower(p_val > 0.05), '\n')
cat('  },\n')
cat('  "correlation": {\n')
cat('    "value": ', as.numeric(cor_val), '\n')
cat('  }\n')
cat('}\n')
