use std::fs;

fn main() {
    let ledger = fs::read_to_string("../input.txt").unwrap();
    let mut elves = ledger
        .split("\n\n")
        .map(|elf| {
            elf.split('\n')
            .map(|item| str::parse::<u64>(item).unwrap_or(0))
            .sum()
        })
        .collect::<Vec<u64>>();

    elves.sort_by(|a, b| b.cmp(a));

    println!("best: {:?}", elves[0]);
    println!("second: {:?}", elves[1]);
    println!("third: {:?}", elves[2]);
    println!("Top 3: {}", elves[0] + elves[1] + elves[2]);
}