use std::time::Instant;

mod euler;

fn main() {
    let now = Instant::now();
    let elapsed = now.elapsed();
    println!("time elapsed: {:?}", elapsed);
}
