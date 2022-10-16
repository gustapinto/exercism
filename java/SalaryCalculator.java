public class SalaryCalculator {
    private final double BASE_SALARY = 1000.0;
    private final double MAX_SALARY = 2000.0;

    public double multiplierPerDaysSkipped(int daysSkipped) {
        return daysSkipped >= 5 ? 0.85 : 1.0;
    }

    public int multiplierPerProductsSold(int productsSold) {
        return productsSold >= 20 ? 13 : 10;
    }

    public double bonusForProductSold(int productsSold) {
        return productsSold * this.multiplierPerProductsSold(productsSold);
    }

    public double finalSalary(int daysSkipped, int productsSold) {
        double multiplier = this.multiplierPerDaysSkipped(daysSkipped);
        double bonus = this.bonusForProductSold(productsSold);

        double salary = (this.BASE_SALARY * multiplier) + bonus;

        return salary > this.MAX_SALARY ? this.MAX_SALARY : salary;
    }

}
