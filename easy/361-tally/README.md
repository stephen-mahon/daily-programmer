# #361 Tally Program

## Description
5 Friends (let's call them a, b, c, d and e) are playing a game and need to keep track of the scores.
Each time someone scores a point, the letter of his name is typed in lowercase.
If someone loses a point, the letter of his name is typed in uppercase.
Give the resulting score from highest to lowest.

## Input Description
A series of characters indicating who scored a point. Examples:
```
abcde
dbbaCEDbdAacCEAadcB
```

## Output Description

The score of every player, sorted from highest to lowest. Examples:

```
a:1, b:1, c:1, d:1, e:1
b:2, d:2, a:1, c:0, e:-2
```

## Challenge Input
```
EbAAdbBEaBaaBBdAccbeebaec
```