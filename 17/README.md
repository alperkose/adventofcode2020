# AdventOfCode: Day#17

Description: https://adventofcode.com/2020/day/17

## Usage

```
cat puzzle.input | go run .
```

## Solution

**Part 1:** used a map as an index of active cubes, then go over the existing active cubes determine if they live or not and then another loop over the neighbors of the active cube to determine if they should become active

**Part 2:** Quite similar to part 1, but with additional dimension. Redesigned the code to accept a function that provides list of neighbors of an active cube. Part 2 provides neighbors based on 4 dimensions and part 1 in 3 dimensions

## Notes
Started using acceptance tests that has the sample information from the description.
