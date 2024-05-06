pub fn reverse(input: &str) -> String {
    return input
        .chars()
        .rev() // <Chars>.rev() -> Reverte o iterador
        .collect::<String>(); // <Rev>.collect::<T>() -> Coleta todos os itens do iterador em uma coleção, nesse caso String
}
