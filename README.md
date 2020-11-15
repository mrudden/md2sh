# md2sh
Convert markdown lists to shell script arrays

## How it works
Reads every bullet point into an array following a header formatted with "###"

## To Run
Run `go run md2sh.go README.md` in your terminal

If you have `code` denoted by backticks, it should appear literally transcribed.

### Test
* Test 1
* Test 2
* Test 3

This should produce a shell script array in stdout.

### Test 2
- Test 4
- Test 5
- Test 6

### Test 3
+ Test 7
+ Test 8
+ Test 9

### Test 4 (The Last One)
* Bullet 1
* Bullet 2
* Bullet 3 (last bullet)

`Example code after last list`

This should be the end of the output.