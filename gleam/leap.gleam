pub fn is_leap_year(year: Int) -> Bool {
  let is_divisible_by_4: Bool = year % 4 == 0
  let is_divisible_by_100: Bool = year % 100 == 0
  let is_divisible_by_400: Bool = year % 400 == 0

  { is_divisible_by_4 && !is_divisible_by_100} || { is_divisible_by_100 && is_divisible_by_400 }
}
