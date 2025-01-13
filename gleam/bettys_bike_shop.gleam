import gleam/int
import gleam/float

const pences_per_pound: Float = 100.0

pub fn pence_to_pounds(pence: Int) -> Float {
  int.to_float(pence) /. pences_per_pound
}

pub fn pounds_to_string(pounds: Float) -> String {
  "£" <> float.to_string(pounds)
}
