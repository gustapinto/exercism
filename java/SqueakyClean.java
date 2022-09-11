class SqueakyClean {

    static String clean(String identifier) {
        var builder = new StringBuilder();

        for (int i = 0; i < identifier.length(); i++) {
            char ch = identifier.charAt(i);

            if (Character.isISOControl(ch)) {
                builder.append("CTRL");
                continue;
            }

            if (Character.isWhitespace(ch)) {
                builder.append("_");
                continue;
            }

            // Converte de kebab-case para CamelCase
            if (ch == '-') {
                // Para converter para CamelCase pega a próxima letra
                // (incrementando o contador) e converte para maiúsculo,
                // continuando com o fluxo de execução
                i++;

                char next = identifier.charAt(i);
                char nextUpper = Character.toUpperCase(next);

                ch = nextUpper;
            }

            // Ignora caracteres do alfabeto Grego
            if (ch >= 'ά' && ch <= 'ω') {
                continue;
            }

            if (Character.isLetter(ch)) {
                builder.append(ch);
            }
        }

        return builder.toString();
    }

}
