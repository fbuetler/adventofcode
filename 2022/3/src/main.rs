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
    let mut priority = 0;
    for bag in lines {
        let mut a = HashSet::new();
        let mut b = HashSet::new();
        for (i, c) in bag.chars().enumerate() {
            if i < bag.chars().count() / 2 {
                a.insert(c);
            } else {
                b.insert(c);
            }
        }
        let common: char = *a.intersection(&b).collect::<Vec<&char>>()[0];
        priority += get_priority(common);
    }
    println!("part 1: {priority}");
}

fn part2(lines: &Vec<&str>) -> () {
    let mut priority = 0;
    let mut i = 0;
    while i < lines.len() {
        let mut bag1 = HashSet::new();
        let mut bag2 = HashSet::new();
        let mut bag3 = HashSet::new();

        for c in lines[i + 0].chars() {
            bag1.insert(c);
        }
        for c in lines[i + 1].chars() {
            bag2.insert(c);
        }
        for c in lines[i + 2].chars() {
            bag3.insert(c);
        }

        let intersection = bag1.iter().filter(|k| bag2.contains(k) & bag3.contains(k));
        let common: char = *intersection.collect::<Vec<&char>>()[0];
        priority += get_priority(common);

        i += 3;
    }
    println!("part 2: {priority}");
}

fn get_priority(c: char) -> u32 {
    let ord: u32 = c.into();
    if ord <= 90 {
        return ord - 38;
    } else {
        return ord - 96;
    }
}
