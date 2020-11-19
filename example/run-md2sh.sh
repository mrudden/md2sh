#!/bin/sh

# # md2sh Example
# This is an example file, to be compared against the output from the script in this same directory. The output file will probably need some variables substituted in, some logic added, and some cleanup to work properly, but it's the start of a shell script!

# ## To Run
# After running `git clone [THIS_REPOSITORY]`, you can run the following commands to test md2sh out and reproduce this file:
git clone [THIS_REPOSITORY]
# `cd example`
cd example
# `go run ../md2sh.go -o -s`
go run ../md2sh.go -o -s

# ## Lists of things
# What type of bullets you use in lists shouldn't matter. The four arrays below should all be formatted properly in output.

TEST_LIST_1=(
  Test-1
  Test-2
  Test-3
)
#TODO: Add logic to process this array

# There should be one total test array in the output above this

TEST_LIST_2=(
  Test-4
  Test-5
  Test-6
)
#TODO: Add logic to process this array

# There should be two total test arrays in the output above this

TEST_LIST_3=(
  Test-7
  Test-8
  Test-9
)
#TODO: Add logic to process this array

# There should be three total test arrays in the output above this

TEST_LIST_4_THE_LAST_ONE=(
  Bullet-1
  Bullet-2
  Bullet-3-last-bullet
)
#TODO: Add logic to process this array

# There should be four total test arrays in the output above this

# After the fourth array, let's look at some sample code that could be used to parse an array:
# `for i in "${ARRAY_NAME_HERE}"; do echo $i; done`
for i in "${ARRAY_NAME_HERE}"; do echo $i; done

# This should be the end of the output. We could `echo "Script complete! Exiting."` here.
echo "Script complete! Exiting."
