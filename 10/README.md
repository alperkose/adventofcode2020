# AdventOfCode: Day#10

Description: https://adventofcode.com/2020/day/10

## Usage

```
cat sample.input | go run .
# OR
cat puzzle.input | go run .
# For using tribonacci alternative
cat puzzle.input | go run . -m tribonacci
```

`-m` parameter defines which method to use for calculation combinations of different jolt adapters

## Solution
For part #1, the solution was a straight forward sorting and counting

For part #2, I first tried recursion. But after running the program for 5 minutes without an output, I decided to change my approach. After some time on paper, I decided to apply recursive combination calculation to the groups of devices separated by 3 jolt difference.

I later find out that there is a formal method called "[tribonacci sequence](https://www.wolframalpha.com/input/?i=tribonacci+sequence)" where the combination of one jolt difference groups can be calculated easily. However, it doesn't work for 2 jolt differences. "Luckily" the problem puzzle input doesn't contain such cases.

Out of these two methods, I picked the first one as default. But the program can be calculated with tribonacci method as well by providing `-m tribonacci`argument (see above)
