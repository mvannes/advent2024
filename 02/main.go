package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var reports []Report
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		report := Report{levels: []int{}}
		for _, l := range levels {
			li, err := strconv.Atoi(l)
			handleError(err)
			report.levels = append(report.levels, li)
		}

		reports = append(reports, report)
	}

	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

func isReportSafe(report Report) bool {
	descending := false
	ascending := false
	for i := 0; i < len(report.levels)-1; i++ {
		curr := report.levels[i]
		next := report.levels[i+1]

		diff := curr - next
		if diff < 0 {
			descending = true
		} else {
			ascending = true
		}

		// All levels must either be ascending or descending.
		if descending && ascending {
			return false
		}

		absDiff := int(math.Abs(float64(diff)))
		// May not be the same ( < 1 = 0 ) or larger diff than 3.
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
