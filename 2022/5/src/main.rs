use regex::Regex;
use std::env;
use std::fs;
use std::io::Read;

#[derive(Debug)]
struct Move {
    moves: usize,
    from: usize,
    to: usize,
}

fn input() -> String {
    let args: Vec<String> = env::args().collect();
    let file_path = &args[1];
    let mut file = fs::File::open(file_path).expect("Failed to open file");
    let mut buf = String::new();
    file.read_to_string(&mut buf).ok();
    return buf;
}

fn main() {
    let input = input();

    let parts: Vec<&str> = input.split("\n\n").collect();
    let stacks_str: Vec<&str> = parts[0].lines().collect();
    let instrs_str: Vec<&str> = parts[1].lines().collect();

    // most tedious parsing ever
    let stack_count: usize = stacks_str[stacks_str.len() - 2]
        .chars()
        .filter(|c| *c == '[')
        .count()
        .try_into()
        .unwrap();
    let mut stacks = vec![vec!["."; 1]; stack_count];
    for line in stacks_str.iter().rev().skip(1) {
        let targets = line.split(" ").collect::<Vec<&str>>();
        let mut i = 0;
        let mut stack_idx = 0;
        while i < targets.len() {
            if targets[i] == "" {
                stack_idx += 1;
                i += 1;
                let mut gaps = 0;
                while i < targets.len() && targets[i] == "" {
                    if gaps != 0 && gaps % 4 == 0 {
                        stack_idx += 1;
                    }
                    i += 1;
                    gaps += 1;
                }
            }
            if i < targets.len() {
                stacks[stack_idx].push(&targets[i][1..2]);
                stack_idx += 1;
                i += 1;
            }
        }
    }

    // not really necessary but lets try some regex
    let re_move = Regex::new(r"move (\d+) from (\d) to (\d)").unwrap();
    let mut instrs: Vec<Move> = Vec::new();
    for instr in instrs_str {
        for cap in re_move.captures_iter(instr) {
            // should only match once
            instrs.push(Move {
                moves: cap[1].parse::<usize>().unwrap(),
                from: cap[2].parse::<usize>().unwrap() - 1,
                to: cap[3].parse::<usize>().unwrap() - 1,
            });
        }
    }

    part1(&mut stacks.clone(), &mut instrs);
    part2(&mut stacks.clone(), &mut instrs);
}

fn part1(stacks: &mut Vec<Vec<&str>>, instrs: &mut Vec<Move>) -> () {
    for instr in instrs {
        for _i in 0..instr.moves {
            let c = if let Some(c) = stacks[instr.from].pop() {
                c
            } else {
                todo!()
            };
            stacks[instr.to].push(c);
        }
    }
    print!("part 1: ");
    for stack in stacks {
        let c = if let Some(c) = stack.last() {
            c
        } else {
            todo!()
        };
        print!("{c}");
    }
    println!();
}

fn part2(stacks: &mut Vec<Vec<&str>>, instrs: &mut Vec<Move>) -> () {
    for instr in instrs {
        let len = stacks[instr.from].len();
        let mut boxes: Vec<_> = stacks[instr.from].drain((len - instr.moves)..).collect();
        stacks[instr.to].append(&mut boxes);
    }
    print!("part 2: ");
    for stack in stacks {
        let c = if let Some(c) = stack.last() {
            c
        } else {
            todo!()
        };
        print!("{c}");
    }
    println!();
}
