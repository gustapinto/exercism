const BASE: u64 = 2;
const CHESSBOARD_RANGE: std::ops::RangeInclusive<u32> = 1..=64;

pub fn square(s: u32) -> u64 {
    if !CHESSBOARD_RANGE.contains(&s) {
        panic!("Square must be between 1 and 64");
    }

    let expoent: u32 = s-1;
    let result: u64 = BASE.pow(expoent);

    return u64::from(result);
}

pub fn total() -> u64 {
    let mut total: u64 = 0;

    for i in CHESSBOARD_RANGE {
        total += square(i);
    }

    return total;
}