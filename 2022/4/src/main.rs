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

    part(&lines);
}

fn part(lines: &Vec<&str>) -> () {
    let mut dups = 0;
    let mut overlaps = 0;
    for line in lines {
        let elfs: Vec<&str> = line.split(",").collect();
        let range1: Vec<&str> = elfs[0].split("-").collect();
        let range2: Vec<&str> = elfs[1].split("-").collect();

        let mut range_a: HashSet<u32> = HashSet::new();
        let mut range_b: HashSet<u32> = HashSet::new();

        let limit_a: u32 = range1[1].parse().unwrap();
        let limit_b: u32 = range2[1].parse().unwrap();

        let mut a: u32 = range1[0].parse().unwrap();
        let mut b: u32 = range2[0].parse().unwrap();
        while a <= limit_a {
            range_a.insert(a);
            a += 1;
        }
        while b <= limit_b {
            range_b.insert(b);
            b += 1;
        }

        let shared: HashSet<&u32> = range_a.intersection(&range_b).collect();
        if shared.len() == range_a.len() || shared.len() == range_b.len() {
            dups += 1;
        }
        if shared.len() != 0 {
            overlaps += 1;
        }
    }
    println!("part 1: {dups}");
    println!("part 2: {overlaps}");
}
