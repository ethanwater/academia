use regex::Regex;
use std::fs;

fn part1(input: &Vec<&str>) {
    let mut sum = 0;
    let mut game_id = 0;

    for game in input {
        game_id += 1;
        let re = Regex::new(r"(\d+)\s*(blue|red|green)").unwrap();
        let mut possible = true;

        for hit in re.captures_iter(game) {
            if let (Some(num), Some(col)) = (hit.get(1), hit.get(2)) {
                let number = num.as_str().parse::<usize>().unwrap_or(0);
                let color = col.as_str();
                match color {
                    "red" => {
                        if number > 12 {
                            possible = false;
                        }
                    }
                    "green" => {
                        if number > 13 {
                            possible = false;
                        }
                    }
                    "blue" => {
                        if number > 14 {
                            possible = false;
                        }
                    }
                    _ => todo!(),
                }
            }
        }
        if possible {
            sum += game_id;
        }
    }
    dbg!(sum);
}

fn part2(input: &Vec<&str>) {
    let mut sum = 0;

    for game in input {
        let re = Regex::new(r"(\d+)\s*(blue|red|green)").unwrap();
        let (mut red, mut green, mut blue) = (0, 0, 0);

        for hit in re.captures_iter(game) {
            if let (Some(num), Some(col)) = (hit.get(1), hit.get(2)) {
                let number = num.as_str().parse::<usize>().unwrap_or(0);
                let color = col.as_str();
                match color {
                    "red" => {
                        if number > red {
                            red = number;
                        }
                    }
                    "green" => {
                        if number > green {
                            green = number;
                        }
                    }
                    "blue" => {
                        if number > blue {
                            blue = number;
                        }
                    }
                    _ => todo!(),
                }
            }
        }

        let cubic_power = red * green * blue;
        sum += cubic_power;
    }
    dbg!(sum);
}

fn main() {
    let file = fs::read_to_string("./cube_conundrum.txt").unwrap();
    let test_cases = file.lines().collect::<Vec<&str>>();
    part1(&test_cases);
    part2(&test_cases);
}
