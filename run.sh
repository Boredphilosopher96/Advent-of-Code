#!/bin/bash
go_files=("day1.go" "day2.go")
input_files=("./inputs/day1.txt" "./inputs/day2.txt")

if (( $# != 1 )); then
    >&2 echo "Illegal parameters. Pass only one parameter to suggest which day's advent of code you are trying to solve"
    exit 1
fi

filename="${go_files[$1-1]}"
input_filename="${input_files[$1-1]}"
echo "running solution file $filename with input file $input_filename"

basename="outputs/${filename%.*}"
go build -o $basename $filename utils.go
echo "$input_filename" | "./${basename}"
