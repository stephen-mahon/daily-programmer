# #388 [Easy] What day was it again?
## Description
What day of the week did hitler get elected on?

What day of the week did the Normans invade Britain on?

What day of the week did Jesus die on?

What day of the week did MacDonald get founded on?

Today we're gonna find out

For todays challenge you are allowed to use your languages built in Calender functions/classes.

But it's more interesting if you do the calculation yourself :)

## Hint
It's leap-year if the year is divisible by 4

Ignore leap-year if the year is divisible by 100

Ignore previous rule if the year is divisible by 400
## Input Example
The input will be 3 integers such as
```
<YEAR> <MONTH> <DAY>
```
Limits for the 3 integers:
```
8000 > YEAR > 0
13 > Month > 0
32 > Date > 0
```
Sorry to anyone starting at 0.

January is 1 and December is 12

Assume all dates to be correct (i.e. no 31th of february)

Input will look like:
```
2017 10 30
```
## Output example
Output is simply the day of the week of the given date, for above it would be:
```
Monday
```
## Challenge input
```
2017 10 30
2016 2 29
2015 2 28
29 4 12
570 11 30
1066 9 25
1776 7 4
1933 1 30
1953 3 6
2100 1 9
2202 12 15
7032 3 26
```
