#include <cmath>

const double BILLABLE_HOURS{8.0};
const double BILLABLE_DAYS{22.0};

double daily_rate(double hourly_rate) {
    return hourly_rate * BILLABLE_HOURS;
}

double apply_discount(double before_discount, double discount) {
    double discount_value = before_discount * (discount / 100.00);

    return before_discount - discount_value;
}

int monthly_rate(double hourly_rate, double discount) {
    double total_monthly_rate = daily_rate(hourly_rate) * BILLABLE_DAYS;
    double discounted_monthly_rate = apply_discount(total_monthly_rate, discount);

    // std::ceil(<double>) -> Arredonda para cima um número decimal para inteiro
    return std::ceil(discounted_monthly_rate);
}

int days_in_budget(int budget, double hourly_rate, double discount) {
    double daily_cost = daily_rate(hourly_rate);
    double discounted_daily_cost = apply_discount(daily_cost, discount);

    // std::ceil(<double>) -> Arredonda para baixo um número decimal para inteiro
    return std::floor(budget / discounted_daily_cost);
}