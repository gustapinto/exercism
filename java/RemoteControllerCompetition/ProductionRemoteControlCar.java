package RemoteControllerCompetition;

class ProductionRemoteControlCar implements RemoteControlCar, Comparable<ProductionRemoteControlCar> {

    private int distanceTravelled = 0;
    private int victories = 0;

    public void drive() {
        this.distanceTravelled += 10;
    }

    public int getDistanceTravelled() {
        return this.distanceTravelled;
    }

    public int getNumberOfVictories() {
        return this.victories;
    }

    public void setNumberOfVictories(int numberofVictories) {
        this.victories = numberofVictories;
    }

    public int compareTo(ProductionRemoteControlCar car) {
        // Integer.compare(<n1>, <n2>) implementa um método de comparação
        // compativel com a interface Comparable<T> para números inteiros
        return Integer.compare(this.getNumberOfVictories(), car.getNumberOfVictories());
    }

}
