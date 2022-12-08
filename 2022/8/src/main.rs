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
    let mut grid: Vec<Vec<char>> = vec![];
    for line in lines {
        grid.push(line.chars().collect());
    }

    let mut max_score = 0;
    let mut visibles = 2 * grid.len() + 2 * (grid.len() - 2);
    for i in 1..grid.len() - 1 {
        for j in 1..grid.len() - 1 {
            let mut visible_from_left = true;
            let mut visible_from_right = true;
            let mut visible_from_top = true;
            let mut visible_from_bottom = true;
            let mut score = 1;
            let height = grid[i][j];

            // look left
            for k in (0..j).rev() {
                if grid[i][k] >= height {
                    visible_from_left = false;
                    score *= j - k;
                    break;
                }
            }
            if visible_from_left {
                score *= j;
            }

            // look right
            for k in j + 1..grid.len() {
                if grid[i][k] >= height {
                    visible_from_right = false;
                    score *= k - j;
                    break;
                }
            }
            if visible_from_right {
                score *= grid.len() - j - 1;
            }

            // look up
            for k in (0..i).rev() {
                if grid[k][j] >= height {
                    visible_from_top = false;
                    score *= i - k;
                    break;
                }
            }
            if visible_from_top {
                score *= i;
            }

            // look down
            for k in i + 1..grid.len() {
                if grid[k][j] >= height {
                    visible_from_bottom = false;
                    score *= k - i;
                    break;
                }
            }
            if visible_from_bottom {
                score *= grid.len() - i - 1;
            }

            if visible_from_left || visible_from_right || visible_from_top || visible_from_bottom {
                visibles += 1;
            }

            if score > max_score {
                max_score = score;
            }
        }
    }

    println!("part 1: {}", visibles);
    println!("part 2: {}", max_score);
}
