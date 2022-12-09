use regex::Regex;
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

    println!("part 1: {}", part(&lines, 2));
    println!("part 2: {}", part(&lines, 10));
}

fn part(lines: &Vec<&str>, rope_len: usize) -> usize {
    let re_move = Regex::new(r"(L|R|U|D) (\d+)").unwrap();

    let mut rope: Vec<(i32, i32)> = vec![(0, 0); rope_len];
    let mut visited: HashSet<(i32, i32)> = HashSet::new();
    visited.insert((0, 0));

    for line in lines {
        for cap in re_move.captures_iter(line) {
            // should only match once
            let dir = &cap[1];
            let steps = cap[2].parse::<usize>().unwrap();
            for _ in 0..steps {
                // update head
                let mut head = rope[0];
                match dir {
                    "L" => head.0 -= 1,
                    "R" => head.0 += 1,
                    "U" => head.1 += 1,
                    "D" => head.1 -= 1,
                    _ => println!("invalid"),
                }
                rope[0] = head;

                // update tail
                for i in 0..rope_len - 1 {
                    let head = rope[i];
                    let mut tail = rope[i + 1];

                    let x_dist = i32::abs(head.0 - tail.0);
                    let y_dist = i32::abs(head.1 - tail.1);
                    if x_dist <= 1 && y_dist <= 1 {
                        // overlapping or adjacent
                        break;
                    }
                    tail.0 += (head.0 - tail.0).signum();
                    tail.1 += (head.1 - tail.1).signum();

                    rope[i + 1] = tail;
                }

                visited.insert(rope.last().cloned().unwrap());
            }
        }
    }
    return visited.len();
}
