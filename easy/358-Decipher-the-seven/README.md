# #358 [Easy] Decipher the seven segments

## Description
Today's challenge will be to create a program to decipher a seven segment display, commonly seen on many older electronic devices.

## Input Description
For this challenge, you will receive 3 lines of input, with each line being 27 characters long (representing 9 total numbers), with the digits spread across the 3 lines.
Your job is to return the represented digits.
You don't need to account for odd spacing or missing segments.

## Output Description
Your program should print the numbers contained in the display.

## Challenge Inputs
```
    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|

    _  _  _  _  _  _  _  _ 
|_| _| _||_|| ||_ |_| _||_ 
  | _| _||_||_| _||_||_  _|

 _  _  _  _  _  _  _  _  _ 
|_  _||_ |_| _|  ||_ | ||_|
 _||_ |_||_| _|  ||_||_||_|

 _  _        _  _  _  _  _ 
|_||_ |_|  || ||_ |_ |_| _|
 _| _|  |  ||_| _| _| _||_ 
```

## Challenge Outputs
```
123456789
433805825
526837608
954105592
```