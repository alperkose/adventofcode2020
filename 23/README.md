# AdventOfCode: Day#23

Description: https://adventofcode.com/2020/day/23

## Usage

```
cat puzzle.input | go run .
```

## Solution

**Part 1:** using linked lists worked well

**Part 2:** using linked lists did not scale at 1M items and 10M iterations. Estimated execution time was 6-7 hours. I tried adding a pointer to reference previous link in the ring. Estimated execution time dropped to 4 hours but it was not enough.  After going over the reddit posts, using a map as an index for faster access made the solution much faster (2-3 seconds)

## Benchmarks
```
Running tool: /usr/local/bin/go test -benchmem -run=^$ -bench ^(BenchmarkSolverWithIncreasingValues)$ github.com/alperkose/adventofcode2020/23

goos: darwin
goarch: amd64

# linked list with next
BenchmarkSolverWithIncreasingValues-8   	     156	   6721395 ns/op	   40192 B/op	    1001 allocs/op

# linked list with next and prev
BenchmarkSolverWithIncreasingValues-8   	     235	   4802527 ns/op	   40192 B/op	    1001 allocs/op

# linked list with maps to have faster random access
BenchmarkSolverWithIncreasingValues-8   	    1989	    580547 ns/op	  126041 B/op	    1075 allocs/op

# linked list with maps to have faster random access (without fmt.Printf(...))
BenchmarkSolverWithIncreasingValues-8   	    2211	    522761 ns/op	  125706 B/op	    1036 allocs/op

```
## Note
Code requires a serious clean-up