# AdventOfCode: Day#16

Description: https://adventofcode.com/2020/day/16

## Usage

```
cat puzzle.input | go run .
```

## Solution

**Part 1:** Iteration over tickets and checking if the numbers conform to ranges

**Part 2:** Started by mapping list of fields conforming to overlapping ranges. Then checked if the tickets with the relevant field is in the same range of mapped list of fields, and eliminating that is not in the range. Finally, another round of elimination of remaining field by checking if the field is the single one in that group then deleting it from other groups

Code requires refactoring