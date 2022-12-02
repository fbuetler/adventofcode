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

// needed for both parts
const ROCK_SCORE: u32 = 1;
const PAPER_SCORE: u32 = 2;
const SCISSORS_SCORE: u32 = 3;

const WIN_SCORE: u32 = 6;
const DRAW_SCORE: u32 = 3;
const LOSE_SCORE: u32 = 0;

const ROCK_OPP: &str = "A";
const PAPER_OPP: &str = "B";
const SCISSORS_OPP: &str = "C";

fn part1(lines: &Vec<&str>) -> () {
    const ROCK_ME: &str = "X";
    const PAPER_ME: &str = "Y";
    const SCISSORS_ME: &str = "Z";

    let mut score = 0u32;
    for line in lines.iter() {
        let round: Vec<&str> = line.split(" ").collect();
        let (opponent, me) = (round[0], round[1]);
        match opponent {
            ROCK_OPP => match me {
                ROCK_ME => score += DRAW_SCORE + ROCK_SCORE,
                PAPER_ME => score += WIN_SCORE + PAPER_SCORE,
                SCISSORS_ME => score += LOSE_SCORE + SCISSORS_SCORE,
                _ => println!("invalid"),
            },
            PAPER_OPP => match me {
                ROCK_ME => score += LOSE_SCORE + ROCK_SCORE,
                PAPER_ME => score += DRAW_SCORE + PAPER_SCORE,
                SCISSORS_ME => score += WIN_SCORE + SCISSORS_SCORE,
                _ => println!("invalid"),
            },
            SCISSORS_OPP => match me {
                ROCK_ME => score += WIN_SCORE + ROCK_SCORE,
                PAPER_ME => score += LOSE_SCORE + PAPER_SCORE,
                SCISSORS_ME => score += DRAW_SCORE + SCISSORS_SCORE,
                _ => println!("invalid"),
            },
            _ => println!("invalid"),
        }
    }
    println!("part 1: {score}");
}

fn part2(lines: &Vec<&str>) -> () {
    const LOSS_NEEDED: &str = "X";
    const DRAW_NEEDED: &str = "Y";
    const WIN_NEEDED: &str = "Z";

    let mut score = 0u32;
    for line in lines.iter() {
        let round: Vec<&str> = line.split(" ").collect();
        let (opponent, me) = (round[0], round[1]);
        match opponent {
            ROCK_OPP => match me {
                LOSS_NEEDED => score += LOSE_SCORE + SCISSORS_SCORE,
                DRAW_NEEDED => score += DRAW_SCORE + ROCK_SCORE,
                WIN_NEEDED => score += WIN_SCORE + PAPER_SCORE,
                _ => println!("invalid"),
            },
            PAPER_OPP => match me {
                LOSS_NEEDED => score += LOSE_SCORE + ROCK_SCORE,
                DRAW_NEEDED => score += DRAW_SCORE + PAPER_SCORE,
                WIN_NEEDED => score += WIN_SCORE + SCISSORS_SCORE,
                _ => println!("invalid"),
            },
            SCISSORS_OPP => match me {
                LOSS_NEEDED => score += LOSE_SCORE + PAPER_SCORE,
                DRAW_NEEDED => score += DRAW_SCORE + SCISSORS_SCORE,
                WIN_NEEDED => score += WIN_SCORE + ROCK_SCORE,
                _ => println!("invalid"),
            },
            _ => println!("invalid"),
        }
    }
    println!("part 2: {score}");
}
