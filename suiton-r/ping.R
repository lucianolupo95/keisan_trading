# Suiton (Agua) - módulo de estadística y análisis
# ping.R

ping <- function() {
  return("Suiton OK")
}

simple_stat <- function(x) {
  return(mean(x))
}

# Main execution
cat("Suiton activo\n")
response <- ping()
cat(sprintf("Ping response: %s\n", response))

# Test simple_stat
test_values <- c(10, 20, 30)
result <- simple_stat(test_values)
cat(sprintf("simple_stat(c(10,20,30)) = %f\n", result))
