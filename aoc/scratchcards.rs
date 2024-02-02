use regex::Regex;
use std::fs;

fn part1(scratchcards: &Vec<&str>) -> i32 {
    let re = Regex::new(r"\s*\|?\s+").unwrap();
    let mut sum = 0;

    for card in scratchcards {
        let mut cardsum = 0;
        let split = card.split('|').collect::<Vec<&str>>();

        let mut left: Vec<&str> = re.split(split[0]).collect();
        left.retain(|&x| x.to_string().parse::<i32>().unwrap_or(-1) >= 0);

        let mut right: Vec<&str> = re.split(split[1]).collect();
        right.retain(|&x| x.to_string().parse::<i32>().unwrap_or(-1) >= 0);

        for guess in right {
            if left.contains(&guess) {
                if cardsum < 2 {
                    cardsum += 1;
                } else {
                    cardsum = cardsum * 2;
                }
            }
        }

        sum += cardsum;
    }
    sum
}

fn part2(scratchcards: &Vec<&str>) -> i32 {
    let mut cards = vec![1; scratchcards.len()];

    for i in 0..scratchcards.len() {
        let n = matches(scratchcards[i]);
        for j in (i + 1)..std::cmp::min(i + 1 + n as usize, scratchcards.len()) {
            cards[j] += cards[i];
        }
    }

    cards.iter().sum()
}

fn matches(card: &str) -> i32 {
    let re = Regex::new(r"\s*\|?\s+").unwrap();
    let split = card.split('|').collect::<Vec<&str>>();

    let left: Vec<i32> = re
        .split(split[0])
        .filter_map(|x| x.parse::<i32>().ok())
        .filter(|&parsed| parsed >= 0)
        .collect();

    let right: Vec<i32> = re
        .split(split[1])
        .filter_map(|x| x.parse::<i32>().ok())
        .filter(|&parsed| parsed >= 0)
        .collect();

    let winning_numbers = right.iter().filter(|&guess| left.contains(guess)).count() as i32;
    winning_numbers
}

fn main() {
    let file = fs::read_to_string("./scratchcards.txt").unwrap();
    let scratchcards = file.lines().collect::<Vec<&str>>();
    dbg!(part2(&scratchcards));
}
