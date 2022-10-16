public class ElonsToyCar {
    private int batteryPercent = 100;
    private int drivenDistance = 0;

    public static ElonsToyCar buy() {
        return new ElonsToyCar();
    }

    public String distanceDisplay() {
        return "Driven " + this.drivenDistance + " meters";
    }

    public String batteryDisplay() {
        if (this.batteryPercent == 0) {
            return "Battery empty";
        }

        return "Battery at " + this.batteryPercent + "%";
    }

    public void drive() {
        if (this.batteryPercent > 0) {
            this.drivenDistance += 20;
            this.batteryPercent -= 1;
        }
    }

}
