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
    let mut signal_strength = 0;
    let mut cycles: i32 = 1;
    let mut register_x: i32 = 1;
    let mut crt = vec![vec!["_"; 41]; 6];

    for line in lines {
        let parts: Vec<&str> = line.split(" ").collect();
        let instr = parts[0];
        match instr {
            "noop" => {
                signal_strength += calc_signal_strength(cycles, register_x);
                draw_crt(&mut crt, cycles, register_x);
                cycles += 1;
            }
            "addx" => {
                signal_strength += calc_signal_strength(cycles, register_x);
                draw_crt(&mut crt, cycles, register_x);
                cycles += 1;

                signal_strength += calc_signal_strength(cycles, register_x);
                draw_crt(&mut crt, cycles, register_x);
                cycles += 1;

                let value = parts[1].parse::<i32>().unwrap();
                register_x += value;
            }
            _ => println!("invalid"),
        }
    }
    println!("part 1: {signal_strength}");
    println!("part 2:");
    print_crt(&crt);
}

fn calc_signal_strength(cycles: i32, register_x: i32) -> i32 {
    if (cycles + 20) % 40 == 0 {
        return cycles * register_x;
    }
    return 0;
}

fn draw_crt(crt: &mut Vec<Vec<&str>>, cycles: i32, register_x: i32) -> () {
    let row: usize = ((cycles - 1) / 40).try_into().unwrap();
    let col: usize = ((cycles - 1) % 40 + 1).try_into().unwrap();

    let col_wrapper: i32 = (col).try_into().unwrap();
    if col_wrapper == register_x || col_wrapper == register_x + 1 || col_wrapper == register_x + 2 {
        crt[row][col] = "#";
    } else {
        crt[row][col] = ".";
    }
}

fn print_crt(crt: &Vec<Vec<&str>>) -> () {
    // draw crt
    for i in 0..6 {
        for j in 1..41 {
            print!("{}", crt[i][j]);
        }
        println!()
    }
}
