# Anscombe Quartet Linear Regression Benchmark

## 📊 Overview

This project compares the performance and accuracy of simple linear regression using **Go**, **Python**, and **R** on the [Anscombe Quartet](https://en.wikipedia.org/wiki/Anscombe%27s_quartet). The goal is to evaluate whether a Go statistics package (specifically, [`montanaflynn/stats`](https://github.com/montanaflynn/stats)) can produce regression results comparable to those from Python and R — and to benchmark their respective runtime and memory usage.

---

## 🧪 Objective

- Validate correctness of slope and intercept from Go regression
- Compare regression results from Go, Python, and R
- Benchmark CPU time and memory usage for each language
- Demonstrate Go’s potential for backend analytical tasks

---

## 🗂️ Project Structure
anscombe-regression/
├── go/
│ ├── main.go
│ ├── linear.go
│ ├── linear_test.go
│ ├── benchmark_test.go
│ ├── anscombe.csv
│ └── go.mod
├── python/
│ └── anscombe_regression.py
├── r/
│ └── anscombe_regression.R
├── img/
│ └── miller-fig-anscombe-from-R.png
└── README.md

# How to Run

## Go
cd go
go run .
go test -v
go test -bench=. -benchmem

### Eample and Result
![Go Output](img/go_res.png)
![Go Output](img/go_bench.png)

## Python

cd python
python3 miller-mtpa-chapter-1-program.py

### Eample and Result
![Python Output](img/py_res.png)

## R

cd r
Rscript miller-mtpa-chapter-1-program.R

### Eample and Result
![R Output](img/r_res.png)

## References
Anscombe, F. J. (1973). Graphs in statistical analysis. The American Statistician, 27(1), 17–21.

Miller, T. W. (2015). Modeling Techniques in Predictive Analytics.

Bates & LaNou (2023), Bodner (2024) — Go Testing References
