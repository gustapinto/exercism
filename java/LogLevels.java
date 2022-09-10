public class LogLevels {

    public static String message(String logLine) {
        // String.split(...) separa uma String em uma lista de Strings a partir
        // de uma substring ou de uma expressão Regex
        String[] splitted = logLine.split(":");
        String message = splitted[1];

        // String.trim(...) remove os espaços em branco no começo e no final da
        // String, atuando também com caracteres unicode
        return message.trim();
    }

    public static String logLevel(String logLine) {
        // String.indexOf(...) obtém o indíce de uma substring
        int start = logLine.indexOf("[") + 1; // +1 pega o próximo elemento
        int end = logLine.indexOf("]");

        // String.substring(...) obtém uma substring a partir do indice inicial
        // e do indice final
        String level = logLine.substring(start, end);

        // String.toLowerCase() converte a String para letras minúsculas
        String lowerCaseLevel = level.toLowerCase();

        return lowerCaseLevel.trim();
    }

    public static String reformat(String logLine) {
        String message = LogLevels.message(logLine);
        String level = LogLevels.logLevel(logLine);

        return message + " (" + level + ")";
    }

}
