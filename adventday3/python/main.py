#!/usr/bin/env python3
import string

def chunk(list, size):
    for i in range(0, len(list), size):
        yield list[i:i + size]

def calc(lists):
    priority = genPriority()

    score = 0
    for list in lists:
        for item in list:
            score += priority[item]
    return score

def genPriority():
    values = dict()
    for index, letter in enumerate(string.ascii_lowercase):
        values[letter] = index + 1
    for index, letter in enumerate(string.ascii_uppercase):
        values[letter] = index + 27
    return values

def part1(list):
    res = []
    for line in list:
        left = set([x for x in line[:len(line)//2]])
        right = set([x for x in line[len(line)//2:]])

        diff = left & right
        res.append(diff)

    print("\npart 1 result: {}".format(calc(res)))

def part2(list):
    res = []

    groups = chunk(list, 3)

    for group in groups:
        diff = set(group[0]) & set(group[1]) & set(group[2])
        res.append(diff)

    print("part 2 result: {}".format(calc(res)))
        

# Main
ruck_sack = open('../input.txt').read()

part1(ruck_sack.splitlines())
part2(ruck_sack.splitlines())
