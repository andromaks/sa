# Programming challenge

This is a Go program that takes two inputs:
• a set of include intervals
• a set of exclude intervals

The sets of intervals can be given in any order, and they may be empty or overlapping.
The program should output the result of taking all the includes and “remove” the excludes.
The output should be given as non overlapping intervals in a sorted order.
Intervals will contain Integers only.

**Example 1:** \
Includes: 10-100 \
Excludes: 20-30 \
Output should be: 10-19, 31-100 \
\
**Example 2:** \
Includes: 50-5000, 10-100 \
Excludes: (none) \
Output: 10-5000 \
\
**Example 3:** \
Includes: 10-100, 200-300 \ 
Excludes: 95-205 \
Output: 10-94,206-300 \
\
**Example 4:** \
Includes: 10-100, 200-300, 400-500 \
Excludes: 95-205, 410-420 \
Output: 10-94, 206-300, 400-409, 421-500



### Actions have 2 scripts:
- golangci-lint
  Auntomatically runs on every push/pull requests and checks script code
- Test Accross Matrix
  Has manual start and checks that code runs on different OSes with different GO versions.



### Unit tests were executed manually according to **Examples** specification
