use std::cmp;
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
    let v = lines.len() * lines[0].len();
    let mut climbable = vec![vec![false; v]; v];
    let mut possible_starts: Vec<usize> = vec![];

    let mut start: usize = 0;
    let mut end: usize = 0;
    let n = lines.len();
    let m = lines[0].chars().collect::<Vec<char>>().len();
    for i in 0..n {
        let line: Vec<char> = lines[i].chars().collect();
        for j in 0..m {
            let rock = i * m + j;

            // right
            if j + 1 < m {
                let right = i * m + j + 1;
                climbable[rock][right] |= is_climbable(height_of(line[j]), height_of(line[j + 1]));
                climbable[right][rock] |= is_climbable(height_of(line[j + 1]), height_of(line[j]));
            }

            // down
            if i + 1 < n {
                let line_next: Vec<char> = lines[i + 1].chars().collect();
                let down = (i + 1) * m + j;
                climbable[rock][down] |= is_climbable(height_of(line[j]), height_of(line_next[j]));
                climbable[down][rock] |= is_climbable(height_of(line_next[j]), height_of(line[j]));
            }

            // start and end marker positions
            if line[j] == 'S' {
                start = rock;
            }
            if line[j] == 'E' {
                end = rock;
            }
            if line[j] == 'a' {
                // S omitted even tough its actually an 'a'
                possible_starts.push(rock);
            }
        }
    }

    println!("part 1: {}", shortest_path(&climbable, start, end));

    let mut min_steps = u32::MAX;
    for p in possible_starts {
        min_steps = cmp::min(min_steps, shortest_path(&climbable, p, end));
    }
    println!("part 2: {}", min_steps);
}

fn height_of(c: char) -> u32 {
    let n: u32 = c.try_into().unwrap();
    if 97 <= n && n <= 122 {
        return n - 97 + 1;
    } else if c == 'S' {
        return 1;
    } else if c == 'E' {
        return 26;
    } else {
        panic!("invalid: {}/{}", c, n);
    }
}

fn is_climbable(a: u32, b: u32) -> bool {
    return a == b || a + 1 == b || a > b;
}

fn shortest_path(climbable: &Vec<Vec<bool>>, start: usize, end: usize) -> u32 {
    let n = climbable.len();
    let mut visited = vec![false; n];
    let mut dist = vec![u32::MAX; n];
    dist[start] = 0;

    for _ in 0..n {
        let mut u = 0;
        let mut min = u32::MAX;
        let mut found = false;
        for w in 0..n {
            if dist[w] < min && !visited[w] {
                min = dist[w];
                u = w;
                found = true;
            }
        }
        if !found {
            // no new minimum found this means we cannot reach the top from here
            break;
        }

        visited[u] = true;
        for v in 0..n {
            if climbable[u][v] && !visited[v] {
                dist[v] = dist[u] + 1;
            }
        }
    }
    return dist[end];
}
