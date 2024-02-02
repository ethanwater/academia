#![allow(dead_code)]
pub fn euler1(n: u32) -> u32 {
    let mut sum = 0;
    for num in 1..n {
        if num % 3 == 0 || num % 5 == 0 {
            sum += num;
        }
    }
    sum
}

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

pub fn euler4a(n: u32) -> bool {
    let (mut c, mut state) = (0, true);
    let nstr = n.to_string();
    while state == true {
        while c <= nstr.len() / 2 {
            if nstr.as_bytes()[c] == nstr.as_bytes()[nstr.len() - c - 1] {
                state = true;
                c += 1;
            } else {
                state = false;
                break;
            }
        }
        break;
    }
    if state == true {
        true
    } else {
        false
    }
}

pub fn euler4b(n: u32) -> u32 {
    //holy fuck, this bug was tough to figure out...
    //an annoying bug occured when i tried to declare upper as ((10**(&n))-1), this was because "**" is not power in Rust
    //explanation:
    //"*" is the arithmetic multiplcation operator,
    //HOWEVER "*" can also be used as a dereference pointer
    //this would cause upper to be 19 since:
    //10*2 => 20 - 1 => 19
    //
    //instead i had to use the pow function to perform 10^2
    //i also had to inclue the _u32 because pow requires explicit types
    let upper: u32 = 10_u32.pow(n);
    let mut max = 0;
    let lower: u32 = 1 + upper / 10;
    for x in (lower - 1..upper).rev() {
        for y in (lower - 1..upper).rev() {
            let product: u32 = x * y;
            if product < max {
                break;
            }
            if euler4a(product) == true && product > max {
                max = product;
            }
        }
    }
    max
}

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

pub fn euler6(n: u32) -> u32 {
    let vecx: Vec<u32> = (1..n + 1).collect();
    let vecy: Vec<u32> = vecx.clone().into_iter().map(|x| x * x).collect();
    let mut sumx: u32 = vecx.iter().sum();
    sumx = sumx.pow(2);
    let sumy: u32 = vecy.iter().sum();
    sumx - sumy
}

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

//pub fn euler8() -> Vec<u8> {}

//pub fn euler9a(mut a: u32, mut b: u32, mut c:u32) -> bool {
//    //determines if pythagorean triple or not
//    a = a.pow(2);
//    b = b.pow(2);
//    c = c.pow(2);
//    if a+b == c { true } else { false }
//}

//pub fn euler9b() {
//    let mut a: i32 = 3;
//    let mut b: i32 = 4;
//    let mut c: i32 = 5;
//    for n in 1..50{
//        a = 3*n;
//        b = 4*n;
//        c = 5*n;
//    }
//}

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

pub fn euler11() {}

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

pub fn euler13() {}

pub fn collatz(mut n: u128) -> u128 {
    let mut seq = 0;
    loop {
        if n.rem_euclid(2) == 0 {
            seq += 1;
            n = n / 2;
        } else {
            seq += 1;
            n = (n * 3) + 1;
        }
        if n == 2 {
            seq += 1; // 1 isnt included
            break;
        }
    }

    seq
}

pub fn euler14(n: u128) {
    let mut max_seq_len: u128 = 0;
    let mut max = 0;
    let mut i = 2;
    while i < n {
        i += 1;
        let len = collatz(i);
        if len >= max_seq_len {
            max_seq_len = len;
            max = i;
        }
        if i > n {
            break;
        }
    }

    println!("max_seq:{max}");
}

pub fn factorial(n: usize) -> usize {
    let mut ans = 1;
    for i in 1..n {
        ans += i * ans;
    }
    ans
}
pub fn euler15(n: u32) {
    //lattice levels are all even numbers, skip all odds
    //lattice level values = n*2 of odd level values
    //row 1: 1 and 1 has no value
    //row 2: 1, 2, 1 has value 2 (row 1*2)

    //solution one: works in theory, but overflows
    //let x = n * 2;
    //let routes = factorial(x) / (factorial(x - n) * factorial(n));

    //solution two:

    //let mut init: u32 = 2;
    //let mut level: u32 = 0;
    //let mut current_neighbor = 1;
    //while level <= 6 {
    //    if level.rem_euclid(2) == 0 {
    //        println!("{level}| init:{init} current_neighbor:{current_neighbor}");
    //        init += current_neighbor;
    //        current_neighbor += init;
    //    } else {
    //        init = current_neighbor * 2;
    //        println!("{level}| init:{init} current_neighbor:{current_neighbor}");
    //        current_neighbor += init;
    //    }

    //    level += 1;
    //}
}
