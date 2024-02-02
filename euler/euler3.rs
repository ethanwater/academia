pub fn euler3(temp: u32) -> u32 {
    let (mut x, mut n, mut factors) = (2, temp, Vec::new());
    while n > 1 {
        while n % x == 0 {
            factors.push(x);
            n /= x;
        }
        x += 1;
        if x * x > n {
            if n > 1 {
                factors.push(n);
            }
            break;
        }
    }
    factors[factors.len() - 1]
}
