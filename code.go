package main

import (
	"fmt"
	"sort"
	"strings"
)

// Interval is a custom type to store the start and end of an interval
type Interval struct {
	start int
	end   int
}

// mergeIntervals takes a slice of intervals and returns a new slice of non-overlapping intervals
func mergeIntervals(intervals []Interval) []Interval {
	// sort the intervals by their start values
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	// create a new slice to store the merged intervals
	merged := make([]Interval, 0)

	// iterate over the sorted intervals
	for _, interval := range intervals {
		// if the merged slice is empty or the current interval does not overlap with the last merged interval, append it to the merged slice
		if len(merged) == 0 || interval.start > merged[len(merged)-1].end {
			merged = append(merged, interval)
		} else {
			// otherwise, update the end of the last merged interval to the maximum of its own end and the current interval's end
			merged[len(merged)-1].end = max(merged[len(merged)-1].end, interval.end)
		}
	}

	return merged
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// define the include and exclude intervals as slices of Interval
	include := []Interval{}
	exclude := []Interval{{20, 30}}

	// merge the include and exclude intervals separately
	include = mergeIntervals(include)
	exclude = mergeIntervals(exclude)

	// create a new slice to store the output intervals
	output := make([]Interval, 0)

	// use two pointers to iterate over the include and exclude slices
	i := 0 // pointer for include
	j := 0 // pointer for exclude

	for i < len(include) && j < len(exclude) {
		// if the current include interval is completely before the current exclude interval, append it to the output and increment i
		if include[i].end < exclude[j].start {
			output = append(output, include[i])
			i++
		} else if include[i].start > exclude[j].end {
			// if the current include interval is completely after the current exclude interval, increment j
			j++
		} else {
			// otherwise, there is some overlap between the current include and exclude intervals
			// check if there is a gap between the start of the include interval and the start of the exclude interval, and append it to the output if so
			if include[i].start < exclude[j].start {
				output = append(output, Interval{include[i].start, exclude[j].start - 1})
			}
			// check if there is a gap between the end of the exclude interval and the end of the include interval, and update the start of the include interval to that gap if so
			if include[i].end > exclude[j].end {
				include[i].start = exclude[j].end + 1
				j++
			} else {
				// otherwise, the include interval is completely covered by the exclude interval, so increment i
				i++
			}
		}
	}

	// append any remaining include intervals to the output
	for i < len(include) {
    output = append(output,include[i])
    i++
  }

	fmt.Println("Output:")
	var sb strings.Builder // use a string builder to concatenate output intervals in one line
	for _, interval := range output {
    sb.WriteString(fmt.Sprintf("%d-%d ",interval.start,interval.end)) // write each interval as start-end followed by a space
  }
	fmt.Println(sb.String()[:sb.Len()-1]) // print the string builder without the last comma 
}

