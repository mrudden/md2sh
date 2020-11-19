# md2sh
Convert markdown README.md files into a bootstrapped shell script. Make self-documenting shell scripts easier to build and work with.

## How it works
Reads every bullet point into an array following a header formatted with "###".
If you have any code on one line denoted by backticks, it should appear literally transcribed, with the original line commented out above it.
If you have a code block denoted by either three backticks or three tildes, it will transcribe the block with comments to note the beginning and end of the code block.
Any other lines will be transposed with a "#" in front as a comment.

## To Run
Run `go run md2sh.go -f [INPUTFILENAME.md]` in your terminal, substituting INPUTFILENAME.md with your path and filename
If no filename is provided, it will default to "README.md"

Flags:
* -f = filename to read in, defaults to "README.md"
* -s = silent mode; won't print output to terminal, defaults to "false"
* -l = type of markdown header to denote a list that should be converted into an array (currently supports "h3" or "###" and "h4" or "####"), defaults to "###"
* -o = creates a file called "run-md2sh.sh" in present directory and outputs to it, defaults to "false"