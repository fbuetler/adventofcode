use std::env;
use std::fs;

// my first rust program ever!

fn main() {
    let args: Vec<String> = env::args().collect();
    let file_path = &args[1];
    let content = fs::read_to_string(file_path).expect("Failed to read file");

    let mut i = 0usize;
    let lines: Vec<&str> = content.lines().collect();
    let mut cals = vec![];
    while i < lines.len() {
        let mut c = 0u32;
        while i < lines.len() && lines[i].chars().count() != 0 {
            c += lines[i].parse::<u32>().unwrap();
            i += 1;
        }
        cals.push(c);
        i += 1;
    }

    cals.sort_by(|a, b| b.cmp(a));

    println!(
        "part 1: {}\npart 2: {}",
        cals[0],
        cals[0] + cals[1] + cals[2]
    )
}
