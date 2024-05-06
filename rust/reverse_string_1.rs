pub fn reverse(input: &str) -> String {
    // String::from(<str>) -> Cria um novo tipo String a partir do string slice <str>
    let mut reversed_input: String = String::from("");

    // <str>.chars() -> Retorna um iterador sobre os caracteres da string
    for c in input.chars() {
        // <char>.to_string() -> Converte um char para String
        let char_string: String = c.to_string();
        // <String>[..] -> Converte uma String para &str
        let char_str: &str = &char_string[..];

        // <String>.insert_str(indíce, <str>) -> Adiciona a str no indíce especificado
        reversed_input.insert_str(0, char_str)
    }

    return reversed_input;
}
