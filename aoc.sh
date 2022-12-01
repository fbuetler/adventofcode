#!/bin/bash

set -e

# run with
# cp -r $(date +'%Y')/template $(date +'%Y')/$(date +'%-d') \
# && ./aoc.sh input \
# && cd $(date +'%Y')/$(date +'%-d')

cli_help() {
    cli_name=${0##*/}
    echo "
Advent of Code CLI
Usage: $cli_name [command]
Commands:
  input     Get input
  submit    Submit puzzle
  *         Help
"
    exit 1
}

get_input() {
    echo "Get input of '$DAY.12.$YEAR'"

    mkdir -p "./$YEAR/$DAY"

    curl --silent --cookie "session=$AOC_SESSION" \
        "$URL/$YEAR/day/$DAY/input" |
        sed -z '$ s/\n$//' >"./$YEAR/$DAY/input.txt"

    echo "Saved input in './$YEAR/$DAY/input.txt'"
}

do_submission() {
    PART=$1
    ANSWER=$2

    if [ $PART = "a" ]; then
        LEVEL=1
    elif [ $PART = "b" ]; then
        LEVEL=2
    else
        echo "Invalid part: $PART"
        exit 1
    fi

    echo "Submit answer '$ANSWER' for part '$PART' of '$DAY.12.$YEAR'"
    curl --cookie "session=$AOC_SESSION" \
        -X POST \
        -d "{\"level\": $LEVEL, \"answer\": $ANSWER}" \
        "$URL/$YEAR/day/$DAY/answer"
}

URL="https://adventofcode.com"

DAY=$(date +'%-d')
YEAR=$(date +'%Y')

AOC_SESSION=$(cat ./token)

case "$1" in
input | i)
    get_input
    ;;
submit | s)
    do_submission $2 $3
    ;;
*)
    cli_help
    ;;
esac
