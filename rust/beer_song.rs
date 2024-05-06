pub fn verse(n: u32) -> String {
    return match n {
        0 => "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n".to_string(),
        1 => "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n".to_string(),
        2 => format!("{} bottles of beer on the wall, {} bottles of beer.\nTake one down and pass it around, {} bottle of beer on the wall.\n", n, n, (n-1)).to_string(),
        _ => format!("{} bottles of beer on the wall, {} bottles of beer.\nTake one down and pass it around, {} bottles of beer on the wall.\n", n, n, (n-1)).to_string(),
    };
}

pub fn sing(start: u32, end: u32) -> String {
    let mut song = String::from("");

    // <int>..<int> -> Cria um iterador numérico exclusivo entre o primeiro e o segundo valor
    // <int>..=<int> -> Cria um iterador numérico inclusivo entre o primeiro e o segundo valor
    for n in (end..=start).rev() {
        let mut _verse = verse(n);

        if (n != start) {
            // format!(<str>) -> Formata uma string slice, retornando uma String
            _verse = format!("\n{}", _verse);
        }

        // <String>.push_str(&<str>) -> Appenda um <str> no final da String
        song.push_str(&_verse);
    }

    return song;
}
