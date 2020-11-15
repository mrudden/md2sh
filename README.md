# md2sh
Convert markdown README.md files into bootstraps for shell scripts. Make self-documenting shell scripts easier to build and work with.

## How it works
Reads every bullet point into an array following a header formatted with "###"
If you have code denoted by backticks, it should appear literally transcribed.

## To Run
Run `go run md2sh.go README.md` in your terminal

## Test section

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