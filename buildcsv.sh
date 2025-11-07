#!/usr/bin/env bash

# This script generate file named `csvFile` whose contents are in format like `index, random number`
# e.g.
# 1, 798
# 2, 345
# 3, 76
# 
# Usage: `./buildcsv.sh`

set -o errexit; set -o nounset; set -o pipefail

CURDIR="$(pwd)"

function cleanup()
{
  if [[ -f "$CURDIR/csvFile" ]]; then
     rm $CURDIR/csvFile
  fi
}

function validateInput()
{
  re='^[0-9]+$'
  if ! [[ $1 =~ $re ]] ; then
   echo "numOfEntries should be a positive integer only..." 1>&2; exit 1
  fi
}

# Initilze numOfEntries with default value as 10 if no other value is provided while running this script from terminal
numOfEntries=${1:-10}
validateInput $numOfEntries

# Remove older csvFile if exists
cleanup

# Create new csvFile in current directory
touch $CURDIR/csvFile

# Generate `numOfEntries` random numbers and store them into `$CURDIR/csvFile`
RANDOM=$$
for i in `seq $numOfEntries`
do
        echo "$i,$RANDOM" >> $CURDIR/csvFile
done
