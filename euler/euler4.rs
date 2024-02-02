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
