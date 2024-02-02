pub fn euler1(n: u32) -> u32 {
    let mut sum = 0;
    for num in 1..n {
        if num % 3 == 0 || num % 5 == 0 {
            sum += num;
        }
    }
    sum
}
