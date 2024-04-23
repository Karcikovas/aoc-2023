# Advent of Code λ2023

[**Advent of Code**](http://adventofcode.com/) is an Advent calendar of small programming puzzles for a
variety of skill sets and skill levels that can be solved in written in [Go](https://goland.org).
People use them as a [speed contest](https://adventofcode.com/2023/leaderboard), interview prep, company training,
university coursework, practice problems,
or to [challenge each other](https://www.reddit.com/r/adventofcode/search?q=flair%3Aupping&restrict_sr=on).

This repository contains solutions for puzzles and cli tool to run solutions to get answers for input on site.

## Implemented solutions

<!--- advent_readme_stars table [2023] --->
### 2023 Results

|                      Day                       |    Part 1     | Part 2 |
|:----------------------------------------------:|:-------------:|:------:|
|  [Day 1](https://adventofcode.com/2023/day/1)  |       ⭐       |   -    |
|  [Day 2](https://adventofcode.com/2023/day/2)  |       ⭐       |   -    |
|  [Day 3](https://adventofcode.com/2023/day/3)  |       ⭐       |   -    |
|  [Day 4](https://adventofcode.com/2023/day/4)  |       ⭐       |   -    |
|  [Day 5](https://adventofcode.com/2023/day/5)  |       ⭐       |   -    |
|  [Day 6](https://adventofcode.com/2023/day/6)  | working on it |   -    |
|  [Day 7](https://adventofcode.com/2023/day/7)  |       -       |   -    |
|  [Day 8](https://adventofcode.com/2023/day/8)  |       -       |   -    |
|  [Day 9](https://adventofcode.com/2023/day/9)  |       -       |   -    |
| [Day 10](https://adventofcode.com/2023/day/10) |       -       |   -    |
| [Day 11](https://adventofcode.com/2023/day/11) |       -       |   -    |


## Project Structure

The project is structured in a way that each day of the Advent of Code challenge has its own package under the `internal`
directory. For example, the solution for Day 10 can be found in [`internal/day10/day10.go`](internal/day1/day1.go).

The [`cmd/aoc2023`](cmd/aoc2023) package is the main entry point of the application.
It combines all the days into a single application.

All my personalised inputs are stored in the [`inputs`](inputs) directory, and each day has a corresponding text file
in there. I've also included expected outputs for each day in the day's package using a call to `WithExpectedAnswers`,
this is to allow for regression testing after refactoring.

## Running the Code

To run the code, you need to have Go installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

Once you have Go installed, you can run the code for all days by navigating to the root directory of the project and running:

```bash
go run ./cmd/aoc2023
```

If you want to run the code for a specific day only, you can do so by providing the `--day` flag followed by the day number. For example, to run the code for Day 10, you would do:

```bash
go run ./cmd/aoc2023 --day 10
```

## Contributing

While this is primarily a personal project, contributions are welcome. If you see an issue or have a suggestion for improvement, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [`LICENSE`](LICENSE) file for more details.