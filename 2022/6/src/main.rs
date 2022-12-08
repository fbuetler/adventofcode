use std::collections::HashSet;
use std::env;
use std::fs;
use std::io::Read;

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
    let lines: Vec<&str> = input.lines().collect();

    part1(&lines);
    part2(&lines);
}
fn part1(lines: &Vec<&str>) -> () {
    for line in lines {
        println!("part 1: {}", solve(line, 4));
    }
}

fn part2(lines: &Vec<&str>) -> () {
    for line in lines {
        println!("part 2: {}", solve(line, 14));
    }
}

fn solve(line: &str, unique: usize) -> usize {
    let mut seen: Vec<char> = Vec::new();
    for (i, c) in line.chars().enumerate() {
        seen.push(c);
        if seen.len() > unique {
            seen.remove(0);
        }

        let s: HashSet<char> = HashSet::from_iter(seen.iter().cloned());
        if s.len() == unique {
            return i + 1;
        }
    }
    return 0;
}
