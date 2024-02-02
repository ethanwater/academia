pub fn euler12slow(input: u128) {
    let mut c: u128 = 1;
    let mut divs: u128 = 0;
    let mut sum: u128 = 0;
    loop {
        for i in 1..=c {
            sum += i;
        }
        for x in 1..=sum {
            if sum.rem_euclid(x) == 0 {
                divs += 1;
            }
        }
        if divs > input {
            break;
        }
        c += 1;
        sum = 0;
        divs = 0;
    }
    println!("{sum}:{divs}")
}
