public class CarsAssemble {

    private final int CARS_PER_HOUR = 221;

    public double productionRatePerHour(int speed) {
        double totalProducedCars = this.CARS_PER_HOUR * speed;

        if (speed >= 5 && speed <= 8) {
            totalProducedCars *= 0.9;
        } else if (speed == 9) {
            totalProducedCars *= 0.8;
        } else if (speed == 10) {
            totalProducedCars *= 0.77;
        }

        return totalProducedCars;
    }

    public int workingItemsPerMinute(int speed) {
        double itemsPerHour = this.productionRatePerHour(speed);
        double itemsPerMinute = itemsPerHour / 60;

        return (int) itemsPerMinute;
    }

}
