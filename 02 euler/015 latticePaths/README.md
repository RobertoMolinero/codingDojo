# Lattice paths

## Problem 15

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

![Sample](p015.png)

How many such routes are there through a 20×20 grid?

## Solution

There are (at least) two good solutions.

### Dynamic

The first solution can be thought of without much previous knowledge. You look at how many possibilities there are from the penultimate point of the journey and note them down. Then you go backwards step by step to the beginning and always add up the possibilities of the successor.

For further simplification you can initialize all right and lower border points with the value 1 before the actual calculation.

This solution is already quite fast. However, it needs some memory. This becomes problematic for grids much larger than the one under consideration here.

### Combinatorics

The second solution requires somewhat more previous knowledge and considerations.

Each path can be considered as a sequence of 'R' (Right) and 'D' (Down). The question is how many combinatorial possibilities there are.

It is known that the number of 'R' and 'D' must be identical because the grid is square. The answer is the binominal coefficient, more precisely '2N choose N'.

A more detailed explanation can be found in the [Mathblog]((https://www.mathblog.dk/project-euler-15/)) and the blog entry of [Samuel Jack](http://blog.functionalfun.net/2008/07/project-euler-problem-15-city-grids-and.html) linked to it.

### The comparison

To compare the runtime behavior of both solutions I have created a test for each solution. The solution is calculated for the same grid sizes. The size 20 required in the task is in fifth place.

To increase the readability I have added two empty lines to the output. 

```console
$ go test -bench=. -benchmem ./...

goos: linux
goarch: amd64
BenchmarkLatticePathsDynamicWithTable/gridSize=001-8             2876251               413 ns/op             256 B/op          2 allocs/op
BenchmarkLatticePathsDynamicWithTable/gridSize=002-8              966789              1291 ns/op             673 B/op          3 allocs/op
BenchmarkLatticePathsDynamicWithTable/gridSize=005-8              173020              7377 ns/op            3687 B/op          9 allocs/op
BenchmarkLatticePathsDynamicWithTable/gridSize=010-8               38763             29362 ns/op           16515 B/op         20 allocs/op
BenchmarkLatticePathsDynamicWithTable/gridSize=020-8                9740            109912 ns/op           62007 B/op         43 allocs/op
BenchmarkLatticePathsDynamicWithTable/gridSize=050-8                2150            556688 ns/op          247733 B/op        120 allocs/op
BenchmarkLatticePathsDynamicWithTable/gridSize=100-8                 519           2258815 ns/op          986728 B/op        345 allocs/op
PASS
ok      _/.../00_dynamic  9.419s

goos: linux
goarch: amd64
BenchmarkLatticePathsCombinatorial/gridSize=001-8               111631365               10.6 ns/op             0 B/op          0 allocs/op
BenchmarkLatticePathsCombinatorial/gridSize=002-8               57919926                20.8 ns/op             0 B/op          0 allocs/op
BenchmarkLatticePathsCombinatorial/gridSize=005-8               19246789                60.6 ns/op             0 B/op          0 allocs/op
BenchmarkLatticePathsCombinatorial/gridSize=010-8                9326984               129 ns/op               0 B/op          0 allocs/op
BenchmarkLatticePathsCombinatorial/gridSize=020-8                4224976               283 ns/op               0 B/op          0 allocs/op
BenchmarkLatticePathsCombinatorial/gridSize=050-8                1508026               799 ns/op               0 B/op          0 allocs/op
BenchmarkLatticePathsCombinatorial/gridSize=100-8                 707361              1633 ns/op               0 B/op          0 allocs/op
PASS
ok      _/.../01_combinatorial    11.330s
```