pub fn euler7a(n: u32) -> bool {
    let mut state: bool = true;
    if n <= 1 {
        state = false;
    }
    for i in 2..f64::sqrt(n.into()) as u32 + 1 {
        if n % i == 0 {
            state = false;
        }
    }
    state
}

//pub fn euler7b(n: usize) -> u32 {
//    //find the n-th prime number
//    let mut primes: Vec<u32> = Vec::new();
//    let mut x = 2;
//    while primes.len() < n{
//        if euler7a(x) == true{
//            primes.push(x);
//        }
//        x+=1;
//    }
//
//    primes[n-1]
//}
