# AdventOfCode: Day#10

Description: https://adventofcode.com/2020/day/13

## Usage

```
cat puzzle.input | go run .
```

## Solution

I tried to use a brute force approach by parallelizing the calculation. However, it was still very slow. Searching google helped me find about [Chinese remainder theorem](https://en.wikipedia.org/wiki/Chinese_remainder_theorem). The solution is implemented with the Sieve approach of Chinese remainder theorem. I still keep the brute force approach in a separate file.