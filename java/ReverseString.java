class ReverseString {

    String reverse(String inputString) {
        var builder = new StringBuilder(inputString);
        // StringBuilder.reverse() inverte a String usada dentro do construtor
        // do builder
        var reversed = builder.reverse();

        return reversed.toString();
    }

}
