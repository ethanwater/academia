pub fn euler10() -> u32 {
    let mut sum: u32 = 0;
    let mut n: u32 = 2;
    while n < 2_000_000 {
        if euler7a(n) == true {
            sum += n;
            println!("{}", n);
        }
        n += 1;
    }
    sum
}
