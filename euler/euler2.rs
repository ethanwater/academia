pub fn euler2a(n: u32) -> u32 {
    if n <= 1 {
        return n;
    } else {
        return euler2a(n - 1) + euler2a(n - 2);
    }
}

pub fn euler2b(n: u32) -> u32 {
    let mut sum = 0;
    for i in 2..=n {
        let x: u32 = euler2a(i);
        if x % 2 == 0 {
            sum += x;
        }
    }
    sum
}
