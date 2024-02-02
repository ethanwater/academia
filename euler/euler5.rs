pub fn euler5(n: u32) -> u32 {
    let (mut count, mut x) = (2, 0);
    while 1 > 0 {
        for i in 1..=n {
            if count % i == 0 {
                x += 1;
            } else {
                count += 1;
                x = 0
            }
        }
        if x >= n - 1 {
            break;
        }
    }
    count
}
