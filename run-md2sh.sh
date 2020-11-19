#!/bin/sh# # md2sh
# Convert markdown README.md files into a skeleton for bootstrapping a shell script. Make self-documenting shell scripts easier to build and work with.

# ## How it works
# Reads every bullet point into an array following a header formatted with "###".
# If you have any code on a line denoted by backticks, it should appear literally transcribed, with the original line commented out above it.
# Any other lines will be transposed with a "#" in front as a comment.

# ## To Run
# Run `go run md2sh.go -f [INPUTFILENAME.md]` in your terminal, substituting INPUTFILENAME.md with your path and filename
go run md2sh.go -f [INPUTFILENAME.md]
# If no filename is provided, it will default to "README.md"

# Flags:
# * -f = filename to read in
# * -s = silent mode; won't print output to terminal
# * -l = type of markdown header to denote a list that should be converted into an array (currently supports "h3" or "###" and "h4" or "####")
# * -o = creates a file called "run-md2sh.sh" in present directory and outputs to it 

# ## Test section

TEST_LIST_1=(
  Test-1
  Test-2
  Test-3
)
#TODO: Add logic to process this array

# This should produce a shell script array in stdout.

TEST_LIST_2=(
  Test-4
  Test-5
  Test-6
)
#TODO: Add logic to process this array

TEST_LIST_3=(
  Test-7
  Test-8
  Test-9
)
#TODO: Add logic to process this array

TEST_LIST_4_THE_LAST_ONE=(
  Bullet-1
  Bullet-2
  Bullet-3-last-bullet
)
#TODO: Add logic to process this array

# `Example code after last list`
Example code after last list

# This should be the end of the output.
