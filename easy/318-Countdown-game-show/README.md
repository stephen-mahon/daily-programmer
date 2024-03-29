# #318 [Easy] Countdown Game Show
Description
This challenge is based off the British tv game show "Countdown".
The rules are pretty simple: Given a set of numbers X1-X5, calculate using mathematical operations to solve for Y.
You can use addition, subtraction, multiplication, or division.

Unlike "real math", the standard order of operations (PEMDAS) is not applied here.
Instead, the order is determined left to right.

## Example Input
The user should input any 6 whole numbers and the target number. E.g.
```
1 3 7 6 8 3 250
```

## Example Output
The output should be the order of numbers and operations that will compute the target number. E.g.
```
3+8*7+6*3+1=250
```

Note that if follow PEMDAS you get:
```
3+8*7+6*3+1 = 78
```

But remember our rule - go left to right and operate.
So the solution is found by:
```
(((((3+8)*7)+6)*3)+1) = 250
```

If you're into functional progamming, this is essentially a fold to the right using the varied operators.

## Challenge Input
```
25 100 9 7 3 7 881
6 75 3 25 50 100 952
```

## Challenge Output
```
7 * 3 + 100 * 7 + 25 + 9 = 881
100 + 6 * 3 * 75 - 50 / 25 = 952
```