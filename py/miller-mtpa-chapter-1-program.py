import pandas as pd
import statsmodels.api as sm
import time
import tracemalloc

# Define the Anscombe dataset
anscombe = pd.DataFrame({
    'x1': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x2': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x3': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x4': [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
    'y1': [8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68],
    'y2': [9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74],
    'y3': [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73],
    'y4': [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89]
})

def run_regression(x, y, label):
    tracemalloc.start()
    start_time = time.perf_counter()

    X = sm.add_constant(x)
    model = sm.OLS(y, X).fit()

    end_time = time.perf_counter()
    current, peak = tracemalloc.get_traced_memory()
    tracemalloc.stop()

    slope = model.params.iloc[1]
    intercept = model.params.iloc[0]

    elapsed_us = (end_time - start_time) * 1e6  # microseconds

    print(f"Dataset {label} => Slope: {slope:.4f}, Intercept: {intercept:.4f}, Time: {elapsed_us:.0f}µs, Mem Used: {peak / 1024:.4f} KB")

# Run for all four datasets
run_regression(anscombe['x1'], anscombe['y1'], "I")
run_regression(anscombe['x2'], anscombe['y2'], "II")
run_regression(anscombe['x3'], anscombe['y3'], "III")
run_regression(anscombe['x4'], anscombe['y4'], "IV")
