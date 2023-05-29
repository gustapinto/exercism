public class Hamming {
    private int hammingDistance = 0;
    private String leftStrand;
    private String rightStrand;

    public Hamming(String leftStrand, String rightStrand) throws IllegalArgumentException {
        if (leftStrand.length() != rightStrand.length()) {
            if (leftStrand.isEmpty()) {
                throw new IllegalArgumentException("left strand must not be empty.");
            }

            if (rightStrand.isEmpty()) {
                throw new IllegalArgumentException("right strand must not be empty.");
            }

            throw new IllegalArgumentException("leftStrand and rightStrand must be of equal length.");
        }

        this.leftStrand = leftStrand;
        this.rightStrand = rightStrand;
    }

    public int getHammingDistance() {
        for (int i = 0; i < this.leftStrand.length(); i++) {
            if (this.leftStrand.charAt(i) != this.rightStrand.charAt(i)) {
                this.hammingDistance++;
            }
        }

        return this.hammingDistance;
    }
}
