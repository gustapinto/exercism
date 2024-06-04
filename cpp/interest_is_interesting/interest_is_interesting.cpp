double interest_rate(double balance) {
    if (balance < 0) {
        return 3.213;
    }

    if (balance < 1000) {
        return 0.5;
    }

    if (balance < 5000) {
        return 1.621;
    }

    return 2.475;
}

double yearly_interest(double balance) {
    double rate = interest_rate(balance);
    double rate_percentage = rate / 100;

    return balance * rate_percentage;
}

double annual_balance_update(double balance) {
    return balance + yearly_interest(balance);
}

int years_until_desired_balance(double balance, double target_balance) {
    double b{balance};
    int years{};

    while (b < target_balance) {
        b = annual_balance_update(b);
        years++;
    }

    return years;
}