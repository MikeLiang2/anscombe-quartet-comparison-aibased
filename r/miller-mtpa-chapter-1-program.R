# Load memory usage package
if (!requireNamespace("pryr", quietly = TRUE)) {
  install.packages("pryr")
}
library(pryr)

# Define Anscombe dataset
anscombe <- data.frame(
  x1 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x2 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x3 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x4 = c(8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8),
  y1 = c(8.04, 6.95,  7.58, 8.81, 8.33, 9.96, 7.24, 4.26,10.84, 4.82, 5.68),
  y2 = c(9.14, 8.14,  8.74, 8.77, 9.26, 8.1, 6.13, 3.1,  9.13, 7.26, 4.74),
  y3 = c(7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73),
  y4 = c(6.58, 5.76,  7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89)
)

# Function to run regression, capture time & memory, and print in Go-style format
run_regression <- function(x, y, label) {
  gc()
  start_mem <- mem_used()
  start_time <- proc.time()

  model <- lm(y ~ x)

  end_time <- proc.time()
  end_mem <- mem_used()

  slope <- coef(model)[2]
  intercept <- coef(model)[1]
  elapsed_time <- (end_time - start_time)[["elapsed"]] * 1e6  # µs
  mem_used_kb <- as.numeric(end_mem - start_mem) / 1024

  cat(sprintf("Dataset %s => Slope: %.4f, Intercept: %.4f, Time: %.0fµs, Mem Used: %.4f KB\n",
              label, slope, intercept, elapsed_time, mem_used_kb))
}

# Run all datasets
run_regression(anscombe$x1, anscombe$y1, "I")
run_regression(anscombe$x2, anscombe$y2, "II")
run_regression(anscombe$x3, anscombe$y3, "III")
run_regression(anscombe$x4, anscombe$y4, "IV")