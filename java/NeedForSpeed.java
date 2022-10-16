class NeedForSpeed {

    private int batteryDrain;
    private int speed;

    private int battery = 100;
    private int distanceDriven = 0;

    public NeedForSpeed(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public int getBatteryDrain() {
        return this.batteryDrain;
    }

    public int getSpeed() {
        return this.speed;
    }

    public boolean batteryDrained() {
        return this.battery == 0;
    }

    public int distanceDriven() {
        return this.distanceDriven;
    }

    public void drive() {
        if (!this.batteryDrained()) {
            this.distanceDriven += this.speed;
            this.battery -= this.batteryDrain;
        }
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }

}

class RaceTrack {

    private int distance;

    public RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean tryFinishTrack(NeedForSpeed car) {
        int possibleDrives = 100 / car.getBatteryDrain();
        int driveableDistance = possibleDrives * car.getSpeed();

        return driveableDistance >= this.distance;
    }

}
