pub fn euler6(n: u32) -> u32 {
    let vecx: Vec<u32> = (1..n + 1).collect();
    let vecy: Vec<u32> = vecx.clone().into_iter().map(|x| x * x).collect();
    let mut sumx: u32 = vecx.iter().sum();
    sumx = sumx.pow(2);
    let sumy: u32 = vecy.iter().sum();
    sumx - sumy
}
