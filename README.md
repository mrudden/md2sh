# md2sh
Convert markdown README.md files into a bootstrapped shell script. Make self-documenting shell scripts easier to build and work with.

## How it works
Reads every bullet point into an array following a header formatted with "###".
If you have any code on a line denoted by backticks, it should appear literally transcribed, with the original line commented out above it.
Any other lines will be transposed with a "#" in front as a comment.

## To Run
Run `go run md2sh.go -f [INPUTFILENAME.md]` in your terminal, substituting INPUTFILENAME.md with your path and filename
If no filename is provided, it will default to "README.md"

Flags:
* -f = filename to read in
* -s = silent mode; won't print output to terminal
* -l = type of markdown header to denote a list that should be converted into an array (currently supports "h3" or "###" and "h4" or "####")
* -o = creates a file called "run-md2sh.sh" in present directory and outputs to it 

## Test section

### Test List 1
* Test 1
* Test 2
* Test 3

This should produce a shell script array in stdout.

### Test List 2
- Test 4
- Test 5
- Test 6

### Test List 3
+ Test 7
+ Test 8
+ Test 9

### Test List 4 (The Last One)
* Bullet 1
* Bullet 2
* Bullet 3 (last bullet)

`Example code after last list`

This should be the end of the output.