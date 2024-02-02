pub fn collatz(mut n: u128) -> Vec<u128> {
    let mut seq: Vec<u128> = Vec::new();
    loop {
        if n.rem_euclid(2) == 0 {
            seq.push(n);
            n = n / 2;
        } else {
            seq.push(n);
            n = (n * 3) + 1;
        }
        if n == 2 {
            seq.push(n);
            seq.push(1);
            break;
        }
    }

    seq
}

pub fn euler14(n: u128) -> u128 {
    let mut max_seq_len: usize = 0;
    let mut max_seq = Vec::new();
    for i in 1..=n {
        let curr = collatz(i);
        if curr.len() > max_seq_len {
            max_seq_len = curr.len() - 1; //we dont consider the 1
            max_seq = curr;
        }
    }

    println!("seq_len:{max_seq_len}\nseq:{:?}", max_seq);
    return max_seq[0];
}
