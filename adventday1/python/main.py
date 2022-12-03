#!/usr/bin/env python3

ledger = open('../input.txt').read()
elvesRaw = ledger.split("\n\n")
elves = []

for elf in elvesRaw:
    e = sum(int(i) for i in elf.splitlines())
    elves.append(e)

elves.sort(reverse=True)

print("best: {}".format(elves[0]))
print("second: {}".format(elves[1]))
print("third: {}".format(elves[2]))
print("Top 3: {}".format(sum(elves[:3])))