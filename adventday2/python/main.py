#!/usr/bin/env python3


def total(lines, scoreSheet):
    score = 0
    for line in lines:
        row = int(ord(line[0])) - int(ord('A'))
        column = int(ord(line[2])) - int(ord('X'))
        #print("row {} column {}".format(row, column))
        #print("scoresheet: {}".format(scoreSheet[row][column]))
        score += scoreSheet[row][column]
    return score

def part1(lines):
    scoreSheet = [
        [4,8,3],
        [1,5,9],
        [7,2,6]
    ]
    return total(lines, scoreSheet)

def part2(lines):
    scoreSheet = [
        [3,4,8],
        [1,5,9],
        [2,6,7]
    ]
    return total(lines, scoreSheet)


rps = open('../input.txt').read()

lines = []
for line in rps.splitlines():
    lines.append(line)

print("Part 1: {}".format(part1(lines)))
print("Part 1: {}".format(part2(lines)))



