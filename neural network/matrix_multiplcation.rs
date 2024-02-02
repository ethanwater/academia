#![allow(unused_imports)]
use core::ops::AddAssign;
use rand;
use std::ops::Index;

pub fn matmul(m: &mut [[f32; 2]; 2], x: &mut [[f32; 2]; 2]) -> [[f32; 2]; 2] {
    let mut product = [[0f32; 2]; 2];

    for i in 0..m.len() {
        for j in 0..2 {
            product[i][j] = 0.;
            for k in 0..x.len() {
                product[i][j] += m[i][k] * x[k][j];
            }
        }
    }

    product
}

//random initialization of matrices
//pub fn uniform(m: &mut [[f32; 2]; 3]) {
//    for (_i, row) in m.iter_mut().enumerate() {
//        for (_j, col) in row.iter_mut().enumerate() {
//            *col = rand::random::<u8>() as f32;
//        }
//    }
//}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn matmul_test() {
        let mut m = [[0f32; 2]; 2];
        m[0][0] = 72.;
        m[0][1] = 93.;
        m[1][0] = 133.;
        m[1][1] = 31.;

        let mut x = [[0f32; 2]; 2];
        x[0][0] = 13.;
        x[0][1] = 142.;
        x[1][0] = 24.;
        x[1][1] = 169.;
        x[2][0] = 34.;

        let mut expected = [[0f32; 2]; 2];
        expected[0][0] = 3168.;
        expected[0][1] = 25941.;
        expected[1][0] = 2473.;
        expected[1][1] = 24125.;

        let dot = matmul(&mut m, &mut x);
        assert_eq!(dot, expected);
    }
}
