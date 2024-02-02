use std::collections::BTreeMap;
use std::fs;

//i use BTreeMap due to its trait of organizing my key

fn part1(input: &str) -> usize {
    let valid: Vec<&str> = vec![
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];
    let mut map: BTreeMap<usize, usize> = BTreeMap::new();

    let mut index = 0;
    for element in input.chars() {
        if element.is_numeric() {
            map.insert(index, element.to_string().parse::<usize>().unwrap());
        }
        index += 1;
    }

    let alpha = map.first_key_value().unwrap().1;
    let omega = map.last_key_value().unwrap().1;

    (alpha * 10) + omega
}

fn part2(input: &str) -> usize {
    let valid: Vec<&str> = vec![
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    let mut map: BTreeMap<usize, usize> = BTreeMap::new();

    let mut index = 0;
    for element in input.chars() {
        if element.is_numeric() {
            map.insert(index, element.to_string().parse::<usize>().unwrap());
        }
        index += 1;
    }

    for word in &valid {
        if input.contains(word) {
            //this gathers all occurences of the "word"
            let v: Vec<(usize, &str)> = input.match_indices(word).collect();

            for entry in &v {
                map.insert(
                    entry.0,
                    //we then get the index of the occurence and increment by 1
                    valid.iter().position(|&r| r == entry.1).unwrap() + 1,
                );
            }
        }
    }

    let alpha = map.first_key_value().unwrap().1;
    let omega = map.last_key_value().unwrap().1;

    (alpha * 10) + omega
}

fn main() {
    let file = fs::read_to_string("./calibration_test_inputs.txt").unwrap();
    let test_cases = file.lines().collect::<Vec<&str>>();

    let mut sum: usize = 0;
    for test in test_cases {
        sum += part1(test);
    }

    let mut sumx: usize = 0;
    for test in test_cases {
        sumx += part2(test);
    }

    dbg!(sum);
}
