package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	var mStart, mEnd runtime.MemStats
	runtime.GC() // Force GC to clean up
	runtime.ReadMemStats(&mStart)

	file, err := os.Open("anscombe.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	datasets := map[string][][]float64{}
	for _, record := range records[1:] {
		label := record[0]
		xVal, _ := strconv.ParseFloat(record[1], 64)
		yVal, _ := strconv.ParseFloat(record[2], 64)
		datasets[label] = append(datasets[label], []float64{xVal, yVal})
	}

	for label, points := range datasets {
		var x, y []float64
		for _, pair := range points {
			x = append(x, pair[0])
			y = append(y, pair[1])
		}

		start := time.Now()
		slope, intercept, err := LinearRegression(x, y)
		elapsed := time.Since(start)

		runtime.ReadMemStats(&mEnd)

		if err != nil {
			log.Fatalf("Dataset %s error: %v", label, err)
		}

		memUsed := mEnd.Alloc - mStart.Alloc

		fmt.Printf("Dataset %s => Slope: %.4f, Intercept: %.4f, Time: %s, Mem Used: %.4f KB\n",
			label, slope, intercept, elapsed, float64(memUsed)/1024)
	}
}
